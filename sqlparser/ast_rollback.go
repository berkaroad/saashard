package sqlparser

// Rollback statement
type Rollback struct {
}

func (node *Rollback) Format(buf *TrackedBuffer) {
	buf.Fprintf("rollback")
}

func (node *Rollback) IStatement() {}
