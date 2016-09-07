package route

import (
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/utils/golog"
)

// Plan execute plain.
type Plan struct {
	Statement sqlparser.Statement
	DataNode  string
	IsSlave   bool
}

// Router used to build plan.
type Router struct {
	SchemaName   string
	Schemas      map[string]*config.SchemaConfig
	Nodes        map[string]*config.NodeConfig
	ConnectionID uint32
}

// NewRouter to create router.
func NewRouter(schemaName string, schemas map[string]*config.SchemaConfig, nodes map[string]*config.NodeConfig, connectionID uint32) *Router {
	r := new(Router)
	r.SchemaName = schemaName
	r.Schemas = schemas
	r.Nodes = nodes
	r.ConnectionID = connectionID
	return r
}

// BuildPlan to build plan
func (r *Router) BuildPlan(stmt sqlparser.Statement) (plan *Plan, result *mysql.Result, err error) {
	originalSQL := sqlparser.String(stmt)
	planSQL := originalSQL
	switch v := stmt.(type) {
	case *sqlparser.ShowEngines:
		plan, result, err = r.buildShowEnginesPlan(v)
	case *sqlparser.ShowPlugins:
		plan, result, err = r.buildShowPluginsPlan(v)
	case *sqlparser.ShowProcessList:
		plan, result, err = r.buildShowProcessListPlan(v)
	case *sqlparser.ShowFullProcessList:
		plan, result, err = r.buildShowFullProcessListPlan(v)
	case *sqlparser.ShowVariables:
		plan, result, err = r.buildShowVariablesPlan(v)
	case *sqlparser.ShowStatus:
		plan, result, err = r.buildShowStatusPlan(v)

	case *sqlparser.ShowTables:
		plan, result, err = r.buildShowTablesPlan(v)
	case *sqlparser.ShowFullTables:
		plan, result, err = r.buildShowFullTablesPlan(v)
	case *sqlparser.ShowColumns:
		plan, result, err = r.buildShowColumnsPlan(v)
	case *sqlparser.ShowFullColumns:
		plan, result, err = r.buildShowFullColumnsPlan(v)
	case *sqlparser.ShowIndex:
		plan, result, err = r.buildShowIndexPlan(v)
	case *sqlparser.ShowTriggers:
		plan, result, err = r.buildShowTriggersPlan(v)
	case *sqlparser.ShowProcedureStatus:
		plan, result, err = r.buildShowProcedureStatusPlan(v)
	case *sqlparser.ShowFunctionStatus:
		plan, result, err = r.buildShowFunctionStatusPlan(v)

	case *sqlparser.ShowCreateDatabase:
		plan, result, err = r.buildShowCreateDatabasePlan(v)
	case *sqlparser.ShowCreateTable:
		plan, result, err = r.buildShowCreateTablePlan(v)
	case *sqlparser.ShowCreateView:
		plan, result, err = r.buildShowCreateViewPlan(v)
	case *sqlparser.ShowCreateTrigger:
		plan, result, err = r.buildShowCreateTriggerPlan(v)
	case *sqlparser.ShowCreateProcedure:
		plan, result, err = r.buildShowCreateProcedurePlan(v)
	case *sqlparser.ShowCreateFunction:
		plan, result, err = r.buildShowCreateFunctionPlan(v)

	case *sqlparser.Insert:

	case *sqlparser.Replace:

	case *sqlparser.Select:
		plan, result, err = r.buildSelectPlan(v)
	case *sqlparser.Update:

	case *sqlparser.Delete:

	default:
		plan, result, err = nil, nil, errors.ErrNoPlan
	}

	planSQL = sqlparser.String(stmt)
	golog.Debug("route", "BuildPlan", "", r.ConnectionID,
		"originalSQL", originalSQL,
		"planSQL", planSQL)
	return
}
