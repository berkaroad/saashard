package proxy

import (
	"strings"

	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) ShowVariables(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW GLOBAL VARIABLES") ||
		strings.HasPrefix(strings.ToUpper(sql), "SHOW SESSION VARIABLES") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 2)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("SESSION_VARIABLES"),
			OrgTable:     []byte("SESSION_VARIABLES"),
			Name:         []byte("Variable_name"),
			OrgName:      []byte("VARIABLE_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("SESSION_VARIABLES"),
			OrgTable:     []byte("SESSION_VARIABLES"),
			Name:         []byte("Value"),
			OrgName:      []byte("VARIABLE_VALUE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 6144,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		if strings.HasPrefix(strings.ToUpper(sql), "SHOW GLOBAL VARIABLES") {
			result.Rows = make([]*mysql.Row, 2)

			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("auto_increment_increment")
			row.AppendStringValue("1")
			result.Rows[0] = row

			row = mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("auto_increment_offset")
			row.AppendStringValue("1")
			result.Rows[1] = row
		} else if strings.Contains(strings.ToUpper(sql), "'LOWER_CASE_TABLE_NAMES'") {
			result.Rows = make([]*mysql.Row, 1)
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("lower_case_table_names")
			row.AppendUIntValue(1)
			result.Rows[0] = row
		} else if strings.Contains(strings.ToUpper(sql), "'SQL_MODE'") {
			result.Rows = make([]*mysql.Row, 1)
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("sql_mode")
			row.AppendStringValue("NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION")
			result.Rows[0] = row
		} else if strings.Contains(strings.ToUpper(sql), "'VERSION_COMMENT'") {
			result.Rows = make([]*mysql.Row, 1)
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("version_comment")
			row.AppendStringValue(mysql.SourceInfo)
			result.Rows[0] = row
		} else if strings.Contains(strings.ToUpper(sql), "'VERSION'") {
			result.Rows = make([]*mysql.Row, 1)
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("version")
			row.AppendStringValue(mysql.ServerVersion)
			result.Rows[0] = row
		} else if strings.Contains(strings.ToUpper(sql), "'VERSION_COMPILE_OS'") {
			result.Rows = make([]*mysql.Row, 1)
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("version_compile_os")
			row.AppendStringValue("debian-linux-gnu")
			result.Rows[0] = row
		}

		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
