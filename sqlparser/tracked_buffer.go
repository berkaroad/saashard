// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

package sqlparser

import (
	"bytes"
	"fmt"
	"reflect"
)

// TrackedBuffer is used to rebuild a query from the ast.
// bindLocations keeps track of locations in the buffer that
// use bind variables for efficient future substitutions.
// nodeFormatter is the formatting function the buffer will
// use to format a node. By default(nil), it's FormatNode.
// But you can supply a different formatting function if you
// want to generate a query that's different from the default.
type TrackedBuffer struct {
	*bytes.Buffer
	nodeFormatter func(buf *TrackedBuffer, node SQLNode)
}

// NewTrackedBuffer create TrackedBuffer
func NewTrackedBuffer(nodeFormatter func(buf *TrackedBuffer, node SQLNode)) *TrackedBuffer {
	buf := &TrackedBuffer{
		Buffer:        bytes.NewBuffer(make([]byte, 0, 128)),
		nodeFormatter: nodeFormatter,
	}
	return buf
}

// Fprintf mimics fmt.Fprintf, but limited to Node(%v), Node.Value(%s) and string(%s).
// It also allows a %a for a value argument, in which case it adds tracking info for
// future substitutions.
func (buf *TrackedBuffer) Fprintf(format string, values ...interface{}) {
	end := len(format)
	fieldnum := 0
	for i := 0; i < end; {
		lasti := i
		for i < end && format[i] != '%' {
			i++
		}
		if i > lasti {
			buf.WriteString(format[lasti:i])
		}
		if i >= end {
			break
		}
		i++ // '%'
		switch format[i] {
		case 'c':
			switch v := values[fieldnum].(type) {
			case byte:
				buf.WriteByte(v)
			case rune:
				buf.WriteRune(v)
			default:
				panic(fmt.Sprintf("unexpected type %T", v))
			}

		case 's':
			switch v := values[fieldnum].(type) {
			case []byte:
				buf.Write(v)
			case string:
				buf.WriteString(v)
			default:
				panic(fmt.Sprintf("unexpected type %T", v))
			}
		case 'v':
			currentVal := values[fieldnum]
			if currentVal != nil && !reflect.ValueOf(currentVal).IsNil() {
				node := currentVal.(SQLNode)
				if buf.nodeFormatter == nil {
					node.Format(buf)
				} else {
					buf.nodeFormatter(buf, node)
				}
			}
		case 'a':
			buf.WriteArg(values[fieldnum].(string))
		default:
			panic("unexpected")
		}
		fieldnum++
		i++
	}
}

// WriteArg write argument.
func (buf *TrackedBuffer) WriteArg(arg string) {
	buf.WriteString(arg)
}
