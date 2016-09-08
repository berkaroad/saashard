package route

import "github.com/berkaroad/saashard/sqlparser"

func (r *Router) buildUseDBPlan(statement *sqlparser.UseDB) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(normalPlan)

	plan.DataNode = schemaConfig.Nodes[0]
	plan.IsSlave = true
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}
