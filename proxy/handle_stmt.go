package proxy

import (
	"fmt"
	"strings"

	"github.com/berkaroad/saashard/backend"
	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/utils/golog"
)

func (c *ClientConn) handleStmtPrepare(sql string) error {
	println("handleStmtPrepare", sql)
	var err error
	s := mysql.NewStmt(c.pkg, c.capability, &c.status)
	sql = strings.TrimRight(sql, ";")

	var statement sqlparser.Statement
	statement, err = sqlparser.Parse(sql)
	if err != nil {
		golog.Error("proxy", "handle_stmt", err.Error(), 0, "sql", sql)
		return mysql.NewError(mysql.ER_SYNTAX_ERROR, fmt.Sprintf("Syntax error or not supported for '%s'", sql))
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
	if conn = c.backendMasterConns[node]; conn == nil {
		var dbHost *backend.DBHost
		dbHost = node.DataHost.Master
		if conn, err = dbHost.GetConnection(node.Database); err != nil {
			return err
		}
		c.backendMasterConns[node] = conn
	}

	var mysqlConn = conn.(*mysqlBackend.Conn)
	mysqlConn.UseDB(node.Database)

	var stmtFromBackend *mysql.Stmt
	stmtFromBackend, err = mysqlConn.Prepare(sqlparser.String(statement))
	if err != nil {
		return fmt.Errorf("prepare error 1 %s, sql1=%s, sql2=%s", err, sql, sqlparser.String(statement))
	}
	if stmtFromBackend == nil {
		return errors.New("prepare error no stmt from backend")
	}
	s.Params = stmtFromBackend.Params
	s.Columns = stmtFromBackend.Columns

	s.ID = c.stmtID
	c.stmtID++

	if err = c.pkg.WriteStmtPrepare(c.capability, c.status, s); err != nil {
		return err
	}

	s.ResetParams()
	c.stmts[s.ID] = s

	err = stmtFromBackend.Close()
	if err != nil {
		return fmt.Errorf("prepare error %s", err)
	}
	return err
}

func (c *ClientConn) handleStmtExecute(data []byte) error {
	var err error
	var s *mysql.Stmt
	println("handleStmtExecute", string(data))
	s, err = c.pkg.WriteStmtExecute(data, func(id uint32) *mysql.Stmt { return c.stmts[id] })

	switch stmt := s.Statement.(type) {
	case *sqlparser.Select:
		err = c.handlePrepareSelect(stmt, s.Query, s.Args)
	case *sqlparser.Insert:
		// err = c.handlePrepareExec(s.Statement, s.Query, s.args)
	case *sqlparser.Update:
		// err = c.handlePrepareExec(s.s, s.sql, s.args)
	case *sqlparser.Delete:
		// err = c.handlePrepareExec(s.s, s.sql, s.args)
	case *sqlparser.Replace:
		// err = c.handlePrepareExec(s.s, s.sql, s.args)
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
	// Get backend conn from slave or master.
	if !c.isInTransaction() && len(node.DataHost.Slaves) > 0 {
		if conn = c.backendSlaveConns[node]; conn == nil {
			var dbHost *backend.DBHost
			dbHost, _ = node.DataHost.GetSlave()
			if conn, err = dbHost.GetConnection(node.Database); err != nil {
				return err
			}
			c.backendSlaveConns[node] = conn
		}
	} else {
		if conn = c.backendMasterConns[node]; conn == nil {
			var dbHost *backend.DBHost
			dbHost = node.DataHost.Master
			if conn, err = dbHost.GetConnection(node.Database); err != nil {
				return err
			}
			c.backendMasterConns[node] = conn
		}
	}

	var mysqlConn = conn.(*mysqlBackend.Conn)
	mysqlConn.UseDB(node.Database)

	var rs *mysql.Result
	rs, err = mysqlConn.Execute(sql, args)
	if err != nil {
		golog.Error("ClientConn", "handlePrepareSelect", err.Error(), c.connectionID)
		return err
	}

	status := c.status | rs.Status
	if rs.Resultset == nil {
		rs.Resultset = c.newEmptyResultset(stmt)
	}
	err = c.pkg.WriteResultSet(c.capability, status, rs)
	return err
}

// func (c *ClientConn) handlePrepareExec(stmt sqlparser.Statement, sql string, args []interface{}) error {
// 	defaultRule := c.schema.rule.DefaultRule
// 	if len(defaultRule.Nodes) == 0 {
// 		return errors.ErrNoDefaultNode
// 	}
// 	defaultNode := c.proxy.GetNode(defaultRule.Nodes[0])

// 	//execute in Master DB
// 	conn, err := c.getBackendConn(defaultNode, false)
// 	defer c.closeConn(conn, false)
// 	if err != nil {
// 		return err
// 	}

// 	if conn == nil {
// 		return c.writeOK(nil)
// 	}

// 	var rs []*mysql.Result
// 	rs, err = c.executeInNode(conn, sql, args)
// 	c.closeConn(conn, false)

// 	if err != nil {
// 		golog.Error("ClientConn", "handlePrepareExec", err.Error(), c.connectionId)
// 		return err
// 	}

// 	status := c.status | rs[0].Status
// 	if rs[0].Resultset != nil {
// 		err = c.writeResultset(status, rs[0].Resultset)
// 	} else {
// 		err = c.writeOK(rs[0])
// 	}

// 	return err
// }

func (c *ClientConn) handleStmtClose(data []byte) error {
	println("handleStmtClose", string(data))
	return nil
}

func (c *ClientConn) handleStmtSendLongData(data []byte) error {
	println("handleStmtSendLongData", string(data))
	return nil
}

func (c *ClientConn) handleStmtReset(data []byte) error {
	println("handleStmtReset", string(data))
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
