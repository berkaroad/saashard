// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

package sqlparser

// http://dev.mysql.com/doc/refman/5.6/en/create-table.html

// DDLStatement ddl statement
type DDLStatement interface {
	IStatement()
	SQLNode
	IDDLStatement()
}

// CreateTable create table statement
type CreateTable struct {
	Comments     Comments
	Table        *TableName
	CreateDefs   CreateDefinitions
	TableOptions OptionKeyValues
}

// Format CreateTable
func (node *CreateTable) Format(buf *TrackedBuffer) {
	buf.Fprintf("create %v table if not exists %v\n(\n%v\n) %v", node.Comments, node.Table, node.CreateDefs, node.TableOptions)
}

func (node *CreateTable) IStatement()    {}
func (node *CreateTable) IDDLStatement() {}

// CreateIndex create index statement
type CreateIndex struct {
	Comments      Comments
	IndexCategory []byte
	Name          []byte
	Table         *TableName
	IndexType     []byte
	IndexColumns  IndexColNames
}

// Format CreateIndex
func (node *CreateIndex) Format(buf *TrackedBuffer) {
	strIndexCategory := ""
	if node.IndexCategory != nil {
		strIndexCategory = " " + string(node.IndexCategory)
	}
	indexType := ""
	if node.IndexType != nil && len(node.IndexType) > 0 {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf("create %v%s index %s%s on %v(%v)", node.Comments, strIndexCategory, node.Name, indexType, node.Table, node.IndexColumns)
}

func (node *CreateIndex) IStatement()    {}
func (node *CreateIndex) IDDLStatement() {}

// AlterTable alter table
type AlterTable struct {
	Comments   Comments
	Ignore     string
	Table      *TableName
	AlterSpecs AlterSpecifications
}

// Format AlterTable
func (node *AlterTable) Format(buf *TrackedBuffer) {
	buf.Fprintf("alter %s table %v\n%v", node.Ignore, node.Table, node.AlterSpecs)
}

func (node *AlterTable) IStatement()    {}
func (node *AlterTable) IDDLStatement() {}

// RenameTable rename table
type RenameTable struct {
	Comments Comments
	OldName  *TableName
	NewName  *TableName
}

// Format RenameTable
func (node *RenameTable) Format(buf *TrackedBuffer) {
	buf.Fprintf("rename %v table %v %v", node.Comments, node.OldName, node.NewName)
}

func (node *RenameTable) IStatement()    {}
func (node *RenameTable) IDDLStatement() {}

// DropTable drop table
type DropTable struct {
	Comments  Comments
	Name      *TableName
	RefOption []byte
}

// Format DropTable
func (node *DropTable) Format(buf *TrackedBuffer) {
	buf.Fprintf("drop %v table if exists %v %s", node.Comments, node.Name, node.RefOption)
}

func (node *DropTable) IStatement()    {}
func (node *DropTable) IDDLStatement() {}

// DropIndex drop index
type DropIndex struct {
	Comments Comments
	Name     []byte
	Table    *TableName
}

// Format DropIndex
func (node *DropIndex) Format(buf *TrackedBuffer) {
	buf.Fprintf("drop %v index %s on %v", node.Comments, node.Name, node.Table)
}

func (node *DropIndex) IStatement()    {}
func (node *DropIndex) IDDLStatement() {}

// CreateDefinitions create definition list.
type CreateDefinitions []CreateDefinition

// Format CreateDefinitions
func (node CreateDefinitions) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Fprintf("%s%v", prefix, n)
		prefix = ",\n"
	}
}

// CreateDefinition create definition
type CreateDefinition interface {
	SQLNode
	ICreateDefinition()
}

// CreateColumnDefinition create column definition
type CreateColumnDefinition struct {
	ColumnName *ColName
	ColumnDef  *ColumnDefinition
}

// Format CreateColumnDefinition
func (node *CreateColumnDefinition) Format(buf *TrackedBuffer) {
	buf.Fprintf("\t%v%v", node.ColumnName, node.ColumnDef)
}

func (node *CreateColumnDefinition) ICreateDefinition() {}

// CreatePrimaryKeyDefinition create primary key definition
type CreatePrimaryKeyDefinition struct {
	Symbol       []byte
	IndexType    []byte
	IndexColumns IndexColNames
}

// Format CreatePrimaryKeyDefinition
func (node *CreatePrimaryKeyDefinition) Format(buf *TrackedBuffer) {
	buf.Fprintf("\t", nil)
	if node.Symbol != nil {
		buf.Fprintf("constraint ", nil)
		escape(buf, node.Symbol)
		buf.Fprintf(" ", nil)
	}
	indexType := ""
	if node.IndexType != nil {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf("primary key%s(%v)", indexType, node.IndexColumns)
}

func (node *CreatePrimaryKeyDefinition) ICreateDefinition() {}

// CreateUniqueIndexDefinition create unique index definition
type CreateUniqueIndexDefinition struct {
	Symbol       []byte
	Name         []byte
	IndexType    []byte
	IndexColumns IndexColNames
}

// Format CreateUniqueIndexDefinition
func (node *CreateUniqueIndexDefinition) Format(buf *TrackedBuffer) {
	buf.Fprintf("\t", nil)
	if node.Symbol != nil {
		buf.Fprintf("constraint ", nil)
		escape(buf, node.Symbol)
		buf.Fprintf(" ", nil)
	}
	buf.Fprintf("unique key", nil)
	if node.Name != nil {
		buf.Fprintf(" ", nil)
		escape(buf, node.Name)
	}
	indexType := ""
	if node.IndexType != nil {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf("%s(%v)", indexType, node.IndexColumns)
}

func (node *CreateUniqueIndexDefinition) ICreateDefinition() {}

// CreateIndexDefinition create index definition
type CreateIndexDefinition struct {
	Name         []byte
	IndexType    []byte
	IndexColumns IndexColNames
}

// Format CreateIndexDefinition
func (node *CreateIndexDefinition) Format(buf *TrackedBuffer) {
	buf.Fprintf("\tindex", nil)
	if node.Name != nil {
		buf.Fprintf(" ", nil)
		escape(buf, node.Name)
	}
	indexType := ""
	if node.IndexType != nil {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf("%s(%v)", indexType, node.IndexColumns)
}

func (node *CreateIndexDefinition) ICreateDefinition() {}

// CreateForeignKeyDefinition create foreign key definition
type CreateForeignKeyDefinition struct {
	Symbol       []byte
	IndexColumns IndexColNames
	ReferenceDef []byte
}

// Format CreateForeignKeyDefinition
func (node *CreateForeignKeyDefinition) Format(buf *TrackedBuffer) {
	buf.Fprintf("\t", nil)
	if node.Symbol != nil {
		buf.Fprintf("constraint ", nil)
		escape(buf, node.Symbol)
		buf.Fprintf(" ", nil)
	}
	buf.Fprintf("foreign key(%v) %s", node.IndexColumns, node.ReferenceDef)
}

func (node *CreateForeignKeyDefinition) ICreateDefinition() {}

// ColumnDefinition column definition
type ColumnDefinition struct {
	Type            *DataType
	IsNotNull       bool
	DefaultValue    ValExpr
	IsAutoIncrement bool
	UniqueOrKey     []byte
	ColumnComment   ValExpr
	ColumnFormat    []byte
	ColumnStorage   []byte
	ReferenceDef    []byte
}

// Format ColumnDefinition
func (node *ColumnDefinition) Format(buf *TrackedBuffer) {
	strNullOrNotNull := ""
	if node.IsNotNull {
		strNullOrNotNull = " not null"
	} else {
		strNullOrNotNull = " null"
	}
	strDefaultValue := ""
	if node.DefaultValue != nil {
		strDefaultValue = " default " + String(node.DefaultValue)
	}
	strAutoIncrement := ""
	if node.IsAutoIncrement {
		strAutoIncrement = " auto_increment"
	}
	strComment := ""
	if node.ColumnComment != nil {
		strComment = " comment " + String(node.ColumnComment)
	}
	strColumnFormat := ""
	if node.ColumnFormat != nil {
		strColumnFormat = " column_format " + string(node.ColumnFormat)
	}
	strColumnStorage := ""
	if node.ColumnStorage != nil {
		strColumnStorage = " storage " + string(node.ColumnStorage)
	}
	strReferenceDef := ""
	if node.ReferenceDef != nil {
		strReferenceDef = " " + string(node.ReferenceDef)
	}
	buf.Fprintf("%v%s%s%s%s%s%s%s%s", node.Type, strNullOrNotNull, strDefaultValue, strAutoIncrement, node.UniqueOrKey, strComment, strColumnFormat, strColumnStorage, strReferenceDef)
}

// DataType data type.
type DataType struct {
	TypeName   string
	IsUnsigned bool
	IsZeroFill bool
	IsBinary   bool
	Charset    []byte
	Collate    []byte
}

// Format DataType
func (node *DataType) Format(buf *TrackedBuffer) {
	unsigned := ""
	if node.IsUnsigned {
		unsigned = " unsigned"
	}

	zeroFill := ""
	if node.IsZeroFill {
		zeroFill = " zerofill"
	}

	strBinary := ""
	if node.IsBinary {
		strBinary = " binary"
	}

	strCharset := ""
	if node.Charset != nil {
		strCharset = " character set " + string(node.Charset)
	}

	strCollate := ""
	if node.Collate != nil {
		strCollate = " collate " + string(node.Collate)
	}

	buf.Fprintf(" %s%s%s%s%s%s", node.TypeName, unsigned, zeroFill, strBinary, strCharset, strCollate)
}

// IndexColNames index column name list
type IndexColNames []*IndexColName

// Format IndexColNames
func (node IndexColNames) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Fprintf("%s%v", prefix, n)
		prefix = ","
	}
}

// IndexColName index column name
type IndexColName struct {
	ColumnName *ColName
	Length     ValExpr
	AscOrDesc  string
}

// Format IndexColName
func (node *IndexColName) Format(buf *TrackedBuffer) {
	strLength := ""
	if node.Length != nil {
		strLength = "(" + String(node.Length) + ")"
	}
	strAscOrDesc := ""
	if node.AscOrDesc != "" {
		strAscOrDesc = " " + node.AscOrDesc
	}
	buf.Fprintf("%v%s%s", node.ColumnName, strLength, strAscOrDesc)
}

// OptionKeyValues option key value list
type OptionKeyValues []*OptionKeyValue

// Format OptionKeyValues
func (node OptionKeyValues) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Fprintf("%s%v", prefix, n)
		prefix = " "
	}
}

func (node OptionKeyValues) IAlterSpecification() {}

// OptionKeyValue option key value
type OptionKeyValue struct {
	Key   string
	Value string
}

// Format OptionKeyValue
func (node *OptionKeyValue) Format(buf *TrackedBuffer) {
	buf.Fprintf("%s=%s", node.Key, node.Value)
}

// FirstOrAfterColumn first or after column
type FirstOrAfterColumn struct {
	FirstOrAfter string
	ColumnName   *ColName
}

// Format FirstOrAfterColumn
func (node *FirstOrAfterColumn) Format(buf *TrackedBuffer) {
	buf.Fprintf(" %s %v", node.FirstOrAfter, node.ColumnName)
}

// AlterSpecifications alter specification list
type AlterSpecifications []AlterSpecification

// Format AlterSpecifications
func (node AlterSpecifications) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Fprintf("%s%v", prefix, n)
		prefix = ",\n"
	}
}

// AlterSpecification alter specification
type AlterSpecification interface {
	IAlterSpecification()
	SQLNode
}

// AddOrModifyColumnSpec add column specification
type AddOrModifyColumnSpec struct {
	Action             string
	ColumnName         *ColName
	ColumnDef          *ColumnDefinition
	FirstOrAfterColumn *FirstOrAfterColumn
}

// Format AddOrModifyColumnSpec
func (node *AddOrModifyColumnSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("%s column %v%v%v", node.Action, node.ColumnName, node.ColumnDef, node.FirstOrAfterColumn)
}

func (node *AddOrModifyColumnSpec) IAlterSpecification() {}

// ChangeColumnSpec add column specification
type ChangeColumnSpec struct {
	OldColumnName      *ColName
	ColumnName         *ColName
	ColumnDef          *ColumnDefinition
	FirstOrAfterColumn *FirstOrAfterColumn
}

// Format ChangeColumnSpec
func (node *ChangeColumnSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("change column %v %v%v%v", node.OldColumnName, node.ColumnName, node.ColumnDef, node.FirstOrAfterColumn)
}

func (node *ChangeColumnSpec) IAlterSpecification() {}

// DropColumnSpec drop column specification
type DropColumnSpec struct {
	ColumnName *ColName
}

// Format DropColumnSpec
func (node *DropColumnSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("drop column %v", node.ColumnName)
}

func (node *DropColumnSpec) IAlterSpecification() {}

// AddIndexSpec add index specification
type AddIndexSpec struct {
	Name         []byte
	IndexType    []byte
	IndexColumns IndexColNames
}

// Format AddIndexSpec
func (node *AddIndexSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("add index", nil)
	if node.Name != nil {
		buf.Fprintf(" ", nil)
		escape(buf, node.Name)
	}
	indexType := ""
	if node.IndexType != nil {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf("%s(%v)", indexType, node.IndexColumns)
}

func (node *AddIndexSpec) IAlterSpecification() {}

// DropIndexSpec drop index specification
type DropIndexSpec struct {
	Name []byte
}

// Format DropIndexSpec
func (node *DropIndexSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("drop index ", nil)
	escape(buf, node.Name)
}

func (node *DropIndexSpec) IAlterSpecification() {}

// AddPrimaryKeySpec add primary key specification
type AddPrimaryKeySpec struct {
	Symbol       []byte
	IndexType    []byte
	IndexColumns IndexColNames
}

// Format AddPrimaryKeySpec
func (node *AddPrimaryKeySpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("add", nil)
	if node.Symbol != nil {
		buf.Fprintf(" constraint ", nil)
		escape(buf, node.Symbol)
		buf.Fprintf(" ", nil)
	}
	indexType := ""
	if node.IndexType != nil {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf(" primary key%s(%v)", indexType, node.IndexColumns)
}

func (node *AddPrimaryKeySpec) IAlterSpecification() {}

// DropPrimaryKeySpec drop primary key specification
type DropPrimaryKeySpec struct {
}

// Format DropPrimaryKeySpec
func (node *DropPrimaryKeySpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("drop primary key", nil)
}

func (node *DropPrimaryKeySpec) IAlterSpecification() {}

// AddUniqueIndexSpec add unique index specification
type AddUniqueIndexSpec struct {
	Symbol       []byte
	Name         []byte
	IndexType    []byte
	IndexColumns IndexColNames
}

// Format AddUniqueIndexSpec
func (node *AddUniqueIndexSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("add ", nil)
	if node.Symbol != nil {
		buf.Fprintf("constraint ", nil)
		escape(buf, node.Symbol)
		buf.Fprintf(" ", nil)
	}
	buf.Fprintf("unique key", nil)
	if node.Name != nil {
		buf.Fprintf(" ", nil)
		escape(buf, node.Name)
	}
	indexType := ""
	if node.IndexType != nil {
		indexType = " using " + string(node.IndexType)
	}
	buf.Fprintf("%s(%v)", indexType, node.IndexColumns)
}

func (node *AddUniqueIndexSpec) IAlterSpecification() {}

// AddForeignKeySpec add foreign key specification
type AddForeignKeySpec struct {
	Symbol       []byte
	IndexColumns IndexColNames
	ReferenceDef []byte
}

// Format AddForeignKeySpec
func (node *AddForeignKeySpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("add ", nil)
	if node.Symbol != nil {
		buf.Fprintf("constraint ", nil)
		escape(buf, node.Symbol)
		buf.Fprintf(" ", nil)
	}
	buf.Fprintf("foreign key(%v) %s", node.IndexColumns, node.ReferenceDef)
}

func (node *AddForeignKeySpec) IAlterSpecification() {}

// DropForeignKeySpec drop foreign key specification
type DropForeignKeySpec struct {
	Name []byte
}

// Format DropForeignKeySpec
func (node *DropForeignKeySpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("drop foreign key ", nil)
	escape(buf, node.Name)
}

func (node *DropForeignKeySpec) IAlterSpecification() {}

// DisableKeysSpec disable keys specification
type DisableKeysSpec struct {
}

// Format DisableKeysSpec
func (node *DisableKeysSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("disable keys", nil)
}

func (node *DisableKeysSpec) IAlterSpecification() {}

// EnableKeysSpec enable keys specification
type EnableKeysSpec struct {
}

// Format EnableKeysSpec
func (node *EnableKeysSpec) Format(buf *TrackedBuffer) {
	buf.Fprintf("enable keys", nil)
}

func (node *EnableKeysSpec) IAlterSpecification() {}
