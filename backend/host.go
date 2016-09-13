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

	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
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
	h.Master = NewDBHost(hostCfg.Master, hostCfg.User, hostCfg.Password, 0)

	if len(hostCfg.Slaves) > 0 {
		h.Slaves = make([]*DBHost, len(hostCfg.Slaves))

		for i, slave := range hostCfg.Slaves {
			slaveConfig := strings.Split(slave, "@")
			var slaveWeight int
			if len(slaveConfig) > 1 {
				slaveWeight, _ = strconv.Atoi(slaveConfig[1])
			}
			h.Slaves[i] = NewDBHost(slaveConfig[0], hostCfg.User, hostCfg.Password, slaveWeight)
		}
	}

	return h
}

// DBHost db host.
type DBHost struct {
	addr     string
	user     string
	password string
	weight   int
}

// NewDBHost new db host.
func NewDBHost(addr string, user, password string, weight int) *DBHost {
	h := new(DBHost)
	h.addr = addr
	h.user = user
	h.password = password
	h.weight = weight
	return h
}

// Connect a backend conn.
// Need mantain connection pool.
func (h *DBHost) Connect(database string) (conn Connection, err error) {
	conn = new(mysqlBackend.Conn)
	err = conn.Connect(h.addr, h.user, h.password, database)
	return
}
