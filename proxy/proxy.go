package proxy

import (
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/berkaroad/saashard/backend"
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/statistic"
	"github.com/berkaroad/saashard/utils/golog"
)

// Server proxy daemon
type Server struct {
	cfg     *config.Config
	bindIP  net.IP
	port    int
	hosts   map[string]*backend.PhysicalDBPool
	nodes   map[string]*backend.PhysicalDBNode
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
}

// NewServer create proxy.
func NewServer(cfg *config.Config) (*Server, error) {
	p := new(Server)
	p.cfg = cfg
	p.bindIP = net.ParseIP(cfg.BindIP)
	p.port = cfg.ProxyPort

	p.hosts = make(map[string]*backend.PhysicalDBPool)
	p.nodes = make(map[string]*backend.PhysicalDBNode)
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
		return nil, err
	}

	if err := p.parseNodes(); err != nil {
		return nil, err
	}

	if err := p.parseSchemas(); err != nil {
		return nil, err
	}

	if err := p.parseAllowIps(); err != nil {
		return nil, err
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

	c.connectionId = atomic.AddUint32(&baseConnId, 1)

	c.status = mysql.SERVER_STATUS_AUTOCOMMIT

	c.salt, _ = mysql.RandomBuf(20)

	// c.txConns = make(map[*backend.Node]*backend.BackendConn)

	c.closed = false

	c.charset = mysql.DEFAULT_CHARSET
	c.collation = mysql.DEFAULT_COLLATION_ID

	c.stmtId = 0
	// c.stmts = make(map[uint32]*Stmt)

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
	if len(cfg.AllowIps) == 0 {
		return nil
	}
	ipVec := strings.Split(cfg.AllowIps, ",")
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
		host := hostConfig
		if p.hosts[host.Name] == nil {
			p.hosts[host.Name] = backend.NewPhysicalDBPool(host)
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
		node := nodeConfig
		if p.nodes[node.Name] == nil {
			if p.hosts[node.Host] != nil {
				p.nodes[node.Name] = backend.NewPhysicalDBNode(node)
			} else {
				return errors.ErrNoDataHost
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
				return errors.ErrNoDataNode
			}
			for _, nodeInSchema := range schema.Nodes {
				if p.nodes[nodeInSchema] == nil {
					return errors.ErrNoDataNode
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
