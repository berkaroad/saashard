package server

import (
	"sync/atomic"

	"github.com/berkaroad/saashard/admin"
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/proxy"
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

	statusIndex int32
	status      [2]int32
	running     bool

	proxy *proxy.Server
	admin *admin.Server
}

// NewServer is to create a server.
func NewServer(cfg *config.Config) (*Server, error) {
	var err error
	s := new(Server)
	s.cfg = cfg

	atomic.StoreInt32(&s.statusIndex, 0)
	s.status[s.statusIndex] = Online

	s.proxy, err = proxy.NewServer(cfg)
	s.admin, err = admin.NewServer(cfg)
	return s, err
}

// Run server.
func (s *Server) Run() {
	s.running = true

	// proxy
	go s.proxy.Run()

	// admin
	s.admin.Run()
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
	s.proxy.Close()
	s.admin.Close()
}
