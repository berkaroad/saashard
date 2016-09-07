package route

import (
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
)

func (r *Router) buildSelectPlan(stmt *sqlparser.Select) (*Plan, *mysql.Result, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(Plan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.Statement = stmt

	return plan, nil, nil
}
