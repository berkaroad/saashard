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

// WriteOK is to write OK packet.
func (p *PacketIO) WriteOK(capability uint32, status uint16, r *Result) error {
	if r == nil {
		r = &Result{Status: status}
	}
	data := make([]byte, 4, 32)

	data = append(data, OK_HEADER)

	data = append(data, NumberToLenencInt(r.AffectedRows)...)
	data = append(data, NumberToLenencInt(r.InsertId)...)

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
	data = append(data, NumberToLenencInt(r.InsertId)...)

	if capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	return p.WritePacketBatch(total, data, direct)
}

// WriteError is to write Error packet.
func (p *PacketIO) WriteError(capability uint32, e error) error {
	var m *errors.SqlError
	var ok bool
	if m, ok = e.(*errors.SqlError); !ok {
		m = NewError(ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if capability&CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	return p.WritePacket(data)
}

// WriteResultSet is to write Result Set packet.
func (p *PacketIO) WriteResultSet(capability uint32, status uint16, r *Result) error {
	total := make([]byte, 0, 1024)
	var err error
	total, err = p.writeResultSetHeader(total, r)
	if err != nil {
		return err
	}
	for _, f := range r.Fields {
		total, err = p.writeResultSetField(total, f)
		if err != nil {
			return err
		}
	}
	if capability&CLIENT_DEPRECATE_EOF == 0 {
		total, err = p.WriteEOFBatch(total, capability, status, false)
		if err != nil {
			return err
		}
	}
	if r.RowDatas != nil && len(r.RowDatas) > 0 {
		for _, r := range r.RowDatas {
			total, err = p.writeResultSetRowData(total, r)
			if err != nil {
				return err
			}
		}
	}
	if capability&CLIENT_DEPRECATE_EOF > 0 {
		total, err = p.WriteOKBatch(total, capability, status, r, true)
		if err != nil {
			return err
		}
	} else {
		total, err = p.WriteEOFBatch(total, capability, status, true)
		if err != nil {
			return err
		}
	}
	return nil
}

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

func (p *PacketIO) writeResultSetHeader(total []byte, r *Result) ([]byte, error) {
	data := make([]byte, 4, 5)
	data = append(data, byte(r.ColumnNumber()))
	return p.WritePacketBatch(total, data, false)
}

func (p *PacketIO) writeResultSetField(total []byte, f *Field) ([]byte, error) {
	f.Data = f.Dump()
	data := make([]byte, 4, 4+len(f.Data))
	data = append(data, f.Data...)
	return p.WritePacketBatch(total, data, false)
}

func (p *PacketIO) writeResultSetRowData(total []byte, r *RowData) ([]byte, error) {
	r.Data = r.Dump()
	data := make([]byte, 4, 4+len(r.Data))
	data = append(data, r.Data...)
	return p.WritePacketBatch(total, data, false)
}
