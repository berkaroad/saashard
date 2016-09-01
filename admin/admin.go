package admin

import (
	"net"
	"strconv"

	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/utils/golog"
)

// Server Admin
type Server struct {
	cfg    *config.Config
	bindIP net.IP
	port   int

	listener net.Listener
	running  bool
}

// NewServer create admin.
func NewServer(cfg *config.Config) (*Server, error) {
	admin := new(Server)
	admin.cfg = cfg
	admin.bindIP = net.ParseIP(cfg.BindIP)
	admin.port = cfg.AdminPort

	var err error
	netProto := "tcp"
	addr := admin.bindIP.String() + ":" + strconv.Itoa(admin.port)

	admin.listener, err = net.Listen(netProto, addr)
	if err != nil {
		return nil, err
	}

	golog.Info("server/admin", "NewServer", "Server running", 0,
		"netProto",
		netProto,
		"address",
		addr)
	return admin, nil
}

// Run admin server.
func (admin *Server) Run() {
	admin.running = true

	// proxy
	for admin.running {
		conn, err := admin.listener.Accept()
		if err != nil {
			golog.Error("server/admin", "Run", err.Error(), 0)
			continue
		}
		go admin.onConn(conn)
	}
}

// Close admin server.
func (admin *Server) Close() {
	admin.running = false
	if admin.listener != nil {
		admin.listener.Close()
	}
}

func (admin *Server) onConn(c net.Conn) {
	golog.Info("server/admin", "onConn", c.RemoteAddr().String(), 0)
}
