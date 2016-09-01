// Package sqltypes implements interfaces and types that represent SQL values.
package sqltypes

import (
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"
)

var (
	// DONTESCAPE dont escape char.
	DONTESCAPE = byte(255)
	nullstr    = []byte("null")
)

// SQLEncodeMap specifies how to escape binary data with '\'.
// Complies to http://dev.mysql.com/doc/refman/5.1/en/string-syntax.html
var SQLEncodeMap [256]byte

// SQLDecodeMap is the reverse of SQLEncodeMap
var SQLDecodeMap [256]byte

var encodeRef = map[byte]byte{
	'\x00': '0',
	'\'':   '\'',
	'"':    '"',
	'\b':   'b',
	'\n':   'n',
	'\r':   'r',
	'\t':   't',
	26:     'Z', // ctl-Z
	'\\':   '\\',
}

func init() {
	for i := range SQLEncodeMap {
		SQLEncodeMap[i] = DONTESCAPE
		SQLDecodeMap[i] = DONTESCAPE
	}
	for i := range SQLEncodeMap {
		if to, ok := encodeRef[byte(i)]; ok {
			SQLEncodeMap[byte(i)] = to
			SQLDecodeMap[to] = byte(i)
		}
	}
	gob.Register(Numeric(nil))
	gob.Register(Fractional(nil))
	gob.Register(String(nil))
}

// BinWriter interface is used for encoding values.
// Types like bytes.Buffer conform to this interface.
// We expect the writer objects to be in-memory buffers.
// So, we don't expect the write operations to fail.
type BinWriter interface {
	Write([]byte) (int, error)
	WriteByte(byte) error
}

// InnerValue defines methods that need to be supported by all non-null value types.
type InnerValue interface {
	raw() []byte
	encodeSQL(BinWriter)
	encodeASCII(BinWriter)
}

// Value can store any SQL value. NULL is stored as nil.
type Value struct {
	Inner InnerValue
}

// MakeNumeric makes a Numeric from a []byte without validation.
func MakeNumeric(b []byte) Value {
	return Value{Numeric(b)}
}

// MakeFractional makes a Fractional value from a []byte without validation.
func MakeFractional(b []byte) Value {
	return Value{Fractional(b)}
}

// MakeString makes a String value from a []byte.
func MakeString(b []byte) Value {
	return Value{String(b)}
}

// BuildValue build value.
func BuildValue(goval interface{}) (v Value, err error) {
	switch bindVal := goval.(type) {
	case nil:
		// no op
	case int:
		v = Value{Numeric(strconv.AppendInt(nil, int64(bindVal), 10))}
	case int32:
		v = Value{Numeric(strconv.AppendInt(nil, int64(bindVal), 10))}
	case int64:
		v = Value{Numeric(strconv.AppendInt(nil, int64(bindVal), 10))}
	case uint:
		v = Value{Numeric(strconv.AppendUint(nil, uint64(bindVal), 10))}
	case uint32:
		v = Value{Numeric(strconv.AppendUint(nil, uint64(bindVal), 10))}
	case uint64:
		v = Value{Numeric(strconv.AppendUint(nil, uint64(bindVal), 10))}
	case float64:
		v = Value{Fractional(strconv.AppendFloat(nil, bindVal, 'f', -1, 64))}
	case string:
		v = Value{String([]byte(bindVal))}
	case []byte:
		v = Value{String(bindVal)}
	case time.Time:
		v = Value{String([]byte(bindVal.Format("'2006-01-02 15:04:05'")))}
	case Numeric, Fractional, String:
		v = Value{bindVal.(InnerValue)}
	case Value:
		v = bindVal
	default:
		return Value{}, fmt.Errorf("unsupported bind variable type %T: %v", goval, goval)
	}
	return v, nil
}

// BuildNumeric builds a Numeric type that represents any whole number.
// It normalizes the representation to ensure 1:1 mapping between the
// number and its representation.
func BuildNumeric(val string) (n Value, err error) {
	if val[0] == '-' || val[0] == '+' {
		signed, err := strconv.ParseInt(val, 0, 64)
		if err != nil {
			return Value{}, err
		}
		n = Value{Numeric(strconv.AppendInt(nil, signed, 10))}
	} else {
		unsigned, err := strconv.ParseUint(val, 0, 64)
		if err != nil {
			return Value{}, err
		}
		n = Value{Numeric(strconv.AppendUint(nil, unsigned, 10))}
	}
	return n, nil
}

// Raw returns the raw bytes. All types are currently implemented as []byte.
func (v Value) Raw() []byte {
	if v.Inner == nil {
		return nil
	}
	return v.Inner.raw()
}

// String returns the raw value as a string
func (v Value) String() string {
	if v.Inner == nil {
		return ""
	}
	return string(v.Inner.raw())
}

// ParseInt64 will parse a Numeric value into an int64
func (v Value) ParseInt64() (val int64, err error) {
	if v.Inner == nil {
		return 0, fmt.Errorf("value is null")
	}
	n, ok := v.Inner.(Numeric)
	if !ok {
		return 0, fmt.Errorf("value is not Numeric")
	}
	return strconv.ParseInt(string(n.raw()), 10, 64)
}

// ParseUint64 will parse a Numeric value into a uint64
func (v Value) ParseUint64() (val uint64, err error) {
	if v.Inner == nil {
		return 0, fmt.Errorf("value is null")
	}
	n, ok := v.Inner.(Numeric)
	if !ok {
		return 0, fmt.Errorf("value is not Numeric")
	}
	return strconv.ParseUint(string(n.raw()), 10, 64)
}

// EncodeSQL encode sql.
func (v Value) EncodeSQL(b BinWriter) {
	if v.Inner == nil {
		if _, err := b.Write(nullstr); err != nil {
			panic(err)
		}
	} else {
		v.Inner.encodeSQL(b)
	}
}

// EncodeASCII encodes the value using 7-bit clean ascii bytes.
func (v Value) EncodeASCII(b BinWriter) {
	if v.Inner == nil {
		if _, err := b.Write(nullstr); err != nil {
			panic(err)
		}
	} else {
		v.Inner.encodeASCII(b)
	}
}

// IsNull check is null.
func (v Value) IsNull() bool {
	return v.Inner == nil
}

// IsNumeric check is numeric.
func (v Value) IsNumeric() (ok bool) {
	if v.Inner != nil {
		_, ok = v.Inner.(Numeric)
	}
	return ok
}

// IsFractional check is fractional
func (v Value) IsFractional() (ok bool) {
	if v.Inner != nil {
		_, ok = v.Inner.(Fractional)
	}
	return ok
}

// IsString check is string.
func (v Value) IsString() (ok bool) {
	if v.Inner != nil {
		_, ok = v.Inner.(String)
	}
	return ok
}

// Numeric represents non-fractional SQL number.
type Numeric []byte

func (n Numeric) raw() []byte {
	return []byte(n)
}

func (n Numeric) encodeSQL(b BinWriter) {
	if _, err := b.Write(n.raw()); err != nil {
		panic(err)
	}
}

func (n Numeric) encodeASCII(b BinWriter) {
	if _, err := b.Write(n.raw()); err != nil {
		panic(err)
	}
}

// Fractional represents fractional types like float and decimal
// It's functionally equivalent to Numeric other than how it's constructed
type Fractional []byte

func (f Fractional) raw() []byte {
	return []byte(f)
}

func (f Fractional) encodeSQL(b BinWriter) {
	if _, err := b.Write(f.raw()); err != nil {
		panic(err)
	}
}

func (f Fractional) encodeASCII(b BinWriter) {
	if _, err := b.Write(f.raw()); err != nil {
		panic(err)
	}
}

// String represents any SQL type that needs to be represented using quotes.
type String []byte

func (s String) raw() []byte {
	return []byte(s)
}

func (s String) encodeSQL(b BinWriter) {
	writebyte(b, '\'')
	for _, ch := range s.raw() {
		if encodedChar := SQLEncodeMap[ch]; encodedChar == DONTESCAPE {
			writebyte(b, ch)
		} else {
			writebyte(b, '\\')
			writebyte(b, encodedChar)
		}
	}
	writebyte(b, '\'')
}

func (s String) encodeASCII(b BinWriter) {
	writebyte(b, '\'')
	encoder := base64.NewEncoder(base64.StdEncoding, b)
	encoder.Write(s.raw())
	encoder.Close()
	writebyte(b, '\'')
}

func writebyte(b BinWriter, c byte) {
	if err := b.WriteByte(c); err != nil {
		panic(err)
	}
}
