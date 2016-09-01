package mysql

import (
	"encoding/binary"

	"github.com/berkaroad/saashard/errors"
)

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
	if r.Rows != nil && len(r.Rows) > 0 {
		for _, row := range r.Rows {
			total, err = p.writeResultSetRowData(total, row)
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

// ReadResultSet read result set.
func (p *PacketIO) ReadResultSet(capability uint32, status *uint16, binary bool) (*Result, error) {
	data, err := p.ReadPacket()
	if err != nil {
		return nil, err
	}

	if data[0] == OK_HEADER {
		return p.handleOKPacket(capability, status, data)
	} else if data[0] == ERR_HEADER {
		return nil, p.handleErrorPacket(capability, data)
	} else if data[0] == LocalInFile_HEADER {
		return nil, errors.ErrMalformPacket
	}

	return p.handleResultsetPacket(capability, status, data, binary)
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

func (p *PacketIO) writeResultSetRowData(total []byte, r *Row) ([]byte, error) {
	rData := r.Dump()
	data := make([]byte, 4, 4+len(rData))
	data = append(data, rData...)
	return p.WritePacketBatch(total, data, false)
}

func (p *PacketIO) handleResultsetPacket(capability uint32, status *uint16, data []byte, binary bool) (*Result, error) {
	result := &Result{
		Status:       0,
		InsertId:     0,
		AffectedRows: 0,

		Resultset: &Resultset{},
	}

	// column count
	count, _, n := LenencIntToNumber(data)

	if n-len(data) != 0 {
		return nil, errors.ErrMalformPacket
	}

	result.Fields = make([]*Field, count)
	result.FieldNames = make(map[string]int, count)

	if err := p.handleResultColumns(capability, status, result); err != nil {
		return nil, err
	}

	if err := p.handleResultRows(capability, status, result, binary); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PacketIO) handleResultColumns(capability uint32, status *uint16, result *Result) (err error) {
	var i = 0
	var data []byte

	for {
		data, err = p.ReadPacket()
		if err != nil {
			return
		}

		// EOF Packet
		if p.isEOFPacket(data) {
			if capability&CLIENT_PROTOCOL_41 > 0 {
				//result.Warnings = binary.LittleEndian.Uint16(data[1:])
				//todo add strict_mode, warning will be treat as error
				result.Status = binary.LittleEndian.Uint16(data[3:])
				*status = result.Status
			}

			if i != len(result.Fields) {
				err = errors.ErrMalformPacket
			}

			return
		}

		result.Fields[i], err = FieldData(data).Parse()
		if err != nil {
			return
		}

		result.FieldNames[string(result.Fields[i].Name)] = i

		i++
	}
}

func (p *PacketIO) handleResultRows(capability uint32, status *uint16, result *Result, isBinary bool) (err error) {
	var data []byte

	for {
		data, err = p.ReadPacket()

		if err != nil {
			return
		}

		// EOF Packet
		if p.isEOFPacket(data) {
			if capability&CLIENT_PROTOCOL_41 > 0 {
				//result.Warnings = binary.LittleEndian.Uint16(data[1:])
				//todo add strict_mode, warning will be treat as error
				result.Status = binary.LittleEndian.Uint16(data[3:])
				*status = result.Status
			}

			break
		}
		var row *Row
		row, err = RowData(data).Parse(isBinary, result.Fields)
		if err != nil {
			return err
		}
		result.Rows = append(result.Rows, row)
	}

	result.Values = make([][]interface{}, len(result.Rows))

	for i := range result.Values {
		result.Values[i] = result.Rows[i].fieldValues

		if err != nil {
			return err
		}
	}

	return nil
}
