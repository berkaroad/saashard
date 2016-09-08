package mysql

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/berkaroad/saashard/utils/golog"
)

// WriteInitialHandshake write initial handshake
// connectionID:
// salt: mysql.mysql.RandomBuf(20)
// collationID:
// capability:
// status:
func (p *PacketIO) WriteInitialHandshake(connectionID uint32, salt []byte, collationID CollationID, capability uint32, status uint16) error {
	data := make([]byte, 4, 128)

	//min version 10
	data = append(data, 10)

	//server version[00]
	data = append(data, ServerVersion...)
	data = append(data, 0)

	//connection id
	data = append(data, byte(connectionID), byte(connectionID>>8), byte(connectionID>>16), byte(connectionID>>24))

	//auth-plugin-data-part-1
	data = append(data, salt[0:8]...)

	//filter [00]
	data = append(data, 0)

	//capability flag lower 2 bytes, using default capability here
	data = append(data, byte(capability), byte(capability>>8))
	// fmt.Printf("Server write capability part1 to client %d, %d \r\n", byte(capability), byte(capability>>8))

	//charset, utf-8 default
	data = append(data, uint8(collationID))

	//status
	data = append(data, byte(status), byte(status>>8))

	//below 13 byte may not be used
	//capability flag upper 2 bytes, using default capability here
	data = append(data, byte(capability>>16), byte(capability>>24))
	// fmt.Printf("Server write capability part2 to client %d, %d \r\n", byte(capability>>16), byte(capability>>24))

	//filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x15)

	//reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	//auth-plugin-data-part-2
	data = append(data, salt[8:]...)

	//filter [00]
	data = append(data, 0)

	return p.WritePacket(data)
}

// ReadInitialHandshake read initial handshake.
// salt:
func (p *PacketIO) ReadInitialHandshake(salt *[]byte) (capability uint32, status uint16, collationID CollationID, err error) {
	var data []byte
	data, err = p.ReadPacket()
	if err != nil {
		return
	}

	if data[0] == ERR_HEADER {
		err = errors.New("read initial handshake error")
		return
	}

	if data[0] < MinProtocolVersion {
		err = fmt.Errorf("invalid protocol version %d, must >= 10", data[0])
		return
	}

	//skip mysql version and connection id
	//mysql version end with 0x00
	//connection id length is 4
	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1 + 4

	*salt = append(*salt, data[pos:pos+8]...)

	//skip filter
	pos += 8 + 1

	//capability lower 2 bytes
	capability = uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))
	// fmt.Printf("Read capability part1 from server %d, %d \r\n", data[pos:pos+1], data[pos+1:pos+2])
	pos += 2

	if len(data) > pos {
		//skip server charset
		collationID = CollationID(data[pos])
		pos++

		status = binary.LittleEndian.Uint16(data[pos : pos+2])
		pos += 2

		capability = uint32(binary.LittleEndian.Uint16(data[pos:pos+2]))<<16 | capability
		// fmt.Printf("Read capability part2 from server %d, %d \r\n", data[pos:pos+1], data[pos+1:pos+2])

		pos += 2

		//skip auth data len or [00]
		//skip reserved (all [00])
		pos += 10 + 1

		// The documentation is ambiguous about the length.
		// The official Python library uses the fixed length 12
		// mysql-proxy also use 12
		// which is not documented but seems to work.
		*salt = append(*salt, data[pos:pos+12]...)
	}
	return
}

// WriteAuthHandshake write auth handshake
func (p *PacketIO) WriteAuthHandshake(capability *uint32, user, password, db string, salt []byte, collationID CollationID) error {
	// Adjust client capability flags based on server support
	*capability &= DEFAULT_CAPABILITY

	//packet length
	//capbility 4
	//max-packet size 4
	//charset 1
	//reserved all[0] 23
	length := 4 + 4 + 1 + 23

	//username
	length += len(user) + 1

	//we only support secure connection
	auth := CalcPassword(salt, []byte(password))

	length += 1 + len(auth)

	if len(db) > 0 {
		*capability |= CLIENT_CONNECT_WITH_DB

		length += len(db) + 1
	}

	data := make([]byte, length+4)

	//capability [32 bit]
	data[4] = byte(*capability)
	data[5] = byte(*capability >> 8)
	data[6] = byte(*capability >> 16)
	data[7] = byte(*capability >> 24)

	// fmt.Printf("Client write capability to server %d, %d, %d, %d \r\n", byte(*capability), byte(*capability>>8), byte(*capability>>16), byte(*capability>>24))

	//MaxPacketSize [32 bit] (none)
	//data[8] = 0x00
	//data[9] = 0x00
	//data[10] = 0x00
	//data[11] = 0x00

	//Charset [1 byte]
	data[12] = byte(collationID)

	//Filler [23 bytes] (all 0x00)
	pos := 13 + 23

	//User [null terminated string]
	if len(user) > 0 {
		pos += copy(data[pos:], user)
	}
	//data[pos] = 0x00
	pos++

	// auth [length encoded integer]
	data[pos] = byte(len(auth))
	pos += 1 + copy(data[pos+1:], auth)

	// db [null terminated string]
	if len(db) > 0 {
		pos += copy(data[pos:], db)
		//data[pos] = 0x00
	}

	return p.WritePacket(data)
}

// ReadHandshakeResponse read handshake response
func (p *PacketIO) ReadHandshakeResponse(getDefaultSchemaByUser func(user string) (string, error), remoteAddr string, salt []byte, getCredentialsConfigBySchema func(db string) (user, password string, err error)) (capability uint32, collationID CollationID, user, db string, err error) {
	var data []byte
	data, err = p.ReadPacket()

	if err != nil {
		return
	}

	pos := 0

	//capability
	capability = binary.LittleEndian.Uint32(data[:4])
	capability = capability & DEFAULT_CAPABILITY
	pos += 4

	//skip max packet size
	pos += 4

	//charset, skip, if you want to use another charset, use set names
	collationID = CollationID(data[pos])
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])

	pos += len(user) + 1

	//auth length and auth
	authLen := int(data[pos])
	pos++

	auth := data[pos : pos+authLen]
	pos += authLen

	if capability&CLIENT_CONNECT_WITH_DB > 0 {
		if len(data[pos:]) == 0 {
			//if connect with non-name database, use default db
			db, err = getDefaultSchemaByUser(user)
		}

		db = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(db) + 1

	} else {
		//if connect without database, use default db
		db, err = getDefaultSchemaByUser(user)
	}
	if err != nil {
		return
	}
	db = strings.ToLower(db)

	var configUser, configPassword string
	configUser, configPassword, err = getCredentialsConfigBySchema(db)
	if err != nil {
		return
	}
	checkAuth := CalcPassword(salt, []byte(configPassword))
	if user != configUser || !bytes.Equal(auth, checkAuth) {
		golog.Error("ClientConn", "readHandshakeResponse", "error", 0,
			"auth", auth,
			"checkAuth", checkAuth,
			"client_user", user,
			"config_set_user", configUser,
			"passworld", configPassword)
		err = NewDefaultError(ER_ACCESS_DENIED_ERROR, user, remoteAddr, "Yes")
		return
	}
	return
}
