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
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"runtime"
	"time"
	"unicode/utf8"
)

var encodeMap [256]byte
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
	for i := range encodeRef {
		encodeRef[i] = DONTESCAPE
	}
	for i := range encodeRef {
		if to, ok := encodeRef[byte(i)]; ok {
			encodeRef[byte(i)] = to
		}
	}
}

// CurrentStack format a stack trace of the calling goroutine.
func CurrentStack() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return string(buf[0:n])
}

// CalcPassword calc password hash.
func CalcPassword(scramble, password []byte) []byte {
	if len(password) == 0 {
		return nil
	}

	// stage1 = SHA1(password)
	crypt := sha1.New()
	crypt.Write(password)
	stage1 := crypt.Sum(nil)

	// stage1Hash = SHA1(stage1)
	crypt.Reset()
	crypt.Write(stage1)
	stage1Hash := crypt.Sum(nil)

	// scrambleHash = SHA1(scramble + stage1Hash)
	crypt.Reset()
	crypt.Write(scramble)
	crypt.Write(stage1Hash)
	scrambleHash := crypt.Sum(nil)

	// token = scrambleHash XOR stage1
	for i := range scrambleHash {
		scrambleHash[i] ^= stage1[i]
	}
	return scrambleHash
}

// RandomBuf random ascii array with specific size.
func RandomBuf(size int) ([]byte, error) {
	buf := make([]byte, size)
	rand.Seed(time.Now().UTC().UnixNano())
	min, max := 30, 127
	for i := 0; i < size; i++ {
		buf[i] = byte(min + rand.Intn(max-min))
	}
	return buf, nil
}

// LenencIntToNumber convert lenenc_int to number.
func LenencIntToNumber(b []byte) (num uint64, isNull bool, n int) {
	switch b[0] {

	// 251: NULL
	case 0xfb:
		n = 1
		isNull = true
		return

	// 252: value of following 2
	case 0xfc:
		num = uint64(b[1]) | uint64(b[2])<<8
		n = 3
		return

	// 253: value of following 3
	case 0xfd:
		num = uint64(b[1]) | uint64(b[2])<<8 | uint64(b[3])<<16
		n = 4
		return

	// 254: value of following 8
	case 0xfe:
		num = uint64(b[1]) | uint64(b[2])<<8 | uint64(b[3])<<16 |
			uint64(b[4])<<24 | uint64(b[5])<<32 | uint64(b[6])<<40 |
			uint64(b[7])<<48 | uint64(b[8])<<56
		n = 9
		return
	}

	// 0-250: value of first byte
	num = uint64(b[0])
	n = 1
	return
}

// NumberToLenencInt convert number to lenenc_int.
func NumberToLenencInt(n uint64) []byte {
	switch {
	case n <= 250:
		return []byte{byte(n)}

	case n <= 0xffff:
		return []byte{0xfc, byte(n), byte(n >> 8)}

	case n <= 0xffffff:
		return []byte{0xfd, byte(n), byte(n >> 8), byte(n >> 16)}

	case n <= 0xffffffffffffffff:
		return []byte{0xfe, byte(n), byte(n >> 8), byte(n >> 16), byte(n >> 24),
			byte(n >> 32), byte(n >> 40), byte(n >> 48), byte(n >> 56)}
	}
	return nil
}

// LenencStrToString convert lenenc_str to string.
func LenencStrToString(b []byte) ([]byte, bool, int, error) {
	// Get length
	num, isNull, n := LenencIntToNumber(b)
	if num < 1 {
		return nil, isNull, n, nil
	}

	n += int(num)

	// Check data length
	if len(b) >= n {
		return b[n-int(num) : n], false, n, nil
	}
	return nil, false, n, io.EOF
}

// LengthOfLenencStr length of lenenc_str.
func LengthOfLenencStr(b []byte) (int, error) {
	// Get length
	num, _, n := LenencIntToNumber(b)
	if num < 1 {
		return n, nil
	}

	n += int(num)

	// Check data length
	if len(b) >= n {
		return n, nil
	}
	return n, io.EOF
}

// StringToLenencStr convert string to lenen_str.
func StringToLenencStr(b []byte) []byte {
	data := make([]byte, 0, len(b)+9)
	data = append(data, NumberToLenencInt(uint64(len(b)))...)
	data = append(data, b...)
	return data
}

// Uint16ToBytes convert uint16 to byte array.
func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

// Uint32ToBytes convert uint32 to byte array.
func Uint32ToBytes(n uint32) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
	}
}

// Uint64ToBytes convert uint64 to byte array.
func Uint64ToBytes(n uint64) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
		byte(n >> 16),
		byte(n >> 24),
		byte(n >> 32),
		byte(n >> 40),
		byte(n >> 48),
		byte(n >> 56),
	}
}

// FormatBinaryDate format binary date.
func FormatBinaryDate(n int, data []byte) ([]byte, error) {
	switch n {
	case 0:
		return []byte("0000-00-00"), nil
	case 4:
		return []byte(fmt.Sprintf("%04d-%02d-%02d",
			binary.LittleEndian.Uint16(data[:2]),
			data[2],
			data[3])), nil
	default:
		return nil, fmt.Errorf("invalid date packet length %d", n)
	}
}

// FormatBinaryDateTime format binary datetime.
func FormatBinaryDateTime(n int, data []byte) ([]byte, error) {
	switch n {
	case 0:
		return []byte("0000-00-00 00:00:00"), nil
	case 4:
		return []byte(fmt.Sprintf("%04d-%02d-%02d 00:00:00",
			binary.LittleEndian.Uint16(data[:2]),
			data[2],
			data[3])), nil
	case 7:
		return []byte(fmt.Sprintf(
			"%04d-%02d-%02d %02d:%02d:%02d",
			binary.LittleEndian.Uint16(data[:2]),
			data[2],
			data[3],
			data[4],
			data[5],
			data[6])), nil
	case 11:
		return []byte(fmt.Sprintf(
			"%04d-%02d-%02d %02d:%02d:%02d.%06d",
			binary.LittleEndian.Uint16(data[:2]),
			data[2],
			data[3],
			data[4],
			data[5],
			data[6],
			binary.LittleEndian.Uint32(data[7:11]))), nil
	default:
		return nil, fmt.Errorf("invalid datetime packet length %d", n)
	}
}

// FormatBinaryTime format binary time
func FormatBinaryTime(n int, data []byte) ([]byte, error) {
	if n == 0 {
		return []byte("00:00:00"), nil
	}

	var sign byte
	if data[0] == 1 {
		sign = byte('-')
	}

	switch n {
	case 8:
		return []byte(fmt.Sprintf(
			"%c%02d:%02d:%02d",
			sign,
			uint16(data[1])*24+uint16(data[5]),
			data[6],
			data[7],
		)), nil
	case 12:
		return []byte(fmt.Sprintf(
			"%c%02d:%02d:%02d.%06d",
			sign,
			uint16(data[1])*24+uint16(data[5]),
			data[6],
			data[7],
			binary.LittleEndian.Uint32(data[8:12]),
		)), nil
	default:
		return nil, fmt.Errorf("invalid time packet length %d", n)
	}
}

// Escape char in string.
func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))

	for i, w := 0, 0; i < len(sql); i += w {
		runeValue, width := utf8.DecodeRuneInString(sql[i:])
		if c := encodeRef[byte(runeValue)]; c == DONTESCAPE {
			dest = append(dest, sql[i:i+width]...)
		} else {
			dest = append(dest, '\\', c)
		}
		w = width
	}

	return string(dest)
}

// PrintPacketData print packet data. just for debug.
func PrintPacketData(data []byte) {
	str := ""
	for _, byt := range data {
		if byt < 16 {
			str += fmt.Sprintf(" 0%x", byt)

		} else {
			str += fmt.Sprintf(" %x", byt)
		}
	}
	fmt.Printf("Packet Data:%s\r\n", str)
}
