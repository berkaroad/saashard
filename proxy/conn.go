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
	"sync"

	"github.com/berkaroad/saashard/backend"
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/utils/golog"
)

// ClientConn client <-> proxy
type ClientConn struct {
	sync.Mutex
	pkg                *mysql.PacketIO
	c                  net.Conn
	clientIP           net.IP
	proxy              *Server
	capability         uint32
	connectionID       uint32
	status             uint16
	collation          mysql.CollationID
	charset            string
	user               string
	db                 string
	salt               []byte
	schemas            map[string]*config.SchemaConfig
	backendMasterConns map[*backend.DataNode]backend.Connection
	backendSlaveConns  map[*backend.DataNode]backend.Connection
	nodeInTrans        *backend.DataNode
	closed             bool
	lastInsertID       int64
	affectedRows       int64
	stmtID             uint32
	//stmts map[uint32]*Stmt //prepare相关,client端到proxy的stmt
}

// IsAllowConnect check ip in whitelist.
func (c *ClientConn) IsAllowConnect() bool {
	clientHost, _, err := net.SplitHostPort(c.c.RemoteAddr().String())
	if err != nil {
		fmt.Println(err)
	}
	c.clientIP = net.ParseIP(clientHost)

	ipVec := c.proxy.allowips[c.proxy.allowipsIndex]
	if ipVecLen := len(ipVec); ipVecLen == 0 {
		return true
	}
	for _, ip := range ipVec {
		if ip.Equal(c.clientIP) {
			return true
		}
	}

	golog.Error("server", "IsAllowConnect", "error", mysql.ER_ACCESS_DENIED_ERROR,
		"ip address", c.c.RemoteAddr().String(), " access denied by kindshard.")
	return false
}

// Handshake between client and proxy.
func (c *ClientConn) Handshake() error {
	var err error
	if err = c.pkg.WriteInitialHandshake(c.connectionID, c.salt, mysql.DEFAULT_COLLATION_ID, mysql.DEFAULT_CAPABILITY, c.status); err != nil {
		golog.Error("server", "Handshake", err.Error(),
			c.connectionID, "msg", "send initial handshake error")
		return err
	}

	getDefaultSchemaByUser := func(user string) (string, error) {
		c.schemas = c.proxy.getSchemasByUser(user)
		if len(c.schemas) == 0 {
			return "", errors.ErrNoSchema
		}

		var name string
		for name = range c.schemas {
			break
		}
		return name, nil
	}
	getCredentialsConfigBySchema := func(schema string) (string, string, error) {
		schemaConfig := c.proxy.schemas[schema]
		if schemaConfig == nil {
			return "", "", mysql.NewDefaultError(mysql.ER_BAD_DB_ERROR, c.db)
		}
		return schemaConfig.User, schemaConfig.Password, nil
	}
	c.capability, c.collation, c.user, c.db, err = c.pkg.ReadHandshakeResponse(getDefaultSchemaByUser, c.c.RemoteAddr().String(), c.salt, getCredentialsConfigBySchema)
	if err != nil {
		golog.Error("server", "readHandshakeResponse",
			err.Error(), c.connectionID,
			"msg", "read Handshake Response error")

		c.pkg.WriteError(c.capability, err)

		return err
	}
	c.schemas = c.proxy.getSchemasByUser(c.user)

	if err := c.pkg.WriteOK(c.capability, c.status, nil); err != nil {
		golog.Error("server", "readHandshakeResponse",
			"write ok fail",
			c.connectionID, "error", err.Error())
		return err
	}

	c.pkg.Sequence = 0

	return nil
}

// Run after handshake.
func (c *ClientConn) Run() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]

			golog.Error("ClientConn", "Run",
				err.Error(), 0,
				"stack", string(buf))
		}

		c.Close()
	}()

	for {
		data, err := c.pkg.ReadPacket()

		if err != nil {
			golog.Error("server", "Run",
				err.Error(), c.connectionID,
			)
			return
		}
		if err := c.dispatch(data); err != nil {
			c.proxy.counter.IncrErrLogTotal()
			if len(data) > 1 {
				golog.Error("server", "Run",
					err.Error(), c.connectionID,
					"sql",
					string(data[1:]))
			} else {
				golog.Error("server", "Run",
					err.Error(), c.connectionID,
				)
			}
			c.pkg.WriteError(c.capability, err)
			if err == errors.ErrBadConn {
				c.Close()
			}
		}

		if c.closed {
			return
		}

		c.pkg.Sequence = 0
		c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
	}
}

// Close client
func (c *ClientConn) Close() error {
	if c.closed {
		return nil
	}
	c.nodeInTrans = nil
	for _, conn := range c.backendMasterConns {
		conn.Rollback()
		conn.ReturnConnection()
	}
	for _, conn := range c.backendSlaveConns {
		conn.ReturnConnection()
	}

	c.c.Close()

	c.closed = true

	return nil
}

func (c *ClientConn) dispatch(data []byte) error {
	c.proxy.counter.IncrClientQPS()
	cmd := data[0]
	data = data[1:]

	switch cmd {
	case mysql.COM_QUIT:
		c.Close()
		return nil
	case mysql.COM_QUERY:
		return c.handleQuery(string(data))
	case mysql.COM_PING:
		return c.pkg.WriteOK(c.capability, c.status, nil)
	case mysql.COM_INIT_DB:
		return c.handleInitDB(string(data))
	case mysql.COM_FIELD_LIST:
		return c.handleFieldList(data)
	// case mysql.COM_STMT_PREPARE:
	// return c.handleStmtPrepare(hack.String(data))
	// case mysql.COM_STMT_EXECUTE:
	// return c.handleStmtExecute(data)
	// case mysql.COM_STMT_CLOSE:
	// return c.handleStmtClose(data)
	// case mysql.COM_STMT_SEND_LONG_DATA:
	// return c.handleStmtSendLongData(data)
	// case mysql.COM_STMT_RESET:
	// return c.handleStmtReset(data)
	case mysql.COM_SET_OPTION:
		return c.pkg.WriteEOF(c.capability, 0)
	default:
		msg := fmt.Sprintf("command %d not supported now", cmd)
		golog.Error("ClientConn", "dispatch", msg, 0)
		return mysql.NewError(mysql.ER_UNKNOWN_ERROR, msg)
	}
}
func (c *ClientConn) isInTransaction() bool {
	return c.status&mysql.SERVER_STATUS_IN_TRANS > 0 ||
		!c.isAutoCommit()
}

func (c *ClientConn) isAutoCommit() bool {
	return c.status&mysql.SERVER_STATUS_AUTOCOMMIT > 0
}
