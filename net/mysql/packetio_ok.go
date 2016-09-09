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
	"errors"
)

// WriteOK is to write OK packet.
func (p *PacketIO) WriteOK(capability uint32, status uint16, r *Result) error {
	if r == nil {
		r = &Result{Status: status}
	}
	data := make([]byte, 4, 32)

	data = append(data, OK_HEADER)

	data = append(data, NumberToLenencInt(r.AffectedRows)...)
	data = append(data, NumberToLenencInt(r.InsertID)...)

	if capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	return p.WritePacket(data)
}

// WriteOKBatch is to write OK packet in batch.
func (p *PacketIO) WriteOKBatch(total []byte, capability uint32, status uint16, r *Result, direct bool) ([]byte, error) {
	if r == nil {
		r = &Result{Status: status}
	}
	data := make([]byte, 4, 32)

	data = append(data, OK_HEADER)

	data = append(data, NumberToLenencInt(r.AffectedRows)...)
	data = append(data, NumberToLenencInt(r.InsertID)...)

	if capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	return p.WritePacketBatch(total, data, direct)
}

// ReadOK is to read OK package.
func (p *PacketIO) ReadOK(capability uint32, status *uint16) (*Result, error) {
	data, err := p.ReadPacket()
	if err != nil {
		return nil, err
	}

	if data[0] == OK_HEADER {
		return p.handleOKPacket(capability, status, data)
	} else if data[0] == ERR_HEADER {
		return nil, p.handleErrorPacket(capability, data)
	} else {
		return nil, errors.New("invalid ok packet")
	}
}

func (p *PacketIO) handleOKPacket(capability uint32, status *uint16, data []byte) (*Result, error) {
	var n int
	var pos = 1

	r := new(Result)

	r.AffectedRows, _, n = LenencIntToNumber(data[pos:])
	pos += n
	r.InsertID, _, n = LenencIntToNumber(data[pos:])
	pos += n

	if capability&CLIENT_PROTOCOL_41 > 0 {
		r.Status = binary.LittleEndian.Uint16(data[pos:])
		*status = r.Status
		pos += 2

		//todo:strict_mode, check warnings as error
		//Warnings := binary.LittleEndian.Uint16(data[pos:])
		//pos += 2
	} else if capability&CLIENT_TRANSACTIONS > 0 {
		r.Status = binary.LittleEndian.Uint16(data[pos:])
		*status = r.Status
		pos += 2
	}

	//info
	return r, nil
}
