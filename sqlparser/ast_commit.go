package sqlparser

// Commit statement
type Commit struct {
}

func (node *Commit) Format(buf *TrackedBuffer) {
	buf.Fprintf("commit")
}

func (node *Commit) IStatement() {}
