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

package mysql

import "github.com/berkaroad/saashard/sqlparser"

// Stmt for stmt.
type Stmt struct {
	pkg        *PacketIO
	capability uint32
	status     *uint16

	ID        uint32
	Query     string
	Statement sqlparser.Statement
	ParamNum  int
	Params    []*Field
	ColumnNum int
	Columns   []*Field
	Args      []interface{}
}

// NewStmt new stmt.
func NewStmt(pkg *PacketIO, capability uint32, status *uint16) *Stmt {
	s := new(Stmt)
	s.pkg = pkg
	s.capability = capability
	s.status = status
	return s
}

// Execute stmt.
func (s *Stmt) Execute(args []interface{}) (*Result, error) {
	if err := s.pkg.StmtExecute(s.ID, args); err != nil {
		return nil, err
	}
	return s.pkg.ReadResultSet(s.capability, s.status, true)
}

// Close stmt.
func (s *Stmt) Close() error {
	if err := s.pkg.StmtClose(s.ID); err != nil {
		return err
	}

	return nil
}

// ResetParams reset params.
func (s *Stmt) ResetParams() {
	s.Args = make([]interface{}, s.ParamNum)
}
