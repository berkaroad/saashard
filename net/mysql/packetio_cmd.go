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

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/sqlparser"
)

// WriteCommand write command.
func (p *PacketIO) WriteCommand(command byte) error {
	p.Sequence = 0

	return p.WritePacket([]byte{
		0x01, //1 bytes long
		0x00,
		0x00,
		0x00, //sequence
		command,
	})
}

// WriteCommandBuf write command buffer.
func (p *PacketIO) WriteCommandBuf(command byte, arg []byte) error {
	p.Sequence = 0
	length := len(arg) + 1
	data := make([]byte, length+4)

	data[4] = command

	copy(data[5:], arg)

	return p.WritePacket(data)
}

// WriteCommandStr write command, with string argument
func (p *PacketIO) WriteCommandStr(command byte, arg string) error {
	p.Sequence = 0
	length := len(arg) + 1
	data := make([]byte, length+4)

	data[4] = command

	copy(data[5:], arg)

	return p.WritePacket(data)
}

// WriteCommandUint32 write command, with uint32 argument
func (p *PacketIO) WriteCommandUint32(command byte, arg uint32) error {
	p.Sequence = 0

	return p.WritePacket([]byte{
		0x05, //5 bytes long
		0x00,
		0x00,
		0x00, //sequence

		command,

		byte(arg),
		byte(arg >> 8),
		byte(arg >> 16),
		byte(arg >> 24),
	})
}

// WriteCommandStrStr write command ,with 2 string arguments.
func (p *PacketIO) WriteCommandStrStr(command byte, arg1 string, arg2 string) error {
	p.Sequence = 0
	data := make([]byte, 4, 6+len(arg1)+len(arg2))

	data = append(data, command)
	data = append(data, arg1...)
	data = append(data, 0)
	data = append(data, arg2...)

	return p.WritePacket(data)
}

// Sleep use command COM_SLEEP
//

// Quit use command COM_QUIT
func (p *PacketIO) Quit(capability uint32, status *uint16) error {
	if err := p.WriteCommand(COM_QUIT); err != nil {
		return err
	}

	if _, err := p.ReadOK(capability, status); err != nil {
		return err
	}
	return nil
}

// InitDB use command COM_INIT_DB
func (p *PacketIO) InitDB(capability uint32, status *uint16, dbName string) error {
	if err := p.WriteCommandStr(COM_INIT_DB, dbName); err != nil {
		return err
	}

	if _, err := p.ReadOK(capability, status); err != nil {
		return err
	}
	return nil
}

// Query use command COM_QUERY
func (p *PacketIO) Query(capability uint32, status *uint16, query string) (*Result, error) {
	var result *Result
	var err error
	if capability&CLIENT_MULTI_STATEMENTS == 0 {
		if err := p.WriteCommandStr(COM_QUERY, query); err != nil {
			return nil, err
		}
		result, err = p.ReadResultSet(capability, status, false)
	} else {
		sqlStatementList := sqlparser.SplitSQLStatement(query)
		for _, sqlStatement := range sqlStatementList {
			if err := p.WriteCommandStr(COM_QUERY, sqlStatement); err != nil {
				return nil, err
			}
			result, err = p.ReadResultSet(capability, status, false)
		}
	}

	return result, err
}

// FieldList use command COM_FIELD_LIST
func (p *PacketIO) FieldList(capability uint32, table string, wildcard string) ([]*Field, error) {
	if err := p.WriteCommandStrStr(COM_FIELD_LIST, table, wildcard); err != nil {
		return nil, err
	}

	data, err := p.ReadPacket()
	if err != nil {
		return nil, err
	}

	fs := make([]*Field, 0, 4)
	var f *Field
	if data[0] == ERR_HEADER {
		return nil, p.handleErrorPacket(capability, data)
	}
	for {
		if data, err = p.ReadPacket(); err != nil {
			return nil, err
		}

		// EOF Packet
		if p.isEOFPacket(data) {
			return fs, nil
		}

		if f, err = FieldData(data).Parse(); err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
}

// CreateDB use command COM_CREATE_DB
//

// DropDB use command COM_DROP_DB
//

// Refresh use command COM_REFRESH
//

// Shutdown use command COM_SHUTDOWN -> SHUTDOWN_DEFAULT
func (p *PacketIO) Shutdown(capability uint32, status *uint16) (*Result, error) {
	var result *Result
	var err error
	if err := p.WriteCommandBuf(COM_SHUTDOWN, []byte{SHUTDOWN_DEFAULT}); err != nil {
		return nil, err
	}
	result, err = p.ReadOK(capability, status)

	return result, err
}

// KillConnection use command COM_SHUTDOWN -> KILL_CONNECTION
func (p *PacketIO) KillConnection(capability uint32, status *uint16) (*Result, error) {
	var result *Result
	var err error
	if err := p.WriteCommandBuf(COM_SHUTDOWN, []byte{KILL_CONNECTION}); err != nil {
		return nil, err
	}
	result, err = p.ReadOK(capability, status)

	return result, err
}

// Statistics use command COM_STATISTICS
//

// ProcessInfo use command COM_PROCESS_INFO
//

// Connect use command COM_CONNECT
//

// ProcessKill use command COM_PROCESS_KILL
//

// Debug use command COM_DEBUG
//

// Ping use command COM_PING
func (p *PacketIO) Ping(capability uint32, status *uint16) error {
	if err := p.WriteCommand(COM_PING); err != nil {
		return err
	}

	if _, err := p.ReadOK(capability, status); err != nil {
		return err
	}

	return nil
}

// Time use command COM_TIME
//

// DelayedInsert use command COM_DELAYED_INSERT
//

// ChangeUser use command COM_CHANGE_USER
//

// BinlogDump use command COM_BINLOG_DUMP
//

// TableDump use command COM_TABLE_DUMP
//

// ConnectOut use command COM_CONNECT_OUT
//

// RegisterSlave use command COM_REGISTER_SLAVE
//

// StmtPrepare use command COM_STMT_PREPARE
func (p *PacketIO) StmtPrepare(capability uint32, query string) (stmtID uint32, columnNumber, paramNumber int, err error) {
	if err = p.WriteCommandStr(COM_STMT_PREPARE, query); err != nil {
		return 0, 0, 0, err
	}
	var data []byte
	data, err = p.ReadPacket()
	if err != nil {
		return 0, 0, 0, err
	}

	if data[0] == ERR_HEADER {
		return 0, 0, 0, p.handleErrorPacket(capability, data)
	}
	if data[0] != OK_HEADER {
		return 0, 0, 0, errors.ErrMalformPacket
	}

	pos := 1

	//for statement id
	stmtID = binary.LittleEndian.Uint32(data[pos:])
	pos += 4

	//number columns
	columnNumber = int(binary.LittleEndian.Uint16(data[pos:]))
	pos += 2

	//number params
	paramNumber = int(binary.LittleEndian.Uint16(data[pos:]))
	pos += 2

	//warnings
	//warnings = binary.LittleEndian.Uint16(data[pos:])

	if paramNumber > 0 {
		if err := p.readUntilEOF(); err != nil {
			return 0, 0, 0, err
		}
	}

	if columnNumber > 0 {
		if err := p.readUntilEOF(); err != nil {
			return 0, 0, 0, err
		}
	}
	return
}

// StmtExecute use command COM_STMT_EXECUTE
func (p *PacketIO) StmtExecute(stmtID uint32, args ...interface{}) error {
	paramsNum := len(args)

	paramTypes := make([]byte, paramsNum<<1)
	paramValues := make([][]byte, paramsNum)

	//NULL-bitmap, length: (num-params+7)
	nullBitmap := make([]byte, (paramsNum+7)>>3)

	var length = int(1 + 4 + 1 + 4 + ((paramsNum + 7) >> 3) + 1 + (paramsNum << 1))

	var newParamBoundFlag byte

	for i := range args {
		if args[i] == nil {
			nullBitmap[i/8] |= (1 << (uint(i) % 8))
			paramTypes[i<<1] = MYSQL_TYPE_NULL
			continue
		}

		newParamBoundFlag = 1

		switch v := args[i].(type) {
		case int8:
			paramTypes[i<<1] = MYSQL_TYPE_TINY
			paramValues[i] = []byte{byte(v)}
		case int16:
			paramTypes[i<<1] = MYSQL_TYPE_SHORT
			paramValues[i] = Uint16ToBytes(uint16(v))
		case int32:
			paramTypes[i<<1] = MYSQL_TYPE_LONG
			paramValues[i] = Uint32ToBytes(uint32(v))
		case int:
			paramTypes[i<<1] = MYSQL_TYPE_LONGLONG
			paramValues[i] = Uint64ToBytes(uint64(v))
		case int64:
			paramTypes[i<<1] = MYSQL_TYPE_LONGLONG
			paramValues[i] = Uint64ToBytes(uint64(v))
		case uint8:
			paramTypes[i<<1] = MYSQL_TYPE_TINY
			paramTypes[(i<<1)+1] = 0x80
			paramValues[i] = []byte{v}
		case uint16:
			paramTypes[i<<1] = MYSQL_TYPE_SHORT
			paramTypes[(i<<1)+1] = 0x80
			paramValues[i] = Uint16ToBytes(uint16(v))
		case uint32:
			paramTypes[i<<1] = MYSQL_TYPE_LONG
			paramTypes[(i<<1)+1] = 0x80
			paramValues[i] = Uint32ToBytes(uint32(v))
		case uint:
			paramTypes[i<<1] = MYSQL_TYPE_LONGLONG
			paramTypes[(i<<1)+1] = 0x80
			paramValues[i] = Uint64ToBytes(uint64(v))
		case uint64:
			paramTypes[i<<1] = MYSQL_TYPE_LONGLONG
			paramTypes[(i<<1)+1] = 0x80
			paramValues[i] = Uint64ToBytes(uint64(v))
		case bool:
			paramTypes[i<<1] = MYSQL_TYPE_TINY
			if v {
				paramValues[i] = []byte{1}
			} else {
				paramValues[i] = []byte{0}
			}
		case float32:
			paramTypes[i<<1] = MYSQL_TYPE_FLOAT
			paramValues[i] = Uint32ToBytes(math.Float32bits(v))
		case float64:
			paramTypes[i<<1] = MYSQL_TYPE_DOUBLE
			paramValues[i] = Uint64ToBytes(math.Float64bits(v))
		case string:
			paramTypes[i<<1] = MYSQL_TYPE_STRING
			paramValues[i] = append(NumberToLenencInt(uint64(len(v))), v...)
		case []byte:
			paramTypes[i<<1] = MYSQL_TYPE_STRING
			paramValues[i] = append(NumberToLenencInt(uint64(len(v))), v...)
		default:
			return fmt.Errorf("invalid argument type %T", args[i])
		}

		length += len(paramValues[i])
	}

	data := make([]byte, 4, 4+length)

	data = append(data, COM_STMT_EXECUTE)
	data = append(data, byte(stmtID), byte(stmtID>>8), byte(stmtID>>16), byte(stmtID>>24))

	//flag: CURSOR_TYPE_NO_CURSOR
	data = append(data, 0x00)

	//iteration-count, always 1
	data = append(data, 1, 0, 0, 0)

	if paramsNum > 0 {
		data = append(data, nullBitmap...)

		//new-params-bound-flag
		data = append(data, newParamBoundFlag)

		if newParamBoundFlag == 1 {
			//type of each parameter, length: num-params * 2
			data = append(data, paramTypes...)

			//value of each parameter
			for _, v := range paramValues {
				data = append(data, v...)
			}
		}
	}
	p.Sequence = 0
	return p.WritePacket(data)
}

// StmtSendLongData use command COM_STMT_SEND_LONG_DATA
//

// StmtClose use command COM_STMT_CLOSE
func (p *PacketIO) StmtClose(stmtID uint32) error {
	if err := p.WriteCommandUint32(COM_STMT_CLOSE, stmtID); err != nil {
		return err
	}
	return nil
}

// StmtReset use command COM_STMT_RESET
//

// SetOption use command COM_SET_OPTION
//

// StmtFetch use command COM_STMT_FETCH
//

// Daemon use command COM_DAEMON
//

// BinlogDumpGTID use command COM_BINLOG_DUMP_GTID
//

// ResetConnection use command COM_RESET_CONNECTION
//
