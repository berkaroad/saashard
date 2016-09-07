package proxy

import (
	"bytes"

	"github.com/berkaroad/saashard/backend"
	mysqlBackend "github.com/berkaroad/saashard/backend/mysql"
	"github.com/berkaroad/saashard/net/mysql"
)

func (c *ClientConn) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	nodeName := "db1_node1" // get from router
	node := c.proxy.nodes[nodeName]

	var err error
	var conn backend.Connection
	if conn, err = node.DataHost.Slaves[0].Connect(node.Database); err != nil {
		return err
	}
	mysqlConn := conn.(*mysqlBackend.Conn)

	defer func() {
		if !c.isInTransaction() {
			conn.Close()
		}
	}()

	if err = conn.UseDB(c.db); err != nil {
		return err
	}

	var fields []*mysql.Field
	if fields, err = mysqlConn.FieldList(table, wildcard); err != nil {
		return err
	}
	c.affectedRows = int64(-1)

	if err = c.pkg.WriteFieldList(c.capability, c.status, fields); err != nil {
		return err
	}
	return nil
}
