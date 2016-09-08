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
	var onlySingleSQL = true
	var plan *route.MergedPlan
	var result *mysql.Result

	sql = strings.TrimSpace(sql)

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

		switch v := stmt.(type) {
		case *sqlparser.UseDB:
			return c.handleInitDB(v.DB)

		case *sqlparser.ShowDatabases:
			result = new(mysql.Result)
			result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
			result.Resultset = new(mysql.Resultset)
			result.Resultset.Fields = make([]*mysql.Field, 1)
			result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
				Table:        []byte("SCHEMATA"),
				OrgTable:     []byte("SCHEMATA"),
				Name:         []byte("Database"),
				OrgName:      []byte("SCHEMA_NAME"),
				Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
				ColumnLength: 192,
				ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
				Flags:        mysql.NOT_NULL_FLAG,
				Decimals:     0}

			result.Rows = make([]*mysql.Row, 0, len(c.schemas))
			for name := range c.schemas {
				row := mysql.NewTextRow(result.Resultset.Fields)
				row.AppendStringValue(name)
				result.Rows = append(result.Rows, row)
			}
			return c.pkg.WriteResultSet(c.capability, c.status, result)

		default:
			router := route.NewRouter(c.db, c.schemas, c.proxy.cfg.GetNodes(), c.connectionID, c.user)
			plan, err = router.BuildMergedPlan(v)
			if err != nil {
				return
			}
			return c.executePlan(plan)
		}
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
	return c.executePlan(plan)
}

func (c *ClientConn) executePlan(plan *route.MergedPlan) (err error) {
	resultCount := len(plan.Statements)
	var direct = false
	var total = make([]byte, 0, 1024)

	node := c.proxy.nodes[plan.DataNode]
	var mysqlConn *mysqlBackend.Conn
	var conn backend.Connection
	var dbHost *backend.DBHost
	if plan.IsSlave {
		dbHost = node.DataHost.Slaves[0]

	} else {
		dbHost = node.DataHost.Master
	}
	if conn, err = dbHost.Connect(node.Database); err != nil {
		return
	}
	mysqlConn = conn.(*mysqlBackend.Conn)

	defer mysqlConn.Close()

	//mysqlConn.SetAutoCommit(1)
	mysqlConn.UseDB(node.Database)

	var result *mysql.Result
	for i := 0; i < resultCount; i++ {
		if i == resultCount-1 {
			direct = true
		}

		if plan.Results[i] != nil {
			total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, plan.Results[i], direct)
			if err != nil {
				return
			}
		} else if plan.Statements[i] != nil {
			statement := plan.Statements[i]
			sql := sqlparser.String(statement)
			if result, err = mysqlConn.Query(sql); err != nil {
				return
			}
			if result.Resultset == nil {
				total, err = c.pkg.WriteOKBatch(total, c.capability, c.status, result, direct)
			} else {
				total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, result, direct)
			}
			if err != nil {
				return
			}
			// switch statement.(type) {
			// case sqlparser.SelectStatement:
			// 	total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, result, direct)
			// 	if err != nil {
			// 		return
			// 	}
			// case sqlparser.ShowStatement:
			// 	total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, result, direct)
			// 	if err != nil {
			// 		return
			// 	}

			// case *sqlparser.Explain:
			// 	total, err = c.pkg.WriteResultSetBatch(total, c.capability, c.status, result, direct)
			// 	if err != nil {
			// 		return
			// 	}

			// default:
			// 	total, err = c.pkg.WriteOKBatch(total, c.capability, c.status, result, direct)
			// 	if err != nil {
			// 		return
			// 	}
			// }
		}
	}
	return nil
}
