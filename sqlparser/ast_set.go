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

// http://dev.mysql.com/doc/refman/5.6/en/set-statement.html

// SetStatement set statement
type SetStatement interface {
	IStatement()
	ISetStatement()
	SQLNode
}

// SetVariable set variable assignment
type SetVariable struct {
	Comments Comments
	Scope    string // SESSION or GLOBAL
	Exprs    UpdateExprs
}

func (node *SetVariable) Format(buf *TrackedBuffer) {
	if len(node.Scope) == 0 {
		buf.Fprintf("set %v %v", node.Comments, node.Exprs)
	} else {
		buf.Fprintf("set %v %s %v", node.Comments, node.Scope, node.Exprs)
	}
}

func (node *SetVariable) IStatement()    {}
func (node *SetVariable) ISetStatement() {}

// SetCharset set charset.
type SetCharset struct {
	Comments Comments
	Scope    string // SESSION or GLOBAL
	Charset  string
}

func (node *SetCharset) Format(buf *TrackedBuffer) {
	if len(node.Scope) == 0 {
		buf.Fprintf("set %v character set %s", node.Comments, node.Charset)
	} else {
		buf.Fprintf("set %v %s character set %s", node.Comments, node.Scope, node.Charset)
	}
}

func (node *SetCharset) IStatement()    {}
func (node *SetCharset) ISetStatement() {}

// SetNames set NAMES.
type SetNames struct {
	Comments Comments
	Names    string
	Collate  string
}

func (node *SetNames) Format(buf *TrackedBuffer) {
	if node.Collate == "" {
		buf.Fprintf("set %v names %s", node.Comments, node.Names)
	} else {
		buf.Fprintf("set %v names %s collate %s", node.Comments, node.Names, node.Collate)
	}
}

func (node *SetNames) IStatement()    {}
func (node *SetNames) ISetStatement() {}

// SetTransactionIsolationLevel represents a SET TRANSACTION ISOLATION LEVEL statement.
type SetTransactionIsolationLevel struct {
	Comments       Comments
	Scope          string // SESSION or GLOBAL
	IsolationLevel string
}

func (node *SetTransactionIsolationLevel) Format(buf *TrackedBuffer) {
	if len(node.Scope) == 0 {
		buf.Fprintf("set %v transaction isolation level %s", node.Comments, node.IsolationLevel)
	} else {
		buf.Fprintf("set %v %s transaction isolation level %s", node.Comments, node.Scope, node.IsolationLevel)
	}
}

func (node *SetTransactionIsolationLevel) IStatement()    {}
func (node *SetTransactionIsolationLevel) ISetStatement() {}
