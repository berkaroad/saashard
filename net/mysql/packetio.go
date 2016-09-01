package mysql

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/berkaroad/saashard/errors"
)

const (
	defaultReaderSize = 8 * 1024
)

// PacketIO is a packet transfer on network.
type PacketIO struct {
	rb *bufio.Reader
	wb io.Writer

	Sequence uint8
}

// NewPacketIO is to create PacketIO
func NewPacketIO(conn net.Conn) *PacketIO {
	p := new(PacketIO)

	p.rb = bufio.NewReaderSize(conn, defaultReaderSize)
	p.wb = conn

	p.Sequence = 0

	return p
}

// ReadPacket is to read packet.
func (p *PacketIO) ReadPacket() ([]byte, error) {
	header := []byte{0, 0, 0, 0}

	if _, err := io.ReadFull(p.rb, header); err != nil {
		return nil, errors.ErrBadConn
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		return nil, fmt.Errorf("invalid payload length %d", length)
	}

	sequence := uint8(header[3])

	if sequence != p.Sequence {
		return nil, fmt.Errorf("invalid sequence %d != %d", sequence, p.Sequence)
	}

	p.Sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(p.rb, data); err != nil {
		return nil, errors.ErrBadConn
	}
	if length < MaxPayloadLen {
		return data, nil
	}

	buf, err := p.ReadPacket()
	if err != nil {
		return nil, errors.ErrBadConn
	}
	return append(data, buf...), nil
}

// WritePacket is to write packet.
func (p *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence

		if n, err := p.wb.Write(data[:4+MaxPayloadLen]); err != nil {
			return errors.ErrBadConn
		} else if n != (4 + MaxPayloadLen) {
			return errors.ErrBadConn
		} else {
			p.Sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	if n, err := p.wb.Write(data); err != nil {
		return errors.ErrBadConn
	} else if n != len(data) {
		return errors.ErrBadConn
	} else {
		p.Sequence++
		return nil
	}
}

// WritePacketBatch is to write packet in batch
func (p *PacketIO) WritePacketBatch(total, data []byte, direct bool) ([]byte, error) {
	if data == nil {
		//only flush the buffer
		if direct == true {
			n, err := p.wb.Write(total)
			if err != nil {
				return nil, errors.ErrBadConn
			}
			if n != len(total) {
				return nil, errors.ErrBadConn
			}
		}
		return total, nil
	}

	length := len(data) - 4
	for length >= MaxPayloadLen {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence
		total = append(total, data[:4+MaxPayloadLen]...)

		p.Sequence++
		length -= MaxPayloadLen
		data = data[MaxPayloadLen:]
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	total = append(total, data...)
	p.Sequence++

	if direct {
		if n, err := p.wb.Write(total); err != nil {
			return nil, errors.ErrBadConn
		} else if n != len(total) {
			return nil, errors.ErrBadConn
		}
	}
	return total, nil
}
