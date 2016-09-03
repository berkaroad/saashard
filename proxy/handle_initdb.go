package proxy

import (
	"strings"

	"github.com/berkaroad/saashard/errors"
)

func (c *ClientConn) handleInitDB(schema string) error {
	newdb := strings.ToLower(schema)
	if _, ok := c.proxy.getSchemasByUser(c.user)[newdb]; ok {
		c.db = newdb
		return nil
	}
	return errors.ErrNoSchema
}
