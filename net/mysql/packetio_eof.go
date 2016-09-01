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

func (p *PacketIO) readUntilEOF() (err error) {
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
	}
}
