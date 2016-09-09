// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

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

package sqlparser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/berkaroad/saashard/sqlparser/sqltypes"
)

var (
	regSQLStatement = regexp.MustCompile(`([^';]*|('[^']*')*)+;`)
)

func handleError(err *error) {
	if x := recover(); x != nil {
		*err = x.(error)
	}
}

// IsColInEqualConditionExists check if exists or not.
func IsColInEqualConditionExists(expr BoolExpr, colName string) bool {
	if expr == nil {
		return false
	}
	switch boolExpr := expr.(type) {
	case *AndExpr:
		exp1 := boolExpr.Left
		exp2 := boolExpr.Right
		return IsColInEqualConditionExists(exp1, colName) || IsColInEqualConditionExists(exp2, colName)
	case *OrExpr:
		exp1 := boolExpr.Left
		exp2 := boolExpr.Right
		return IsColInEqualConditionExists(exp1, colName) && IsColInEqualConditionExists(exp2, colName)
	case *ParenBoolExpr:
		return IsColInEqualConditionExists(boolExpr.Expr, colName)
	case *ComparisonExpr:
		switch boolExpr.Operator {
		case AST_EQ:
			if GetColName(boolExpr.Left) == colName {
				return true
			}
			if GetColName(boolExpr.Right) == colName {
				return true
			}
			return false
		default:
			return false
		}
	default:
		return false
	}
}

// GetColName returns the column name, only if
// it's a simple expression. Otherwise, it returns "".
func GetColName(node Expr) string {
	if n, ok := node.(*ColName); ok {
		return strings.Trim(strings.ToLower(string(n.Name)), "`")
	}
	return ""
}

// IsColName returns true if the ValExpr is a *ColName.
func IsColName(node ValExpr) bool {
	_, ok := node.(*ColName)
	return ok
}

// IsValue returns true if the ValExpr is a string, number or value arg.
// NULL is not considered to be a value.
func IsValue(node ValExpr) bool {
	switch node.(type) {
	case StrVal, NumVal, ValArg:
		return true
	}
	return false
}

// HasINClause returns true if an yof the conditions has an IN clause.
func HasINClause(conditions []BoolExpr) bool {
	for _, node := range conditions {
		if c, ok := node.(*ComparisonExpr); ok && c.Operator == AST_IN {
			return true
		}
	}
	return false
}

// IsSimpleTuple returns true if the ValExpr is a ValTuple that
// contains simple values.
func IsSimpleTuple(node ValExpr) bool {
	list, ok := node.(ValTuple)
	if !ok {
		// It's a subquery.
		return false
	}
	for _, n := range list {
		if !IsValue(n) {
			return false
		}
	}
	return true
}

// AsInterface converts the ValExpr to an interface. It converts
// ValTuple to []interface{}, ValArg to string, StrVal to sqltypes.String,
// NumVal to sqltypes.Numeric. Otherwise, it returns an error.
func AsInterface(node ValExpr) (interface{}, error) {
	switch node := node.(type) {
	case ValTuple:
		vals := make([]interface{}, 0, len(node))
		for _, val := range node {
			v, err := AsInterface(val)
			if err != nil {
				return nil, err
			}
			vals = append(vals, v)
		}
		return vals, nil
	case ValArg:
		return string(node), nil
	case StrVal:
		return sqltypes.MakeString(node), nil
	case NumVal:
		n, err := sqltypes.BuildNumeric(string(node))
		if err != nil {
			return nil, fmt.Errorf("type mismatch: %s", err)
		}
		return n, nil
	}
	return nil, fmt.Errorf("unexpected node %v", node)
}

// StringIn is a convenience function that returns
// true if str matches any of the values.
func StringIn(str string, values ...string) bool {
	for _, val := range values {
		if str == val {
			return true
		}
	}
	return false
}

// SplitSQLStatement is to split multi-sql to single-sql.
func SplitSQLStatement(multiSQL string) []string {
	sqlStatementList := regSQLStatement.FindAll([]byte(strings.TrimRight(multiSQL, ";")+";"), -1)
	var sqlStrArray = make([]string, len(sqlStatementList))
	for i, sqlStatement := range sqlStatementList {
		query := string(sqlStatement[:len(sqlStatement)-1])
		sqlStrArray[i] = query
	}
	return sqlStrArray
}
