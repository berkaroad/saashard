package route

import "github.com/berkaroad/saashard/sqlparser"

func (r *Router) buildSetPlan(statement sqlparser.SetStatement) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(normalPlan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = false
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}
