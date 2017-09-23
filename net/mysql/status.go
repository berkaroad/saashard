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

// https://dev.mysql.com/doc/internals/en/status-flags.html
const (
	SERVER_STATUS_IN_TRANS uint16 = uint16(1 << iota)
	SERVER_STATUS_AUTOCOMMIT
)

const (
	SERVER_MORE_RESULTS_EXISTS uint16 = uint16(1 << (iota + 3))
	SERVER_STATUS_NO_GOOD_INDEX_USED
	SERVER_STATUS_NO_INDEX_USED
	SERVER_STATUS_CURSOR_EXISTS
	SERVER_STATUS_LAST_ROW_SEND
	SERVER_STATUS_DB_DROPPED
	SERVER_STATUS_NO_BACKSLASH_ESCAPED
	SERVER_STATUS_METADATA_CHANGED
	SERVER_QUERY_WAS_SLOW
	SERVER_PS_OUT_PARAMS
	SERVER_STATUS_IN_TRANS_READONLY
	SERVER_SESSION_STATE_CHANGED
)
