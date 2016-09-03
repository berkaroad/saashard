package mysql

import (
	"github.com/berkaroad/saashard/net/mysql"
)

type Stmt struct {
	conn  *Conn
	id    uint32
	query string

	params  int
	columns int
}

func (s *Stmt) Execute(args ...interface{}) (*mysql.Result, error) {
	if err := s.conn.pkg.StmtExecute(s.id, args...); err != nil {
		return nil, err
	}
	return s.conn.pkg.ReadResultSet(s.conn.capability, &(s.conn.status), true)
}

func (s *Stmt) Close() error {
	if err := s.conn.pkg.StmtClose(s.id); err != nil {
		return err
	}

	return nil
}
