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

	sql = strings.TrimSpace(sql)

	var onlySingleSQL = true
	var plan route.Plan
	var sqls []string
	if c.capability&mysql.CLIENT_MULTI_STATEMENTS > 0 {
		sqls = sqlparser.SplitSQLStatement(sql)
		onlySingleSQL = len(sqls) == 1
		sql = sqls[0]
	}

	if onlySingleSQL {
		var stmt sqlparser.Statement
		stmt, err = sqlparser.Parse(sql)
		if err != nil {
			golog.Error("proxy", "handleQuery", err.Error(), 0, "sql", sql)
			err = mysql.NewError(mysql.ER_SYNTAX_ERROR, fmt.Sprintf("Syntax error or not supported for '%s'", sql))
			return
		}
		router := route.NewRouter(c.db, c.schemas, c.proxy.cfg.GetNodes(), c.connectionID, c.user)
		plan, err = router.BuildMergedPlan(stmt)
		if err != nil {
			return
		}
		return plan.Execute(c.executePlan)
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
	var mysqlConn *mysqlBackend.Conn
	var conn backend.Connection
	var dbHost *backend.DBHost
	if isSlave {
		dbHost = node.DataHost.Slaves[0]

	} else {
		dbHost = node.DataHost.Master
	}
	if conn, err = dbHost.Connect(node.Database); err != nil {
		return
	}
	mysqlConn = conn.(*mysqlBackend.Conn)

	defer mysqlConn.Close()

	mysqlConn.SetAutoCommit(1)
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
				err = c.handleInitDB(v.DB)
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
