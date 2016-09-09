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
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
)

func (r *Router) buildExplainPlan(statement *sqlparser.Explain) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(normalPlan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = statement
	plan.anyNode = true

	result := new(mysql.Result)
	result.Status = mysql.SERVER_STATUS_AUTOCOMMIT
	result.Resultset = new(mysql.Resultset)
	result.Resultset.Fields = make([]*mysql.Field, 4)
	result.Resultset.Fields[0] = &mysql.Field{Schema: []byte(""),
		Table:        []byte(""),
		OrgTable:     []byte(""),
		Name:         []byte("data_node"),
		OrgName:      []byte(""),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 192,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}
	result.Resultset.Fields[1] = &mysql.Field{Schema: []byte(""),
		Table:        []byte(""),
		OrgTable:     []byte(""),
		Name:         []byte("is_slave"),
		OrgName:      []byte(""),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 32,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}
	result.Resultset.Fields[2] = &mysql.Field{Schema: []byte(""),
		Table:        []byte(""),
		OrgTable:     []byte(""),
		Name:         []byte("any_node"),
		OrgName:      []byte(""),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 32,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}
	result.Resultset.Fields[3] = &mysql.Field{Schema: []byte(""),
		Table:        []byte(""),
		OrgTable:     []byte(""),
		Name:         []byte("plan_sql"),
		OrgName:      []byte(""),
		Charset:      uint16(mysql.DEFAULT_COLLATION_ID),
		ColumnLength: 6144,
		ColumnType:   mysql.MYSQL_TYPE_VAR_STRING,
		Flags:        mysql.NOT_NULL_FLAG,
		Decimals:     0}

	result.Rows = make([]*mysql.Row, 1)
	row := mysql.NewTextRow(result.Fields)

	if innerPlan, err := r.BuildNormalPlan(statement.Statement); err != nil {
		row.AppendNullValue()
		row.AppendNullValue()
		row.AppendNullValue()
		row.AppendStringValue(err.Error())
	} else {
		innerNormalPlan := innerPlan.(*normalPlan)
		row.AppendStringValue(innerNormalPlan.DataNode)
		row.AppendBooleanValue(innerNormalPlan.IsSlave)
		row.AppendBooleanValue(innerNormalPlan.anyNode)
		row.AppendStringValue(sqlparser.String(innerNormalPlan.Statement))
		result.Resultset.Rows[0] = row
	}
	result.Resultset.Rows[0] = row
	plan.Result = result

	return plan, nil
}
