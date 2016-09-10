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
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/route"
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/utils/golog"

	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
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

	var stmts = make([]sqlparser.Statement, len(sqls))
	for i, sql := range sqls {
		stmt, err := sqlparser.Parse(sql)
		if err != nil {
			golog.Error("proxy", "handleQuery", err.Error(), 0, "sql", sql)
			return mysql.NewError(mysql.ER_SYNTAX_ERROR, fmt.Sprintf("Syntax error or not supported for '%s'", sql))
		}
		stmts[i] = stmt
	}

	router := route.NewRouter(c.db, c.schemas, c.proxy.cfg.GetNodes(), c.connectionID, c.user)
	plan, err = router.BuildMergedPlan(stmts...)
	if err != nil {
		return
	}
	return plan.Execute(c.executePlan)
}

func (c *ClientConn) executePlan(statements []sqlparser.Statement, results []*mysql.Result, dataNode string, isSlave bool) (err error) {
	resultCount := len(statements)
	var direct = false
	var total = make([]byte, 0, 1024)

	node := c.proxy.nodes[dataNode]
	var conn backend.Connection
	if conn = c.backendConns[node]; conn == nil {
		var dbHost *backend.DBHost
		if isSlave && len(node.DataHost.Slaves) > 0 {
			dbHost = node.DataHost.Slaves[0]
		} else {
			dbHost = node.DataHost.Master
		}
		if conn, err = dbHost.Connect(node.Database); err != nil {
			return
		}
		c.backendConns[node] = conn
	}

	var mysqlConn = conn.(*mysqlBackend.Conn)
	mysqlConn.UseDB(node.Database)

	var result *mysql.Result
	for i := 0; i < resultCount; i++ {
		if i == resultCount-1 {
			direct = true
		}

		if results[i] != nil {
			total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, results[i], direct)
			if err != nil {
				return
			}
		} else if statements[i] != nil {
			statement := statements[i]
			switch v := statement.(type) {
			case *sqlparser.UseDB:
				return c.handleInitDB(v.DB)
			case *sqlparser.Begin:
				c.status |= mysql.SERVER_STATUS_IN_TRANS
				sql := sqlparser.String(statement)
				if result, err = mysqlConn.Query(sql); err != nil {
					return
				}
				return c.pkg.WriteOK(c.capability, c.status, result)
			case *sqlparser.Commit:
				c.status &= ^mysql.SERVER_STATUS_IN_TRANS
				sql := sqlparser.String(statement)
				if result, err = mysqlConn.Query(sql); err != nil {
					return
				}
				return c.pkg.WriteOK(c.capability, c.status, result)
			case *sqlparser.Rollback:
				c.status &= ^mysql.SERVER_STATUS_IN_TRANS
				sql := sqlparser.String(statement)
				if result, err = mysqlConn.Query(sql); err != nil {
					return
				}
				return c.pkg.WriteOK(c.capability, c.status, result)
			default:
				sql := sqlparser.String(statement)
				if result, err = mysqlConn.Query(sql); err != nil {
					return
				}
				if result.Resultset == nil {
					total, err = c.pkg.WriteOKBatch(total, c.capability, c.status, result, direct)
				} else {
					total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, result, direct)
				}
			}
			if err != nil {
				return
			}
		}
	}
	return nil
}
