// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package proxy

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/berkaroad/saashard/backend"
	// Import mysql backend
	_ "github.com/berkaroad/saashard/backend/mysql"
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/statistic"
	"github.com/berkaroad/saashard/utils/golog"
)

var (
	baseConnID uint32 = 10000
)

// Server proxy daemon
type Server struct {
	sync.Mutex
	cfg     *config.Config
	bindIP  net.IP
	port    int
	hosts   map[string]*backend.DataHost
	nodes   map[string]*backend.DataNode
	schemas map[string]*config.SchemaConfig

	logSQLIndex      int32
	logSQL           [2]string
	slowLogTimeIndex int32
	slowLogTime      [2]int
	allowipsIndex    int32
	allowips         [2][]net.IP

	counter  *statistic.Counter
	listener net.Listener
	running  bool
	conns    map[uint32]*ClientConn
}

// NewServer create proxy.
func NewServer(cfg *config.Config) (*Server, error) {
	p := new(Server)
	p.cfg = cfg
	p.bindIP = net.ParseIP(cfg.BindIP)
	p.port = cfg.ProxyPort
	p.conns = make(map[uint32]*ClientConn)

	p.hosts = make(map[string]*backend.DataHost)
	p.nodes = make(map[string]*backend.DataNode)
	p.schemas = make(map[string]*config.SchemaConfig)

	p.counter = new(statistic.Counter)
	atomic.StoreInt32(&p.logSQLIndex, 0)
	p.logSQL[p.logSQLIndex] = cfg.LogSQL
	atomic.StoreInt32(&p.slowLogTimeIndex, 0)
	p.slowLogTime[p.slowLogTimeIndex] = cfg.SlowLogTime
	if len(cfg.Charset) != 0 {
		cid, ok := mysql.CharsetIds[cfg.Charset]
		if !ok {
			return nil, errors.ErrInvalidCharset
		}
		//change the default charset
		mysql.DEFAULT_CHARSET = cfg.Charset
		mysql.DEFAULT_COLLATION_ID = cid
		mysql.DEFAULT_COLLATION_NAME = mysql.Collations[cid]
	}

	if err := p.parseHosts(); err != nil {
		panic(err)
	}

	if err := p.parseNodes(); err != nil {
		panic(err)
	}

	if err := p.parseSchemas(); err != nil {
		panic(err)
	}

	if err := p.parseAllowIps(); err != nil {
		panic(err)
	}

	var err error
	netProto := "tcp"
	addr := p.bindIP.String() + ":" + strconv.Itoa(p.port)

	p.listener, err = net.Listen(netProto, addr)
	if err != nil {
		return nil, err
	}

	golog.Info("server/proxy", "NewServer", "Server running", 0,
		"netProto",
		netProto,
		"address",
		addr)
	return p, nil
}

// Run proxy server.
func (p *Server) Run() {
	p.running = true

	// flush counter
	go p.flushCounter()

	// proxy
	for p.running {
		conn, err := p.listener.Accept()
		if err != nil {
			golog.Error("server/proxy", "Run", err.Error(), 0)
			continue
		}
		go p.onConn(conn)
	}
}

// Close proxy server.
func (p *Server) Close() {
	p.running = false
	if p.listener != nil {
		p.listener.Close()
	}
}

// GetConnection get connection
func (p *Server) GetConnection(connectionID uint32) *ClientConn {
	defer p.Unlock()

	p.Lock()
	if conn, ok := p.conns[connectionID]; ok {
		return conn
	}
	return nil
}

// RemoveConnection remove connection
func (p *Server) RemoveConnection(connectionID uint32) {
	defer p.Unlock()

	p.Lock()
	if conn, ok := p.conns[connectionID]; ok {
		conn.Close()
		delete(p.conns, connectionID)
	}
}

func (p *Server) onConn(c net.Conn) {
	p.counter.IncrClientConns()
	conn := p.newClientConn(c) //新建一个conn

	defer func() {
		err := recover()
		if err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
			golog.Error("server", "onProxyConn", "error", 0,
				"remoteAddr", c.RemoteAddr().String(),
				"stack", string(buf),
			)
		}

		conn.Close()
		p.counter.DecrClientConns()
	}()

	if allowConnect := conn.IsAllowConnect(); allowConnect == false {
		err := mysql.NewError(mysql.ER_ACCESS_DENIED_ERROR, "ip address access denied by SaaShard.")
		conn.pkg.WriteError(conn.capability, err)
		conn.Close()
		return
	}
	if err := conn.Handshake(); err != nil {
		golog.Error("server/proxy", "onConn", err.Error(), 0)
		c.Close()
		return
	}

	conn.schemas = p.getSchemasByUser(conn.user)
	p.Lock()
	p.conns[conn.connectionID] = conn // save conn
	p.Unlock()
	conn.Run()
}

func (p *Server) newClientConn(co net.Conn) *ClientConn {
	c := new(ClientConn)
	tcpConn := co.(*net.TCPConn)

	//SetNoDelay controls whether the operating system should delay packet transmission
	// in hopes of sending fewer packets (Nagle's algorithm).
	// The default is true (no delay),
	// meaning that data is sent as soon as possible after a Write.
	//I set this option false.
	tcpConn.SetNoDelay(false)
	c.c = tcpConn

	c.pkg = mysql.NewPacketIO(tcpConn)
	c.proxy = p
	c.pkg.Sequence = 0
	c.connectionID = atomic.AddUint32(&baseConnID, 1)
	c.status = mysql.SERVER_STATUS_AUTOCOMMIT
	c.salt, _ = mysql.RandomBuf(20)
	c.backendMasterConns = make(map[*backend.DataNode]backend.Connection)
	c.backendSlaveConns = make(map[*backend.DataNode]backend.Connection)
	c.closed = false
	c.charset = mysql.DEFAULT_CHARSET
	c.collation = mysql.DEFAULT_COLLATION_ID
	c.stmtID = 0
	c.stmts = make(map[uint32]*mysql.Stmt)
	return c
}

func (p *Server) flushCounter() {
	for {
		p.counter.FlushCounter()
		time.Sleep(1 * time.Second)
	}
}

func (p *Server) parseAllowIps() error {
	atomic.StoreInt32(&p.allowipsIndex, 0)
	cfg := p.cfg
	if cfg.AllowIps == nil || len(cfg.AllowIps) == 0 {
		return nil
	}
	ipVec := cfg.AllowIps
	p.allowips[p.allowipsIndex] = make([]net.IP, 0, 10)
	p.allowips[1] = make([]net.IP, 0, 10)
	for _, ip := range ipVec {
		p.allowips[p.allowipsIndex] = append(p.allowips[p.allowipsIndex], net.ParseIP(strings.TrimSpace(ip)))
	}
	return nil
}

func (p *Server) parseHosts() error {
	cfg := p.cfg
	if len(cfg.Hosts) == 0 {
		return errors.ErrNoDataHost
	}
	for _, hostConfig := range cfg.Hosts {
		hostCfg := hostConfig
		if p.hosts[hostCfg.Name] == nil {
			p.hosts[hostCfg.Name] = backend.NewDataHost(hostCfg)
		}
	}
	return nil
}

func (p *Server) parseNodes() error {
	cfg := p.cfg
	if len(cfg.Nodes) == 0 {
		return errors.ErrNoDataNode
	}
	for _, nodeConfig := range cfg.Nodes {
		nodeCfg := nodeConfig
		if p.nodes[nodeCfg.Name] == nil {
			if host, ok := p.hosts[nodeCfg.Host]; ok {
				p.nodes[nodeCfg.Name] = backend.NewDataNode(nodeCfg, host)
			} else {
				return fmt.Errorf("data host '%s' not exists", nodeCfg.Host)
			}
		}
	}
	return nil
}

func (p *Server) parseSchemas() error {
	cfg := p.cfg
	if len(cfg.Schemas) == 0 {
		return errors.ErrNoSchema
	}
	for _, schemaConfig := range cfg.Schemas {
		schema := schemaConfig
		if p.schemas[schema.Name] == nil {
			if len(schema.Nodes) == 0 {
				return fmt.Errorf("no data node in schema '%s'", schema.Name)
			}
			for _, nodeInSchema := range schema.Nodes {
				if p.nodes[nodeInSchema] == nil {
					return fmt.Errorf("data node '%s' not exists", nodeInSchema)
				}
			}
			p.schemas[schema.Name] = &schema
		}
	}
	return nil
}

func (p *Server) getSchemasByUser(user string) map[string]*config.SchemaConfig {
	schemas := make(map[string]*config.SchemaConfig)
	for _, schema := range p.schemas {
		if schema.User == strings.ToLower(user) {
			schemas[schema.Name] = schema
		}
	}
	return schemas
}
