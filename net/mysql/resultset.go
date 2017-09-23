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
	"fmt"
	"strconv"
)

// Result is query's result
type Result struct {
	Status       uint16
	InsertID     uint64
	AffectedRows uint64
	*Resultset
}

// Resultset is result set of result.
type Resultset struct {
	Fields     []*Field
	FieldNames map[string]int
	Values     [][]interface{}

	Rows []*Row
}

// RowNumber row number
func (r *Resultset) RowNumber() int {
	return len(r.Values)
}

// ColumnNumber column number
func (r *Resultset) ColumnNumber() int {
	return len(r.Fields)
}

// GetValue get value
func (r *Resultset) GetValue(row, column int) (interface{}, error) {
	if row >= len(r.Values) || row < 0 {
		return nil, fmt.Errorf("invalid row index %d", row)
	}

	if column >= len(r.Fields) || column < 0 {
		return nil, fmt.Errorf("invalid column index %d", column)
	}

	return r.Values[row][column], nil
}

// GetIndexByName get index by name.
func (r *Resultset) GetIndexByName(name string) (int, error) {
	if column, ok := r.FieldNames[name]; ok {
		return column, nil
	}
	return 0, fmt.Errorf("invalid field name %s", name)
}

// GetValueByName get value by name.
func (r *Resultset) GetValueByName(row int, name string) (interface{}, error) {
	var column int
	var err error
	if column, err = r.GetIndexByName(name); err != nil {
		return nil, err
	}
	return r.GetValue(row, column)
}

// IsNull is null.
func (r *Resultset) IsNull(row, column int) (bool, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return false, err
	}
	return d == nil, nil
}

// IsNullByName is the column null.
func (r *Resultset) IsNullByName(row int, name string) (bool, error) {
	var column int
	var err error
	if column, err = r.GetIndexByName(name); err != nil {
		return false, err
	}
	return r.IsNull(row, column)
}

// GetUint get value as uint.
func (r *Resultset) GetUint(row, column int) (uint64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case uint64:
		return v, nil
	case int64:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case string:
		return strconv.ParseUint(v, 10, 64)
	case []byte:
		return strconv.ParseUint(string(v), 10, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

// GetUintByName get value as uint.
func (r *Resultset) GetUintByName(row int, name string) (uint64, error) {
	column, err := r.GetIndexByName(name)
	if err != nil {
		return 0, err
	}
	return r.GetUint(row, column)
}

// GetInt get value as int.
func (r *Resultset) GetInt(row, column int) (int64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case uint64:
		return int64(v), nil
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	case []byte:
		return strconv.ParseInt(string(v), 10, 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

// GetIntByName get value as int.
func (r *Resultset) GetIntByName(row int, name string) (int64, error) {
	column, err := r.GetIndexByName(name)
	if err != nil {
		return 0, err
	}
	return r.GetInt(row, column)
}

// GetFloat get value as float.
func (r *Resultset) GetFloat(row, column int) (float64, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case float64:
		return v, nil
	case uint64:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

// GetFloatByName get value as float.
func (r *Resultset) GetFloatByName(row int, name string) (float64, error) {
	column, err := r.GetIndexByName(name)
	if err != nil {
		return 0, err
	}
	return r.GetFloat(row, column)
}

// GetString get value as string.
func (r *Resultset) GetString(row, column int) (string, error) {
	d, err := r.GetValue(row, column)
	if err != nil {
		return "", err
	}

	switch v := d.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("data type is %T", v)
	}
}

// GetStringByName get value as string.
func (r *Resultset) GetStringByName(row int, name string) (string, error) {
	column, err := r.GetIndexByName(name)
	if err != nil {
		return "", err
	}
	return r.GetString(row, column)
}
