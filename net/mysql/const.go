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

const (
	MinProtocolVersion byte   = 10
	MaxPayloadLen      int    = 1<<24 - 1
	ServerVersion      string = "5.6.0-SaaShard"
	SourceInfo         string = "github.com/berkaroad/saashard"
)

const (
	OK_HEADER          byte = 0x00
	ERR_HEADER         byte = 0xff
	EOF_HEADER         byte = 0xfe
	LocalInFile_HEADER byte = 0xfb
)

const (
	NOT_NULL_FLAG uint16 = 1 << iota
	PRI_KEY_FLAG
	UNIQUE_KEY_FLAG
	MUlTIPLE_KEY_FLAG
	BLOB_FLAG
	UNSIGNED_FLAG
	ZERO_FILL_FLAG
	BINARY_FLAG
	ENUM_FLAG
	AUTO_INCREMENT_FLAG
	TIMESTAMP_FLAG
	SET_FLAG

	PART_KEY_FLAG int = 1 << (14 + iota)
	NUM_OR_GROUP_FLAG
	UNIQUE_FLAG
)

const (
	AUTH_NAME = "mysql_native_password"
)

var (
	TK_ID_INSERT   = 1
	TK_ID_UPDATE   = 2
	TK_ID_DELETE   = 3
	TK_ID_REPLACE  = 4
	TK_ID_SET      = 5
	TK_ID_BEGIN    = 6
	TK_ID_COMMIT   = 7
	TK_ID_ROLLBACK = 8
	TK_ID_ADMIN    = 9
	TK_ID_USE      = 10

	TK_ID_SELECT      = 11
	TK_ID_START       = 12
	TK_ID_TRANSACTION = 13
	TK_ID_SHOW        = 14

	PARSE_TOKEN_MAP = map[string]int{
		"insert":      TK_ID_INSERT,
		"update":      TK_ID_UPDATE,
		"delete":      TK_ID_DELETE,
		"replace":     TK_ID_REPLACE,
		"set":         TK_ID_SET,
		"begin":       TK_ID_BEGIN,
		"commit":      TK_ID_COMMIT,
		"rollback":    TK_ID_ROLLBACK,
		"admin":       TK_ID_ADMIN,
		"select":      TK_ID_SELECT,
		"use":         TK_ID_USE,
		"start":       TK_ID_START,
		"transaction": TK_ID_TRANSACTION,
		"show":        TK_ID_SHOW,
	}
	// '*'
	COMMENT_PREFIX uint8 = 42
	COMMENT_STRING       = "*"

	//
	TK_STR_FROM = "from"
	TK_STR_INTO = "into"
	TK_STR_SET  = "set"

	TK_STR_TRANSACTION    = "transaction"
	TK_STR_LAST_INSERT_ID = "last_insert_id()"
	TK_STR_MASTER_HINT    = "*master*"

	SET_KEY_WORDS = map[string]struct{}{
		"names": struct{}{},

		"character_set_results":           struct{}{},
		"@@character_set_results":         struct{}{},
		"@@session.character_set_results": struct{}{},

		"character_set_client":           struct{}{},
		"@@character_set_client":         struct{}{},
		"@@session.character_set_client": struct{}{},

		"character_set_connection":           struct{}{},
		"@@character_set_connection":         struct{}{},
		"@@session.character_set_connection": struct{}{},

		"autocommit":           struct{}{},
		"@@autocommit":         struct{}{},
		"@@session.autocommit": struct{}{},
	}
)

const (
	DONTESCAPE = byte(255)
)
