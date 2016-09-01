package mysql

const (
	MinProtocolVersion byte   = 10
	MaxPayloadLen      int    = 1<<24 - 1
	TimeFormat         string = "2016-08-25 14:30:00"
	ServerVersion      string = "5.5.5-SaaShard"
	SourceInfo         string = "github.com/berkaroad/saashard"
)

const (
	OK_HEADER          byte = 0x00
	ERR_HEADER         byte = 0xff
	EOF_HEADER         byte = 0xfe
	LocalInFile_HEADER byte = 0xfb
)

const (
	NOT_NULL_FLAG uint16 = 1 << iota
	PRI_KEY_FLAG
	UNIQUE_KEY_FLAG
	MUlTIPLE_KEY_FLAG
	BLOB_FLAG
	UNSIGNED_FLAG
	ZERO_FILL_FLAG
	BINARY_FLAG
	ENUM_FLAG
	AUTO_INCREMENT_FLAG
	TIMESTAMP_FLAG
	SET_FLAG

	PART_KEY_FLAG int = 1 << (14 + iota)
	NUM_OR_GROUP_FLAG
	UNIQUE_FLAG
)

const (
	AUTH_NAME = "mysql_native_password"
)

var (
	TK_ID_INSERT   = 1
	TK_ID_UPDATE   = 2
	TK_ID_DELETE   = 3
	TK_ID_REPLACE  = 4
	TK_ID_SET      = 5
	TK_ID_BEGIN    = 6
	TK_ID_COMMIT   = 7
	TK_ID_ROLLBACK = 8
	TK_ID_ADMIN    = 9
	TK_ID_USE      = 10

	TK_ID_SELECT      = 11
	TK_ID_START       = 12
	TK_ID_TRANSACTION = 13
	TK_ID_SHOW        = 14

	PARSE_TOKEN_MAP = map[string]int{
		"insert":      TK_ID_INSERT,
		"update":      TK_ID_UPDATE,
		"delete":      TK_ID_DELETE,
		"replace":     TK_ID_REPLACE,
		"set":         TK_ID_SET,
		"begin":       TK_ID_BEGIN,
		"commit":      TK_ID_COMMIT,
		"rollback":    TK_ID_ROLLBACK,
		"admin":       TK_ID_ADMIN,
		"select":      TK_ID_SELECT,
		"use":         TK_ID_USE,
		"start":       TK_ID_START,
		"transaction": TK_ID_TRANSACTION,
		"show":        TK_ID_SHOW,
	}
	// '*'
	COMMENT_PREFIX uint8 = 42
	COMMENT_STRING       = "*"

	//
	TK_STR_FROM = "from"
	TK_STR_INTO = "into"
	TK_STR_SET  = "set"

	TK_STR_TRANSACTION    = "transaction"
	TK_STR_LAST_INSERT_ID = "last_insert_id()"
	TK_STR_MASTER_HINT    = "*master*"

	SET_KEY_WORDS = map[string]struct{}{
		"names": struct{}{},

		"character_set_results":           struct{}{},
		"@@character_set_results":         struct{}{},
		"@@session.character_set_results": struct{}{},

		"character_set_client":           struct{}{},
		"@@character_set_client":         struct{}{},
		"@@session.character_set_client": struct{}{},

		"character_set_connection":           struct{}{},
		"@@character_set_connection":         struct{}{},
		"@@session.character_set_connection": struct{}{},

		"autocommit":           struct{}{},
		"@@autocommit":         struct{}{},
		"@@session.autocommit": struct{}{},
	}
)

const (
	DONTESCAPE = byte(255)
)
