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

	"github.com/berkaroad/saashard/errors"
)

// FieldData is field's dump data.
type FieldData []byte

// Field struct.
type Field struct {
	Data FieldData

	Schema   []byte
	Table    []byte
	OrgTable []byte
	Name     []byte
	OrgName  []byte

	Charset      uint16
	ColumnLength uint32
	ColumnType   uint8
	Flags        uint16
	Decimals     uint8

	DefaultValueLength uint64
	DefaultValue       []byte
}

// Parse FieldData as Field.
func (p FieldData) Parse() (f *Field, err error) {
	f = new(Field)

	f.Data = p

	var n int
	pos := 0
	//skip catelog(always def)
	n, err = LengthOfLenencStr(p)
	if err != nil {
		return
	}
	pos += n

	//schema(lenenc_str)
	f.Schema, _, n, err = LenencStrToString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//table(lenenc_str)
	f.Table, _, n, err = LenencStrToString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//org_table(lenenc_str)
	f.OrgTable, _, n, err = LenencStrToString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//name(lenenc_str)
	f.Name, _, n, err = LenencStrToString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//org_name(lenenc_str)
	f.OrgName, _, n, err = LenencStrToString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//skip next_length(0x0c)
	pos++

	//character_set
	f.Charset = binary.LittleEndian.Uint16(p[pos:])
	pos += 2

	//column_length
	f.ColumnLength = binary.LittleEndian.Uint32(p[pos:])
	pos += 4

	//column_type
	f.ColumnType = p[pos]
	pos++

	//flags
	f.Flags = binary.LittleEndian.Uint16(p[pos:])
	pos += 2

	//decimals 1
	f.Decimals = p[pos]
	pos++

	//filter [0x00][0x00]
	pos += 2

	f.DefaultValue = nil
	//if more data, command was field list
	if len(p) > pos {
		//length of default value lenenc-int
		f.DefaultValueLength, _, n = LenencIntToNumber(p[pos:])
		pos += n

		if pos+int(f.DefaultValueLength) > len(p) {
			err = errors.ErrMalformPacket
			return
		}

		//default value string[$len]
		f.DefaultValue = p[pos:(pos + int(f.DefaultValueLength))]
	}

	return
}

// Dump Field as byte array.
func (f *Field) Dump() []byte {
	if f.Data != nil {
		return []byte(f.Data)
	}

	l := len(f.Schema) + len(f.Table) + len(f.OrgTable) + len(f.Name) + len(f.OrgName) + len(f.DefaultValue) + 48

	data := make([]byte, 0, l)

	data = append(data, StringToLenencStr([]byte("def"))...)

	data = append(data, StringToLenencStr(f.Schema)...)

	data = append(data, StringToLenencStr(f.Table)...)
	data = append(data, StringToLenencStr(f.OrgTable)...)

	data = append(data, StringToLenencStr(f.Name)...)
	data = append(data, StringToLenencStr(f.OrgName)...)

	data = append(data, 0x0c)

	data = append(data, Uint16ToBytes(f.Charset)...)
	data = append(data, Uint32ToBytes(f.ColumnLength)...)
	data = append(data, f.ColumnType)
	data = append(data, Uint16ToBytes(f.Flags)...)
	data = append(data, f.Decimals)
	data = append(data, 0, 0)

	if f.DefaultValue != nil {
		data = append(data, Uint64ToBytes(f.DefaultValueLength)...)
		data = append(data, f.DefaultValue...)
	}

	return data
}
