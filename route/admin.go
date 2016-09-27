package route

import "github.com/berkaroad/saashard/sqlparser"

func (r *Router) buildKillConnection(statement *sqlparser.KillConnection) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(normalPlan)

	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.onSlave = true && !r.InTrans
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildKillQuery(statement *sqlparser.KillQuery) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	plan := new(normalPlan)

	plan.nodeNames = []string{schemaConfig.Nodes[0]}
	plan.onSlave = true && !r.InTrans
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}
