package sqlparser

// Begin statement
type Begin struct {
}

func (node *Begin) Format(buf *TrackedBuffer) {
	buf.Fprintf("begin")
}

func (node *Begin) IStatement() {}
