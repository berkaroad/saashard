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
	"container/list"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/utils/golog"
)

// CreateConnection create connection function, this is default value, must replace it.
var CreateConnection = func(dbHost *DBHost) Connection { return new(nilConnection) }

// Connection is backend connection.
type Connection interface {
	// GetConnectionID get connection id
	GetConnectionID() uint32

	// SetConnectionID set connection id
	SetConnectionID(id uint32)

	// Connect db server.
	Connect(dbHost *DBHost, db string) error

	// Reconnect db server.
	Reconnect() error

	// Close db server.
	Close() error

	// ReturnConnection give back connection.
	ReturnConnection()

	// UseDB to set current db.
	UseDB(database string) error

	// Ping db server.
	Ping() error

	// Begin transaction
	Begin() error

	// Commit transaction
	Commit() error

	// Rollback transaction
	Rollback() error

	// SetAutoCommit
	SetAutoCommit(autocommit bool) error

	// GetAddr Get addr info
	GetAddr() string
}

// nilConnection default connection.
type nilConnection struct{ connectionID uint32 }

// GetConnectionID get connection id
func (c *nilConnection) GetConnectionID() uint32 {
	return c.connectionID
}

// SetConnectionID set connection id
func (c *nilConnection) SetConnectionID(id uint32) { c.connectionID = id }

// Connect db server.
func (c *nilConnection) Connect(dbHost *DBHost, db string) error {
	return nil
}

// Reconnect db server.
func (c *nilConnection) Reconnect() error { return nil }

// Close db server.
func (c *nilConnection) Close() error { return nil }

// ReturnConnection give back connection.
func (c *nilConnection) ReturnConnection() {}

// UseDB to set current db.
func (c *nilConnection) UseDB(database string) error { return nil }

// Ping db server.
func (c *nilConnection) Ping() error { return nil }

// Begin transaction
func (c *nilConnection) Begin() error { return nil }

// Commit transaction
func (c *nilConnection) Commit() error { return nil }

// Rollback transaction
func (c *nilConnection) Rollback() error { return nil }

// SetAutoCommit
func (c *nilConnection) SetAutoCommit(autocommit bool) error { return nil }

// GetAddr Get addr info
func (c *nilConnection) GetAddr() string { return "127.0.0.1:6051" }

// ConnectionPool to manage connection pool.
type ConnectionPool struct {
	locker      *sync.Mutex
	MaxPoolSize uint32
	used        uint32
	dbHost      *DBHost
	connections *list.List
	connids     map[uint32]interface{}
}

// NewConnectionPool create connection pool
func NewConnectionPool(maxPoolSize uint32, dbHost *DBHost) *ConnectionPool {
	p := new(ConnectionPool)
	p.locker = new(sync.Mutex)
	p.MaxPoolSize = maxPoolSize
	if p.MaxPoolSize == 0 {
		p.MaxPoolSize = math.MaxUint32
	}
	p.dbHost = dbHost
	p.connections = list.New()
	p.connids = make(map[uint32]interface{})
	return p
}

// GetIdleCount Get Idle count.
func (p *ConnectionPool) GetIdleCount() uint32 {
	return p.MaxPoolSize - p.used
}

// GetConnection get connection from pool
func (p *ConnectionPool) GetConnection(database string) (Connection, error) {
	defer p.locker.Unlock()
	defer p.logConnIdleInfo()

	p.locker.Lock()
	var conn Connection
	var err error
	var retryCount int
	f := func() (Connection, error) {
		var conn Connection
		var err error
		if p.GetIdleCount() > 0 {
			atomic.AddUint32(&p.used, 1)
			if p.connections.Len() > 0 {
				elem := p.connections.Back()
				p.connections.Remove(elem)
				conn = elem.Value.(Connection)
				delete(p.connids, conn.GetConnectionID())
				err = conn.Reconnect()
				if err != nil {
					atomic.AddUint32(&p.used, ^uint32(0))
					conn = nil
				}
			} else {
				conn = CreateConnection(p.dbHost)
				err = conn.Connect(p.dbHost, database)
				if err != nil {
					atomic.AddUint32(&p.used, ^uint32(0))
					conn = nil
				} else if conn.GetConnectionID() == 0 {
					conn.SetConnectionID(p.used)
				}
			}
		} else {
			time.Sleep(time.Second)
			err = errors.ErrNoIdleConn
		}
		return conn, err
	}
	conn, err = f()
	retryCount++
	for ; err != nil && retryCount <= 3; retryCount++ {
		conn, err = f()
	}

	return conn, err
}

// ReturnConnection give back connection to pool
func (p *ConnectionPool) ReturnConnection(conn Connection) {
	defer p.locker.Unlock()
	defer p.logConnIdleInfo()

	p.locker.Lock()
	if conn != nil && conn.GetConnectionID() > 0 {
		if _, exists := p.connids[conn.GetConnectionID()]; !exists {
			p.connections.PushFront(conn)
			p.connids[conn.GetConnectionID()] = nil
			atomic.AddUint32(&p.used, ^uint32(0))
		}
	}
}

func (p *ConnectionPool) logConnIdleInfo() {
	idleCount := p.GetIdleCount()
	// idleCount is zero, or less or equal then 20%, then warn
	if idleCount*100/p.MaxPoolSize <= 20 {
		golog.Warn("backend", "NewConnectionPool", "Idle count is less or equal then 20 pecent", 0,
			"DBHost",
			p.dbHost.Addr,
			"used",
			p.used,
			"cached",
			p.connections.Len(),
			"idle",
			idleCount)
	} else {
		golog.Info("backend", "NewConnectionPool", "Idle info", 0,
			"DBHost",
			p.dbHost.Addr,
			"used",
			p.used,
			"cached",
			p.connections.Len(),
			"idle",
			idleCount)
	}
}
