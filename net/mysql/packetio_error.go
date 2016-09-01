package mysql

import (
	"encoding/binary"

	"github.com/berkaroad/saashard/errors"
)

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

func (p *PacketIO) handleErrorPacket(capability uint32, data []byte) error {
	if data[0] != ERR_HEADER {
		return nil
	}

	e := new(errors.SqlError)
	var pos = 1

	e.Code = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	if capability&CLIENT_PROTOCOL_41 > 0 {
		//skip '#'
		pos++
		e.State = string(data[pos : pos+5])
		pos += 5
	}

	e.Message = string(data[pos:])
	return e
}
