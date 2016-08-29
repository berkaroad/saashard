package server

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
			result.RowDatas = make([]*mysql.RowData, 3)

			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("buy_id"))
			rowData.AppendStringValue([]byte("bigint(20)"))
			rowData.AppendStringValue([]byte("NO"))
			rowData.AppendStringValue([]byte("PRI"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte("auto_increment"))
			result.RowDatas[0] = rowData

			rowData = mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("client_id"))
			rowData.AppendStringValue([]byte("int(11)"))
			rowData.AppendStringValue([]byte("NO"))
			rowData.AppendStringValue([]byte("MUL"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte(""))
			result.RowDatas[1] = rowData

			rowData = mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("supplier"))
			rowData.AppendStringValue([]byte("varchar(100)"))
			rowData.AppendStringValue([]byte("NO"))
			rowData.AppendStringValue([]byte(""))
			rowData.AppendStringValue([]byte(""))
			rowData.AppendStringValue([]byte(""))
			result.RowDatas[2] = rowData
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
			result.RowDatas = make([]*mysql.RowData, 3)

			rowData := mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("buy_id"))
			rowData.AppendStringValue([]byte("bigint(20)"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte("NO"))
			rowData.AppendStringValue([]byte("PRI"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte("auto_increment"))
			rowData.AppendStringValue([]byte("select,insert,update,references"))
			rowData.AppendStringValue([]byte(""))
			result.RowDatas[0] = rowData

			rowData = mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("client_id"))
			rowData.AppendStringValue([]byte("int(11)"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte("NO"))
			rowData.AppendStringValue([]byte("MUL"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte(""))
			rowData.AppendStringValue([]byte("select,insert,update,references"))
			rowData.AppendStringValue([]byte("租户ID"))
			result.RowDatas[1] = rowData

			rowData = mysql.NewRowData(false, result.Resultset.Fields)
			rowData.AppendStringValue([]byte("supplier"))
			rowData.AppendStringValue([]byte("varchar(100)"))
			rowData.AppendNullValue()
			rowData.AppendStringValue([]byte("NO"))
			rowData.AppendStringValue([]byte(""))
			rowData.AppendStringValue([]byte(""))
			rowData.AppendStringValue([]byte(""))
			rowData.AppendStringValue([]byte("select,insert,update,references"))
			rowData.AppendStringValue([]byte("供应商"))
			result.RowDatas[2] = rowData
		}
		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
