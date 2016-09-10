// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	"reflect"
	"regexp"
	"strings"
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
func IsColInEqualConditionExists(expr BoolExpr, colName string) (exists bool, strOrNumValue ValExpr) {
	if expr == nil {
		return
	}
	switch boolExpr := expr.(type) {
	case *AndExpr:
		exp1 := boolExpr.Left
		exp2 := boolExpr.Right
		exists1, colValue1 := IsColInEqualConditionExists(exp1, colName)
		exists2, colValue2 := IsColInEqualConditionExists(exp2, colName)

		// If exists, get column value
		if exists1 {
			if colValue1 == nil {
				strOrNumValue = nil
				exists1 = false
			} else if _, ok := colValue1.(*ColName); !ok {
				strOrNumValue = colValue1
			}
		}

		// If exists, get column value; if not equal, set nil.
		if exists2 {
			if colValue2 == nil {
				strOrNumValue = nil
				exists2 = false
			} else if _, ok := colValue2.(*ColName); !ok {
				if strOrNumValue == nil {
					strOrNumValue = colValue2
				} else if String(strOrNumValue) != String(colValue2) {
					strOrNumValue = nil
				}
			}
		}
		exists = exists1 || exists2
		return
	case *OrExpr:
		exp1 := boolExpr.Left
		exp2 := boolExpr.Right
		exists1, colValue1 := IsColInEqualConditionExists(exp1, colName)
		exists2, colValue2 := IsColInEqualConditionExists(exp2, colName)

		// If exists, get column value
		if exists1 {
			if colValue1 == nil {
				strOrNumValue = nil
				exists1 = false
			} else if _, ok := colValue1.(*ColName); !ok {
				strOrNumValue = colValue1
			}
		}

		// If exists, get column value; if not equal, set nil.
		if exists2 {
			if colValue2 == nil {
				strOrNumValue = nil
				exists2 = false
			} else if _, ok := colValue2.(*ColName); !ok {
				if strOrNumValue == nil {
					strOrNumValue = colValue2
				} else if String(strOrNumValue) != String(colValue2) {
					strOrNumValue = nil
				}
			}
		}
		exists = exists1 && exists2
		return
	case *ParenBoolExpr:
		exists, strOrNumValue = IsColInEqualConditionExists(boolExpr.Expr, colName)
		return
	case *ComparisonExpr:
		switch boolExpr.Operator {
		case AST_EQ:
			if GetColName(boolExpr.Left) == colName {
				exists = true
				strOrNumValue = boolExpr.Right
			} else if GetColName(boolExpr.Right) == colName {
				exists = true
				strOrNumValue = boolExpr.Left
			}
			if strOrNumValue != nil {
				switch strOrNumValue.(type) {
				case StrVal:
				case NumVal:
				case *ColName:
				default:
					strOrNumValue = nil
				}
			}
			return
		default:
			return
		}
	default:
		exists = false
		return
	}
}

// CheckTableInSelect remove db and check table's name.
func CheckTableInSelect(statement SelectStatement, tableNames interface{}) bool {
	var isValid = true
	switch selStmt := statement.(type) {
	case *Select:
		for _, tabExpr := range selStmt.From {
			switch realTabExpr := tabExpr.(type) {
			case *AliasedTableExpr:
				switch simpExpr := realTabExpr.Expr.(type) {
				case *TableName:
					simpExpr.Qualifier = nil
					tableName := strings.Trim(strings.ToLower(string(simpExpr.Name)), "`")
					if isValid = Contains(tableNames, tableName); !isValid {
						break
					}
				case *Subquery:
					if isValid = CheckTableInSelect(simpExpr.Select, tableNames); !isValid {
						break
					}
				}
			}
		}
	case *Union:
		isValid = CheckTableInSelect(selStmt.Left, tableNames) && CheckTableInSelect(selStmt.Right, tableNames)
	}
	return isValid
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

// Contains is a convenience function that returns
// true if str matches any of the values.
func Contains(arr interface{}, item interface{}) bool {
	arrValue := reflect.ValueOf(arr)
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrValue.Len(); i++ {
			if arrValue.Index(i).Interface() == item {
				return true
			}
		}
	case reflect.Map:
		if arrValue.MapIndex(reflect.ValueOf(item)).IsValid() {
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
