package server

import (
	"strings"

	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) ShowStatus(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW SESSION STATUS") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 2)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("SESSION_STATUS"),
			OrgTable:     []byte("SESSION_STATUS"),
			Name:         []byte("Variable_name"),
			OrgName:      []byte("VARIABLE_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("SESSION_STATUS"),
			OrgTable:     []byte("SESSION_STATUS"),
			Name:         []byte("Value"),
			OrgName:      []byte("VARIABLE_VALUE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 6144,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		if strings.Contains(strings.ToUpper(sql), "'SSL_CIPHER'") {
			result.RowDatas = make([]*mysql.RowData, 1)
			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("Ssl_cipher"))
			rowData.AppendStringValue([]byte(""))
			result.RowDatas[0] = rowData
		} else if strings.Contains(strings.ToUpper(sql), "'SSL_CIPHER'") {

		}
		return true, c.pkg.WriteResultSet(c.capability, c.status, result)

	}
	return false, nil
}
