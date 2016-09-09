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

// http://dev.mysql.com/doc/refman/5.6/en/show.html

// ShowStatement show statement
type ShowStatement interface {
	IStatement()
	IShowStatement()
	SQLNode
}

// ShowCharset statement
type ShowCharset struct {
	Comments    Comments
	LikeOrWhere Expr
}

func (node *ShowCharset) Format(buf *TrackedBuffer) {
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v character set", node.Comments)
	} else {
		buf.Fprintf("show %v character set%v", node.Comments, node.LikeOrWhere)
	}
}

func (node *ShowCharset) IStatement()     {}
func (node *ShowCharset) IShowStatement() {}

// ShowCollation statement
type ShowCollation struct {
	Comments    Comments
	LikeOrWhere Expr
}

func (node *ShowCollation) Format(buf *TrackedBuffer) {
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v collation", node.Comments)
	} else {
		buf.Fprintf("show %v collation%v", node.Comments, node.LikeOrWhere)
	}
}

func (node *ShowCollation) IStatement()     {}
func (node *ShowCollation) IShowStatement() {}

// ShowVariables statement
type ShowVariables struct {
	Comments    Comments
	Scope       string // GLOBAL or SESSION
	LikeOrWhere Expr
}

func (node *ShowVariables) Format(buf *TrackedBuffer) {
	scope := ""
	if node.Scope != "" {
		scope = " " + node.Scope
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v%s variables", node.Comments, scope)
	} else {
		buf.Fprintf("show %v%s variables%v", node.Comments, scope, node.LikeOrWhere)
	}
}

func (node *ShowVariables) IStatement()     {}
func (node *ShowVariables) IShowStatement() {}

// ShowStatus statement
type ShowStatus struct {
	Comments    Comments
	Scope       string // GLOBAL or SESSION
	LikeOrWhere Expr
}

func (node *ShowStatus) Format(buf *TrackedBuffer) {
	scope := ""
	if node.Scope != "" {
		scope = " " + node.Scope
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v%s status", node.Comments, scope)
	} else {
		buf.Fprintf("show %v%s status%v", node.Comments, scope, node.LikeOrWhere)
	}
}

func (node *ShowStatus) IStatement()     {}
func (node *ShowStatus) IShowStatement() {}

// ShowDatabases statement
type ShowDatabases struct {
	Comments    Comments
	LikeOrWhere Expr
}

func (node *ShowDatabases) Format(buf *TrackedBuffer) {
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v databases", node.Comments)
	} else {
		buf.Fprintf("show %v databases%v", node.Comments, node.LikeOrWhere)
	}
}

func (node *ShowDatabases) IStatement()     {}
func (node *ShowDatabases) IShowStatement() {}

// ShowTables statement
type ShowTables struct {
	Comments    Comments
	From        *TableName
	LikeOrWhere Expr
}

func (node *ShowTables) Format(buf *TrackedBuffer) {
	table := ""
	if node.From != nil {
		table = " from " + String(node.From)
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v tables%s", node.Comments, table)
	} else {
		buf.Fprintf("show %v tables%s%v", node.Comments, table, node.LikeOrWhere)
	}
}

func (node *ShowTables) IStatement()     {}
func (node *ShowTables) IShowStatement() {}

// ShowFullTables statement
type ShowFullTables struct {
	Comments    Comments
	From        *TableName
	LikeOrWhere Expr
}

func (node *ShowFullTables) Format(buf *TrackedBuffer) {
	table := ""
	if node.From != nil {
		table = " from " + String(node.From)
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v full tables%s", node.Comments, table)
	} else {
		buf.Fprintf("show %v full tables%s%v", node.Comments, table, node.LikeOrWhere)
	}
}

func (node *ShowFullTables) IStatement()     {}
func (node *ShowFullTables) IShowStatement() {}

// ShowColumns statement
type ShowColumns struct {
	Comments    Comments
	From        *TableName
	LikeOrWhere Expr
}

func (node *ShowColumns) Format(buf *TrackedBuffer) {
	table := ""
	if node.From != nil {
		table = " from " + String(node.From)
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v columns%s", node.Comments, table)
	} else {
		buf.Fprintf("show %v columns%s%v", node.Comments, table, node.LikeOrWhere)
	}
}

func (node *ShowColumns) IStatement()     {}
func (node *ShowColumns) IShowStatement() {}

// ShowFullColumns statement
type ShowFullColumns struct {
	Comments    Comments
	From        *TableName
	LikeOrWhere Expr
}

func (node *ShowFullColumns) Format(buf *TrackedBuffer) {
	table := ""
	if node.From != nil {
		table = " from " + String(node.From)
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v full columns%s", node.Comments, table)
	} else {
		buf.Fprintf("show %v full columns%s%v", node.Comments, table, node.LikeOrWhere)
	}
}

func (node *ShowFullColumns) IStatement()     {}
func (node *ShowFullColumns) IShowStatement() {}

// ShowProcedureStatus statement
type ShowProcedureStatus struct {
	Comments    Comments
	LikeOrWhere Expr
}

func (node *ShowProcedureStatus) Format(buf *TrackedBuffer) {
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v procedure status", node.Comments)
	} else {
		buf.Fprintf("show %v procedure status%v", node.Comments, node.LikeOrWhere)
	}
}

func (node *ShowProcedureStatus) IStatement()     {}
func (node *ShowProcedureStatus) IShowStatement() {}

// ShowFunctionStatus statement
type ShowFunctionStatus struct {
	Comments    Comments
	LikeOrWhere Expr
}

func (node *ShowFunctionStatus) Format(buf *TrackedBuffer) {
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v function status", node.Comments)
	} else {
		buf.Fprintf("show %v function status%v", node.Comments, node.LikeOrWhere)
	}
}

func (node *ShowFunctionStatus) IStatement()     {}
func (node *ShowFunctionStatus) IShowStatement() {}

// ShowIndex statement
type ShowIndex struct {
	Comments Comments
	From     *TableName
	Where    BoolExpr
}

func (node *ShowIndex) Format(buf *TrackedBuffer) {
	table := ""
	if node.From != nil {
		table = " from " + String(node.From)
	}
	if node.Where == nil {
		buf.Fprintf("show %v index%s", node.Comments, table)
	} else {
		buf.Fprintf("show %v index%s where %v", node.Comments, table, node.Where)
	}
}

func (node *ShowIndex) IStatement()     {}
func (node *ShowIndex) IShowStatement() {}

// ShowEngines statement
type ShowEngines struct {
}

func (node *ShowEngines) Format(buf *TrackedBuffer) {
	buf.Fprintf("show engines")
}

func (node *ShowEngines) IStatement()     {}
func (node *ShowEngines) IShowStatement() {}

// ShowPlugins statement
type ShowPlugins struct {
}

func (node *ShowPlugins) Format(buf *TrackedBuffer) {
	buf.Fprintf("show plugins")
}

func (node *ShowPlugins) IStatement()     {}
func (node *ShowPlugins) IShowStatement() {}

// ShowTriggers statement
type ShowTriggers struct {
	Comments    Comments
	From        *TableName
	LikeOrWhere Expr
}

func (node *ShowTriggers) Format(buf *TrackedBuffer) {
	table := ""
	if node.From != nil {
		table = " from " + String(node.From)
	}
	if node.LikeOrWhere == nil {
		buf.Fprintf("show %v triggers%s", node.Comments, table)
	} else {
		buf.Fprintf("show %v triggers%s%v", node.Comments, table, node.LikeOrWhere)
	}
}

func (node *ShowTriggers) IStatement()     {}
func (node *ShowTriggers) IShowStatement() {}

// ShowCreateDatabase statement
type ShowCreateDatabase struct {
	Comments Comments
	Name     *TableName
}

func (node *ShowCreateDatabase) Format(buf *TrackedBuffer) {
	name := ""
	if node.Name != nil {
		name = " " + String(node.Name)
	}
	buf.Fprintf("show %v create database%s", node.Comments, name)
}

func (node *ShowCreateDatabase) IStatement()     {}
func (node *ShowCreateDatabase) IShowStatement() {}

// ShowCreateTable statement
type ShowCreateTable struct {
	Comments Comments
	Name     *TableName
}

func (node *ShowCreateTable) Format(buf *TrackedBuffer) {
	name := ""
	if node.Name != nil {
		name = " " + String(node.Name)
	}
	buf.Fprintf("show %v create table%s", node.Comments, name)
}

func (node *ShowCreateTable) IStatement()     {}
func (node *ShowCreateTable) IShowStatement() {}

// ShowCreateView statement
type ShowCreateView struct {
	Comments Comments
	Name     *TableName
}

func (node *ShowCreateView) Format(buf *TrackedBuffer) {
	name := ""
	if node.Name != nil {
		name = " " + String(node.Name)
	}
	buf.Fprintf("show %v create view%s", node.Comments, name)
}

func (node *ShowCreateView) IStatement()     {}
func (node *ShowCreateView) IShowStatement() {}

// ShowCreateProcedure statement
type ShowCreateProcedure struct {
	Comments Comments
	Name     *TableName
}

func (node *ShowCreateProcedure) Format(buf *TrackedBuffer) {
	name := ""
	if node.Name != nil {
		name = " " + String(node.Name)
	}
	buf.Fprintf("show %v create procedure%s", node.Comments, name)
}

func (node *ShowCreateProcedure) IStatement()     {}
func (node *ShowCreateProcedure) IShowStatement() {}

// ShowCreateFunction statement
type ShowCreateFunction struct {
	Comments Comments
	Name     *TableName
}

func (node *ShowCreateFunction) Format(buf *TrackedBuffer) {
	name := ""
	if node.Name != nil {
		name = " " + String(node.Name)
	}
	buf.Fprintf("show %v create function%s", node.Comments, name)
}

func (node *ShowCreateFunction) IStatement()     {}
func (node *ShowCreateFunction) IShowStatement() {}

// ShowCreateTrigger statement
type ShowCreateTrigger struct {
	Comments Comments
	Name     *TableName
}

func (node *ShowCreateTrigger) Format(buf *TrackedBuffer) {
	name := ""
	if node.Name != nil {
		name = " " + String(node.Name)
	}
	buf.Fprintf("show %v create trigger%s", node.Comments, name)
}

func (node *ShowCreateTrigger) IStatement()     {}
func (node *ShowCreateTrigger) IShowStatement() {}

// ShowProcessList statement
type ShowProcessList struct {
	Comments Comments
}

func (node *ShowProcessList) Format(buf *TrackedBuffer) {
	buf.Fprintf("show %v processlist", node.Comments)
}

func (node *ShowProcessList) IStatement()     {}
func (node *ShowProcessList) IShowStatement() {}

// ShowFullProcessList statement
type ShowFullProcessList struct {
	Comments Comments
}

func (node *ShowFullProcessList) Format(buf *TrackedBuffer) {
	buf.Fprintf("show %v full processlist", node.Comments)
}

func (node *ShowFullProcessList) IStatement()     {}
func (node *ShowFullProcessList) IShowStatement() {}

// ShowSlaveStatus statement
type ShowSlaveStatus struct {
	Comments Comments
}

func (node *ShowSlaveStatus) Format(buf *TrackedBuffer) {
	buf.Fprintf("show %v slave status", node.Comments)
}

func (node *ShowSlaveStatus) IStatement()     {}
func (node *ShowSlaveStatus) IShowStatement() {}
