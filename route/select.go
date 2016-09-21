// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package route

import (
	"strings"

	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
)

var currentUserField = &mysql.Field{Schema: []byte(""),
	Table:        []byte(""),
	OrgTable:     []byte(""),
	Name:         []byte("current_user()"),
	OrgName:      []byte(""),
	Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
	ColumnLength: 423,
	ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
	Flags:        mysql.NOT_NULL_FLAG,
	Decimals:     31}

var versionField = &mysql.Field{Schema: []byte(""),
	Table:        []byte(""),
	OrgTable:     []byte(""),
	Name:         []byte("version()"),
	OrgName:      []byte(""),
	Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
	ColumnLength: 72,
	ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
	Flags:        mysql.NOT_NULL_FLAG,
	Decimals:     31}

var connectionIDField = &mysql.Field{Schema: []byte(""),
	Table:        []byte(""),
	OrgTable:     []byte(""),
	Name:         []byte("CONNECTION_ID()"),
	OrgName:      []byte(""),
	Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
	ColumnLength: 10,
	ColumnType:   mysql.MYSQL_TYPE_LONGLONG,
	Flags:        mysql.NOT_NULL_FLAG | mysql.BINARY_FLAG,
	Decimals:     0}

var databaseField = &mysql.Field{Schema: []byte(""),
	Table:        []byte(""),
	OrgTable:     []byte(""),
	Name:         []byte("DATABASE()"),
	OrgName:      []byte(""),
	Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
	ColumnLength: 102,
	ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
	Flags:        mysql.NOT_NULL_FLAG,
	Decimals:     31}

func (r *Router) buildSimpleSelectPlan(statement *sqlparser.SimpleSelect) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	supportedFieldNames := map[string]*mysql.Field{
		"current_user()":  currentUserField,
		"version()":       versionField,
		"connection_id()": connectionIDField,
		"database()":      databaseField}

	supportedFieldValues := map[string]func(*mysql.Row){
		"current_user()":  func(row *mysql.Row) { row.AppendStringValue(r.User) },
		"version()":       func(row *mysql.Row) { row.AppendStringValue(mysql.ServerVersion) },
		"connection_id()": func(row *mysql.Row) { row.AppendUIntValue(uint64(r.ConnectionID)) },
		"database()":      func(row *mysql.Row) { row.AppendStringValue(r.SchemaName) }}

	allFieldsSupported := true
	for _, fieldExpr := range statement.SelectExprs {
		fieldName := strings.ToLower(sqlparser.String(fieldExpr))
		if _, ok := supportedFieldNames[fieldName]; !ok {
			allFieldsSupported = false
			break
		}
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)

	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.onSlave = true && !hint.OnMaster && !r.InTrans
	plan.Statement = statement
	plan.anyNode = true

	if allFieldsSupported {
		result := new(mysql.Result)
		result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
		result.Resultset = new(mysql.Resultset)
		result.Resultset.Fields = make([]*mysql.Field, len(statement.SelectExprs))
		for i, fieldExpr := range statement.SelectExprs {
			fieldName := strings.ToLower(sqlparser.String(fieldExpr))
			result.Resultset.Fields[i] = supportedFieldNames[fieldName]
		}
		result.Rows = make([]*mysql.Row, 1)
		row := mysql.NewTextRow(result.Resultset.Fields)
		for _, fieldExpr := range statement.SelectExprs {
			fieldName := strings.ToLower(sqlparser.String(fieldExpr))
			supportedFieldValues[fieldName](row)
		}
		result.Rows[0] = row
		plan.Result = result
	}
	return plan, nil
}

func (r *Router) buildSelectPlan(statement *sqlparser.Select) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	isOnlySystemDB := false
	nodeIndex := 0
	if schemaConfig.ShardEnabled() {
		if isOnlySystemDB = sqlparser.IsOnlySystemDBInTableExprs(statement.From); !isOnlySystemDB {
			var err error
			if err = sqlparser.CheckTableExprsInSelect(statement, schemaConfig.GetTables()); err != nil {
				return nil, err
			}
			var colValue sqlparser.ValExpr
			if colValue, err = sqlparser.CheckColumnInSelect(statement, schemaConfig.ShardKey); err != nil {
				return nil, err
			} else if colValue == nil {
				return nil, errors.ErrWhereOrJoinOnKey
			}

			algo := ParseShardAlgorithm(schemaConfig.ShardAlgo)
			nodeIndex, err = algo(sqlparser.String(colValue), len(schemaConfig.Nodes))
			if err != nil {
				return nil, err
			}
		}
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)

	plan.nodeNames = []string{schemaConfig.Nodes[nodeIndex]}
	plan.onSlave = true && !hint.OnMaster && !r.InTrans
	if isOnlySystemDB {
		plan.anyNode = true
	}
	plan.Statement = statement

	return plan, nil
}

func (r *Router) buildUnionPlan(statement *sqlparser.Union) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	nodeIndex := 0
	if schemaConfig.ShardEnabled() {
		var err error
		if err = sqlparser.CheckTableExprsInSelect(statement, schemaConfig.GetTables()); err != nil {
			return nil, err
		}
		var colValue sqlparser.ValExpr
		if colValue, err = sqlparser.CheckColumnInSelect(statement, schemaConfig.ShardKey); err != nil {
			return nil, err
		} else if colValue == nil {
			return nil, errors.ErrWhereOrJoinOnKey
		}

		algo := ParseShardAlgorithm(schemaConfig.ShardAlgo)
		nodeIndex, err = algo(sqlparser.String(colValue), len(schemaConfig.Nodes))
		if err != nil {
			return nil, err
		}
	}
	var hint *Hint
	switch left := statement.Left.(type) {
	case *sqlparser.SimpleSelect:
		hint = ReadHint(&left.Comments)
	case *sqlparser.Select:
		hint = ReadHint(&left.Comments)
	}

	plan := new(normalPlan)
	plan.nodeNames = []string{schemaConfig.Nodes[nodeIndex]}
	plan.onSlave = true && !r.InTrans
	if hint != nil {
		plan.onSlave = plan.onSlave && !hint.OnMaster
	}
	plan.Statement = statement

	return plan, nil
}
