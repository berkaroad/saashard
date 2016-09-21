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
	"github.com/berkaroad/saashard/sqlparser"
)

func (r *Router) buildInsertPlan(statement *sqlparser.Insert) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
		// check table exists or not.
		table := string(statement.Table.Name)
		table = strings.Trim(strings.ToLower(table), "`")
		tables := schemaConfig.GetTables()
		if _, ok := tables[table]; !ok {
			return nil, errors.ErrTableNotExists
		}

		colValue, err := sqlparser.CheckColumnInInsertOrReplace(statement.Columns, statement.Rows, statement.OnDup, schemaConfig.ShardKey)
		if err != nil {
			return nil, err
		}
		println("Insert ShardKey=", sqlparser.String(colValue))
	}

	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildUpdatePlan(statement *sqlparser.Update) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
		// check table exists or not.
		table := string(statement.Table.Name)
		table = strings.Trim(strings.ToLower(table), "`")
		tables := schemaConfig.GetTables()
		if _, ok := tables[table]; !ok {
			return nil, errors.ErrTableNotExists
		}

		// UPDATE expression, couldn't contain shardkey.
		for _, setExpr := range statement.Exprs {
			colName := strings.ToLower(string(setExpr.Name.Name))
			colName = strings.Trim(colName, "`")
			if colName == schemaConfig.ShardKey {
				return nil, errors.ErrUpdateKey
			}
		}

		// WHERE expression, should contain shardkey.
		if statement.Where == nil || statement.Where.Expr == nil {
			return nil, errors.ErrWhereOrJoinOnKey
		}
		var err error
		var colValue sqlparser.ValExpr
		if colValue, err = sqlparser.CheckColumnInBoolExpr(statement.Where.Expr, schemaConfig.ShardKey); err != nil {
			return nil, err
		} else if colValue == nil {
			return nil, errors.ErrWhereOrJoinOnKey
		}
		println("Update ShardKey=", sqlparser.String(colValue))
	}
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.Statement = statement

	return plan, nil
}

func (r *Router) buildDeletePlan(statement *sqlparser.Delete) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
		// check table exists or not.
		table := string(statement.Table.Name)
		table = strings.Trim(strings.ToLower(table), "`")
		tables := schemaConfig.GetTables()
		if _, ok := tables[table]; !ok {
			return nil, errors.ErrTableNotExists
		}
		// WHERE expression, should contain shardkey.
		if statement.Where == nil || statement.Where.Expr == nil {
			return nil, errors.ErrWhereOrJoinOnKey
		}
		var err error
		var colValue sqlparser.ValExpr
		if colValue, err = sqlparser.CheckColumnInBoolExpr(statement.Where.Expr, schemaConfig.ShardKey); err != nil {
			return nil, err
		} else if colValue == nil {
			return nil, errors.ErrWhereOrJoinOnKey
		}
		println("Delete ShardKey=", sqlparser.String(colValue))
	}
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.Statement = statement

	return plan, nil
}

func (r *Router) buildReplacePlan(statement *sqlparser.Replace) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
		// check table exists or not.
		table := string(statement.Table.Name)
		table = strings.Trim(strings.ToLower(table), "`")
		tables := schemaConfig.GetTables()
		if _, ok := tables[table]; !ok {
			return nil, errors.ErrTableNotExists
		}

		colValue, err := sqlparser.CheckColumnInInsertOrReplace(statement.Columns, statement.Rows, nil, schemaConfig.ShardKey)
		if err != nil {
			return nil, err
		}
		println("Replace ShardKey=", sqlparser.String(colValue))
	}
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.Statement = statement

	return plan, nil
}
