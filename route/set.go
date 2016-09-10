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

import "github.com/berkaroad/saashard/sqlparser"

func (r *Router) buildSetCharsetPlan(statement *sqlparser.SetCharset) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeName = schemaConfig.Nodes[0]
	plan.onSlave = false
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildSetNamesPlan(statement *sqlparser.SetNames) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeName = schemaConfig.Nodes[0]
	plan.onSlave = false
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildSetVariablePlan(statement *sqlparser.SetVariable) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeName = schemaConfig.Nodes[0]
	plan.onSlave = false
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildSetTransactionIsolationLevelPlan(statement *sqlparser.SetTransactionIsolationLevel) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	ReadHint(&statement.Comments)

	plan := new(normalPlan)
	plan.nodeName = schemaConfig.Nodes[0]
	plan.onSlave = false
	plan.anyNode = true
	plan.Statement = statement
	return plan, nil
}
