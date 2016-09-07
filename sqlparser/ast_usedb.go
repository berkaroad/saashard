package sqlparser

// UseDB statement
type UseDB struct {
	DB string
}

func (node *UseDB) Format(buf *TrackedBuffer) {
	buf.Fprintf("use %s", node.DB)
}

func (node *UseDB) IStatement() {}
