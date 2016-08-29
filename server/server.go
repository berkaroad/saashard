package server

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

const (
	// Offline = 0
	Offline = iota
	// Online = 1
	Online
	// Unknown = 2
	Unknown
)

// Server is startup endpoint.
type Server struct {
	cfg *config.Config

	bindIP    net.IP
	proxyPort int
	adminPort int
	hosts     map[string]*backend.PhysicalDBPool
	nodes     map[string]*backend.PhysicalDBNode
	schemas   map[string]*config.SchemaConfig

	statusIndex      int32
	status           [2]int32
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

// NewServer is to create a server.
func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)
	s.cfg = cfg
	s.bindIP = net.ParseIP(cfg.BindIP)
	s.proxyPort = cfg.ProxyPort
	s.adminPort = cfg.AdminPort
	s.hosts = make(map[string]*backend.PhysicalDBPool)
	s.nodes = make(map[string]*backend.PhysicalDBNode)
	s.schemas = make(map[string]*config.SchemaConfig)

	s.counter = new(statistic.Counter)
	atomic.StoreInt32(&s.statusIndex, 0)
	s.status[s.statusIndex] = Online
	atomic.StoreInt32(&s.logSQLIndex, 0)
	s.logSQL[s.logSQLIndex] = cfg.LogSQL
	atomic.StoreInt32(&s.slowLogTimeIndex, 0)
	s.slowLogTime[s.slowLogTimeIndex] = cfg.SlowLogTime
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

	if err := s.parseHosts(); err != nil {
		return nil, err
	}

	if err := s.parseNodes(); err != nil {
		return nil, err
	}

	if err := s.parseSchemas(); err != nil {
		return nil, err
	}

	if err := s.parseAllowIps(); err != nil {
		return nil, err
	}

	var err error
	netProto := "tcp"
	proxyAddr := s.bindIP.String() + ":" + strconv.Itoa(s.proxyPort)
	s.listener, err = net.Listen(netProto, proxyAddr)

	if err != nil {
		return nil, err
	}

	golog.Info("server", "NewServer", "Server running", 0,
		"netProto",
		netProto,
		"address",
		proxyAddr)
	return s, nil
}

// Run server.
func (s *Server) Run() error {
	s.running = true

	// flush counter
	go s.flushCounter()

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			golog.Error("server", "Run", err.Error(), 0)
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

// Status of server.
func (s *Server) Status() string {
	var status string
	switch s.status[s.statusIndex] {
	case Online:
		status = "online"
	case Offline:
		status = "offline"
	case Unknown:
		status = "unknown"
	default:
		status = "unknown"
	}
	return status
}

// Close server.
func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) onConn(c net.Conn) {
	s.counter.IncrClientConns()
	conn := s.newClientConn(c) //新建一个conn

	defer func() {
		err := recover()
		if err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
			golog.Error("server", "onConn", "error", 0,
				"remoteAddr", c.RemoteAddr().String(),
				"stack", string(buf),
			)
		}

		conn.Close()
		s.counter.DecrClientConns()
	}()

	if allowConnect := conn.IsAllowConnect(); allowConnect == false {
		err := mysql.NewError(mysql.ER_ACCESS_DENIED_ERROR, "ip address access denied by SaaShard.")
		conn.pkg.WriteError(conn.capability, err)
		conn.Close()
		return
	}
	if err := conn.Handshake(); err != nil {
		golog.Error("server", "onConn", err.Error(), 0)
		c.Close()
		return
	}

	conn.schemas = s.getSchemasByUser(conn.user)
	conn.Run()
}

//TODO
func (s *Server) parseAllowIps() error {
	atomic.StoreInt32(&s.allowipsIndex, 0)
	cfg := s.cfg
	if len(cfg.AllowIps) == 0 {
		return nil
	}
	ipVec := strings.Split(cfg.AllowIps, ",")
	s.allowips[s.allowipsIndex] = make([]net.IP, 0, 10)
	s.allowips[1] = make([]net.IP, 0, 10)
	for _, ip := range ipVec {
		s.allowips[s.allowipsIndex] = append(s.allowips[s.allowipsIndex], net.ParseIP(strings.TrimSpace(ip)))
	}
	return nil
}

func (s *Server) parseHosts() error {
	cfg := s.cfg
	if len(cfg.Hosts) == 0 {
		return errors.ErrNoDataHost
	}
	for _, hostConfig := range cfg.Hosts {
		host := hostConfig
		if s.hosts[host.Name] == nil {
			s.hosts[host.Name] = backend.NewPhysicalDBPool(host)
		}
	}
	return nil
}

func (s *Server) parseNodes() error {
	cfg := s.cfg
	if len(cfg.Nodes) == 0 {
		return errors.ErrNoDataNode
	}
	for _, nodeConfig := range cfg.Nodes {
		node := nodeConfig
		if s.nodes[node.Name] == nil {
			if s.hosts[node.Host] != nil {
				s.nodes[node.Name] = backend.NewPhysicalDBNode(node)
			} else {
				return errors.ErrNoDataHost
			}
		}
	}
	return nil
}

func (s *Server) parseSchemas() error {
	cfg := s.cfg
	if len(cfg.Schemas) == 0 {
		return errors.ErrNoSchema
	}
	for _, schemaConfig := range cfg.Schemas {
		schema := schemaConfig
		if s.schemas[schema.Name] == nil {
			if len(schema.Nodes) == 0 {
				return errors.ErrNoDataNode
			}
			for _, nodeInSchema := range schema.Nodes {
				if s.nodes[nodeInSchema] == nil {
					return errors.ErrNoDataNode
				}
			}
			s.schemas[schema.Name] = &schema
		}
	}
	return nil
}

func (s *Server) flushCounter() {
	for {
		s.counter.FlushCounter()
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) newClientConn(co net.Conn) *ClientConn {
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
	c.proxy = s

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

func (s *Server) getSchemasByUser(user string) map[string]*config.SchemaConfig {
	schemas := make(map[string]*config.SchemaConfig)
	for _, schema := range s.schemas {
		if schema.User == strings.ToLower(user) {
			schemas[schema.Name] = schema
		}
	}
	return schemas
}
