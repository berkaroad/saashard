package proxy

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

		schema := c.schemas[c.db]
		result.Rows = make([]*mysql.Row, 0, len(schema.Tables)+len(schema.Views))
		for _, table := range schema.Tables {
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue(table.Name)
			result.Rows = append(result.Rows, row)
		}

		for _, view := range schema.Views {
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue(view.Name)
			result.Rows = append(result.Rows, row)
		}

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

		schema := c.schemas[c.db]
		result.Rows = make([]*mysql.Row, 0, len(schema.Tables)+len(schema.Views))
		for _, table := range schema.Tables {
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue(table.Name)
			row.AppendStringValue("BASE TABLE")
			result.Rows = append(result.Rows, row)
		}

		for _, view := range schema.Views {
			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue(view.Name)
			row.AppendStringValue("VIEW")
			result.Rows = append(result.Rows, row)
		}

		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
