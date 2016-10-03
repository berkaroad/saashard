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

package proxy

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/berkaroad/saashard/backend"
	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
)

func (c *ClientConn) handleStmtPrepare(sql string) error {
	var err error
	s := mysql.NewStmt(c.pkg, c.capability, &c.status)
	sql = strings.TrimRight(sql, ";")

	var statement sqlparser.Statement
	statement, err = sqlparser.Parse(sql)
	if err != nil {
		return mysql.NewError(mysql.ER_SYNTAX_ERROR, fmt.Sprintf("Syntax error or not supported for '%s': '%s'", sql, err.Error()))
	}

	s.Query = sql
	s.Statement = statement

	node := c.proxy.nodes[c.schemas[c.db].Nodes[0]]
	// If in transaction, must exec in the same node.
	if c.isInTransaction() && node != c.nodeInTrans {
		return errors.ErrTransInMulti
	}

	var conn backend.Connection
	// Get backend conn from master.
	conn, err = c.getOrCreateMasterConn(node)
	if err != nil {
		return err
	}

	var mysqlConn = conn.(*mysqlBackend.Conn)
	mysqlConn.UseDB(node.Database)

	var stmtFromBackend *mysql.Stmt
	stmtFromBackend, err = mysqlConn.Prepare(sqlparser.String(statement))
	if err != nil {
		return err
	}
	if stmtFromBackend == nil {
		return errors.New("prepare error no stmt from backend")
	}
	s.ParamNum = stmtFromBackend.ParamNum
	s.Params = stmtFromBackend.Params

	s.ColumnNum = stmtFromBackend.ColumnNum
	s.Columns = stmtFromBackend.Columns
	for i := 0; i < len(s.Columns); i++ {
		s.Columns[i].Schema = []byte(c.db)
	}
	atomic.AddUint32(&c.stmtID, 1)
	s.ID = c.stmtID

	if err = c.pkg.WriteStmtPrepareResponse(c.capability, c.status, s); err != nil {
		return err
	}

	s.ResetParams()
	c.stmts[s.ID] = s

	err = stmtFromBackend.Close()
	if err != nil {
		return err
	}
	return err
}

func (c *ClientConn) handleStmtExecute(data []byte) error {
	var err error
	var s *mysql.Stmt
	s, err = c.pkg.ReadStmtExecuteRequest(data, func(id uint32) *mysql.Stmt { return c.stmts[id] })

	switch stmt := s.Statement.(type) {
	case *sqlparser.Select:
		err = c.handlePrepareSelect(stmt, s.Query, s.Args)
	case *sqlparser.Insert:
		err = c.handlePrepareExec(s.Statement, s.Query, s.Args)
	case *sqlparser.Update:
		err = c.handlePrepareExec(s.Statement, s.Query, s.Args)
	case *sqlparser.Delete:
		err = c.handlePrepareExec(s.Statement, s.Query, s.Args)
	case *sqlparser.Replace:
		err = c.handlePrepareExec(s.Statement, s.Query, s.Args)
	case *sqlparser.Commit:
		err = c.handlePrepareExec(s.Statement, s.Query, s.Args)
	default:
		err = fmt.Errorf("command %T not supported now", stmt)
	}
	s.ResetParams()
	return err
}

func (c *ClientConn) handlePrepareSelect(stmt *sqlparser.Select, sql string, args []interface{}) error {
	var err error
	node := c.proxy.nodes[c.schemas[c.db].Nodes[0]]
	// If in transaction, must exec in the same node.
	if c.isInTransaction() && node != c.nodeInTrans {
		return errors.ErrTransInMulti
	}

	var conn backend.Connection
	// Get backend conn from master.
	conn, err = c.getOrCreateMasterConn(node)
	// // Get backend conn from slave or master.
	// if !c.isInTransaction() && len(node.DataHost.Slaves) > 0 {
	// 	conn, err = c.getOrCreateSlaveConn(node)
	// } else {
	// 	conn, err = c.getOrCreateMasterConn(node)
	// }
	if err != nil {
		return err
	}

	var mysqlConn = conn.(*mysqlBackend.Conn)
	mysqlConn.UseDB(node.Database)

	var rs *mysql.Result
	rs, err = mysqlConn.Execute(sql, args)
	if err != nil {
		return err
	}

	status := c.status | rs.Status
	if rs.Resultset == nil {
		rs.Resultset = c.newEmptyResultset(stmt)
	}
	err = c.pkg.WriteResultSet(c.capability, status, rs)
	return err
}

func (c *ClientConn) handlePrepareExec(stmt sqlparser.Statement, sql string, args []interface{}) error {
	var err error
	node := c.proxy.nodes[c.schemas[c.db].Nodes[0]]
	// If in transaction, must exec in the same node.
	if c.isInTransaction() && node != c.nodeInTrans {
		return errors.ErrTransInMulti
	}

	var conn backend.Connection
	// Get backend conn from master.
	conn, err = c.getOrCreateMasterConn(node)
	// // Get backend conn from slave or master.
	// if !c.isInTransaction() && len(node.DataHost.Slaves) > 0 {
	// 	conn, err = c.getOrCreateSlaveConn(node)
	// } else {
	// 	conn, err = c.getOrCreateMasterConn(node)
	// }
	if err != nil {
		return err
	}

	var mysqlConn = conn.(*mysqlBackend.Conn)
	mysqlConn.UseDB(node.Database)

	var rs *mysql.Result
	rs, err = mysqlConn.Execute(sql, args)
	if err != nil {
		return err
	}

	status := c.status | rs.Status
	if rs.Resultset != nil {
		err = c.pkg.WriteResultSet(c.capability, status, rs)
	} else {
		err = c.pkg.WriteOK(c.capability, status, rs)
	}

	return err
}

func (c *ClientConn) handleStmtClose(data []byte) error {
	mysql.PrintPacketData("handleStmtClose", data)
	return nil
}

func (c *ClientConn) handleStmtSendLongData(data []byte) error {
	mysql.PrintPacketData("handleStmtSendLongData", data)
	return nil
}

func (c *ClientConn) handleStmtReset(data []byte) error {
	mysql.PrintPacketData("handleStmtReset", data)
	return nil
}

func (c *ClientConn) newEmptyResultset(stmt *sqlparser.Select) *mysql.Resultset {
	r := new(mysql.Resultset)
	r.Fields = make([]*mysql.Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &mysql.Field{}
		switch e := expr.(type) {
		case *sqlparser.StarExpr:
			r.Fields[i].Name = []byte("*")
		case *sqlparser.NonStarExpr:
			if e.As != nil {
				r.Fields[i].Name = e.As
				r.Fields[i].OrgName = []byte(sqlparser.String(e.Expr))
			} else {
				r.Fields[i].Name = []byte(sqlparser.String(e.Expr))
			}
		default:
			r.Fields[i].Name = []byte(sqlparser.String(e))
		}
	}

	r.Values = make([][]interface{}, 0)
	r.Rows = make([]*mysql.Row, 0)

	return r
}

// func (c *ClientConn) executePlanWithStmtPrepareCommand(statements []sqlparser.Statement, results []*mysql.Result,
// 	dataNodes []string, isSlave bool,
// 	queryDataNodes map[sqlparser.Statement][]string) (stmt *mysql.Stmt, err error) {
// 	if len(dataNodes) == 1 {
// 		node := c.proxy.nodes[dataNodes[0]]
// 		// If in transaction, must exec in the same node.
// 		if c.isInTransaction() && node != c.nodeInTrans {
// 			return nil, errors.ErrTransInMulti
// 		}

// 		var conn backend.Connection
// 		// Get backend conn from slave or master.
// 		if !c.isInTransaction() && isSlave && len(node.DataHost.Slaves) > 0 {
// 			if conn = c.backendSlaveConns[node]; conn == nil {
// 				var dbHost *backend.DBHost
// 				dbHost, _ = node.DataHost.GetSlave()
// 				if conn, err = dbHost.GetConnection(node.Database); err != nil {
// 					return
// 				}
// 				c.backendSlaveConns[node] = conn
// 			}
// 		} else {
// 			if conn = c.backendMasterConns[node]; conn == nil {
// 				var dbHost *backend.DBHost
// 				dbHost = node.DataHost.Master
// 				if conn, err = dbHost.GetConnection(node.Database); err != nil {
// 					return
// 				}
// 				c.backendMasterConns[node] = conn
// 			}
// 		}

// 		var mysqlConn = conn.(*mysqlBackend.Conn)
// 		mysqlConn.UseDB(node.Database)

// 		stmt, err = mysqlConn.Prepare(sqlparser.String(statements[0]))
// 		if err == nil {
// 			return nil, err
// 		}
// 		err = stmt.Close()
// 	}
// 	return stmt, nil
// }
