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
