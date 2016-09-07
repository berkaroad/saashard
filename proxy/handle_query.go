package proxy

import (
	"fmt"
	"runtime"
	"strings"

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
	var stmt sqlparser.Statement
	stmt, err = sqlparser.Parse(sql)
	if err != nil {
		golog.Error("proxy", "handleQuery", err.Error(), 0, "sql", sql)
		return c.pkg.WriteError(c.capability, mysql.NewError(mysql.ER_SYNTAX_ERROR, fmt.Sprintf("Syntax error or not supported for '%s'", sql)))
	}
	if stmt == nil {
		golog.Debug("proxy", "handle_query",
			sql,
			c.connectionID)
		return c.pkg.WriteOK(c.capability, c.status, nil)
	}

	var plan *route.Plan

	var result *mysql.Result

	switch v := stmt.(type) {
	case *sqlparser.UseDB:
		return c.handleInitDB(v.DB)

	case *sqlparser.SimpleSelect:
		if len(v.SelectExprs) == 1 {
			switch strings.ToLower(sqlparser.String(v.SelectExprs[0])) {
			case "current_user()":
				result = new(mysql.Result)
				result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
				result.Resultset = new(mysql.Resultset)
				result.Resultset.Fields = make([]*mysql.Field, 1)
				result.Resultset.Fields[0] = &mysql.Field{Schema: []byte(""),
					Table:        []byte(""),
					OrgTable:     []byte(""),
					Name:         []byte("current_user()"),
					OrgName:      []byte(""),
					Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
					ColumnLength: 423,
					ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
					Flags:        mysql.NOT_NULL_FLAG,
					Decimals:     31}

				result.Rows = make([]*mysql.Row, 1)
				row := mysql.NewTextRow(result.Resultset.Fields)
				row.AppendStringValue(c.user)
				result.Rows[0] = row
				return c.pkg.WriteResultSet(c.capability, c.status, result)

			case "version()":
				result = new(mysql.Result)
				result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
				result.Resultset = new(mysql.Resultset)
				result.Resultset.Fields = make([]*mysql.Field, 1)
				result.Resultset.Fields[0] = &mysql.Field{Schema: []byte(""),
					Table:        []byte(""),
					OrgTable:     []byte(""),
					Name:         []byte("version()"),
					OrgName:      []byte(""),
					Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
					ColumnLength: 72,
					ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
					Flags:        mysql.NOT_NULL_FLAG,
					Decimals:     31}

				result.Rows = make([]*mysql.Row, 1)
				row := mysql.NewTextRow(result.Resultset.Fields)
				row.AppendStringValue(mysql.ServerVersion)
				result.Rows[0] = row
				return c.pkg.WriteResultSet(c.capability, c.status, result)

			case "connection_id()":
				result = new(mysql.Result)
				result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
				result.Resultset = new(mysql.Resultset)
				result.Resultset.Fields = make([]*mysql.Field, 1)
				result.Resultset.Fields[0] = &mysql.Field{Schema: []byte(""),
					Table:        []byte(""),
					OrgTable:     []byte(""),
					Name:         []byte("CONNECTION_ID()"),
					OrgName:      []byte(""),
					Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
					ColumnLength: 10,
					ColumnType:   mysql.MYSQL_TYPE_LONGLONG,
					Flags:        mysql.NOT_NULL_FLAG | mysql.BINARY_FLAG,
					Decimals:     0}

				result.Rows = make([]*mysql.Row, 1)
				row := mysql.NewTextRow(result.Resultset.Fields)
				row.AppendUIntValue(uint64(c.connectionID))
				result.Rows[0] = row
				return c.pkg.WriteResultSet(c.capability, c.status, result)

			case "database()":
				result = new(mysql.Result)
				result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
				result.Resultset = new(mysql.Resultset)
				result.Resultset.Fields = make([]*mysql.Field, 1)
				result.Resultset.Fields[0] = &mysql.Field{Schema: []byte(""),
					Table:        []byte(""),
					OrgTable:     []byte(""),
					Name:         []byte("DATABASE()"),
					OrgName:      []byte(""),
					Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
					ColumnLength: 102,
					ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
					Flags:        mysql.NOT_NULL_FLAG,
					Decimals:     31}

				result.Rows = make([]*mysql.Row, 1)
				row := mysql.NewTextRow(result.Resultset.Fields)
				row.AppendStringValue(c.db)
				result.Rows[0] = row
				return c.pkg.WriteResultSet(c.capability, c.status, result)

			default:
				return c.pkg.WriteError(c.capability, mysql.NewError(mysql.ER_TABLEACCESS_DENIED_ERROR, fmt.Sprintf("SELECT command denied to user '%s'@'%v' for '%s'", c.user, c.clientIP, sql)))
			}
		} else {
			return c.pkg.WriteError(c.capability, mysql.NewError(mysql.ER_TABLEACCESS_DENIED_ERROR, fmt.Sprintf("SELECT command denied to user '%s'@'%v' for '%s'", c.user, c.clientIP, sql)))
		}

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

	case sqlparser.SetStatement:
		return c.pkg.WriteOK(c.capability, c.status, result)

	default:
		router := route.NewRouter(c.db, c.schemas, c.proxy.cfg.GetNodes(), c.connectionID)
		plan, result, err = router.BuildPlan(v)
		if err != nil {
			return c.pkg.WriteError(c.capability, err)
		}
		return c.executePlain(plan, result)
		// if strings.HasPrefix(strings.ToUpper(sql), "SHOW") ||
		// 	strings.HasPrefix(strings.ToUpper(sql), "SELECT") ||
		// 	strings.HasPrefix(strings.ToUpper(sql), "EXPLAIN") {

		// 	golog.Error("proxy", "handle_query", fmt.Sprintf("Command denied to user '%s'@'%v' for '%s'", c.user, c.clientIP, sql), c.connectionID)
		// 	return c.pkg.WriteError(c.capability, mysql.NewError(mysql.ER_TABLEACCESS_DENIED_ERROR, fmt.Sprintf("Command denied for '%s'", sql)))
		// }
		// return c.pkg.WriteOK(c.capability, c.status, nil)
	}
}

func (c *ClientConn) executePlain(plan *route.Plan, result *mysql.Result) (err error) {
	if result != nil {
		return c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	sql := sqlparser.String(plan.Statement)
	node := c.proxy.nodes[plan.DataNode]

	var mysqlConn *mysqlBackend.Conn
	if conn, err := node.DataHost.Slaves[0].Connect(node.Database); err != nil {
		return err
	} else {
		mysqlConn = conn.(*mysqlBackend.Conn)
	}

	defer mysqlConn.Close()

	mysqlConn.SetAutoCommit(1)
	mysqlConn.UseDB(node.Database)
	if result, err = mysqlConn.Query(sql); err != nil {
		return c.pkg.WriteError(c.capability, err)
	}

	switch plan.Statement.(type) {
	case sqlparser.SelectStatement:
		return c.pkg.WriteResultSet(c.capability, c.status, result)

	case sqlparser.ShowStatement:
		return c.pkg.WriteResultSet(c.capability, c.status, result)

	case *sqlparser.Explain:
		return c.pkg.WriteResultSet(c.capability, c.status, result)

	default:
		return c.pkg.WriteOK(c.capability, c.status, result)
	}
}
