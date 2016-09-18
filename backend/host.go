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

package backend

import (
	"strconv"
	"strings"

	"github.com/berkaroad/saashard/config"
)

// DataHost is data host.
type DataHost struct {
	Name             string
	MaxConnNum       int
	DownAfterNoAlive int
	PingInterval     int
	Master           *DBHost
	Slaves           []*DBHost
}

// NewDataHost new host.
func NewDataHost(hostCfg config.HostConfig) *DataHost {
	h := new(DataHost)
	h.Name = hostCfg.Name
	h.MaxConnNum = hostCfg.MaxConnNum
	h.DownAfterNoAlive = hostCfg.DownAfterNoAlive
	h.PingInterval = hostCfg.PingInterval
	h.Master = NewDBHost(hostCfg.Master, hostCfg.User, hostCfg.Password, 0, h.MaxConnNum)

	if len(hostCfg.Slaves) > 0 {
		h.Slaves = make([]*DBHost, len(hostCfg.Slaves))

		for i, slave := range hostCfg.Slaves {
			slaveConfig := strings.Split(slave, "@")
			var slaveWeight int
			if len(slaveConfig) > 1 {
				slaveWeight, _ = strconv.Atoi(slaveConfig[1])
			}
			h.Slaves[i] = NewDBHost(slaveConfig[0], hostCfg.User, hostCfg.Password, slaveWeight, h.MaxConnNum)
		}
	}

	return h
}

// DBHost db host.
type DBHost struct {
	Addr     string
	User     string
	Password string
	Weight   int
	Pool     *ConnectionPool
}

// NewDBHost new db host.
func NewDBHost(addr string, user, password string, weight int, maxConnNum int) *DBHost {
	h := new(DBHost)
	h.Addr = addr
	h.User = user
	h.Password = password
	h.Weight = weight
	h.Pool = NewConnectionPool(uint32(maxConnNum), h)
	return h
}

// GetConnection to connect a backend conn.
func (h *DBHost) GetConnection(database string) (Connection, error) {
	return h.Pool.GetConnection(database)
}

// ReturnConnection tu give back a backend conn.
func (h *DBHost) ReturnConnection(conn Connection) {
	h.Pool.ReturnConnection(conn)
}
