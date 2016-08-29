package mysql

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"

	"github.com/berkaroad/saashard/errors"
)

// FieldValue is field's value in row.
type FieldValue []byte

// RowDataData is rowData's dump data.
type RowDataData []byte

// RowData is row data.
type RowData struct {
	Data RowDataData

	isBinary    bool
	fields      []*Field
	fieldValues []FieldValue
}

// NewRowData new rowdata.
func NewRowData(isBinary bool, fields []*Field) *RowData {
	r := new(RowData)
	r.isBinary = isBinary
	r.fields = fields
	r.fieldValues = make([]FieldValue, 0, len(fields))
	return r
}

// AppendStringValue append string value.
func (r *RowData) AppendStringValue(str []byte) {
	encodedVal := StringToLenencStr(str)
	r.fieldValues = append(r.fieldValues, encodedVal)
}

// AppendUIntValue append unsigned int value.
func (r *RowData) AppendUIntValue(val uint64) {
	r.AppendStringValue([]byte(strconv.FormatUint(val, 10)))
}

// AppendIntValue append int value.
func (r *RowData) AppendIntValue(val int64) {
	r.AppendStringValue([]byte(strconv.FormatInt(val, 10)))
}

// AppendFloatValue append float value.
func (r *RowData) AppendFloatValue(val float64) {
	r.AppendStringValue([]byte(strconv.FormatFloat(val, 'f', 14, 64)))
}

// AppendNullValue append null value.
func (r *RowData) AppendNullValue() {
	r.fieldValues = append(r.fieldValues, []byte{0xfb})
}

// Dump the RowData as byte array.
func (r *RowData) Dump() []byte {
	if r.Data != nil {
		return []byte(r.Data)
	}
	if r.isBinary {
		data := make([]byte, 1+(len(r.fields)+7+2)>>3, 1+(len(r.fields)+7+2)>>3+r.dataLength())
		data[0] = 0 // struct header
		// fill in null bitmap, not impletement.
		data[1] = 0

		for _, fv := range r.fieldValues {
			data = append(data, fv...)
		}
		return data
	}
	data := make([]byte, 0, r.dataLength())
	for _, fv := range r.fieldValues {
		data = append(data, fv...)
	}
	return data
}

func (r *RowData) dataLength() int {
	l := 0
	for _, fv := range r.fieldValues {
		l += len(fv)
	}
	return l
}

// ParseAsRowData to parse as row data.
func ParseAsRowData(raw []byte, f []*Field) ([]interface{}, error) {
	if raw[0] == 0 {
		return parseAsRowDataBinary(raw, f)
	}
	return parseAsRowDataText(raw, f)
}

func parseAsRowDataText(raw []byte, f []*Field) ([]interface{}, error) {
	data := make([]interface{}, len(f))

	var err error
	var v []byte
	var isNull, isUnsigned bool
	var pos int
	var n int

	for i := range f {
		v, isNull, n, err = LenencStrToString(raw[pos:])
		if err != nil {
			return nil, err
		}

		pos += n

		if isNull {
			data[i] = nil
		} else {
			isUnsigned = (f[i].Flags&UNSIGNED_FLAG > 0)
			switch f[i].ColumnType {
			case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_LONG, MYSQL_TYPE_INT24,
				MYSQL_TYPE_LONGLONG, MYSQL_TYPE_YEAR:
				if isUnsigned {
					data[i], err = strconv.ParseUint(string(v), 10, 64)
				} else {
					data[i], err = strconv.ParseInt(string(v), 10, 64)
				}
			case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE, MYSQL_TYPE_NEWDECIMAL:
				data[i], err = strconv.ParseFloat(string(v), 64)
			case MYSQL_TYPE_VARCHAR, MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING:
				data[i] = string(v)
			default:
				data[i] = v
			}

			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

func parseAsRowDataBinary(raw []byte, f []*Field) ([]interface{}, error) {
	data := make([]interface{}, len(f))

	if raw[0] != OK_HEADER {
		return nil, errors.ErrMalformPacket
	}

	pos := 1 + ((len(f) + 7 + 2) >> 3)

	nullBitmap := raw[1:pos]

	var isUnsigned bool
	var isNull bool
	var n int
	var err error
	var v []byte
	for i := range data {
		if nullBitmap[(i+2)/8]&(1<<(uint(i+2)%8)) > 0 {
			data[i] = nil
			continue
		}

		isUnsigned = f[i].Flags&UNSIGNED_FLAG > 0

		switch f[i].ColumnType {
		case MYSQL_TYPE_NULL:
			data[i] = nil
			continue

		case MYSQL_TYPE_TINY:
			if isUnsigned {
				data[i] = uint64(raw[pos])
			} else {
				data[i] = int64(raw[pos])
			}
			pos++
			continue

		case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
			if isUnsigned {
				data[i] = uint64(binary.LittleEndian.Uint16(raw[pos : pos+2]))
			} else {
				data[i] = int64((binary.LittleEndian.Uint16(raw[pos : pos+2])))
			}
			pos += 2
			continue

		case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
			if isUnsigned {
				data[i] = uint64(binary.LittleEndian.Uint32(raw[pos : pos+4]))
			} else {
				data[i] = int64(binary.LittleEndian.Uint32(raw[pos : pos+4]))
			}
			pos += 4
			continue

		case MYSQL_TYPE_LONGLONG:
			if isUnsigned {
				data[i] = binary.LittleEndian.Uint64(raw[pos : pos+8])
			} else {
				data[i] = int64(binary.LittleEndian.Uint64(raw[pos : pos+8]))
			}
			pos += 8
			continue

		case MYSQL_TYPE_FLOAT:
			data[i] = float64(math.Float32frombits(binary.LittleEndian.Uint32(raw[pos : pos+4])))
			pos += 4
			continue

		case MYSQL_TYPE_DOUBLE:
			data[i] = math.Float64frombits(binary.LittleEndian.Uint64(raw[pos : pos+8]))
			pos += 8
			continue

		case MYSQL_TYPE_DECIMAL, MYSQL_TYPE_NEWDECIMAL, MYSQL_TYPE_VARCHAR,
			MYSQL_TYPE_BIT, MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_TINY_BLOB,
			MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_LONG_BLOB, MYSQL_TYPE_BLOB,
			MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING, MYSQL_TYPE_GEOMETRY:
			v, isNull, n, err = LenencStrToString(raw[pos:])
			pos += n
			if err != nil {
				return nil, err
			}

			if !isNull {
				data[i] = v
				continue
			} else {
				data[i] = nil
				continue
			}
		case MYSQL_TYPE_DATE, MYSQL_TYPE_NEWDATE:
			var num uint64
			num, isNull, n = LenencIntToNumber(raw[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryDate(int(num), raw[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		case MYSQL_TYPE_TIMESTAMP, MYSQL_TYPE_DATETIME:
			var num uint64
			num, isNull, n = LenencIntToNumber(raw[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryDateTime(int(num), raw[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		case MYSQL_TYPE_TIME:
			var num uint64
			num, isNull, n = LenencIntToNumber(raw[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryTime(int(num), raw[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("Stmt Unknown FieldType %d %s", f[i].ColumnType, f[i].Name)
		}
	}

	return data, nil
}
