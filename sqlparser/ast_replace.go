package sqlparser

// Replace represents an REPLACE statement.
type Replace struct {
	Comments Comments
	Table    *TableName
	Columns  Columns
	Rows     InsertRows
}

func (node *Replace) Format(buf *TrackedBuffer) {
	buf.Fprintf("replace %vinto %v%v %v",
		node.Comments,
		node.Table, node.Columns, node.Rows)
}

func (node *Replace) IStatement() {}
