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
	"time"

	"github.com/berkaroad/saashard/net/mysql"
)

var (
	pingPeriod = int64(time.Second * 16)
)

// Conn proxy <-> mysql server
type Conn struct {
	conn net.Conn

	pkg *mysql.PacketIO

	addr     string
	user     string
	password string
	db       string

	capability uint32

	status uint16

	collation mysql.CollationID
	charset   string
	salt      []byte

	pushTimestamp int64
	pkgErr        error
}

// NewConn new conn.
func NewConn(addr string, user string, password string, db string) *Conn {
	c := new(Conn)
	c.addr = addr
	c.user = user
	c.password = password
	c.db = db

	//use utf8
	c.collation = mysql.DEFAULT_COLLATION_ID
	c.charset = mysql.DEFAULT_CHARSET

	return c
}

// Connect to mysql
func (c *Conn) Connect(addr string, user string, password string, db string) error {
	c.addr = addr
	c.user = user
	c.password = password
	c.db = db

	//use utf8
	c.collation = mysql.DEFAULT_COLLATION_ID
	c.charset = mysql.DEFAULT_CHARSET

	return c.Reconnect()
}

// Reconnect reconnect to mysql.
func (c *Conn) Reconnect() error {
	var err error
	if c.conn != nil {
		c.conn.Close()
	}

	n := "tcp"
	if strings.Contains(c.addr, "/") {
		n = "unix"
	}

	netConn, err := net.Dial(n, c.addr)
	if err != nil {
		return err
	}

	tcpConn := netConn.(*net.TCPConn)

	//SetNoDelay controls whether the operating system should delay packet transmission
	// in hopes of sending fewer packets (Nagle's algorithm).
	// The default is true (no delay),
	// meaning that data is sent as soon as possible after a Write.
	//I set this option false.
	tcpConn.SetNoDelay(false)
	tcpConn.SetKeepAlive(true)
	c.conn = tcpConn
	c.pkg = mysql.NewPacketIO(tcpConn)

	if c.capability, c.status, c.collation, err = c.pkg.ReadInitialHandshake(&(c.salt)); err != nil {
		c.conn.Close()
		return err
	}

	if err := c.pkg.WriteAuthHandshake(&(c.capability), c.user, c.password, c.db, c.salt, c.collation); err != nil {
		c.conn.Close()

		return err
	}

	if _, err := c.pkg.ReadOK(c.capability, &(c.status)); err != nil {
		c.conn.Close()

		return err
	}

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
		c.pkgErr = nil
	}

	return nil
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
	return c.addr
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
func (c *Conn) Execute(command string, args ...interface{}) (*mysql.Result, error) {
	if len(args) == 0 {
		return c.pkg.Query(c.capability, &(c.status), command)
	}
	if s, err := c.Prepare(command); err != nil {
		return nil, err
	} else {
		var r *mysql.Result
		r, err = s.Execute(args...)
		s.Close()
		return r, err
	}

}

// Prepare stmt.
func (c *Conn) Prepare(query string) (*Stmt, error) {
	var err error

	stmt := new(Stmt)
	stmt.conn = c
	if stmt.id, stmt.columns, stmt.params, err = c.pkg.StmtPrepare(c.capability, query); err != nil {
		return nil, err
	}

	return stmt, nil
}

func (c *Conn) Begin() error {
	_, err := c.pkg.Query(c.capability, &(c.status), "begin")
	return err
}

func (c *Conn) Commit() error {
	_, err := c.pkg.Query(c.capability, &(c.status), "commit")
	return err
}

func (c *Conn) Rollback() error {
	_, err := c.pkg.Query(c.capability, &(c.status), "rollback")
	return err
}

func (c *Conn) SetAutoCommit(n uint8) error {
	if n == 0 {
		if _, err := c.pkg.Query(c.capability, &(c.status), "set autocommit = 0"); err != nil {
			c.conn.Close()

			return err
		}
	} else {
		if _, err := c.pkg.Query(c.capability, &(c.status), "set autocommit = 1"); err != nil {
			c.conn.Close()

			return err
		}
	}
	return nil
}

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

func (c *Conn) IsAutoCommit() bool {
	return c.status&mysql.SERVER_STATUS_AUTOCOMMIT > 0
}

func (c *Conn) IsInTransaction() bool {
	return c.status&mysql.SERVER_STATUS_IN_TRANS > 0
}

func (c *Conn) GetCharset() string {
	return c.charset
}
