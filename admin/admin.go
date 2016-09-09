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
