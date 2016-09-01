package proxy

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
			result.Rows = make([]*mysql.Row, 1)
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("Ssl_cipher")
			row.AppendStringValue("")
			result.Rows[0] = row
		} else if strings.Contains(strings.ToUpper(sql), "'SSL_CIPHER'") {

		}
		return true, c.pkg.WriteResultSet(c.capability, c.status, result)

	}
	return false, nil
}
