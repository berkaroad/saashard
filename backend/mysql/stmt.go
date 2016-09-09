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

package mysql

import (
	"github.com/berkaroad/saashard/net/mysql"
)

type Stmt struct {
	conn  *Conn
	id    uint32
	query string

	params  int
	columns int
}

func (s *Stmt) Execute(args ...interface{}) (*mysql.Result, error) {
	if err := s.conn.pkg.StmtExecute(s.id, args...); err != nil {
		return nil, err
	}
	return s.conn.pkg.ReadResultSet(s.conn.capability, &(s.conn.status), true)
}

func (s *Stmt) Close() error {
	if err := s.conn.pkg.StmtClose(s.id); err != nil {
		return err
	}

	return nil
}
