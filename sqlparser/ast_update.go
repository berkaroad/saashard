package sqlparser

// Update represents an UPDATE statement.
type Update struct {
	Comments Comments
	Table    *TableName
	Exprs    UpdateExprs
	Where    *Where
	OrderBy  OrderBy
	Limit    *Limit
}

func (node *Update) Format(buf *TrackedBuffer) {
	buf.Fprintf("update %v%v set %v%v%v%v",
		node.Comments, node.Table,
		node.Exprs, node.Where, node.OrderBy, node.Limit)
}

func (node *Update) IStatement() {}
