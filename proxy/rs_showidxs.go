package proxy

import (
	"strings"

	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) ShowIndexes(sql string) (bool, error) {

	if strings.HasPrefix(strings.ToUpper(sql), "SHOW INDEXES FROM") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 13)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Table"),
			OrgName:      []byte("TABLE_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Non_unique"),
			OrgName:      []byte("NON_UNIQUE"),
			Charset:      uint16(0x3f), // COLLATE binary
			ColumnLength: 1,
			ColumnType:   mysql.MYSQL_TYPE_LONGLONG,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[2] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Key_name"),
			OrgName:      []byte("INDEX_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[3] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Seq_in_index"),
			OrgName:      []byte("SEQ_IN_INDEX"),
			Charset:      uint16(0x3f), // COLLATE binary
			ColumnLength: 2,
			ColumnType:   mysql.MYSQL_TYPE_LONGLONG,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[4] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Column_name"),
			OrgName:      []byte("COLUMN_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[5] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Collation"),
			OrgName:      []byte("COLLATION"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 3,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[6] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Cardinality"),
			OrgName:      []byte("CARDINALITY"),
			Charset:      uint16(0x3f),
			ColumnLength: 21,
			ColumnType:   mysql.MYSQL_TYPE_LONGLONG,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[7] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Sub_part"),
			OrgName:      []byte("SUB_PART"),
			Charset:      uint16(0x3f),
			ColumnLength: 3,
			ColumnType:   mysql.MYSQL_TYPE_LONGLONG,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[8] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Packed"),
			OrgName:      []byte("PACKED"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 30,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[9] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Null"),
			OrgName:      []byte("NULLABLE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 9,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[10] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Index_type"),
			OrgName:      []byte("INDEX_TYPE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 48,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[11] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Comment"),
			OrgName:      []byte("COMMENT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 48,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[12] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("STATISTICS"),
			OrgTable:     []byte("STATISTICS"),
			Name:         []byte("Index_comment"),
			OrgName:      []byte("INDEX_COMMENT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 3072,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		if strings.HasSuffix(strings.ToUpper(sql), "`KGW2`.`BUY_LIST`") {
			result.Rows = make([]*mysql.Row, 2)

			row := mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("buy_list")
			row.AppendUIntValue(0)
			row.AppendStringValue("PRIMARY")
			row.AppendUIntValue(1)
			row.AppendStringValue("buy_id")
			row.AppendStringValue("A")
			row.AppendUIntValue(1947)
			row.AppendNullValue()
			row.AppendNullValue()
			row.AppendStringValue("")
			row.AppendStringValue("BTREE")
			row.AppendStringValue("")
			row.AppendStringValue("")
			result.Rows[0] = row

			row = mysql.NewTextRow(result.Resultset.Fields)
			row.AppendStringValue("buy_list")
			row.AppendUIntValue(1)
			row.AppendStringValue("index3")
			row.AppendUIntValue(1)
			row.AppendStringValue("client_id")
			row.AppendStringValue("A")
			row.AppendUIntValue(243)
			row.AppendNullValue()
			row.AppendNullValue()
			row.AppendStringValue("")
			row.AppendStringValue("BTREE")
			row.AppendStringValue("")
			row.AppendStringValue("")
			result.Rows[1] = row
		}
		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
