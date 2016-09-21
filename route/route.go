// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package route

import (
	"net"
	"strings"
	"time"

	"github.com/berkaroad/saashard/config"
	"github.com/berkaroad/saashard/errors"
	"github.com/berkaroad/saashard/net/mysql"
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/statistic"
	"github.com/berkaroad/saashard/utils/golog"
)

// Route info
type Route interface {
	GetPlanSQL() string
	GetNodeNames() []string
	OnSlave() bool
}

// Plan to execute.
type Plan interface {
	Route
	Execute(executor func(statements []sqlparser.Statement, results []*mysql.Result,
		dataNodes []string, isSlave bool,
		queryDataNodes map[sqlparser.Statement][]string) error,
		clientAddr net.Addr, logSQLEnabled bool, slowLogTime int, counter *statistic.Counter) error
}

type normalPlan struct {
	Statement      sqlparser.Statement
	Result         *mysql.Result // If has result then get it, or execute plan.
	nodeNames      []string
	queryNodeNames []string
	onSlave        bool // Execute at slave or master.
	anyNode        bool // Can execute at any node or not.
}

func (plan *normalPlan) GetPlanSQL() string {
	return sqlparser.String(plan.Statement)
}

func (plan *normalPlan) GetNodeNames() []string {
	return plan.nodeNames
}

func (plan *normalPlan) OnSlave() bool {
	return plan.onSlave
}

// Execute the plan
func (plan *normalPlan) Execute(executor func(statements []sqlparser.Statement, results []*mysql.Result,
	dataNodes []string, isSlave bool,
	queryDataNodes map[sqlparser.Statement][]string) error, clientAddr net.Addr, logSQLEnabled bool, slowLogTime int, counter *statistic.Counter) error {
	if executor != nil {
		var state string
		startTime := time.Now().UnixNano()
		err := executor([]sqlparser.Statement{plan.Statement}, []*mysql.Result{plan.Result},
			plan.nodeNames, plan.onSlave,
			map[sqlparser.Statement][]string{plan.Statement: plan.queryNodeNames})
		execTime := float64(time.Now().UnixNano()-startTime) / float64(time.Millisecond)
		if err != nil {
			state = "ERROR"
		} else {
			state = "OK"
		}
		planSQL := plan.GetPlanSQL()
		if logSQLEnabled &&
			execTime > float64(slowLogTime) {
			counter.IncrSlowLogTotal()
			golog.OutputSql(state, "%.1fms - %s->%s:OnSlave=%v:%s",
				execTime,
				clientAddr,
				strings.Join(plan.nodeNames, ","),
				plan.onSlave,
				planSQL,
			)
		}
		return err
	}
	return nil
}

// mergedPlan from plan array.
type mergedPlan struct {
	Statements     []sqlparser.Statement
	Results        []*mysql.Result
	nodeNames      []string
	queryNodeNames map[sqlparser.Statement][]string // select or union will use.
	onSlave        bool                             // Execute at slave or master.
	anyNode        bool                             // Can execute at any node or not.
}

func (plan *mergedPlan) GetPlanSQL() string {
	planSQL := ""
	for _, statement := range plan.Statements {
		planSQL += sqlparser.String(statement) + "; "
	}
	return planSQL
}

func (plan *mergedPlan) GetNodeNames() []string {
	return plan.nodeNames
}

func (plan *mergedPlan) OnSlave() bool {
	return plan.onSlave
}

// Execute the plan
func (plan *mergedPlan) Execute(executor func(statements []sqlparser.Statement, results []*mysql.Result,
	dataNodes []string, isSlave bool,
	queryDataNodes map[sqlparser.Statement][]string) error, clientAddr net.Addr, logSQLEnabled bool, slowLogTime int, counter *statistic.Counter) error {
	if executor != nil {
		var state string
		startTime := time.Now().UnixNano()
		err := executor(plan.Statements, plan.Results,
			plan.nodeNames, plan.onSlave, plan.queryNodeNames)
		execTime := float64(time.Now().UnixNano()-startTime) / float64(time.Millisecond)
		if err != nil {
			state = "ERROR"
		} else {
			state = "OK"
		}
		planSQL := plan.GetPlanSQL()
		if logSQLEnabled &&
			execTime > float64(slowLogTime) {
			counter.IncrSlowLogTotal()
			golog.OutputSql(state, "%.1fms - %s->%s:OnSlave=%v:%s",
				execTime,
				clientAddr,
				strings.Join(plan.nodeNames, ","),
				plan.onSlave,
				planSQL,
			)
		}
		return err
	}
	return nil
}

// Router used to build plan.
type Router struct {
	SchemaName   string
	Schemas      map[string]*config.SchemaConfig
	Nodes        map[string]*config.NodeConfig
	ConnectionID uint32
	User         string
	InTrans      bool
}

// NewRouter to create router.
func NewRouter(schemaName string, schemas map[string]*config.SchemaConfig, nodes map[string]*config.NodeConfig,
	connectionID uint32, user string, inTrans bool) *Router {
	r := new(Router)
	r.SchemaName = schemaName
	r.Schemas = schemas
	r.Nodes = nodes
	r.ConnectionID = connectionID
	r.User = user
	r.InTrans = inTrans
	return r
}

// BuildMergedPlan to merge from plan array.
func (r *Router) BuildMergedPlan(statements ...sqlparser.Statement) (plan Plan, err error) {
	if len(statements) == 0 {
		return nil, errors.ErrNoStatement
	}
	plans := make([]*normalPlan, len(statements))
	for i, stmt := range statements {
		statement := stmt
		var thePlan Plan
		thePlan, err = r.BuildNormalPlan(statement)
		if err != nil {
			return
		}
		plans[i] = thePlan.(*normalPlan)
	}
	planCount := len(plans)

	mergedPlan := new(mergedPlan)
	mergedPlan.Statements = make([]sqlparser.Statement, planCount)
	mergedPlan.Results = make([]*mysql.Result, planCount)
	firstNormalPlan := plans[0]
	mergedPlan.nodeNames = firstNormalPlan.nodeNames
	mergedPlan.queryNodeNames = make(map[sqlparser.Statement][]string)
	mergedPlan.onSlave = firstNormalPlan.onSlave
	mergedPlan.anyNode = firstNormalPlan.anyNode
	mergedPlan.Results[0] = firstNormalPlan.Result
	mergedPlan.Statements[0] = firstNormalPlan.Statement

	if planCount > 1 {
		for i, currentPlan := range plans[1:] {
			if currentPlan.Result != nil {
				mergedPlan.Results[i+1] = currentPlan.Result
			} else {
				// couldn't execute multi-query that exists more than one data node.
				if len(mergedPlan.nodeNames) > 1 {
					return nil, errors.ErrExecInMulti
				}
				// couldn't execute in any node and exists more than one data node.
				if !mergedPlan.anyNode && mergedPlan.nodeNames[0] != currentPlan.nodeNames[0] {
					return nil, errors.ErrExecInMulti
				}
				// if last plan can execute in any node, override it
				if mergedPlan.anyNode {
					mergedPlan.nodeNames = currentPlan.nodeNames
					mergedPlan.anyNode = currentPlan.anyNode
				}
				// if last plan use slave, override it
				if mergedPlan.onSlave {
					mergedPlan.onSlave = currentPlan.onSlave
				}

				// mapping nodes to query statement.
				if len(currentPlan.queryNodeNames) > 0 {
					mergedPlan.queryNodeNames[currentPlan.Statement] = currentPlan.queryNodeNames
				}
			}
			mergedPlan.Statements[i+1] = currentPlan.Statement
		}
	}
	plan = mergedPlan
	return
}

// BuildNormalPlan to build plan
func (r *Router) BuildNormalPlan(statement sqlparser.Statement) (plan Plan, err error) {
	var realPlan *normalPlan
	switch v := statement.(type) {
	case *sqlparser.UseDB:
		realPlan, err = r.buildUseDBPlan(v)

	case *sqlparser.ShowEngines:
		realPlan, err = r.buildShowEnginesPlan(v)
	case *sqlparser.ShowPlugins:
		realPlan, err = r.buildShowPluginsPlan(v)
	case *sqlparser.ShowProcessList:
		realPlan, err = r.buildShowProcessListPlan(v)
	case *sqlparser.ShowFullProcessList:
		realPlan, err = r.buildShowFullProcessListPlan(v)
	case *sqlparser.ShowSlaveStatus:
		realPlan, err = r.buildShowSlaveStatusPlan(v)
	case *sqlparser.ShowProfiles:
		realPlan, err = r.buildShowProfilesPlan(v)
	case *sqlparser.ShowCharset:
		realPlan, err = r.buildShowCharsetPlan(v)
	case *sqlparser.ShowCollation:
		realPlan, err = r.buildShowCollationPlan(v)
	case *sqlparser.ShowVariables:
		realPlan, err = r.buildShowVariablesPlan(v)
	case *sqlparser.ShowStatus:
		realPlan, err = r.buildShowStatusPlan(v)
	case *sqlparser.ShowDatabases:
		realPlan, err = r.buildShowDatabasesPlan(v)
	case *sqlparser.ShowTables:
		realPlan, err = r.buildShowTablesPlan(v)
	case *sqlparser.ShowFullTables:
		realPlan, err = r.buildShowFullTablesPlan(v)
	case *sqlparser.ShowColumns:
		realPlan, err = r.buildShowColumnsPlan(v)
	case *sqlparser.ShowFullColumns:
		realPlan, err = r.buildShowFullColumnsPlan(v)
	case *sqlparser.ShowIndex:
		realPlan, err = r.buildShowIndexPlan(v)
	case *sqlparser.ShowTriggers:
		realPlan, err = r.buildShowTriggersPlan(v)
	case *sqlparser.ShowProcedureStatus:
		realPlan, err = r.buildShowProcedureStatusPlan(v)
	case *sqlparser.ShowFunctionStatus:
		realPlan, err = r.buildShowFunctionStatusPlan(v)
	case *sqlparser.ShowCreateDatabase:
		realPlan, err = r.buildShowCreateDatabasePlan(v)
	case *sqlparser.ShowCreateTable:
		realPlan, err = r.buildShowCreateTablePlan(v)
	case *sqlparser.ShowCreateView:
		realPlan, err = r.buildShowCreateViewPlan(v)
	case *sqlparser.ShowCreateTrigger:
		realPlan, err = r.buildShowCreateTriggerPlan(v)
	case *sqlparser.ShowCreateProcedure:
		realPlan, err = r.buildShowCreateProcedurePlan(v)
	case *sqlparser.ShowCreateFunction:
		realPlan, err = r.buildShowCreateFunctionPlan(v)

	case *sqlparser.SetCharset:
		realPlan, err = r.buildSetCharsetPlan(v)
	case *sqlparser.SetNames:
		realPlan, err = r.buildSetNamesPlan(v)
	case *sqlparser.SetVariable:
		realPlan, err = r.buildSetVariablePlan(v)
	case *sqlparser.SetTransactionIsolationLevel:
		realPlan, err = r.buildSetTransactionIsolationLevelPlan(v)

	case sqlparser.TransactionStatement:
		realPlan, err = r.buildTransactionPlan(v)

	case *sqlparser.Explain:
		realPlan, err = r.buildExplainPlan(v)

	case *sqlparser.SimpleSelect:
		realPlan, err = r.buildSimpleSelectPlan(v)
	case *sqlparser.Select:
		realPlan, err = r.buildSelectPlan(v)
	case *sqlparser.Union:
		realPlan, err = r.buildUnionPlan(v)

	case *sqlparser.Insert:
		realPlan, err = r.buildInsertPlan(v)
	case *sqlparser.Update:
		realPlan, err = r.buildUpdatePlan(v)
	case *sqlparser.Delete:
		realPlan, err = r.buildDeletePlan(v)
	case *sqlparser.Replace:
		realPlan, err = r.buildReplacePlan(v)

	case *sqlparser.CreateTable:
		realPlan, err = r.buildCreateTablePlan(v)
	case *sqlparser.CreateIndex:
		realPlan, err = r.buildCreateIndexPlan(v)
	case *sqlparser.RenameTable:
		realPlan, err = r.buildRenameTablePlan(v)
	case *sqlparser.DropTable:
		realPlan, err = r.buildDropTablePlan(v)
	case *sqlparser.DropIndex:
		realPlan, err = r.buildDropIndexPlan(v)
	default:
		realPlan, err = nil, errors.ErrNoPlan
	}
	plan = realPlan
	return
}
