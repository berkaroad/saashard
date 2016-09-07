package sqlparser

// SelectStatement any SELECT statement.
type SelectStatement interface {
	IStatement()
	ISelectStatement()
	IInsertRows()
	SQLNode
}

//SimpleSelect no from and where statement.
type SimpleSelect struct {
	Comments    Comments
	Distinct    string
	SelectExprs SelectExprs
	Limit       *Limit
}

func (node *SimpleSelect) Format(buf *TrackedBuffer) {
	buf.Fprintf("select %v%s%v%v", node.Comments, node.Distinct, node.SelectExprs, node.Limit)
}

func (*SimpleSelect) IStatement()       {}
func (*SimpleSelect) ISelectStatement() {}
func (*SimpleSelect) IInsertRows()      {}

// Select represents a SELECT statement.
type Select struct {
	Comments    Comments
	Distinct    string
	SelectExprs SelectExprs
	From        TableExprs
	Where       *Where
	GroupBy     GroupBy
	Having      *Where
	OrderBy     OrderBy
	Limit       *Limit
	Lock        string
}

// Select.Distinct
const (
	AST_DISTINCT = "distinct "
)

// Select.Lock
const (
	AST_FOR_UPDATE = " for update"
	AST_SHARE_MODE = " lock in share mode"
)

// Format Select.
func (node *Select) Format(buf *TrackedBuffer) {
	buf.Fprintf("select %v%s%v from %v%v%v%v%v%v%s",
		node.Comments, node.Distinct, node.SelectExprs,
		node.From, node.Where,
		node.GroupBy, node.Having, node.OrderBy,
		node.Limit, node.Lock)
}

func (node *Select) IStatement()       {}
func (node *Select) ISelectStatement() {}
func (node *Select) IInsertRows()      {}

// Union represents a UNION statement.
type Union struct {
	Type        string
	Left, Right SelectStatement
}

// Union.Type
const (
	AST_UNION     = "union"
	AST_UNION_ALL = "union all"
	AST_SET_MINUS = "minus"
	AST_EXCEPT    = "except"
	AST_INTERSECT = "intersect"
)

func (node *Union) Format(buf *TrackedBuffer) {
	buf.Fprintf("%v %s %v", node.Left, node.Type, node.Right)
}

func (node *Union) IStatement()       {}
func (node *Union) ISelectStatement() {}
func (node *Union) IInsertRows()      {}
