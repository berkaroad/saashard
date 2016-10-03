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

package mysql

import (
	"fmt"
	"net"
	"strings"

	"github.com/berkaroad/saashard/backend"
	"github.com/berkaroad/saashard/net/mysql"
)

func init() {
	backend.CreateConnection = func(dbHost *backend.DBHost) backend.Connection {
		return new(Conn)
	}
}

// Conn proxy <-> mysql server
type Conn struct {
	conn net.Conn

	pkg *mysql.PacketIO

	dbHost       *backend.DBHost
	db           string
	connectionID uint32

	capability uint32
	status     uint16

	collation mysql.CollationID
	charset   string
	salt      []byte
}

// GetConnectionID get connection id
func (c *Conn) GetConnectionID() uint32 {
	return c.connectionID
}

// SetConnectionID set connection id
func (c *Conn) SetConnectionID(id uint32) { c.connectionID = id }

// Connect to mysql
func (c *Conn) Connect(dbHost *backend.DBHost, db string) error {
	c.dbHost = dbHost
	c.db = db

	//use utf8
	c.collation = mysql.DEFAULT_COLLATION_ID
	c.charset = mysql.DEFAULT_CHARSET

	return c.Reconnect()
}

// Reconnect reconnect to mysql.
func (c *Conn) Reconnect() error {
	var needReconnect = true
	if c.conn != nil {
		if err := c.Ping(); err == nil {
			needReconnect = false
		} else {
			c.Close()
		}
	}

	if needReconnect {
		n := "tcp"
		if strings.Contains(c.dbHost.Addr, "/") {
			n = "unix"
		}

		netConn, err := net.Dial(n, c.dbHost.Addr)
		if err != nil {
			return err
		}

		if tcpConn, ok := netConn.(*net.TCPConn); ok {
			//SetNoDelay controls whether the operating system should delay packet transmission
			// in hopes of sending fewer packets (Nagle's algorithm).
			// The default is true (no delay),
			// meaning that data is sent as soon as possible after a Write.
			//I set this option false.
			tcpConn.SetNoDelay(false)
			tcpConn.SetKeepAlive(true)
		}

		c.conn = netConn
		c.pkg = mysql.NewPacketIO(netConn)

		if c.capability, c.status, c.collation, err = c.pkg.ReadInitialHandshake(&(c.salt)); err != nil {
			c.conn.Close()
			c.conn = nil
			return err
		}

		if err := c.pkg.WriteAuthHandshake(&(c.capability), c.dbHost.User, c.dbHost.Password, c.db, c.salt, c.collation); err != nil {
			c.conn.Close()
			c.conn = nil
			return err
		}

		if _, err := c.pkg.ReadOK(c.capability, &(c.status)); err != nil {
			c.conn.Close()
			c.conn = nil
			return err
		}
	}
	c.pkg.Sequence = 0

	//we must always use autocommit
	if !c.IsAutoCommit() {
		if _, err := c.pkg.Query(c.capability, &(c.status), "set autocommit = 1"); err != nil {
			c.conn.Close()

			return err
		}
	}

	return nil
}

// Close conn
func (c *Conn) Close() error {
	if c.conn != nil {
		c.pkg.Quit(c.capability, &(c.status))
		c.conn.Close()
		c.conn = nil
		c.salt = nil
		c.pkg = nil
	}

	return nil
}

// ReturnConnection give back connection.
func (c *Conn) ReturnConnection() {
	if c.dbHost != nil {
		c.dbHost.Pool.ReturnConnection(c)
	}
}

// Ping db
func (c *Conn) Ping() error {
	if err := c.pkg.Ping(c.capability, &(c.status)); err != nil {
		return err
	}

	return nil
}

// UseDB use db.
func (c *Conn) UseDB(dbName string) error {
	if c.db == dbName || len(dbName) == 0 {
		return nil
	}

	if err := c.pkg.InitDB(c.capability, &(c.status), dbName); err != nil {
		return err
	}

	c.db = dbName
	return nil
}

// GetDB get db.
func (c *Conn) GetDB() string {
	return c.db
}

// GetAddr get addr.
func (c *Conn) GetAddr() string {
	return c.dbHost.Addr
}

// Query command.
func (c *Conn) Query(query string) (*mysql.Result, error) {
	return c.pkg.Query(c.capability, &(c.status), query)
}

// FieldList return field list.
func (c *Conn) FieldList(table string, wildcard string) ([]*mysql.Field, error) {
	return c.pkg.FieldList(c.capability, table, wildcard)
}

// Execute command.
func (c *Conn) Execute(command string, args []interface{}) (*mysql.Result, error) {
	if len(args) == 0 {
		return c.pkg.Query(c.capability, &(c.status), command)
	}
	s, err := c.Prepare(command)
	if err != nil {
		return nil, err
	}
	var r *mysql.Result
	r, err = s.Execute(args)
	s.Close()
	return r, err
}

// Prepare stmt.
func (c *Conn) Prepare(query string) (*mysql.Stmt, error) {
	var err error

	stmt := mysql.NewStmt(c.pkg, c.capability, &c.status)
	stmt.Query = query
	if err = c.pkg.StmtPrepare(c.capability, stmt); err != nil {
		return nil, err
	}

	return stmt, nil
}

// Begin Tx
func (c *Conn) Begin() error {
	_, err := c.Query("begin")
	return err
}

// Commit Tx
func (c *Conn) Commit() error {
	_, err := c.Query("commit")
	return err
}

// Rollback Tx
func (c *Conn) Rollback() error {
	_, err := c.Query("rollback")
	return err
}

// SetCharset set charset
func (c *Conn) SetCharset(charset string) error {
	charset = strings.Trim(charset, "\"'`")
	if c.charset == charset {
		return nil
	}

	cid, ok := mysql.CharsetIds[charset]
	if !ok {
		return fmt.Errorf("invalid charset %s", charset)
	}

	if _, err := c.pkg.Query(c.capability, &(c.status), fmt.Sprintf("set names %s", charset)); err != nil {
		return err
	}
	c.collation = cid
	c.charset = charset
	return nil

}

// GetCharset get charset.
func (c *Conn) GetCharset() string {
	return c.charset
}

// IsAutoCommit status.
func (c *Conn) IsAutoCommit() bool {
	return c.status&mysql.SERVER_STATUS_AUTOCOMMIT > 0
}

// IsInTransaction status.
func (c *Conn) IsInTransaction() bool {
	return c.status&mysql.SERVER_STATUS_IN_TRANS > 0
}
