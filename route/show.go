package route

import (
	"strings"

	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
)

func (r *Router) buildShowEnginesPlan(stmt *sqlparser.ShowEngines) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowPluginsPlan(stmt *sqlparser.ShowPlugins) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowProcessListPlan(stmt *sqlparser.ShowProcessList) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowFullProcessListPlan(stmt *sqlparser.ShowFullProcessList) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowVariablesPlan(stmt *sqlparser.ShowVariables) (*Plan, *mysql.Result, error) {
	if stmt.Scope == "session" {
		switch v := stmt.LikeOrWhere.(type) {
		case *sqlparser.LikeExpr:
			if val, ok := v.Expr.(sqlparser.StrVal); ok {
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
				switch strings.ToLower(strings.Trim(sqlparser.String(val), "'")) {
				case "lower_case_table_names":
					result.Rows = make([]*mysql.Row, 1)
					row := mysql.NewTextRow(result.Resultset.Fields)
					row.AppendStringValue("lower_case_table_names")
					row.AppendUIntValue(1)
					result.Rows[0] = row
					return nil, result, nil

				case "version_comment":
					result.Rows = make([]*mysql.Row, 1)
					row := mysql.NewTextRow(result.Resultset.Fields)
					row.AppendStringValue("version_comment")
					row.AppendStringValue(mysql.SourceInfo)
					result.Rows[0] = row
					return nil, result, nil

				case "version":
					result.Rows = make([]*mysql.Row, 1)
					row := mysql.NewTextRow(result.Resultset.Fields)
					row.AppendStringValue("version")
					row.AppendStringValue(mysql.ServerVersion)
					result.Rows[0] = row
					return nil, result, nil

				case "version_compile_os":
					result.Rows = make([]*mysql.Row, 1)
					row := mysql.NewTextRow(result.Resultset.Fields)
					row.AppendStringValue("version_compile_os")
					row.AppendStringValue("debian-linux-gnu")
					result.Rows[0] = row
					return nil, result, nil
				}
			}
		}
	}

	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt
	return plan, nil, nil
}

func (r *Router) buildShowStatusPlan(stmt *sqlparser.ShowStatus) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt
	return plan, nil, nil
}

func (r *Router) buildShowTablesPlan(stmt *sqlparser.ShowTables) (*Plan, *mysql.Result, error) {
	db := ""
	if stmt.From != nil {
		db = string(stmt.From.Name)
	}
	if db == "" {
		db = r.SchemaName
	}
	db = strings.Trim(strings.ToLower(db), "`")

	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}

	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt
		return plan, nil, nil
	}
	result := new(mysql.Result)
	result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
	result.Resultset = new(mysql.Resultset)
	result.Resultset.Fields = make([]*mysql.Field, 1)
	result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TABLE_NAMES"),
		OrgTable:     []byte("TABLE_NAMES"),
		Name:         []byte("Tables_in_" + schemaConfig.Name),
		OrgName:      []byte("TABLE_NAME"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 219,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Rows = make([]*mysql.Row, 0, len(schemaConfig.Tables))
	for _, table := range schemaConfig.Tables {
		row := mysql.NewTextRow(result.Resultset.Fields)
		row.AppendStringValue(table.Name)
		result.Rows = append(result.Rows, row)
	}
	return nil, result, nil
}

func (r *Router) buildShowFullTablesPlan(stmt *sqlparser.ShowFullTables) (*Plan, *mysql.Result, error) {
	db := ""
	if stmt.From != nil {
		db = string(stmt.From.Name)
	}
	if db == "" {
		db = r.SchemaName
	}
	db = strings.Trim(strings.ToLower(db), "`")

	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}

	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt
		return plan, nil, nil
	}

	result := new(mysql.Result)
	result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
	result.Resultset = new(mysql.Resultset)
	result.Resultset.Fields = make([]*mysql.Field, 2)
	result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TABLE_NAMES"),
		OrgTable:     []byte("TABLE_NAMES"),
		Name:         []byte("Tables_in_" + schemaConfig.Name),
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

	result.Rows = make([]*mysql.Row, 0, len(schemaConfig.Tables))
	for _, table := range schemaConfig.Tables {
		row := mysql.NewTextRow(result.Resultset.Fields)
		row.AppendStringValue(table.Name)
		row.AppendStringValue("BASE TABLE")
		result.Rows = append(result.Rows, row)
	}
	return nil, result, nil
}

func (r *Router) buildShowColumnsPlan(stmt *sqlparser.ShowColumns) (*Plan, *mysql.Result, error) {
	plan := new(Plan)

	db := string(stmt.From.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	table := string(stmt.From.Name)
	table = strings.Trim(strings.ToLower(table), "`")
	tables := schemaConfig.GetTables()
	if _, ok := tables[table]; schemaConfig.ShardKey != "" && !ok {
		return nil, nil, errors.ErrTableOrViewNotExists
	}

	if stmt.From != nil {
		stmt.From.Qualifier = nil
	}

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowFullColumnsPlan(stmt *sqlparser.ShowFullColumns) (*Plan, *mysql.Result, error) {
	plan := new(Plan)

	db := string(stmt.From.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	table := string(stmt.From.Name)
	table = strings.Trim(strings.ToLower(table), "`")
	tables := schemaConfig.GetTables()
	if _, ok := tables[table]; schemaConfig.ShardKey != "" && !ok {
		return nil, nil, errors.ErrTableOrViewNotExists
	}

	if stmt.From != nil {
		stmt.From.Qualifier = nil
	}

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowIndexPlan(stmt *sqlparser.ShowIndex) (*Plan, *mysql.Result, error) {
	plan := new(Plan)

	db := string(stmt.From.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	table := string(stmt.From.Name)
	table = strings.Trim(strings.ToLower(table), "`")
	tables := schemaConfig.GetTables()
	if _, ok := tables[table]; schemaConfig.ShardKey != "" && !ok {
		return nil, nil, errors.ErrTableOrViewNotExists
	}

	if stmt.From != nil {
		stmt.From.Qualifier = nil
	}

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowTriggersPlan(stmt *sqlparser.ShowTriggers) (*Plan, *mysql.Result, error) {
	db := string(stmt.From.Name)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.ToLower(db)
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		switch v := stmt.LikeOrWhere.(type) {
		case *sqlparser.WhereExpr:
			if comparisonExpr, ok := v.Expr.(*sqlparser.ComparisonExpr); ok {
				if colName, ok := comparisonExpr.Left.(*sqlparser.ColName); ok && strings.Trim(strings.ToLower(string(colName.Name)), "`") == "table" {
					if colValue, ok := comparisonExpr.Right.(*sqlparser.StrVal); ok {
						println("Triggers from table ", sqlparser.String(colValue))
					}
				} else if colName, ok := comparisonExpr.Right.(*sqlparser.ColName); ok && strings.Trim(strings.ToLower(string(colName.Name)), "`") == "table" {
					if colValue, ok := comparisonExpr.Left.(*sqlparser.StrVal); ok {
						println("Triggers from table ", sqlparser.String(colValue))
					}
				}
			}

		case *sqlparser.LikeExpr:
			if table, ok := v.Expr.(*sqlparser.StrVal); ok {
				println("Triggers from table ", sqlparser.String(table))
			}
		}

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt

		return plan, nil, nil
	}
	result := new(mysql.Result)
	result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
	result.Resultset = new(mysql.Resultset)
	result.Resultset.Fields = make([]*mysql.Field, 11)
	result.Resultset.Fields[0] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Trigger"),
		OrgName:      []byte("TRIGGER_NAME"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 192,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[1] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Event"),
		OrgName:      []byte("EVENT_MANIPULATION"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 18,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[2] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Table"),
		OrgName:      []byte("EVENT_OBJECT_TABLE"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 192,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[3] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Statement"),
		OrgName:      []byte("ACTION_STATEMENT"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 589815,
		ColumnType:   mysql.MYSQL_TYPE_BLOB,
		Flags:        mysql.NOT_NULL_FLAG | mysql.BLOB_FLAG,
		Decimals:     0}

	result.Resultset.Fields[4] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Timing"),
		OrgName:      []byte("ACTION_TIMING"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 18,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[5] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Created"),
		OrgName:      []byte("CREATED"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 19,
		ColumnType:   mysql.MYSQL_TYPE_DATETIME,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[6] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("sql_mode"),
		OrgName:      []byte("SQL_MODE"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 24576,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[7] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Definer"),
		OrgName:      []byte("DEFINER"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 567,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[8] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("character_set_client"),
		OrgName:      []byte("CHARACTER_SET_CLIENT"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 96,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[9] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("collation_connection"),
		OrgName:      []byte("COLLATION_CONNECTION"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 96,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Resultset.Fields[10] = &mysql.Field{Schema: []byte("information_schema"),
		Table:        []byte("TRIGGERS"),
		OrgTable:     []byte("TRIGGERS"),
		Name:         []byte("Database Collation"),
		OrgName:      []byte("DATABASE_COLLATION"),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 96,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	return nil, result, nil
}

func (r *Router) buildShowProcedureStatusPlan(stmt *sqlparser.ShowProcedureStatus) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		switch v := stmt.LikeOrWhere.(type) {
		case *sqlparser.WhereExpr:
			if comparisonExpr, ok := v.Expr.(*sqlparser.ComparisonExpr); ok {
				if colName, ok := comparisonExpr.Left.(*sqlparser.ColName); ok && strings.Trim(strings.ToLower(string(colName.Name)), "`") == "db" {
					if colValue, ok := comparisonExpr.Right.(*sqlparser.StrVal); ok {
						println("Procedure from db ", sqlparser.String(colValue))
					}
				} else if colName, ok := comparisonExpr.Right.(*sqlparser.ColName); ok && strings.Trim(strings.ToLower(string(colName.Name)), "`") == "db" {
					if colValue, ok := comparisonExpr.Left.(*sqlparser.StrVal); ok {
						println("Procedure from db ", sqlparser.String(colValue))
					}
				}
			}
		}

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt

		return plan, nil, nil
	}
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
		Name:         []byte("collation_connection"),
		OrgName:      []byte("COLLATION_CONNECTION"),
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

	return nil, result, nil
}

func (r *Router) buildShowFunctionStatusPlan(stmt *sqlparser.ShowFunctionStatus) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		switch v := stmt.LikeOrWhere.(type) {
		case *sqlparser.WhereExpr:
			if comparisonExpr, ok := v.Expr.(*sqlparser.ComparisonExpr); ok {
				if colName, ok := comparisonExpr.Left.(*sqlparser.ColName); ok && strings.Trim(strings.ToLower(string(colName.Name)), "`") == "db" {
					if colValue, ok := comparisonExpr.Right.(*sqlparser.StrVal); ok {
						println("Function from db ", sqlparser.String(colValue))
					}
				} else if colName, ok := comparisonExpr.Right.(*sqlparser.ColName); ok && strings.Trim(strings.ToLower(string(colName.Name)), "`") == "db" {
					if colValue, ok := comparisonExpr.Left.(*sqlparser.StrVal); ok {
						println("Function from db ", sqlparser.String(colValue))
					}
				}
			}
		}

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt

		return plan, nil, nil
	}
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
		Name:         []byte("collation_connection"),
		OrgName:      []byte("COLLATION_CONNECTION"),
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

	return nil, result, nil
}

func (r *Router) buildShowCreateDatabasePlan(stmt *sqlparser.ShowCreateDatabase) (*Plan, *mysql.Result, error) {
	db := string(stmt.Name.Name)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	stmt.Name.Name = []byte(r.Nodes[schemaConfig.Nodes[0]].Database)
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowCreateTablePlan(stmt *sqlparser.ShowCreateTable) (*Plan, *mysql.Result, error) {
	db := string(stmt.Name.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	stmt.Name.Qualifier = []byte(r.Nodes[schemaConfig.Nodes[0]].Database)
	plan := new(Plan)

	table := string(stmt.Name.Name)
	table = strings.Trim(strings.ToLower(table), "`")
	tables := schemaConfig.GetTables()
	if _, ok := tables[table]; schemaConfig.ShardKey != "" && !ok {
		return nil, nil, errors.ErrTableOrViewNotExists
	}

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowCreateViewPlan(stmt *sqlparser.ShowCreateView) (*Plan, *mysql.Result, error) {
	db := string(stmt.Name.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	stmt.Name.Qualifier = []byte(r.Nodes[schemaConfig.Nodes[0]].Database)
	plan := new(Plan)

	table := string(stmt.Name.Name)
	table = strings.Trim(strings.ToLower(table), "`")
	tables := schemaConfig.GetTables()
	if _, ok := tables[table]; schemaConfig.ShardKey != "" && !ok {
		return nil, nil, errors.ErrTableOrViewNotExists
	}

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}

func (r *Router) buildShowCreateTriggerPlan(stmt *sqlparser.ShowCreateTrigger) (*Plan, *mysql.Result, error) {
	db := string(stmt.Name.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	stmt.Name.Qualifier = []byte(r.Nodes[schemaConfig.Nodes[0]].Database)
	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		trigger := string(stmt.Name.Name)
		trigger = strings.Trim(strings.ToLower(trigger), "`")

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt

		return plan, nil, nil
	}
	return nil, nil, errors.ErrNoPlan
}

func (r *Router) buildShowCreateProcedurePlan(stmt *sqlparser.ShowCreateProcedure) (*Plan, *mysql.Result, error) {
	db := string(stmt.Name.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	stmt.Name.Qualifier = []byte(r.Nodes[schemaConfig.Nodes[0]].Database)
	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		procedure := string(stmt.Name.Name)
		procedure = strings.Trim(strings.ToLower(procedure), "`")

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt

		return plan, nil, nil
	}
	return nil, nil, errors.ErrNoPlan
}

func (r *Router) buildShowCreateFunctionPlan(stmt *sqlparser.ShowCreateFunction) (*Plan, *mysql.Result, error) {
	db := string(stmt.Name.Qualifier)
	if db == "" {
		db = r.SchemaName
	} else {
		db = strings.Trim(strings.ToLower(db), "`")
	}
	var schemaConfig *config.SchemaConfig
	if schemaConfig = r.Schemas[db]; schemaConfig == nil {
		return nil, nil, errors.ErrDatabaseNotExists
	}
	stmt.Name.Qualifier = []byte(r.Nodes[schemaConfig.Nodes[0]].Database)
	if schemaConfig.ShardKey == "" {
		plan := new(Plan)

		function := string(stmt.Name.Name)
		function = strings.Trim(strings.ToLower(function), "`")

		plan.DataNode = schemaConfig.Nodes[0]
		plan.IsSlave = true
		plan.Statement = stmt

		return plan, nil, nil
	}
	return nil, nil, errors.ErrNoPlan
}
