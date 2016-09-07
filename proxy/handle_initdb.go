package proxy

import (
	"strings"

	"github.com/berkaroad/saashard/errors"
)

func (c *ClientConn) handleInitDB(db string) error {
	db = strings.ToLower(db)
	if _, ok := c.proxy.getSchemasByUser(c.user)[db]; ok {
		c.db = db
		return c.pkg.WriteOK(c.capability, c.status, nil)
	}
	return errors.ErrNoSchema
}
