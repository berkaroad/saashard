package backend

import (
	"strconv"
	"strings"

	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
	"github.com/berkaroad/saashard/config"
)

// DataHost is data host.
type DataHost struct {
	Name   string
	Master *DBHost
	Slaves []*DBHost
}

// NewDataHost new host.
func NewDataHost(hostCfg config.HostConfig) *DataHost {
	h := new(DataHost)
	h.Name = hostCfg.Name

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
func (h *DBHost) Connect(database string) (Connection, error) {
	conn := new(mysqlBackend.Conn)
	err := conn.Connect(h.addr, h.user, h.password, database)
	return conn, err
}
