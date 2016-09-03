package backend

// Connection is backend connection.
type Connection interface {
	// Connect db server.
	Connect(addr string, user string, password string, db string) error

	// Reconnect db server.
	Reconnect() error

	// Close db server.
	Close() error

	// UseDB to set current db.
	UseDB(database string) error

	// Ping db server.
	Ping() error
}
