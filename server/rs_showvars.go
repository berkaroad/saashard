package server

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
			result.RowDatas = make([]*mysql.RowData, 2)

			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("auto_increment_increment"))
			rowData.AppendStringValue([]byte("1"))
			result.RowDatas[0] = rowData

			rowData = mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("auto_increment_offset"))
			rowData.AppendStringValue([]byte("1"))
			result.RowDatas[1] = rowData
		} else if strings.Contains(strings.ToUpper(sql), "'LOWER_CASE_TABLE_NAMES'") {
			result.RowDatas = make([]*mysql.RowData, 1)
			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("lower_case_table_names"))
			rowData.AppendUIntValue(1)
			result.RowDatas[0] = rowData
		} else if strings.Contains(strings.ToUpper(sql), "'SQL_MODE'") {
			result.RowDatas = make([]*mysql.RowData, 1)
			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("sql_mode"))
			rowData.AppendStringValue([]byte("NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"))
			result.RowDatas[0] = rowData
		} else if strings.Contains(strings.ToUpper(sql), "'VERSION_COMMENT'") {
			result.RowDatas = make([]*mysql.RowData, 1)
			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("version_comment"))
			rowData.AppendStringValue([]byte(mysql.SourceInfo))
			result.RowDatas[0] = rowData
		} else if strings.Contains(strings.ToUpper(sql), "'VERSION'") {
			result.RowDatas = make([]*mysql.RowData, 1)
			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("version"))
			rowData.AppendStringValue([]byte(mysql.ServerVersion))
			result.RowDatas[0] = rowData
		} else if strings.Contains(strings.ToUpper(sql), "'VERSION_COMPILE_OS'") {
			result.RowDatas = make([]*mysql.RowData, 1)
			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("version_compile_os"))
			rowData.AppendStringValue([]byte("debian-linux-gnu"))
			result.RowDatas[0] = rowData
		}

		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
