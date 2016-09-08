package route

import (
	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/utils/golog"
)

// system db.
var sysDBMapping = map[string]string{
	"mysql": "mysql"}

// Plan execute plain.
type Plan struct {
	Statement sqlparser.Statement
	DataNode  string
	IsSlave   bool          // Execute at slave or master.
	Result    *mysql.Result // If has result then get it, or execute plan.
	anyNode   bool          // Can execute at any node or not.
}

// MergedPlan from plan array.
type MergedPlan struct {
	Statements []sqlparser.Statement
	Results    []*mysql.Result
	DataNode   string
	IsSlave    bool // Execute at slave or master.
	anyNode    bool // Can execute at any node or not.
}

// Router used to build plan.
type Router struct {
	SchemaName   string
	Schemas      map[string]*config.SchemaConfig
	Nodes        map[string]*config.NodeConfig
	ConnectionID uint32
	User         string
}

// NewRouter to create router.
func NewRouter(schemaName string, schemas map[string]*config.SchemaConfig, nodes map[string]*config.NodeConfig,
	connectionID uint32, user string) *Router {
	r := new(Router)
	r.SchemaName = schemaName
	r.Schemas = schemas
	r.Nodes = nodes
	r.ConnectionID = connectionID
	r.User = user
	return r
}

// BuildMergedPlan to merge from plan array.
func (r *Router) BuildMergedPlan(statements ...sqlparser.Statement) (plan *MergedPlan, err error) {
	if len(statements) == 0 {
		return nil, errors.ErrNoPlan
	}
	plans := make([]*Plan, len(statements))
	for i, stmt := range statements {
		statement := stmt
		plans[i], err = r.BuildPlan(statement)
		if err != nil {
			return
		}
	}
	planCount := len(plans)
	mergedPlan := new(MergedPlan)
	mergedPlan.Statements = make([]sqlparser.Statement, planCount)
	mergedPlan.Results = make([]*mysql.Result, planCount)
	mergedPlan.DataNode = plans[0].DataNode
	mergedPlan.IsSlave = plans[0].IsSlave
	mergedPlan.anyNode = plans[0].anyNode
	if plans[0].Result != nil {
		mergedPlan.Results[0] = plans[0].Result
	}
	mergedPlan.Statements[0] = plans[0].Statement

	if planCount > 1 {
		for i, currentPlan := range plans[1:] {
			if currentPlan.Result != nil {
				mergedPlan.Results[i+1] = currentPlan.Result
			} else {
				// couldn't execute in any node and exists more than one data node.
				if !mergedPlan.anyNode && mergedPlan.DataNode != currentPlan.DataNode {
					return nil, errors.ErrNoPlan
				}
				// if last plan can execute in any node, override it
				if mergedPlan.anyNode {
					mergedPlan.DataNode = currentPlan.DataNode
					mergedPlan.anyNode = currentPlan.anyNode
				}
				// if last plan use slave, override it
				if mergedPlan.IsSlave {
					mergedPlan.IsSlave = currentPlan.IsSlave
				}
			}
			mergedPlan.Statements[i+1] = currentPlan.Statement
		}
	}
	return mergedPlan, nil
}

// BuildPlan to build plan
func (r *Router) BuildPlan(statement sqlparser.Statement) (plan *Plan, err error) {
	originalSQL := sqlparser.String(statement)
	planSQL := originalSQL
	switch v := statement.(type) {
	case *sqlparser.ShowEngines:
		plan, err = r.buildShowEnginesPlan(v)
	case *sqlparser.ShowPlugins:
		plan, err = r.buildShowPluginsPlan(v)
	case *sqlparser.ShowProcessList:
		plan, err = r.buildShowProcessListPlan(v)
	case *sqlparser.ShowFullProcessList:
		plan, err = r.buildShowFullProcessListPlan(v)
	case *sqlparser.ShowVariables:
		plan, err = r.buildShowVariablesPlan(v)
	case *sqlparser.ShowStatus:
		plan, err = r.buildShowStatusPlan(v)
	case *sqlparser.ShowTables:
		plan, err = r.buildShowTablesPlan(v)
	case *sqlparser.ShowFullTables:
		plan, err = r.buildShowFullTablesPlan(v)
	case *sqlparser.ShowColumns:
		plan, err = r.buildShowColumnsPlan(v)
	case *sqlparser.ShowFullColumns:
		plan, err = r.buildShowFullColumnsPlan(v)
	case *sqlparser.ShowIndex:
		plan, err = r.buildShowIndexPlan(v)
	case *sqlparser.ShowTriggers:
		plan, err = r.buildShowTriggersPlan(v)
	case *sqlparser.ShowProcedureStatus:
		plan, err = r.buildShowProcedureStatusPlan(v)
	case *sqlparser.ShowFunctionStatus:
		plan, err = r.buildShowFunctionStatusPlan(v)
	case *sqlparser.ShowCreateDatabase:
		plan, err = r.buildShowCreateDatabasePlan(v)
	case *sqlparser.ShowCreateTable:
		plan, err = r.buildShowCreateTablePlan(v)
	case *sqlparser.ShowCreateView:
		plan, err = r.buildShowCreateViewPlan(v)
	case *sqlparser.ShowCreateTrigger:
		plan, err = r.buildShowCreateTriggerPlan(v)
	case *sqlparser.ShowCreateProcedure:
		plan, err = r.buildShowCreateProcedurePlan(v)
	case *sqlparser.ShowCreateFunction:
		plan, err = r.buildShowCreateFunctionPlan(v)

	case *sqlparser.SimpleSelect:
		plan, err = r.buildSimpleSelectPlan(v)
	case *sqlparser.Select:
		plan, err = r.buildSelectPlan(v)

	case sqlparser.SetStatement:
		plan, err = r.buildSetPlan(v)

	case *sqlparser.Insert:

	case *sqlparser.Replace:

	case *sqlparser.Update:

	case *sqlparser.Delete:

	default:
		plan, err = nil, errors.ErrNoPlan
	}

	planSQL = sqlparser.String(statement)
	golog.Debug("route", "BuildPlan", "", r.ConnectionID,
		"originalSQL", originalSQL,
		"planSQL", planSQL)
	return
}
