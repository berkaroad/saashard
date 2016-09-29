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
	"runtime"
	"strings"

	"github.com/berkaroad/saashard/backend"
	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/route"
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/utils/golog"
)

func (c *ClientConn) handleQuery(sql string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			golog.OutputSql("Error", "err:%v,sql:%s", e, sql)

			if err, ok := e.(error); ok {
				const size = 4096
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]

				golog.Error("ClientConn", "handleQuery",
					err.Error(), 0,
					"stack", string(buf), "sql", sql)
			}
			return
		}
	}()

	var plan route.Plan
	var sqls []string
	if c.capability&mysql.CLIENT_MULTI_STATEMENTS > 0 {
		sqls = sqlparser.SplitSQLStatement(sql)
	} else {
		sqls = []string{strings.TrimSpace(sql)}
	}

	var stmts = make([]sqlparser.Statement, 0, len(sqls))
	for _, sql := range sqls {
		stmt, err := sqlparser.Parse(sql)
		if err != nil {
			golog.Error("proxy", "handleQuery", err.Error(), 0, "sql", sql)
			return mysql.NewError(mysql.ER_SYNTAX_ERROR, fmt.Sprintf("Syntax error or not supported for '%s': '%s'", sql, err.Error()))
		}
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}

	if len(stmts) > 0 {
		router := route.NewRouter(c.db, c.schemas, c.proxy.cfg.GetNodes(), c.connectionID, c.user, c.isInTransaction())
		plan, err = router.BuildMergedPlan(stmts...)
		if err != nil {
			return
		}
		return plan.Execute(c.executePlanWithQueryCommand, c.c.RemoteAddr(), strings.ToLower(c.proxy.logSQL[c.proxy.logSQLIndex]) != golog.LogSqlOff, c.proxy.slowLogTime[c.proxy.slowLogTimeIndex], c.proxy.counter)
	}
	return c.pkg.WriteOK(c.capability, c.status, nil)
}

func (c *ClientConn) executePlanWithQueryCommand(statements []sqlparser.Statement, results []*mysql.Result,
	dataNodes []string, isSlave bool,
	queryDataNodes map[sqlparser.Statement][]string) (backendConnAddrs []string, err error) {

	backendConnAddrs = []string{}

	if len(dataNodes) == 1 {
		resultCount := len(statements)
		node := c.proxy.nodes[dataNodes[0]]
		// If in transaction, must exec in the same node.
		if c.isInTransaction() && node != c.nodeInTrans {
			err = errors.ErrTransInMulti
			return
		}

		var conn backend.Connection
		// Get backend conn from slave or master.
		if isSlave && len(node.DataHost.Slaves) > 0 {
			if conn, err = c.getOrCreateSlaveConn(node); err != nil {
				return
			}
		} else {
			if conn, err = c.getOrCreateMasterConn(node); err != nil {
				return
			}
		}

		backendConnAddrs = []string{conn.GetAddr()}
		var mysqlConn = conn.(*mysqlBackend.Conn)
		mysqlConn.UseDB(node.Database)
		var moreResult = true

		var result *mysql.Result
		for i := 0; i < resultCount; i++ {
			if i == resultCount-1 {
				moreResult = false
			}

			if results[i] != nil {
				if results[i].Resultset == nil {
					err = c.pkg.WriteOK(c.capability, c.status, results[i])
				} else {
					err = c.pkg.WriteResultSet(c.capability, c.status, results[i])
				}
				if err != nil {
					return
				}
			} else if statements[i] != nil {
				statement := statements[i]

				switch v := statement.(type) {
				case *sqlparser.UseDB:
					c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
					err = c.handleInitDB(v.DB)
					return
				case *sqlparser.Begin:
					sql := sqlparser.String(statement)
					if result, err = mysqlConn.Query(sql); err != nil {
						return
					}
					c.status |= mysql.SERVER_STATUS_IN_TRANS
					c.nodeInTrans = node
					c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
					err = c.pkg.WriteOK(c.capability, c.status, result)
					return
				case *sqlparser.Commit:
					sql := sqlparser.String(statement)
					if result, err = mysqlConn.Query(sql); err != nil {
						return
					}
					c.status &= ^mysql.SERVER_STATUS_IN_TRANS
					c.nodeInTrans = nil
					c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
					err = c.pkg.WriteOK(c.capability, c.status, result)
					return
				case *sqlparser.Rollback:
					sql := sqlparser.String(statement)
					if result, err = mysqlConn.Query(sql); err != nil {
						return
					}
					c.status &= ^mysql.SERVER_STATUS_IN_TRANS
					c.nodeInTrans = nil
					c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
					err = c.pkg.WriteOK(c.capability, c.status, result)
					return
				case *sqlparser.KillQuery:
					connID := v.GetConnectionID()
					if c.proxy.cfg.AllowKillQuery {
						c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
						if connID == c.connectionID {
							c.Close()
						} else if specConn, ok := c.proxy.conns[connID]; ok {
							err = mysql.NewDefaultError(mysql.ER_QUERY_INTERRUPTED)
							specConn.Close()
						} else {
							err = mysql.NewDefaultError(mysql.ER_NO_SUCH_THREAD, connID)
						}
					} else {
						err = mysql.NewDefaultError(mysql.ER_KILL_DENIED_ERROR, connID)
					}
					return
				case *sqlparser.KillConnection:
					connID := v.GetConnectionID()
					if c.proxy.cfg.AllowKillQuery {
						c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
						if connID == c.connectionID {
							c.Close()
						} else if specConn, ok := c.proxy.conns[connID]; ok {
							err = mysql.NewDefaultError(mysql.ER_QUERY_INTERRUPTED)
							specConn.Close()
						} else {
							err = mysql.NewDefaultError(mysql.ER_NO_SUCH_THREAD, connID)
						}
					} else {
						err = mysql.NewDefaultError(mysql.ER_KILL_DENIED_ERROR, connID)
					}
					return
				case *sqlparser.SetVariable:
					sql := sqlparser.String(statement)
					if result, err = mysqlConn.Query(sql); err != nil {
						return
					}
					for _, varNameVal := range v.Exprs {
						if string(varNameVal.Name.Name) == "autocommit" {
							autoCommit := sqlparser.String(varNameVal.Expr)
							if autoCommit == "0" {
								c.status &= ^mysql.SERVER_STATUS_AUTOCOMMIT
								c.nodeInTrans = node
							} else {
								c.status |= mysql.SERVER_STATUS_AUTOCOMMIT
								c.status &= ^mysql.SERVER_STATUS_IN_TRANS
								c.nodeInTrans = nil
							}
							break
						}
					}
					if moreResult {
						c.status |= mysql.SERVER_MORE_RESULTS_EXISTS
					} else {
						c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
					}
					err = c.pkg.WriteOK(c.capability, c.status, result)
				default:
					sql := sqlparser.String(statement)
					if result, err = mysqlConn.Query(sql); err != nil {
						return
					}
					if moreResult {
						c.status |= mysql.SERVER_MORE_RESULTS_EXISTS
					} else {
						c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
					}
					if result.Resultset == nil {
						err = c.pkg.WriteOK(c.capability, c.status, result)
					} else {
						err = c.pkg.WriteResultSet(c.capability, c.status, result)
					}
				}
				if err != nil {
					return
				}
			} else {
				err = errors.ErrNoStatement
				return
			}
		}
	} else {
		var result *mysql.Result
		for _, dataNode := range dataNodes {
			node := c.proxy.nodes[dataNode]
			statement := statements[0]
			// If in transaction, must exec in the same node.
			if c.isInTransaction() && node != c.nodeInTrans {
				return nil, errors.ErrTransInMulti
			}

			var conn backend.Connection
			// Get backend conn from slave or master.
			if isSlave && len(node.DataHost.Slaves) > 0 {
				if conn, err = c.getOrCreateSlaveConn(node); err != nil {
					return
				}
			} else {
				if conn, err = c.getOrCreateMasterConn(node); err != nil {
					return
				}
			}

			backendConnAddrs = append(backendConnAddrs, conn.GetAddr())
			var mysqlConn = conn.(*mysqlBackend.Conn)
			mysqlConn.UseDB(node.Database)

			switch statement.(type) {
			case sqlparser.SelectStatement:
				err = errors.ErrCmdUnsupport
				return
			case sqlparser.DDLStatement:
				sql := sqlparser.String(statement)
				if result, err = mysqlConn.Query(sql); err != nil {
					return
				}
				c.status &= ^mysql.SERVER_MORE_RESULTS_EXISTS
			default:
				err = errors.ErrCmdUnsupport
				return
			}
		}
		if result == nil {
			err = errors.ErrCmdUnsupport
			return
		}
		if result.Resultset == nil {
			err = c.pkg.WriteOK(c.capability, c.status, result)
		} else {
			err = c.pkg.WriteResultSet(c.capability, c.status, result)
		}
	}
	return
}
