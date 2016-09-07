package sqlparser

// AdminStatement admin statement.
type AdminStatement interface {
	IStatement()
	IAdminStatement()
	SQLNode
}
