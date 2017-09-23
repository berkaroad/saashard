// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// SelectStatement any SELECT statement.
type SelectStatement interface {
	IStatement()
	ISelectStatement()
	IInsertRows()
	SQLNode
}

//SimpleSelect no from and where statement.
type SimpleSelect struct {
	Comments    Comments
	Distinct    string
	SelectExprs SelectExprs
	Limit       *Limit
}

func (node *SimpleSelect) Format(buf *TrackedBuffer) {
	buf.Fprintf("select %v%s%v%v", node.Comments, node.Distinct, node.SelectExprs, node.Limit)
}

func (*SimpleSelect) IStatement()       {}
func (*SimpleSelect) ISelectStatement() {}
func (*SimpleSelect) IInsertRows()      {}

// Select represents a SELECT statement.
type Select struct {
	Comments    Comments
	Distinct    string
	SelectExprs SelectExprs
	From        TableExprs
	Where       *Where
	GroupBy     GroupBy
	Having      *Where
	OrderBy     OrderBy
	Limit       *Limit
	Lock        string
}

// Select.Distinct
const (
	AST_DISTINCT = "distinct "
)

// Select.Lock
const (
	AST_FOR_UPDATE = " for update"
	AST_SHARE_MODE = " lock in share mode"
)

// Format Select.
func (node *Select) Format(buf *TrackedBuffer) {
	buf.Fprintf("select %v%s%v from %v%v%v%v%v%v%s",
		node.Comments, node.Distinct, node.SelectExprs,
		node.From, node.Where,
		node.GroupBy, node.Having, node.OrderBy,
		node.Limit, node.Lock)
}

func (node *Select) IStatement()       {}
func (node *Select) ISelectStatement() {}
func (node *Select) IInsertRows()      {}

// Union represents a UNION statement.
type Union struct {
	Type        string
	Left, Right SelectStatement
}

// Union.Type
const (
	AST_UNION     = "union"
	AST_UNION_ALL = "union all"
	AST_SET_MINUS = "minus"
	AST_EXCEPT    = "except"
	AST_INTERSECT = "intersect"
)

func (node *Union) Format(buf *TrackedBuffer) {
	buf.Fprintf("%v %s %v", node.Left, node.Type, node.Right)
}

func (node *Union) IStatement()       {}
func (node *Union) ISelectStatement() {}
func (node *Union) IInsertRows()      {}
