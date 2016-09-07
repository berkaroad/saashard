package sqlparser

// http://dev.mysql.com/doc/refman/5.6/en/set-statement.html

// SetStatement set statement
type SetStatement interface {
	IStatement()
	ISetStatement()
	SQLNode
}

// SetVariable set variable assignment
type SetVariable struct {
	Comments Comments
	Scope    string // SESSION or GLOBAL
	Exprs    UpdateExprs
}

func (node *SetVariable) Format(buf *TrackedBuffer) {
	if len(node.Scope) == 0 {
		buf.Fprintf("set %v %v", node.Comments, node.Exprs)
	} else {
		buf.Fprintf("set %v %s %v", node.Comments, node.Scope, node.Exprs)
	}
}

func (node *SetVariable) IStatement()    {}
func (node *SetVariable) ISetStatement() {}

// SetCharset set charset.
type SetCharset struct {
	Comments Comments
	Scope    string // SESSION or GLOBAL
	Charset  string
}

func (node *SetCharset) Format(buf *TrackedBuffer) {
	if len(node.Scope) == 0 {
		buf.Fprintf("set %v character set %s", node.Comments, node.Charset)
	} else {
		buf.Fprintf("set %v %s character set %s", node.Comments, node.Scope, node.Charset)
	}
}

func (node *SetCharset) IStatement()    {}
func (node *SetCharset) ISetStatement() {}

// SetNames set NAMES.
type SetNames struct {
	Comments Comments
	Names    string
	Collate  string
}

func (node *SetNames) Format(buf *TrackedBuffer) {
	if node.Collate == "" {
		buf.Fprintf("set %v names %s", node.Comments, node.Names)
	} else {
		buf.Fprintf("set %v names %s collate %s", node.Comments, node.Names, node.Collate)
	}
}

func (node *SetNames) IStatement()    {}
func (node *SetNames) ISetStatement() {}

// SetTransactionIsolationLevel represents a SET TRANSACTION ISOLATION LEVEL statement.
type SetTransactionIsolationLevel struct {
	Comments       Comments
	Scope          string // SESSION or GLOBAL
	IsolationLevel string
}

func (node *SetTransactionIsolationLevel) Format(buf *TrackedBuffer) {
	if len(node.Scope) == 0 {
		buf.Fprintf("set %v transaction isolation level %s", node.Comments, node.IsolationLevel)
	} else {
		buf.Fprintf("set %v %s transaction isolation level %s", node.Comments, node.Scope, node.IsolationLevel)
	}
}

func (node *SetTransactionIsolationLevel) IStatement()    {}
func (node *SetTransactionIsolationLevel) ISetStatement() {}
