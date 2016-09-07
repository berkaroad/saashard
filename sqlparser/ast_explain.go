package sqlparser

// Explain select statement.
type Explain struct {
	Statement
}

// Format Explain.
func (node *Explain) Format(buf *TrackedBuffer) {
	buf.Fprintf("explain %v", node.Statement)
}

func (node *Explain) IStatement() {}
