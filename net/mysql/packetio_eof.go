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

// WriteEOF is to write EOF packet
func (p *PacketIO) WriteEOF(capability uint32, status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, EOF_HEADER)
	if capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return p.WritePacket(data)
}

// WriteEOFBatch is to write EOF packet in batch.
func (p *PacketIO) WriteEOFBatch(total []byte, capability uint32, status uint16, direct bool) ([]byte, error) {
	data := make([]byte, 4, 9)

	data = append(data, EOF_HEADER)
	if capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}
	return p.WritePacketBatch(total, data, direct)
}

func (p *PacketIO) isEOFPacket(data []byte) bool {
	return data[0] == EOF_HEADER && len(data) <= 5
}

func (p *PacketIO) readUntilEOF() (datas [][]byte, err error) {
	datas = make([][]byte, 0, 2)
	var data []byte
	for {
		data, err = p.ReadPacket()

		if err != nil {
			return
		}

		// EOF Packet
		if p.isEOFPacket(data) {
			return
		}
		datas = append(datas, data)
	}
}
