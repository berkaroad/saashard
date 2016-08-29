package server

import (
	"strings"

	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) ShowTables(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW TABLES") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 1)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("TABLE_NAMES"),
			OrgTable:     []byte("TABLE_NAMES"),
			Name:         []byte("Tables_name"),
			OrgName:      []byte("TABLE_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 219,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.RowDatas = make([]*mysql.RowData, 1)

		rowData := mysql.NewRowData(false, result.Resultset.Fields)
		rowData.AppendStringValue([]byte("buy_list"))
		result.RowDatas[0] = rowData

		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}

func (c *ClientConn) ShowFullTables(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW FULL TABLES") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 2)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("TABLE_NAMES"),
			OrgTable:     []byte("TABLE_NAMES"),
			Name:         []byte("Table_name"),
			OrgName:      []byte("TABLE_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 219,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}
		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("TABLE_NAMES"),
			OrgTable:     []byte("TABLE_NAMES"),
			Name:         []byte("Table_type"),
			OrgName:      []byte("TABLE_TYPE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.RowDatas = make([]*mysql.RowData, 1)

		rowData := mysql.NewRowData(false, result.Resultset.Fields)
		rowData.AppendStringValue([]byte("buy_list"))
		rowData.AppendStringValue([]byte("BASE TABLE"))
		result.RowDatas[0] = rowData

		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
