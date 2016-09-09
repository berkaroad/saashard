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
	"strconv"

	"github.com/berkaroad/saashard/errors"
)

// RowData is row's dump data.
type RowData []byte

// Parse RowData as Row.
func (p RowData) Parse(isBinary bool, fields []*Field) (r *Row, err error) {
	r = new(Row)
	r.isBinary = isBinary
	r.fields = fields

	if r.isBinary {
		r.fieldValues, r.nullBitmap, r.fieldValuesCache, err = parseAsBinaryRow(p, r.fields)
	}
	r.fieldValues, r.fieldValuesCache, err = parseAsTextRow(p, r.fields)
	return
}

// Row struct
type Row struct {
	Data RowData

	isBinary    bool
	fields      []*Field
	fieldValues []interface{}

	nullBitmap       []byte
	fieldValuesCache [][]byte
}

// NewTextRow new rowdata.
func NewTextRow(fields []*Field) *Row {
	r := new(Row)
	r.isBinary = false
	r.fields = fields
	r.fieldValues = make([]interface{}, 0, len(fields))
	r.fieldValuesCache = make([][]byte, 0, len(fields))
	return r
}

// AppendStringValue append string value.
func (r *Row) AppendStringValue(str string) {
	if len(r.fieldValues) < len(r.fields) {
		encodedVal := StringToLenencStr([]byte(str))
		r.fieldValues = append(r.fieldValues, str)
		r.fieldValuesCache = append(r.fieldValuesCache, encodedVal)
	}
}

// AppendUIntValue append unsigned int value.
func (r *Row) AppendUIntValue(val uint64) {
	if len(r.fieldValues) < len(r.fields) {
		str := strconv.FormatUint(val, 10)
		encodedVal := StringToLenencStr([]byte(str))
		r.fieldValues = append(r.fieldValues, str)
		r.fieldValuesCache = append(r.fieldValuesCache, encodedVal)
	}
}

// AppendIntValue append int value.
func (r *Row) AppendIntValue(val int64) {
	if len(r.fieldValues) < len(r.fields) {
		str := strconv.FormatInt(val, 10)
		encodedVal := StringToLenencStr([]byte(str))
		r.fieldValues = append(r.fieldValues, str)
		r.fieldValuesCache = append(r.fieldValuesCache, encodedVal)
	}
}

// AppendFloatValue append float value.
func (r *Row) AppendFloatValue(val float64) {
	if len(r.fieldValues) < len(r.fields) {
		str := strconv.FormatFloat(val, 'f', 14, 64)
		encodedVal := StringToLenencStr([]byte(str))
		r.fieldValues = append(r.fieldValues, str)
		r.fieldValuesCache = append(r.fieldValuesCache, encodedVal)
	}
}

// AppendBooleanValue append bool value.
func (r *Row) AppendBooleanValue(val bool) {
	if len(r.fieldValues) < len(r.fields) {
		str := strconv.FormatBool(val)
		encodedVal := StringToLenencStr([]byte(str))
		r.fieldValues = append(r.fieldValues, str)
		r.fieldValuesCache = append(r.fieldValuesCache, encodedVal)
	}
}

// AppendNullValue append null value.
func (r *Row) AppendNullValue() {
	if len(r.fieldValues) < len(r.fields) {
		r.fieldValues = append(r.fieldValues, nil)
		r.fieldValuesCache = append(r.fieldValuesCache, []byte{0xfb})
	}
}

// Dump the Row as byte array.
func (r *Row) Dump() []byte {
	if r.Data != nil {
		return r.Data
	}

	data := make([]byte, 0, r.dataLength())
	if r.isBinary {
		data = append(data, OK_HEADER)
		data = append(data, r.nullBitmap...)
	}
	for _, fv := range r.fieldValuesCache {
		data = append(data, fv...)
	}
	r.Data = data
	return data
}

func (r *Row) dataLength() int {
	l := 0
	if r.isBinary {
		l = len(r.nullBitmap) + 1
	}

	for _, fv := range r.fieldValuesCache {
		l += len(fv)
	}
	return l
}

func parseAsTextRow(raw []byte, f []*Field) (fieldValues []interface{}, fieldValuesCache [][]byte, err error) {
	fieldValues = make([]interface{}, len(f))
	fieldValuesCache = make([][]byte, len(f))

	var v []byte
	var isNull, isUnsigned bool
	var pos int
	var n int

	for i := range f {
		v, isNull, n, err = LenencStrToString(raw[pos:])
		if err != nil {
			return nil, nil, err
		}
		fieldValuesCache[i] = StringToLenencStr(v)

		pos += n

		if isNull {
			fieldValues[i] = nil
		} else {
			isUnsigned = (f[i].Flags&UNSIGNED_FLAG > 0)
			switch f[i].ColumnType {
			case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_LONG, MYSQL_TYPE_INT24,
				MYSQL_TYPE_LONGLONG, MYSQL_TYPE_YEAR:
				if isUnsigned {
					fieldValues[i], err = strconv.ParseUint(string(v), 10, 64)
				} else {
					fieldValues[i], err = strconv.ParseInt(string(v), 10, 64)
				}
			case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE, MYSQL_TYPE_NEWDECIMAL:
				fieldValues[i], err = strconv.ParseFloat(string(v), 64)
			case MYSQL_TYPE_VARCHAR, MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING:
				fieldValues[i] = string(v)
			default:
				fieldValues[i] = v
			}

			if err != nil {
				return nil, nil, err
			}
		}
	}
	return
}

func parseAsBinaryRow(raw []byte, f []*Field) (fieldValues []interface{}, nullBitmap []byte, fieldValuesCache [][]byte, err error) {
	fieldValues = make([]interface{}, len(f))
	fieldValuesCache = make([][]byte, len(f))
	if raw[0] != OK_HEADER {
		return nil, nil, nil, errors.ErrMalformPacket
	}
	fieldValuesCache[0] = []byte{raw[0]}
	pos := 1 + ((len(f) + 7 + 2) >> 3)

	nullBitmap = raw[1:pos]

	var isUnsigned bool
	var isNull bool
	var n int

	var v []byte
	for i := range fieldValues {
		if nullBitmap[(i+2)/8]&(1<<(uint(i+2)%8)) > 0 {
			fieldValues[i] = nil
			continue
		}

		isUnsigned = f[i].Flags&UNSIGNED_FLAG > 0

		switch f[i].ColumnType {
		case MYSQL_TYPE_NULL:
			fieldValues[i] = nil
			continue

		case MYSQL_TYPE_TINY:
			if isUnsigned {
				fieldValues[i] = uint64(raw[pos])
			} else {
				fieldValues[i] = int64(raw[pos])
			}
			fieldValuesCache[i] = []byte{raw[pos]}
			pos++
			continue

		case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
			if isUnsigned {
				fieldValues[i] = uint64(binary.LittleEndian.Uint16(raw[pos : pos+2]))
			} else {
				fieldValues[i] = int64((binary.LittleEndian.Uint16(raw[pos : pos+2])))
			}
			fieldValuesCache[i] = raw[pos : pos+2]
			pos += 2
			continue

		case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
			if isUnsigned {
				fieldValues[i] = uint64(binary.LittleEndian.Uint32(raw[pos : pos+4]))
			} else {
				fieldValues[i] = int64(binary.LittleEndian.Uint32(raw[pos : pos+4]))
			}
			fieldValuesCache[i] = raw[pos : pos+4]
			pos += 4
			continue

		case MYSQL_TYPE_LONGLONG:
			if isUnsigned {
				fieldValues[i] = binary.LittleEndian.Uint64(raw[pos : pos+8])
			} else {
				fieldValues[i] = int64(binary.LittleEndian.Uint64(raw[pos : pos+8]))
			}
			fieldValuesCache[i] = raw[pos : pos+8]
			pos += 8
			continue

		case MYSQL_TYPE_FLOAT:
			fieldValues[i] = float64(math.Float32frombits(binary.LittleEndian.Uint32(raw[pos : pos+4])))
			fieldValuesCache[i] = raw[pos : pos+4]
			pos += 4
			continue

		case MYSQL_TYPE_DOUBLE:
			fieldValues[i] = math.Float64frombits(binary.LittleEndian.Uint64(raw[pos : pos+8]))
			fieldValuesCache[i] = raw[pos : pos+8]
			pos += 8
			continue

		case MYSQL_TYPE_DECIMAL, MYSQL_TYPE_NEWDECIMAL, MYSQL_TYPE_VARCHAR,
			MYSQL_TYPE_BIT, MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_TINY_BLOB,
			MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_LONG_BLOB, MYSQL_TYPE_BLOB,
			MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING, MYSQL_TYPE_GEOMETRY:
			v, isNull, n, err = LenencStrToString(raw[pos:])
			if err != nil {
				return nil, nil, nil, err
			}
			fieldValuesCache[i] = raw[pos : pos+n]
			pos += n

			if !isNull {
				fieldValues[i] = v
				continue
			} else {
				fieldValues[i] = nil
				continue
			}
		case MYSQL_TYPE_DATE, MYSQL_TYPE_NEWDATE:
			var num uint64
			num, isNull, n = LenencIntToNumber(raw[pos:])
			if isNull {
				fieldValuesCache[i] = raw[pos : pos+n]
			}
			pos += n

			if isNull {
				fieldValues[i] = nil
				continue
			}

			fieldValues[i], err = FormatBinaryDate(int(num), raw[pos:])

			if err != nil {
				return nil, nil, nil, err
			}
			fieldValuesCache[i] = raw[pos:int(num)]
			pos += int(num)

		case MYSQL_TYPE_TIMESTAMP, MYSQL_TYPE_DATETIME:
			var num uint64
			num, isNull, n = LenencIntToNumber(raw[pos:])
			if isNull {
				fieldValuesCache[i] = raw[pos : pos+n]
			}
			pos += n

			if isNull {
				fieldValues[i] = nil
				continue
			}

			fieldValues[i], err = FormatBinaryDateTime(int(num), raw[pos:])
			fieldValuesCache[i] = raw[pos : pos+int(num)]
			pos += int(num)

			if err != nil {
				return nil, nil, nil, err
			}

		case MYSQL_TYPE_TIME:
			var num uint64
			num, isNull, n = LenencIntToNumber(raw[pos:])
			if isNull {
				fieldValuesCache[i] = raw[pos:n]
			}
			pos += n

			if isNull {
				fieldValues[i] = nil
				continue
			}

			fieldValues[i], err = FormatBinaryTime(int(num), raw[pos:])
			if err != nil {
				return nil, nil, nil, err
			}
			fieldValuesCache[i] = raw[pos:int(num)]
			pos += int(num)

		default:
			return nil, nil, nil, fmt.Errorf("Stmt Unknown FieldType %d %s", f[i].ColumnType, f[i].Name)
		}
	}
	return
}
