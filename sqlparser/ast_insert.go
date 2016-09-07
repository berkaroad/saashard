package sqlparser

// Insert represents an INSERT statement.
type Insert struct {
	Comments Comments
	Ignore   string
	Table    *TableName
	Columns  Columns
	Rows     InsertRows
	OnDup    OnDup
}

func (node *Insert) Format(buf *TrackedBuffer) {
	buf.Fprintf("insert %v%s into %v%v %v%v",
		node.Comments, node.Ignore,
		node.Table, node.Columns, node.Rows, node.OnDup)
}
func (node *Insert) IStatement() {}
