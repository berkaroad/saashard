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
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/utils"
)

var (
	regSQLStatement = regexp.MustCompile(`([^';]*|('[^']*')*)+;`)
	sysdbs          = []string{"information_schema", "mysql", "performance_schema"}
)

func handleError(err *error) {
	if x := recover(); x != nil {
		*err = x.(error)
	}
}

// CheckColumnInBoolExpr check if exists or not.
// strOrNumValue: if nil, mean no value or has different value.
func CheckColumnInBoolExpr(expr BoolExpr, colName string) (strOrNumValue ValExpr, err error) {
	if expr == nil {
		return nil, errors.ErrWhereOrJoinOnKey
	}
	switch boolExpr := expr.(type) {
	case *AndExpr:
		exp1 := boolExpr.Left
		exp2 := boolExpr.Right
		colValue1, err1 := CheckColumnInBoolExpr(exp1, colName)
		colValue2, err2 := CheckColumnInBoolExpr(exp2, colName)

		// If exists, get column value
		if err1 == nil && colValue1 != nil {
			strOrNumValue = colValue1
		}

		// If exists, get column value; if not equal, set nil.
		if err2 == nil && colValue2 != nil {
			if strOrNumValue == nil {
				strOrNumValue = colValue2
			} else if String(strOrNumValue) != String(colValue2) {
				strOrNumValue = nil
			}
		}
		if err1 != nil && err2 != nil {
			strOrNumValue = nil
			err = errors.ErrWhereOrJoinOnKey
		}
		return
	case *OrExpr:
		exp1 := boolExpr.Left
		exp2 := boolExpr.Right
		colValue1, err1 := CheckColumnInBoolExpr(exp1, colName)
		colValue2, err2 := CheckColumnInBoolExpr(exp2, colName)

		// If exists, get column value
		if err1 == nil && colValue1 != nil {
			strOrNumValue = colValue1
		}

		// If exists, get column value; if not equal, set nil.
		if err2 == nil && colValue2 != nil {
			if strOrNumValue == nil {
				strOrNumValue = colValue2
			} else if String(strOrNumValue) != String(colValue2) {
				strOrNumValue = nil
			}
		}
		if err1 != nil || err2 != nil {
			strOrNumValue = nil
			err = errors.ErrWhereOrJoinOnKey
		}
		return
	case *ParenBoolExpr:
		strOrNumValue, err = CheckColumnInBoolExpr(boolExpr.Expr, colName)
		return
	case *ComparisonExpr:
		switch boolExpr.Operator {
		case AST_EQ:
			if GetColName(boolExpr.Left) == colName {
				strOrNumValue = boolExpr.Right
			} else if GetColName(boolExpr.Right) == colName {
				strOrNumValue = boolExpr.Left
			}
			if strOrNumValue != nil {
				switch strOrNumValue.(type) {
				case StrVal:
				case NumVal:
				case *ColName:
					strOrNumValue = nil
				default:
					strOrNumValue = nil
					err = errors.ErrWhereOrJoinOnKey
				}
			} else {
				err = errors.ErrWhereOrJoinOnKey
			}
			return
		default:
			return
		}
	default:
		err = errors.ErrWhereOrJoinOnKey
		return
	}
}

// CheckTableExprs remove db and check table's name.
func CheckTableExprs(tabExprs TableExprs, tableNames interface{}) (err error) {
	for _, tabExpr := range tabExprs {
		switch realTabExpr := tabExpr.(type) {
		case *AliasedTableExpr:
			switch simpExpr := realTabExpr.Expr.(type) {
			case *TableName:
				isSystemDB := false
				if len(simpExpr.Qualifier) > 0 {
					db := strings.ToLower(string(simpExpr.Qualifier))
					if isSystemDB = IsSystemDB(db); !isSystemDB {
						simpExpr.Qualifier = nil
					}
				}
				if !isSystemDB {
					tableName := strings.Trim(strings.ToLower(string(simpExpr.Name)), "`")
					if isValid := utils.Contains(tableNames, tableName); !isValid {
						err = &errors.SqlError{Code: 1146, Message: fmt.Sprintf("Table '%-.192s.%-.192s' doesn't exist", "", tableName), State: "42S02"}
						break
					}
				}
			case *Subquery:
				if err = CheckTableExprsInSelect(simpExpr.Select, tableNames); err != nil {
					break
				}
			}
		case *ParenTableExpr:
			err = CheckTableExprs(TableExprs{realTabExpr.Expr}, tableNames)
		case *JoinTableExpr:
			err = CheckTableExprs(TableExprs{realTabExpr.LeftExpr}, tableNames)
			if err == nil {
				CheckTableExprs(TableExprs{realTabExpr.RightExpr}, tableNames)
			}
		}
	}
	return err
}

// IsOnlySystemDBInTableExprs only system db in table exprs.
func IsOnlySystemDBInTableExprs(tabExprs TableExprs) bool {
	var onlySystemDB = true
	for _, tabExpr := range tabExprs {
		if aliasedTabExpr, ok := tabExpr.(*AliasedTableExpr); ok {
			if tabName, ok := aliasedTabExpr.Expr.(*TableName); ok {
				db := strings.ToLower(string(tabName.Qualifier))
				if !IsSystemDB(db) {
					onlySystemDB = false
					break
				}
			} else {
				onlySystemDB = false
				break
			}
		} else {
			onlySystemDB = false
			break
		}
	}
	return onlySystemDB
}

// CheckTableExprsInSelect remove db and check table's name.
func CheckTableExprsInSelect(stmt SelectStatement, tableNames interface{}) (err error) {
	switch selStmt := stmt.(type) {
	case *Select:
		err = CheckTableExprs(selStmt.From, tableNames)
	case *Union:
		err = CheckTableExprsInSelect(selStmt.Left, tableNames)
		if err == nil {
			err = CheckTableExprsInSelect(selStmt.Right, tableNames)
		}
	}
	return err
}

// CheckColumnInTableExpr check shard key should exists in table expression, and has same shard key's value in it.
func CheckColumnInTableExpr(tabExpr TableExpr, colName string) (strOrNumValue ValExpr, err error) {
	switch realTabExpr := tabExpr.(type) {
	case *AliasedTableExpr:
		if subQuery, ok := realTabExpr.Expr.(*Subquery); ok {
			valInSubQuery, errInSubQuery := CheckColumnInSelect(subQuery.Select, colName)
			strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInSubQuery, errInSubQuery)
			if err != nil {
				return
			}
		}
	case *ParenTableExpr:
		valInParen, errInParen := CheckColumnInTableExpr(realTabExpr.Expr, colName)
		strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInParen, errInParen)
		if err != nil {
			return
		}
	case *JoinTableExpr:
		if realTabExpr.On == nil {
			strOrNumValue = nil
			err = errors.ErrWhereOrJoinOnKey
			return
		}
		valInOn, errInOn := CheckColumnInBoolExpr(realTabExpr.On, colName)
		strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInOn, errInOn)
		if err != nil {
			return
		}

		valInLeft, errInLeft := CheckColumnInTableExpr(realTabExpr.LeftExpr, colName)
		strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInLeft, errInLeft)
		if err != nil {
			return
		}

		valInRight, errInRight := CheckColumnInTableExpr(realTabExpr.RightExpr, colName)
		strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInRight, errInRight)
		if err != nil {
			return
		}
	}
	return
}

// CheckColumnInSelect check shard key should exists in where and join expression, and has same shard key's value in it.
func CheckColumnInSelect(statement SelectStatement, colName string) (strOrNumValue ValExpr, err error) {
	switch selStmt := statement.(type) {
	case *Select:
		if selStmt.Where == nil {
			strOrNumValue = nil
			err = errors.ErrWhereOrJoinOnKey
			return
		}
		strOrNumValue, err = CheckColumnInBoolExpr(selStmt.Where.Expr, colName)
		if err != nil {
			return
		}
		for _, tabExpr := range selStmt.From {
			valInTab, errInTab := CheckColumnInTableExpr(tabExpr, colName)
			strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInTab, errInTab)
			if err != nil {
				return
			}
		}
	case *Union:
		valInLeft, errInLeft := CheckColumnInSelect(selStmt.Left, colName)
		strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInLeft, errInLeft)
		if err != nil {
			return
		}

		valInRight, errInRight := CheckColumnInSelect(selStmt.Right, colName)
		strOrNumValue, err = mergeValExprAndError(strOrNumValue, err, valInRight, errInRight)
		if err != nil {
			return
		}
	}
	return strOrNumValue, err
}

// SetLimitInSelect set limit expression in select statement.
func SetLimitInSelect(statement SelectStatement, maxRowCount int) {
	if statement != nil && maxRowCount > 0 {
		switch v := statement.(type) {
		case *Select:
			if v.Limit == nil {
				v.Limit = &Limit{Offset: NumVal("0"), Rowcount: NumVal(strconv.Itoa(maxRowCount))}
			}
			// case *Union:
			// 	SetLimitInSelect(v.Left, maxRowCount)
			// 	SetLimitInSelect(v.Right, maxRowCount)
		}
	}
}

// CheckColumnInInsertOrReplace check shard key should exists in columns, not in dup expers, has same shard key's value in rows.
func CheckColumnInInsertOrReplace(columns Columns, insertRows InsertRows, onDup OnDup, colName string) (strOrNumValue ValExpr, err error) {
	var colValue ValExpr
	// insert statement must contain shardkey column.
	if columns == nil {
		return nil, errors.ErrInsertColumnsKey
	}
	shardKeyExists := false
	shardKeyPos := -1
	for colIndex, columnExpr := range columns {
		if realColumnExpr, ok := columnExpr.(*NonStarExpr); ok {
			if colNameExpr, ok := realColumnExpr.Expr.(*ColName); ok {
				currentColName := strings.ToLower(string(colNameExpr.Name))
				currentColName = strings.Trim(currentColName, "`")
				if currentColName == colName {
					shardKeyExists = true
					shardKeyPos = colIndex
					break
				}
			}
		}
	}
	// insert statement must contain shardkey column.
	if !shardKeyExists {
		return nil, errors.ErrInsertColumnsKey
	}

	// ON DUPLICATE KEY UPDATE expression, couldn't contain shardkey.
	if onDup != nil {
		for _, dupExpr := range onDup {
			currentColName := strings.ToLower(string(dupExpr.Name.Name))
			currentColName = strings.Trim(currentColName, "`")
			if currentColName == colName {
				return nil, errors.ErrUpdateKey
			}
		}
	}

	// Fetch shard key's value.
	switch values := insertRows.(type) {
	case Values:
		for _, val := range values {
			switch val := val.(type) {
			case ValTuple:
				if len(val) != len(columns) {
					return nil, errors.ErrColsLenNotMatch
				}
				if colValue == nil {
					colValue = val[shardKeyPos]
				} else {
					colValue1 := val[shardKeyPos]
					if String(colValue1) != String(colValue) {
						colValue = nil
						break
					}
				}
			}
		}
	}
	if colValue == nil {
		return nil, errors.ErrInsertValuesKey
	}
	return colValue, nil
}

// GetColName returns the column name, only if
// it's a simple expression. Otherwise, it returns "".
func GetColName(node Expr) string {
	if n, ok := node.(*ColName); ok {
		return strings.Trim(strings.ToLower(string(n.Name)), "`")
	}
	return ""
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

// IsSystemDB is system db or not.
func IsSystemDB(db string) bool {
	return utils.Contains(sysdbs, db)
}

func mergeValExprAndError(val1 ValExpr, err1 error, val2 ValExpr, err2 error) (strOrNumValue ValExpr, err error) {
	strOrNumValue, err = val1, err1
	if err2 != nil {
		strOrNumValue = nil
		err = err2
		return
	} else if val2 != nil {
		if strOrNumValue == nil {
			strOrNumValue = val2
		} else if strOrNumValue != nil && String(val2) != String(strOrNumValue) {
			strOrNumValue = nil
			err = errors.ErrWhereOrJoinOnKey
			return
		}
	}
	return
}
