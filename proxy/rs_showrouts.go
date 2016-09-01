package proxy

import (
	"strings"

	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) ShowRoutines(sql string) (bool, error) {
	if strings.HasPrefix(strings.ToUpper(sql), "SHOW PROCEDURE STATUS") ||
		strings.HasPrefix(strings.ToUpper(sql), "SHOW FUNCTION STATUS") {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, 11)
		result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Db"),
			OrgName:      []byte("ROUTINE_SCHEMA"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Name"),
			OrgName:      []byte("ROUTINE_NAME"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 192,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[2] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Type"),
			OrgName:      []byte("ROUTINE_TYPE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 27,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[3] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Definer"),
			OrgName:      []byte("DEFINER"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 567,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[4] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Modified"),
			OrgName:      []byte("LAST_ALTERED"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 19,
			ColumnType:   mysql.MYSQL_TYPE_DATETIME,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[5] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Created"),
			OrgName:      []byte("CREATED"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 19,
			ColumnType:   mysql.MYSQL_TYPE_DATETIME,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[6] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Security_type"),
			OrgName:      []byte("SECURITY_TYPE"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 21,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[7] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Comment"),
			OrgName:      []byte("ROUTINE_COMMENT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 589815,
			ColumnType:   mysql.MYSQL_TYPE_BLOB,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[8] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("character_set_client"),
			OrgName:      []byte("CHARACTER_SET_CLIENT"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 96,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[9] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Database Collation"),
			OrgName:      []byte("DATABASE_COLLATION"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 96,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		result.Resultset.Fields[10] = &mysql.Field{Schema: []byte("information_schema"),
			Table:        []byte("ROUTINES"),
			OrgTable:     []byte("ROUTINES"),
			Name:         []byte("Database Collation"),
			OrgName:      []byte("DATABASE_COLLATION"),
			Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
			ColumnLength: 96,
			ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
			Flags:        mysql.NOT_NULL_FLAG,
			Decimals:     0}

		return true, c.pkg.WriteResultSet(c.capability, c.status, result)
	}
	return false, nil
}
