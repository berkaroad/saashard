package proxy

import (
	"strings"

	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) ShowColumns(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW COLUMNS FROM") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 6)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Field"),
			OrgName:      []byte("COLUMNS"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Type"),
			OrgName:      []byte("COLUMN_TYPE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 589515,
			ColumnType:   mysql.MYSQL_TYPE_BLOB,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[2] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Null"),
			OrgName:      []byte("IS_NULLABLE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 9,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[3] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Key"),
			OrgName:      []byte("COLUMN_KEY"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 9,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[4] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Default"),
			OrgName:      []byte("COLUMN_DEFAULT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 589815,
			ColumnType:   mysql.MYSQL_TYPE_BLOB,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[5] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Extra"),
			OrgName:      []byte("EXTRA"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 81,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		if strings.HasSuffix(strings.ToUpper(sql), "`KGW2`.`BUY_LIST`") {
			result.Rows = make([]*mysql.Row, 3)

			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("buy_id")
			row.AppendStringValue("bigint(20)")
			row.AppendStringValue("NO")
			row.AppendStringValue("PRI")
			row.AppendNullValue()
			row.AppendStringValue("auto_increment")
			result.Rows[0] = row

			row = mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("client_id")
			row.AppendStringValue("int(11)")
			row.AppendStringValue("NO")
			row.AppendStringValue("MUL")
			row.AppendNullValue()
			row.AppendStringValue("")
			result.Rows[1] = row

			row = mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("supplier")
			row.AppendStringValue("varchar(100)")
			row.AppendStringValue("NO")
			row.AppendStringValue("")
			row.AppendStringValue("")
			row.AppendStringValue("")
			result.Rows[2] = row
		}
		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}

func (c *ClientConn) ShowFullColumns(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW FULL COLUMNS FROM") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 9)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Field"),
			OrgName:      []byte("COLUMNS"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Type"),
			OrgName:      []byte("COLUMN_TYPE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 589515,
			ColumnType:   mysql.MYSQL_TYPE_BLOB,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[2] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Collation"),
			OrgName:      []byte("COLLATION_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 96,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[3] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Null"),
			OrgName:      []byte("IS_NULLABLE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 9,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[4] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Key"),
			OrgName:      []byte("COLUMN_KEY"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 9,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[5] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Default"),
			OrgName:      []byte("COLUMN_DEFAULT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 589815,
			ColumnType:   mysql.MYSQL_TYPE_BLOB,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[6] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Extra"),
			OrgName:      []byte("EXTRA"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 81,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[7] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Privileges"),
			OrgName:      []byte("PRIVILEGES"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 240,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[8] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("COLUMNS"),
			OrgTable:     []byte("COLUMNS"),
			Name:         []byte("Comment"),
			OrgName:      []byte("COLUMN_COMMENT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 3072,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		if strings.HasSuffix(strings.ToUpper(sql), "`KGW2`.`BUY_LIST`") {
			result.Rows = make([]*mysql.Row, 3)

			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("buy_id")
			row.AppendStringValue("bigint(20)")
			row.AppendNullValue()
			row.AppendStringValue("NO")
			row.AppendStringValue("PRI")
			row.AppendNullValue()
			row.AppendStringValue("auto_increment")
			row.AppendStringValue("select,insert,update,references")
			row.AppendStringValue("")
			result.Rows[0] = row

			row = mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("client_id")
			row.AppendStringValue("int(11)")
			row.AppendNullValue()
			row.AppendStringValue("NO")
			row.AppendStringValue("MUL")
			row.AppendNullValue()
			row.AppendStringValue("")
			row.AppendStringValue("select,insert,update,references")
			row.AppendStringValue("租户ID")
			result.Rows[1] = row

			row = mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("supplier")
			row.AppendStringValue("varchar(100)")
			row.AppendNullValue()
			row.AppendStringValue("NO")
			row.AppendStringValue("")
			row.AppendStringValue("")
			row.AppendStringValue("")
			row.AppendStringValue("select,insert,update,references")
			row.AppendStringValue("供应商")
			result.Rows[2] = row
		}
		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
