%{
// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

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

import "bytes"

// SetParseTree to build ast.
func SetParseTree(yylex interface{}, stmt Statement) {
  yylex.(*Tokenizer).ParseTree = stmt
}

// SetAllowComments set allow comments.
func SetAllowComments(yylex interface{}, allow bool) {
  yylex.(*Tokenizer).AllowComments = allow
}

// ForceEOF force EOF.
func ForceEOF(yylex interface{}) {
  yylex.(*Tokenizer).ForceEOF = true
}

var (
  SHARE =        []byte("share")
  MODE  =        []byte("mode")
  IF_BYTES =     []byte("if")
  VALUES_BYTES = []byte("values")
)

%}

%union {
  empty       struct{}
  statement   Statement
  selStmt     SelectStatement
  setStmt     SetStatement
  showStmt    ShowStatement
  ddlStmt     DDLStatement
  byt         byte
  bytes       []byte
  bytes2      [][]byte
  str         string
  boolean     bool
  selectExprs SelectExprs
  selectExpr  SelectExpr
  columns     Columns
  colName     *ColName
  tableExprs  TableExprs
  tableExpr   TableExpr
  smTableExpr SimpleTableExpr
  tableName   *TableName
  indexHints  *IndexHints
  expr        Expr
  boolExpr    BoolExpr
  valExpr     ValExpr
  tuple       Tuple
  valExprs    ValExprs
  values      Values
  subquery    *Subquery
  caseExpr    *CaseExpr
  funcExpr    *FuncExpr
  whens       []*When
  when        *When
  orderBy     OrderBy
  order       *Order
  limit       *Limit
  insRows     InsertRows
  updateExprs UpdateExprs
  updateExpr  *UpdateExpr
  createDef   CreateDefinition
  createDefs  CreateDefinitions
  columnDef   *ColumnDefinition
  dataType    *DataType
  idxColName  *IndexColName
  idxColNames IndexColNames
  optKeyVal   *OptionKeyValue
  optKeyVals  OptionKeyValues
  fiOAfCol    *FirstOrAfterColumn
  alterSpecs  AlterSpecifications
  alterSpec   AlterSpecification
}

%token LEX_ERROR
%token <empty> SELECT INSERT UPDATE DELETE FROM WHERE GROUP HAVING ORDER BY LIMIT FOR
%token <empty> ALL DISTINCT AS EXISTS NULL ASC DESC VALUES INTO DUPLICATE KEY DEFAULT SET LOCK
%token <empty> SHOW EXPLAIN DESCRIBE
%token <bytes> ID STRING NUMBER VALUE_ARG COMMENTS
%token <empty> '(' '~'

%left <empty> UNION MINUS EXCEPT INTERSECT
%left <empty> ','
%left <empty> FULL JOIN STRAIGHT_JOIN LEFT RIGHT INNER OUTER CROSS NATURAL USE FORCE
%left <empty> ON
%left <empty> OR
%left <empty> AND
%right <empty> NOT
%left <empty> BETWEEN CASE WHEN THEN ELSE
//REGEXP 
%left <empty> '=' '<' '>' LE GE NE NULL_SAFE_EQUAL IS LIKE IN
%left <empty> '|'
%left <empty> '&'
//%left <empty> SHIFT_LEFT SHIFT_RIGHT
%left <empty> '+' '-'
%left <empty> '*' '/' '%'
%left <empty> '^'
%nonassoc <empty> '.'
%left <empty> UNARY

%left <empty> END

// Transaction Tokens
%token <empty> BEGIN START TRANSACTION COMMIT ROLLBACK ISOLATION LEVEL READ COMMITTED UNCOMMITTED REPEATABLE SERIALIZABLE

// Charset Tokens
%token <empty> NAMES CHARSET CHARACTER COLLATION
%token <empty> ARMSCII8 ASCII BIG5 BINARY CP1250 CP1251 CP1256 CP1257 CP850 CP852 CP866 CP932
%token <empty> DEC8 EUCJPMS EUCKR GB2312 GBK GEOSTD8 GREEK HEBREW HP8 KEYBCS2 KOI8R KOI8U
%token <empty> LATIN1 LATIN2 LATIN5 LATIN7 MACCE MACROMAN SJIS SWE7 TIS620 UCS2 UJIS
%token <empty> UTF16 UTF16LE UTF32 UTF8 UTF8MB4

// Collation Tokens
%token <empty> ARMSCII8_GENERAL_CI ARMSCII8_BIN ASCII_GENERAL_CI ASCII_BIN BIG5_CHINESE_CI BIG5_BIN
%token <empty> CP1250_GENERAL_CI CP1250_BIN CP1251_GENERAL_CI CP1251_GENERAL_CS CP1251_BIN CP1256_GENERAL_CI CP1256_BIN CP1257_GENERAL_CI CP1257_BIN
%token <empty> CP850_GENERAL_CI CP850_BIN CP852_GENERAL_CI CP852_BIN CP866_GENERAL_CI CP866_BIN CP932_JAPANESE_CI CP932_BIN
%token <empty> DEC8_SWEDISH_CI DEC8_BIN EUCJPMS_JAPANESE_CI EUCJPMS_BIN EUCKR_KOREAN_CI EUCKR_BIN GB2312_CHINESE_CI GB2312_BIN GBK_CHINESE_CI GBK_BIN
%token <empty> GEOSTD8_GENERAL_CI GEOSTD8_BIN GREEK_GENERAL_CI GREEK_BIN HEBREW_GENERAL_CI HEBREW_BIN HP8_ENGLISH_CI HP8_BIN
%token <empty> KEYBCS2_GENERAL_CI KEYBCS2_BIN KOI8R_GENERAL_CI KOI8R_BIN KOI8U_GENERAL_CI KOI8U_BIN
%token <empty> LATIN1_GENERAL_CI LATIN1_GENERAL_CS LATIN1_BIN LATIN2_GENERAL_CI LATIN2_BIN LATIN5_TURKISH_CI LATIN5_BIN LATIN7_GENERAL_CI LATIN7_GENERAL_CS LATIN7_BIN
%token <empty> MACCE_GENERAL_CI MACCE_BIN MACROMAN_GENERAL_CI MACROMAN_BIN SJIS_JAPANESE_CI SJIS_BIN SWE7_SWEDISH_CI SWE7_BIN TIS620_THAI_CI TIS620_BIN
%token <empty> UCS2_GENERAL_CI UCS2_UNICODE_CI UCS2_BIN UJIS_JAPANESE_CI UJIS_BIN UTF16_GENERAL_CI UTF16_UNICODE_CI UTF16_BIN UTF16LE_GENERAL_CI UTF16LE_BIN
%token <empty> UTF32_GENERAL_CI UTF32_UNICODE_CI UTF32_BIN UTF8_GENERAL_CI UTF8_UNICODE_CI UTF8_BIN UTF8MB4_GENERAL_CI UTF8MB4_UNICODE_CI UTF8MB4_BIN


// Scope Tokens
%token <empty> SESSION GLOBAL

%token <empty> VARIABLES STATUS
%token <empty> DATABASES SCHEMAS DATABASE
%token <empty> STORAGE ENGINES
%token <empty> TABLES COLUMNS FIELDS PROCEDURE FUNCTION INDEXES KEYS TRIGGER TRIGGERS
%token <empty> PLUGINS PROCESSLIST SLAVE
%token <empty> PROFILES

// Replace
%token <empty> REPLACE

// admin
%token <empty> ADMIN HELP
//offset
%token <empty> OFFSET
//collate
%token <empty> COLLATE

// DDL Tokens
%token <empty> CREATE ALTER DROP RENAME
%token <empty> TABLE INDEX VIEW TO IGNORE IF UNIQUE FULLTEXT USING
%token <empty> BTREE HASH

// Data Type Tokens
%token <empty> BIT TINYINT SMALLINT MEDIUMINT INT INTEGER BIGINT REAL DOUBLE FLOAT DECIMAL
%token <empty> DATE TIME TIMESTAMP DATETIME YEAR
%token <empty> CHAR VARCHAR TINYTEXT TEXT MEDIUMTEXT LONGTEXT
%token <empty> VARBINARY TINYBLOB BLOB MEDIUMBLOB LONGBLOB


%token <empty> ENUM AUTO_INCREMENT ENGINE PRIMARY REFERENCES COMMENT
%token <empty> COLUMN_FORMAT FIXED DYNAMIC
%token <empty> DISK MEMORY
%token <empty> MATCH PARTIAL SIMPLE
%token <empty> RESTRICT CASCADE
%token <empty> NO ACTION
%token <empty> UNSIGNED ZEROFILL
%token <empty> CONSTRAINT FOREIGN
%token <empty> FIRST AFTER
%token <empty> ADD COLUMN CHANGE MODIFY
%token <empty> ENABLE DISABLE

// Functin
%token <empty> POSITION

%start any_command

%type <statement> command
%type <selStmt> select_statement
%type <setStmt> set_statement
%type <showStmt> show_statement describe_statement
%type <ddlStmt> create_statement alter_statement rename_statement drop_statement
%type <statement> insert_statement update_statement delete_statement replace_statement 
%type <statement> begin_statement commit_statement rollback_statement 
%type <statement> use_statement explain_statement admin_statement

%type <bytes2> comments_list_opt comments_list
%type <str> union_op
%type <str> distinct_opt
%type <str> ignore_opt
%type <selectExprs> select_expression_list
%type <selectExpr> select_expression
%type <bytes> as_lower_opt as_opt
%type <expr> expression
%type <tableExprs> table_expression_list
%type <tableExpr> table_expression
%type <str> join_type
%type <smTableExpr> simple_table_expression
%type <tableName> table_name
%type <indexHints> index_hint_list
%type <bytes2> sql_id_list
%type <boolExpr> where_expression_opt
%type <boolExpr> boolean_expression condition
%type <str> compare
%type <insRows> row_list
%type <valExpr> value value_expression
%type <tuple> tuple
%type <valExprs> value_expression_list
%type <values> tuple_list
%type <bytes> keyword_as_func
%type <subquery> subquery
%type <byt> unary_operator
%type <colName> column_name
%type <caseExpr> case_expression
%type <whens> when_expression_list
%type <when> when_expression
%type <funcExpr> function_expression
%type <valExpr> value_expression_opt else_expression_opt
%type <valExprs> group_by_opt
%type <boolExpr> having_opt
%type <orderBy> order_by_opt order_list
%type <order> order
%type <str> asc_desc_opt
%type <limit> limit_opt
%type <str> lock_opt
%type <columns> column_list_opt column_list
%type <updateExprs> on_dup_opt
%type <updateExprs> update_list
%type <updateExpr> update_expression
%type <expr> where_or_like_opt
%type <empty> exists_opt not_exists_opt
%type <bytes> index_category_opt index_type_opt
%type <idxColName> index_column_name
%type <idxColNames> index_column_list
%type <bytes> sql_id
// %type <empty> force_eof
%type <bytes> charset_words collate_words
%type <bytes> isolation_level
%type <bytes> scope_opt

%type <valExpr> default_value_opt column_comment_opt
%type <bytes> unique_or_primary_opt column_format_opt column_storage_opt
%type <bytes> reference_definition reference_definition_opt reference_match_opt
%type <bytes> reference_on_delete_or_update_opt reference_option reference_option_opt
%type <bytes> data_type_charset_opt data_type_collate_opt
%type <createDef> create_definition
%type <createDefs> create_definition_list
%type <alterSpec> alter_specification
%type <alterSpecs> alter_specification_list alter_specification_list_opt
%type <valExprs> value_list
%type <boolean> not_null_opt auto_increment_opt data_type_unsigned_opt data_type_zerofill_opt data_type_binary_opt
%type <columnDef> column_definition
%type <dataType> data_type
%type <bytes> constraint_opt
%type <optKeyVal> table_option
%type <optKeyVals> table_option_list table_option_list_opt

%type <fiOAfCol> first_or_after_column first_or_after_column_opt


%%

any_command:
  command
  {
    SetParseTree(yylex, $1)
  }

command:
  select_statement
  { $$ = $1 }
| set_statement
  { $$ = $1 }
| show_statement
  { $$ = $1 }
| describe_statement
  { $$ = $1 }
| insert_statement
| update_statement
| delete_statement
| replace_statement
| explain_statement
| create_statement
  { $$ = $1 }
| alter_statement
  { $$ = $1 }
| rename_statement
  { $$ = $1 }
| drop_statement
  { $$ = $1 }
| begin_statement
| commit_statement
| rollback_statement
| use_statement
| admin_statement
| comments_list
  { $$ = nil }

select_statement:
  SELECT comments_list_opt distinct_opt select_expression_list limit_opt
  {
    $$ = &SimpleSelect{Comments: Comments($2), Distinct: $3, SelectExprs: $4, Limit: $5}
  }
| SELECT comments_list_opt distinct_opt select_expression_list FROM table_expression_list where_expression_opt group_by_opt having_opt order_by_opt limit_opt lock_opt
  {
    $$ = &Select{Comments: Comments($2), Distinct: $3, SelectExprs: $4, From: $6, Where: NewWhere(AST_WHERE, $7), GroupBy: GroupBy($8), Having: NewWhere(AST_HAVING, $9), OrderBy: $10, Limit: $11, Lock: $12}
  }
| select_statement union_op select_statement %prec UNION
  {
    $$ = &Union{Type: $2, Left: $1, Right: $3}
  }

insert_statement:
  INSERT comments_list_opt ignore_opt INTO table_name column_list_opt row_list on_dup_opt
  {
    $$ = &Insert{Comments: Comments($2), Ignore:$3, Table: $5, Columns: $6, Rows: $7, OnDup: OnDup($8)}
  }
| INSERT comments_list_opt ignore_opt INTO table_name SET update_list on_dup_opt
  {
    cols := make(Columns, 0, len($7))
    vals := make(ValTuple, 0, len($7))
    for _, col := range $7 {
      cols = append(cols, &NonStarExpr{Expr: col.Name})
      vals = append(vals, col.Expr)
    }
    $$ = &Insert{Comments: Comments($2), Ignore:$3, Table: $5, Columns: cols, Rows: Values{vals}, OnDup: OnDup($8)}
  }

replace_statement:
  REPLACE comments_list_opt INTO table_name column_list_opt row_list
  {
    $$ = &Replace{Comments: Comments($2), Table: $4, Columns: $5, Rows: $6}
  }
| REPLACE comments_list_opt INTO table_name SET update_list
  {
    cols := make(Columns, 0, len($6))
    vals := make(ValTuple, 0, len($6))
    for _, col := range $6 {
      cols = append(cols, &NonStarExpr{Expr: col.Name})
      vals = append(vals, col.Expr)
    }
    $$ = &Replace{Comments: Comments($2), Table: $4, Columns: cols, Rows: Values{vals}}
  }

update_statement:
  UPDATE comments_list_opt table_name SET update_list where_expression_opt order_by_opt limit_opt
  {
    $$ = &Update{Comments: Comments($2), Table: $3, Exprs: $5, Where: NewWhere(AST_WHERE, $6), OrderBy: $7, Limit: $8}
  }

delete_statement:
  DELETE comments_list_opt FROM table_name where_expression_opt order_by_opt limit_opt
  {
    $$ = &Delete{Comments: Comments($2), Table: $4, Where: NewWhere(AST_WHERE, $5), OrderBy: $6, Limit: $7}
  }

explain_statement:
  EXPLAIN select_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN insert_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN update_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN replace_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN delete_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN set_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN show_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN describe_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN create_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN alter_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN rename_statement
  {
    $$= &Explain{Statement: $2}
  }
| EXPLAIN drop_statement
  {
    $$= &Explain{Statement: $2}
  }

set_statement:
  SET comments_list_opt scope_opt update_list
  {
    $$ = &SetVariable{
          Comments: Comments($2), 
          Scope:string($3), 
          Exprs: $4,
        }
  }
| SET comments_list_opt CHARACTER SET charset_words
  {
    $$ = &SetCharset{
	        Comments: Comments($2), 
	        Charset: string($5),
	    }
  }
| SET comments_list_opt CHARSET charset_words
  {
    $$ = &SetCharset{
	        Comments: Comments($2), 
	        Charset: string($4),
	    }
  }
| SET comments_list_opt NAMES charset_words 
  {
    $$ = &SetNames{
          Comments: Comments($2), 
          Names: string($4),
        }
  }
| SET comments_list_opt NAMES charset_words COLLATE STRING
  {
    $$ = &SetNames{
	        Comments: Comments($2), 
	        Names: string($4),
          Collate: string($6),
	    }
  }
| SET comments_list_opt scope_opt TRANSACTION ISOLATION LEVEL isolation_level
  {
    $$ = &SetTransactionIsolationLevel{
	        Comments: Comments($2),
          Scope: string($3),
	        IsolationLevel: string($7),
	    }
  }

begin_statement:
  BEGIN
  {
    $$ = &Begin{}
  }
| START TRANSACTION
  {
    $$ = &Begin{}
  }

commit_statement:
  COMMIT
  {
    $$ = &Commit{}
  }

rollback_statement:
  ROLLBACK
  {
    $$ = &Rollback{}
  }

use_statement:
  USE sql_id
  {
	  $$ = &UseDB{DB : string($2)}
  }

create_statement:
  CREATE comments_list_opt TABLE not_exists_opt table_name '(' create_definition_list ')' table_option_list_opt
  {
    $$ = &CreateTable{Comments : Comments($2), Table: $5, CreateDefs: $7, TableOptions: $9 }
  }
| CREATE comments_list_opt index_category_opt INDEX sql_id index_type_opt ON table_name '(' index_column_list ')'
  {
    $$ = &CreateIndex{Comments : Comments($2), IndexCategory: $3, Name: $5, IndexType: $6, Table: $8, IndexColumns: $10 }
  }

alter_statement:
  ALTER comments_list_opt ignore_opt TABLE table_name alter_specification_list_opt
  {
    $$ = &AlterTable{Comments : Comments($2), Ignore: $3, Table: $5, AlterSpecs: $6}
  }

rename_statement:
  RENAME comments_list_opt TABLE table_name TO table_name
  {
    $$ = &RenameTable{Comments : Comments($2), OldName: $4, NewName: $6}
  }

drop_statement:
  DROP comments_list_opt TABLE exists_opt table_name reference_option_opt
  {
    $$ = &DropTable{Comments : Comments($2), Name: $5, RefOption: $6}
  }
| DROP comments_list_opt INDEX sql_id ON table_name
  {
    $$ = &DropIndex{Comments : Comments($2), Name: $4, Table: $6}
  }

show_statement:
  SHOW comments_list_opt CHARACTER SET where_or_like_opt
  {
    $$ = &ShowCharset{Comments : Comments($2), LikeOrWhere : $5}
  }
| SHOW comments_list_opt CHARSET where_or_like_opt
  {
    $$ = &ShowCharset{Comments : Comments($2), LikeOrWhere : $4}
  }
| SHOW comments_list_opt COLLATION where_or_like_opt
  {
    $$ = &ShowCollation{Comments : Comments($2), LikeOrWhere : $4}
  }
| SHOW comments_list_opt scope_opt VARIABLES where_or_like_opt
  {
    $$ = &ShowVariables{Comments : Comments($2), Scope: string($3), LikeOrWhere : $5}
  }
| SHOW comments_list_opt scope_opt STATUS where_or_like_opt
  {
    $$ = &ShowStatus{Comments : Comments($2), Scope: string($3), LikeOrWhere : $5}
  }
| SHOW comments_list_opt DATABASES where_or_like_opt
  {
    $$ = &ShowDatabases{Comments : Comments($2), LikeOrWhere : $4}
  }
| SHOW comments_list_opt SCHEMAS where_or_like_opt
  {
    $$ = &ShowDatabases{Comments : Comments($2), LikeOrWhere : $4}
  }
| SHOW comments_list_opt TABLES where_or_like_opt
  {
    $$ = &ShowTables{Comments : Comments($2), LikeOrWhere : $4}
  }
| SHOW comments_list_opt TABLES FROM table_name where_or_like_opt
  {
    $$ = &ShowTables{Comments : Comments($2), From : $5, LikeOrWhere : $6}
  }
| SHOW comments_list_opt FULL TABLES where_or_like_opt
  {
    $$ = &ShowFullTables{Comments : Comments($2), LikeOrWhere : $5}
  }
| SHOW comments_list_opt FULL TABLES FROM table_name where_or_like_opt
  {
    $$ = &ShowFullTables{Comments : Comments($2), From : $6, LikeOrWhere : $7}
  }
| SHOW comments_list_opt TABLE STATUS where_or_like_opt
  {
    $$ = &ShowTableStatus{Comments : Comments($2), LikeOrWhere : $5}
  }
| SHOW comments_list_opt TABLE STATUS FROM table_name where_or_like_opt
  {
    $$ = &ShowTableStatus{Comments : Comments($2), From : $6, LikeOrWhere : $7}
  }
| SHOW comments_list_opt COLUMNS FROM table_name where_or_like_opt
  {
    $$ = &ShowColumns{Comments : Comments($2), From : $5, LikeOrWhere : $6}
  }
| SHOW comments_list_opt FIELDS FROM table_name where_or_like_opt
  {
    $$ = &ShowColumns{Comments : Comments($2), From : $5, LikeOrWhere : $6}
  }
| SHOW comments_list_opt FULL COLUMNS FROM table_name where_or_like_opt
  {
    $$ = &ShowFullColumns{Comments : Comments($2), From : $6, LikeOrWhere : $7}
  }
| SHOW comments_list_opt FULL FIELDS FROM table_name where_or_like_opt
  {
    $$ = &ShowFullColumns{Comments : Comments($2), From : $6, LikeOrWhere : $7}
  }
| SHOW comments_list_opt PROCEDURE STATUS where_or_like_opt
  {
    $$ = &ShowProcedureStatus{Comments : Comments($2), LikeOrWhere : $5}
  }
| SHOW comments_list_opt FUNCTION STATUS where_or_like_opt
  {
    $$ = &ShowFunctionStatus{Comments : Comments($2), LikeOrWhere : $5}
  }
| SHOW comments_list_opt INDEX FROM table_name where_expression_opt
  {
    $$ = &ShowIndex{Comments : Comments($2), From : $5, Where : $6}
  }
| SHOW comments_list_opt INDEXES FROM table_name where_expression_opt
  {
    $$ = &ShowIndex{Comments : Comments($2), From : $5, Where : $6}
  }
| SHOW comments_list_opt KEYS FROM table_name where_expression_opt
  {
    $$ = &ShowIndex{Comments : Comments($2), From : $5, Where : $6}
  }
| SHOW comments_list_opt TRIGGERS FROM table_name where_or_like_opt
  {
    $$ = &ShowTriggers{Comments : Comments($2), From : $5, LikeOrWhere : $6}
  }
| SHOW comments_list_opt CREATE DATABASE table_name
  {
    $$ = &ShowCreateDatabase{Comments : Comments($2), Name : $5}
  }
| SHOW comments_list_opt CREATE TABLE table_name
  {
    $$ = &ShowCreateTable{Comments : Comments($2), Name : $5}
  }
| SHOW comments_list_opt CREATE VIEW table_name
  {
    $$ = &ShowCreateView{Comments : Comments($2), Name : $5}
  }
| SHOW comments_list_opt CREATE PROCEDURE table_name
  {
    $$ = &ShowCreateProcedure{Comments : Comments($2), Name : $5}
  }
| SHOW comments_list_opt CREATE FUNCTION table_name
  {
    $$ = &ShowCreateFunction{Comments : Comments($2), Name : $5}
  }
| SHOW comments_list_opt CREATE TRIGGER table_name
  {
    $$ = &ShowCreateTrigger{Comments : Comments($2), Name : $5}
  }
| SHOW comments_list_opt PROCESSLIST
  {
    $$ = &ShowProcessList{Comments : Comments($2)}
  }
| SHOW comments_list_opt FULL PROCESSLIST
  {
    $$ = &ShowFullProcessList{Comments : Comments($2)}
  }
| SHOW comments_list_opt SLAVE STATUS
  {
    $$ = &ShowSlaveStatus{Comments : Comments($2)}
  }
| SHOW ENGINES
  {
    $$ = &ShowEngines{}
  }
| SHOW STORAGE ENGINES
  {
    $$ = &ShowEngines{}
  }
| SHOW PLUGINS
  {
    $$ = &ShowPlugins{}
  }
| SHOW PROFILES
  {
    $$ = &ShowProfiles{}
  }

describe_statement:
  DESCRIBE comments_list_opt table_name
  {
    $$ = &ShowColumns{Comments : Comments($2), From : $3}
  }


admin_statement:
  ADMIN
  { $$ = nil }
| ADMIN HELP
  { $$ = nil }

comments_list_opt:
  {
    SetAllowComments(yylex, true)
  }
  comments_list
  {
    $$ = $2
    SetAllowComments(yylex, false)
  }

comments_list:
  {
    $$ = nil
  }
| comments_list COMMENTS
  {
    $$ = append($1, $2)
  }

union_op:
  UNION
  {
    $$ = AST_UNION
  }
| UNION ALL
  {
    $$ = AST_UNION_ALL
  }
| MINUS
  {
    $$ = AST_SET_MINUS
  }
| EXCEPT
  {
    $$ = AST_EXCEPT
  }
| INTERSECT
  {
    $$ = AST_INTERSECT
  }

distinct_opt:
  {
    $$ = ""
  }
| DISTINCT
  {
    $$ = AST_DISTINCT
  }

select_expression_list:
  select_expression
  {
    $$ = SelectExprs{$1}
  }
| select_expression_list ',' select_expression
  {
    $$ = append($$, $3)
  }

select_expression:
  '*'
  {
    $$ = &StarExpr{}
  }
| expression as_lower_opt
  {
    $$ = &NonStarExpr{Expr: $1, As: $2}
  }
| ID '.' '*'
  {
    $$ = &StarExpr{TableName: $1}
  }

expression:
  boolean_expression
  {
    $$ = $1
  }
| value_expression
  {
    $$ = $1
  }

as_lower_opt:
  {
    $$ = nil
  }
| sql_id
  {
    $$ = $1
  }
| AS sql_id
  {
    $$ = $2
  }

table_expression_list:
  table_expression
  {
    $$ = TableExprs{$1}
  }
| table_expression_list ',' table_expression
  {
    $$ = append($$, $3)
  }

table_expression:
  simple_table_expression as_opt index_hint_list
  {
    $$ = &AliasedTableExpr{Expr:$1, As: $2, Hints: $3}
  }
| '(' table_expression ')'
  {
    $$ = &ParenTableExpr{Expr: $2}
  }
| table_expression join_type table_expression ON boolean_expression %prec JOIN
  {
    $$ = &JoinTableExpr{LeftExpr: $1, Join: $2, RightExpr: $3, On: $5}
  }

as_opt:
  {
    $$ = nil
  }
| ID
  {
    $$ = $1
  }
| AS ID
  {
    $$ = $2
  }

join_type:
  JOIN
  {
    $$ = AST_JOIN
  }
| STRAIGHT_JOIN
  {
    $$ = AST_STRAIGHT_JOIN
  }
| LEFT JOIN
  {
    $$ = AST_LEFT_JOIN
  }
| LEFT OUTER JOIN
  {
    $$ = AST_LEFT_JOIN
  }
| RIGHT JOIN
  {
    $$ = AST_RIGHT_JOIN
  }
| RIGHT OUTER JOIN
  {
    $$ = AST_RIGHT_JOIN
  }
| INNER JOIN
  {
    $$ = AST_JOIN
  }
| CROSS JOIN
  {
    $$ = AST_CROSS_JOIN
  }
| NATURAL JOIN
  {
    $$ = AST_NATURAL_JOIN
  }

simple_table_expression:
ID
  {
    $$ = &TableName{Name: $1}
  }
| ID '.' ID
  {
    $$ = &TableName{Qualifier: $1, Name: $3}
  }
| subquery
  {
    $$ = $1
  }

table_name:
ID
  {
    $$ = &TableName{Name: $1}
  }
| ID '.' ID
  {
    $$ = &TableName{Qualifier: $1, Name: $3}
  }

index_hint_list:
  {
    $$ = nil
  }
| USE INDEX '(' sql_id_list ')'
  {
    $$ = &IndexHints{Type: AST_USE, Indexes: $4}
  }
| IGNORE INDEX '(' sql_id_list ')'
  {
    $$ = &IndexHints{Type: AST_IGNORE, Indexes: $4}
  }
| FORCE INDEX '(' sql_id_list ')'
  {
    $$ = &IndexHints{Type: AST_FORCE, Indexes: $4}
  }

sql_id_list:
  sql_id
  {
    $$ = [][]byte{$1}
  }
| sql_id_list ',' sql_id
  {
    $$ = append($1, $3)
  }

value_list:
  value
  {
    $$ = ValExprs{$1}
  }
| value_list ',' value
  {
    $$ = append($1, $3)
  }

where_expression_opt:
  {
    $$ = nil
  }
| WHERE boolean_expression
  {
    $$ = $2
  }

boolean_expression:
  condition
| boolean_expression AND boolean_expression
  {
    $$ = &AndExpr{Left: $1, Right: $3}
  }
| boolean_expression OR boolean_expression
  {
    $$ = &OrExpr{Left: $1, Right: $3}
  }
| NOT boolean_expression
  {
    $$ = &NotExpr{Expr: $2}
  }
| '(' boolean_expression ')'
  {
    $$ = &ParenBoolExpr{Expr: $2}
  }

condition:
  value_expression compare value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: $2, Right: $3}
  }
| value_expression IN tuple
  {
    $$ = &ComparisonExpr{Left: $1, Operator: AST_IN, Right: $3}
  }
| value_expression NOT IN tuple
  {
    $$ = &ComparisonExpr{Left: $1, Operator: AST_NOT_IN, Right: $4}
  }
| value_expression LIKE value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: AST_LIKE, Right: $3}
  }
| value_expression NOT LIKE value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: AST_NOT_LIKE, Right: $4}
  }
| value_expression BETWEEN value_expression AND value_expression
  {
    $$ = &RangeCond{Left: $1, Operator: AST_BETWEEN, From: $3, To: $5}
  }
| value_expression NOT BETWEEN value_expression AND value_expression
  {
    $$ = &RangeCond{Left: $1, Operator: AST_NOT_BETWEEN, From: $4, To: $6}
  }
| value_expression IS NULL
  {
    $$ = &NullCheck{Operator: AST_IS_NULL, Expr: $1}
  }
| value_expression IS NOT NULL
  {
    $$ = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: $1}
  }
| EXISTS subquery
  {
    $$ = &ExistsExpr{Subquery: $2}
  }

compare:
  '='
  {
    $$ = AST_EQ
  }
| '<'
  {
    $$ = AST_LT
  }
| '>'
  {
    $$ = AST_GT
  }
| LE
  {
    $$ = AST_LE
  }
| GE
  {
    $$ = AST_GE
  }
| NE
  {
    $$ = AST_NE
  }
| NULL_SAFE_EQUAL
  {
    $$ = AST_NSE
  }

row_list:
  VALUES tuple_list
  {
    $$ = $2
  }
| select_statement
  {
    $$ = $1
  }

tuple_list:
  tuple
  {
    $$ = Values{$1}
  }
| tuple_list ',' tuple
  {
    $$ = append($1, $3)
  }

tuple:
  '(' value_expression_list ')'
  {
    $$ = ValTuple($2)
  }
| subquery
  {
    $$ = $1
  }

subquery:
  '(' select_statement ')'
  {
    $$ = &Subquery{$2}
  }

value_expression_list:
  value_expression
  {
    $$ = ValExprs{$1}
  }
| value_expression_list ',' value_expression
  {
    $$ = append($1, $3)
  }

value_expression:
  value
  {
    $$ = $1
  }
| column_name
  {
    $$ = $1
  }
| tuple
  {
    $$ = $1
  }
| value_expression '&' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_BITAND, Right: $3}
  }
| value_expression '|' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_BITOR, Right: $3}
  }
| value_expression '^' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_BITXOR, Right: $3}
  }
| value_expression '+' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_PLUS, Right: $3}
  }
| value_expression '-' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_MINUS, Right: $3}
  }
| value_expression '*' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_MULT, Right: $3}
  }
| value_expression '/' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_DIV, Right: $3}
  }
| value_expression '%' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: AST_MOD, Right: $3}
  }
| unary_operator value_expression %prec UNARY
  {
    if num, ok := $2.(NumVal); ok {
      switch $1 {
      case '-':
        $$ = append(NumVal("-"), num...)
      case '+':
        $$ = num
      default:
        $$ = &UnaryExpr{Operator: $1, Expr: $2}
      }
    } else {
      $$ = &UnaryExpr{Operator: $1, Expr: $2}
    }
  }
| function_expression
  {
    $$ = $1
  }
| case_expression
  {
    $$ = $1
  }

function_expression:
  sql_id '(' ')'
  {
    $$ = &FuncExpr{Name: $1}
  }
| sql_id '(' value_expression_list ')'
  {
    $$ = &FuncExpr{Name: $1, Exprs: $3}
  }
| sql_id '(' DISTINCT value_expression_list ')'
  {
    $$ = &FuncExpr{Name: $1, Distinct: true, Exprs: $4}
  }
| POSITION '(' value_expression IN value_expression ')'
  {
    $$ = &FuncExpr{Name: []byte("locate"), Exprs: ValExprs{$3,$5}}
  }
| keyword_as_func '(' value_expression_list ')'
  {
    $$ = &FuncExpr{Name: $1, Exprs: $3}
  }

keyword_as_func:
  IF
  {
    $$ = IF_BYTES
  }
| VALUES
  {
    $$ = VALUES_BYTES
  }

unary_operator:
  '+'
  {
    $$ = AST_UPLUS
  }
| '-'
  {
    $$ = AST_UMINUS
  }
| '~'
  {
    $$ = AST_TILDA
  }

case_expression:
  CASE value_expression_opt when_expression_list else_expression_opt END
  {
    $$ = &CaseExpr{Expr: $2, Whens: $3, Else: $4}
  }

value_expression_opt:
  {
    $$ = nil
  }
| value_expression
  {
    $$ = $1
  }

when_expression_list:
  when_expression
  {
    $$ = []*When{$1}
  }
| when_expression_list when_expression
  {
    $$ = append($1, $2)
  }

when_expression:
  WHEN boolean_expression THEN value_expression
  {
    $$ = &When{Cond: $2, Val: $4}
  }

else_expression_opt:
  {
    $$ = nil
  }
| ELSE value_expression
  {
    $$ = $2
  }

column_name:
  sql_id
  {
    $$ = &ColName{Name: $1}
  }
| ID '.' sql_id
  {
    $$ = &ColName{Qualifier: $1, Name: $3}
  }

value:
  STRING
  {
    $$ = StrVal($1)
  }
| NUMBER
  {
    $$ = NumVal($1)
  }
| VALUE_ARG
  {
    $$ = ValArg($1)
  }
| NULL
  {
    $$ = &NullVal{}
  }

group_by_opt:
  {
    $$ = nil
  }
| GROUP BY value_expression_list
  {
    $$ = $3
  }

having_opt:
  {
    $$ = nil
  }
| HAVING boolean_expression
  {
    $$ = $2
  }

order_by_opt:
  {
    $$ = nil
  }
| ORDER BY order_list
  {
    $$ = $3
  }

order_list:
  order
  {
    $$ = OrderBy{$1}
  }
| order_list ',' order
  {
    $$ = append($1, $3)
  }

order:
  value_expression asc_desc_opt
  {
    $$ = &Order{Expr: $1, Direction: $2}
  }

asc_desc_opt:
  {
    $$ = AST_ASC
  }
| ASC
  {
    $$ = AST_ASC
  }
| DESC
  {
    $$ = AST_DESC
  }

limit_opt:
  {
    $$ = nil
  }
| LIMIT value_expression
  {
    $$ = &Limit{Rowcount: $2}
  }
| LIMIT value_expression ',' value_expression
  {
    $$ = &Limit{Offset: $2, Rowcount: $4}
  }
| LIMIT value_expression OFFSET value_expression
  {
	$$ = &Limit{Offset: $4, Rowcount: $2}
  }

lock_opt:
  {
    $$ = ""
  }
| FOR UPDATE
  {
    $$ = AST_FOR_UPDATE
  }
| LOCK IN sql_id sql_id
  {
    if !bytes.Equal($3, SHARE) {
      yylex.Error("expecting share")
      return 1
    }
    if !bytes.Equal($4, MODE) {
      yylex.Error("expecting mode")
      return 1
    }
    $$ = AST_SHARE_MODE
  }

column_list_opt:
  {
    $$ = nil
  }
| '(' column_list ')'
  {
    $$ = $2
  }

column_list:
  column_name
  {
    $$ = Columns{&NonStarExpr{Expr: $1}}
  }
| column_list ',' column_name
  {
    $$ = append($$, &NonStarExpr{Expr: $3})
  }

on_dup_opt:
  {
    $$ = nil
  }
| ON DUPLICATE KEY UPDATE update_list
  {
    $$ = $5
  }

update_list:
  update_expression
  {
    $$ = UpdateExprs{$1}
  }
| update_list ',' update_expression
  {
    $$ = append($1, $3)
  }

update_expression:
  column_name '=' value_expression
  {
    $$ = &UpdateExpr{Name: $1, Expr: $3} 
  }

exists_opt:
  { $$ = struct{}{} }
| IF EXISTS
  { $$ = struct{}{} }

not_exists_opt:
  { $$ = struct{}{} }
| IF NOT EXISTS
  { $$ = struct{}{} }

ignore_opt:
  { $$ = "" }
| IGNORE
  { $$ = AST_IGNORE }

index_category_opt:
  { $$ = nil }
| UNIQUE
  { $$ = []byte("unique") }
| FULLTEXT
  { $$ = []byte("fulltext") }

index_column_list:
  index_column_name
  {
    $$ = IndexColNames{$1}
  }
| index_column_list ',' index_column_name
  {
    $$ = append($1, $3)
  }

index_column_name:
  column_name asc_desc_opt
  {
    $$ = &IndexColName{ColumnName: $1, AscOrDesc: $2}
  }
| column_name '(' NUMBER ')' asc_desc_opt
  {
    $$ = &IndexColName{ColumnName: $1, Length: NumVal($3), AscOrDesc: $5}
  }

index_type_opt:
  { $$ = nil }
| USING BTREE
  { $$ = []byte("btree") }
| USING HASH
  { $$ = []byte("hash") }

sql_id:
  ID
  {
    $$ = bytes.ToLower($1)
  }
| DATABASE
  {
    $$ = []byte("database")
  }

// force_eof:
// {
//   ForceEOF(yylex)
// }

charset_words:
  ARMSCII8
  { $$ = []byte("armscii8") }
| ASCII
  { $$ = []byte("ascii") }
| BIG5
  { $$ = []byte("big5") }
| BINARY
  { $$ = []byte("binary") }
| CP1250
  { $$ = []byte("cp1250") }
| CP1251
  { $$ = []byte("cp1251") }
| CP1256
  { $$ = []byte("cp1256") }
| CP1257
  { $$ = []byte("cp1257") }
| CP850
  { $$ = []byte("cp850") }
| CP852
  { $$ = []byte("cp852") }
| CP866
  { $$ = []byte("cp866") }
| CP932
  { $$ = []byte("cp932") }
| DEC8
  { $$ = []byte("dec8") }
| EUCJPMS
  { $$ = []byte("eucjpms") }
| EUCKR
  { $$ = []byte("euckr") }
| GB2312
  { $$ = []byte("gb2312") }
| GBK
  { $$ = []byte("gbk") }
| GEOSTD8
  { $$ = []byte("geostd8") }
| GREEK
  { $$ = []byte("greek") }
| HEBREW
  { $$ = []byte("hebrew") }
| HP8
  { $$ = []byte("hp8") }
| KEYBCS2
  { $$ = []byte("keybcs2") }
| KOI8R
  { $$ = []byte("koi8r") }
| KOI8U
  { $$ = []byte("koi8u") }
| LATIN1
  { $$ = []byte("latin1") }
| LATIN2
  { $$ = []byte("latin2") }
| LATIN5
  { $$ = []byte("latin5") }
| LATIN7
  { $$ = []byte("latin7") }
| MACCE
  { $$ = []byte("macce") }
| MACROMAN
  { $$ = []byte("macroman") }
| SJIS
  { $$ = []byte("sjis") }
| SWE7
  { $$ = []byte("swe7") }
| TIS620
  { $$ = []byte("tis620") }
| UCS2
  { $$ = []byte("ucs2") }
| UJIS
  { $$ = []byte("ujis") }
| UTF16
  { $$ = []byte("utf16") }
| UTF16LE
  { $$ = []byte("utf16le") }
| UTF32
  { $$ = []byte("utf32") }
| UTF8
  { $$ = []byte("utf8") }
| UTF8MB4
  { $$ = []byte("utf8mb4") }

collate_words:
  ARMSCII8_GENERAL_CI
  { $$ = []byte("armscii8_general_ci") }
| ARMSCII8_BIN
  { $$ = []byte("armscii8_bin") }
| ASCII_GENERAL_CI
  { $$ = []byte("ascii_general_ci") }
| ASCII_BIN
  { $$ = []byte("ascii_bin") }
| BIG5_CHINESE_CI
  { $$ = []byte("big5_chinese_ci") }
| BIG5_BIN
  { $$ = []byte("big5_bin") }
| BINARY
  { $$ = []byte("binary") }
| CP1250_GENERAL_CI
  { $$ = []byte("cp1250_general_ci") }
| CP1250_BIN
  { $$ = []byte("cp1250_bin") }
| CP1251_GENERAL_CI
  { $$ = []byte("cp1251_chinese_ci") }
| CP1251_GENERAL_CS
  { $$ = []byte("cp1251_chinese_cs") }
| CP1251_BIN
  { $$ = []byte("cp1251_bin") }
| CP1256_GENERAL_CI
  { $$ = []byte("cp1256_chinese_ci") }
| CP1256_BIN
  { $$ = []byte("cp1256_bin") }
| CP1257_GENERAL_CI
  { $$ = []byte("cp1257_chinese_ci") }
| CP1257_BIN
  { $$ = []byte("cp1257_bin") }
| CP850_GENERAL_CI
  { $$ = []byte("cp850_chinese_ci") }
| CP850_BIN
  { $$ = []byte("cp850_bin") }
| CP852_GENERAL_CI
  { $$ = []byte("cp852_chinese_ci") }
| CP852_BIN
  { $$ = []byte("cp852_bin") }
| CP866_GENERAL_CI
  { $$ = []byte("cp866_chinese_ci") }
| CP866_BIN
  { $$ = []byte("cp866_bin") }
| CP932_JAPANESE_CI
  { $$ = []byte("cp932_japanese_ci") }
| CP932_BIN
  { $$ = []byte("cp932_bin") }
| DEC8_SWEDISH_CI
  { $$ = []byte("dec8_swedish_ci") }
| DEC8_BIN
  { $$ = []byte("dec8_bin") }
| EUCJPMS_JAPANESE_CI
  { $$ = []byte("eucjpms_japanese_ci") }
| EUCJPMS_BIN
  { $$ = []byte("eucjpms_bin") }
| EUCKR_KOREAN_CI
  { $$ = []byte("euckr_korean_ci") }
| EUCKR_BIN
  { $$ = []byte("euckr_bin") }
| GB2312_CHINESE_CI
  { $$ = []byte("gb2312_chinese_ci") }
| GB2312_BIN
  { $$ = []byte("gb2312_bin") }
| GBK_CHINESE_CI
  { $$ = []byte("gbk_chinese_ci") }
| GBK_BIN
  { $$ = []byte("gbk_bin") }
| GEOSTD8_GENERAL_CI
  { $$ = []byte("geostd8_general_ci") }
| GEOSTD8_BIN
  { $$ = []byte("geostd8_bin") }
| GREEK_GENERAL_CI
  { $$ = []byte("greek_general_ci") }
| GREEK_BIN
  { $$ = []byte("greek_bin") }
| HEBREW_GENERAL_CI
  { $$ = []byte("hebrew_general_ci") }
| HEBREW_BIN
  { $$ = []byte("hebrew_bin") }
| HP8_ENGLISH_CI
  { $$ = []byte("hp8_english_ci") }
| HP8_BIN
  { $$ = []byte("hp8_bin") }
| KEYBCS2_GENERAL_CI
  { $$ = []byte("keybcs2_general_ci") }
| KEYBCS2_BIN
  { $$ = []byte("keybcs2_bin") }
| KOI8R_GENERAL_CI
  { $$ = []byte("koi8r_general_ci") }
| KOI8R_BIN
  { $$ = []byte("koi8r_bin") }
| KOI8U_GENERAL_CI
  { $$ = []byte("koi8u_general_ci") }
| KOI8U_BIN
  { $$ = []byte("koi8u_bin") }
| LATIN1_GENERAL_CI
  { $$ = []byte("latin1_general_ci") }
| LATIN1_GENERAL_CS
  { $$ = []byte("latin1_general_cs") }
| LATIN1_BIN
  { $$ = []byte("latin1_bin") }
| LATIN2_GENERAL_CI
  { $$ = []byte("latin2_general_ci") }
| LATIN2_BIN
  { $$ = []byte("latin2_bin") }
| LATIN5_TURKISH_CI
  { $$ = []byte("latin5_turkish_ci") }
| LATIN5_BIN
  { $$ = []byte("latin5_bin") }
| LATIN7_GENERAL_CI
  { $$ = []byte("latin7_general_ci") }
| LATIN7_GENERAL_CS
  { $$ = []byte("latin7_general_cs") }
| LATIN7_BIN
  { $$ = []byte("latin7_bin") }
| MACCE_GENERAL_CI
  { $$ = []byte("macce_general_ci") }
| MACCE_BIN
  { $$ = []byte("macce_bin") }
| MACROMAN_GENERAL_CI
  { $$ = []byte("macroman_general_ci") }
| MACROMAN_BIN
  { $$ = []byte("macroman_bin") }
| SJIS_JAPANESE_CI
  { $$ = []byte("sjis_japanese_ci") }
| SJIS_BIN
  { $$ = []byte("sjis_bin") }
| SWE7_SWEDISH_CI
  { $$ = []byte("swe7_swedish_ci") }
| SWE7_BIN
  { $$ = []byte("swe7_bin") }
| TIS620_THAI_CI
  { $$ = []byte("tis620_thai_ci") }
| TIS620_BIN
  { $$ = []byte("tis620_bin") }
| UCS2_GENERAL_CI
  { $$ = []byte("ucs2_general_ci") }
| UCS2_UNICODE_CI
  { $$ = []byte("ucs2_unicode_ci") }
| UCS2_BIN
  { $$ = []byte("ucs2_bin") }
| UJIS_JAPANESE_CI
  { $$ = []byte("ujis_japanese_ci") }
| UJIS_BIN
  { $$ = []byte("ujis_bin") }
| UTF16_GENERAL_CI
  { $$ = []byte("utf16_general_ci") }
| UTF16_UNICODE_CI
  { $$ = []byte("utf16_unicode_ci") }
| UTF16_BIN
  { $$ = []byte("utf16_bin") }
| UTF16LE_GENERAL_CI
  { $$ = []byte("utf16le_general_ci") }
| UTF16LE_BIN
  { $$ = []byte("utf16le_bin") }
| UTF32_GENERAL_CI
  { $$ = []byte("utf32_general_ci") }
| UTF32_UNICODE_CI
  { $$ = []byte("utf32_unicode_ci") }
| UTF32_BIN
  { $$ = []byte("utf32_bin") }
| UTF8_GENERAL_CI
  { $$ = []byte("utf8_general_ci") }
| UTF8_UNICODE_CI
  { $$ = []byte("utf8_unicode_ci") }
| UTF8_BIN
  { $$ = []byte("utf8_bin") }
| UTF8MB4_GENERAL_CI
  { $$ = []byte("utf8mb4_general_ci") }
| UTF8MB4_UNICODE_CI
  { $$ = []byte("utf8mb4_unicode_ci") }
| UTF8MB4_BIN
  { $$ = []byte("utf8mb4_bin") }


isolation_level:
  READ COMMITTED
  { $$ = []byte("read committed") }
| READ UNCOMMITTED
  { $$ = []byte("read uncommitted") }
| REPEATABLE READ
  { $$ = []byte("repeatable read") }
| SERIALIZABLE
  { $$ = []byte("serializable") }

scope_opt:
  { $$ = nil }
| SESSION
  { $$ = []byte("session") }
| GLOBAL
  { $$ = []byte("global") }

where_or_like_opt:
  { $$ = nil }
| LIKE value_expression
  {
    $$ = &LikeExpr{Expr:$2}
  }
| WHERE boolean_expression
  {
    $$ = &WhereExpr{Expr:$2}
  }

create_definition_list:
  create_definition
  {
    $$ = CreateDefinitions{ $1 }
  }
| create_definition_list ',' create_definition
  {
    $$ = append($1, $3)
  }

create_definition:
  column_name column_definition
  {
    $$ = &CreateColumnDefinition{ColumnName: $1, ColumnDef: $2 }
  }
| constraint_opt PRIMARY KEY index_type_opt '(' index_column_list ')'
  {
    $$ = &CreatePrimaryKeyDefinition{Symbol: $1, IndexType: $4, IndexColumns: $6}
  }
| INDEX sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &CreateIndexDefinition{Name: $2, IndexType: $3, IndexColumns: $5}
  }
| KEY sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &CreateIndexDefinition{Name: $2, IndexType: $3, IndexColumns: $5}
  }
| constraint_opt UNIQUE sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &CreateUniqueIndexDefinition{Symbol: $1, Name: $3, IndexType: $4, IndexColumns: $6}
  }
| constraint_opt UNIQUE INDEX sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &CreateUniqueIndexDefinition{Symbol: $1, Name: $4, IndexType: $5, IndexColumns: $7}
  }
| constraint_opt UNIQUE KEY sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &CreateUniqueIndexDefinition{Symbol: $1, Name: $4, IndexType: $5, IndexColumns: $7}
  }
| constraint_opt FOREIGN KEY '(' index_column_list ')' reference_definition
  {
    $$ = &CreateForeignKeyDefinition{Symbol: $1, IndexColumns: $5, ReferenceDef: $7}
  }

column_definition:
  data_type not_null_opt default_value_opt auto_increment_opt unique_or_primary_opt column_comment_opt column_format_opt column_storage_opt reference_definition_opt
  {
    $$ = &ColumnDefinition{Type: $1,
          IsNotNull: $2,
          DefaultValue: $3,
          IsAutoIncrement: $4,
          UniqueOrKey: $5,
          ColumnComment: $6,
          ColumnFormat: $7,
          ColumnStorage: $8,
          ReferenceDef: $9 }
  }

data_type:
  BIT
  {
    $$ = &DataType{ TypeName: "bit" } 
  }
| BIT '(' NUMBER ')'
  {
    $$ = &DataType{ TypeName: "bit(" + string($3) + ")" } 
  }
| TINYINT data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "tinyint", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| TINYINT '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "tinyint(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| SMALLINT data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "smallint", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| SMALLINT '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "smallint(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| MEDIUMINT data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "mediumint", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| MEDIUMINT '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "mediumint(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| INT data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "int", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| INT '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "int(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| INTEGER data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "integer", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| INTEGER '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "integer(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| BIGINT data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "bigint", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| BIGINT '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "bigint(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| REAL data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "real", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| REAL '(' NUMBER ',' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "real(" + string($3) + "," + string($5) + ")", IsUnsigned: $7, IsZeroFill: $8 } 
  }
| DOUBLE data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "double", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| DOUBLE '(' NUMBER ',' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "double(" + string($3) + "," + string($5) + ")", IsUnsigned: $7, IsZeroFill: $8 } 
  }
| FLOAT data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "float", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| FLOAT '(' NUMBER ',' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "float(" + string($3) + "," + string($5) + ")", IsUnsigned: $7, IsZeroFill: $8 } 
  }
| DECIMAL data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "decimal", IsUnsigned: $2, IsZeroFill: $3 } 
  }
| DECIMAL '(' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "decimal(" + string($3) + ")", IsUnsigned: $5, IsZeroFill: $6 } 
  }
| DECIMAL '(' NUMBER ',' NUMBER ')' data_type_unsigned_opt data_type_zerofill_opt
  { 
    $$ = &DataType{ TypeName: "decimal(" + string($3) + "," + string($5) + ")", IsUnsigned: $7, IsZeroFill: $8 } 
  }
| DATE
  {
    $$ = &DataType{ TypeName: "date" }
  }
| TIME
  {
    $$ = &DataType{ TypeName: "time" }
  }
| TIME '(' NUMBER ')'
  {
    $$ = &DataType{ TypeName: "time(" + string($3) + ")" }
  }
| TIMESTAMP
  {
    $$ = &DataType{ TypeName: "timestamp" }
  }
| TIMESTAMP '(' NUMBER ')'
  {
    $$ = &DataType{ TypeName: "timestamp(" + string($3) + ")" }
  }
| DATETIME
  {
    $$ = &DataType{ TypeName: "datetime" }
  }
| DATETIME '(' NUMBER ')'
  {
    $$ = &DataType{ TypeName: "datetime(" + string($3) + ")" }
  }
| YEAR
  {
    $$ = &DataType{ TypeName: "year" }
  }
| CHAR data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{ TypeName: "char", IsBinary: $2, Charset: $3, Collate: $4 }
  }
| CHAR '(' NUMBER ')' data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{ TypeName: "char(" + string($3) + ")", IsBinary: $5, Charset: $6, Collate: $7 }
  }
| VARCHAR '(' NUMBER ')' data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{ TypeName: "varchar(" + string($3) + ")", IsBinary: $5, Charset: $6, Collate: $7 }
  }
| BINARY
  {
    $$ = &DataType{TypeName: "binary"}
  }
| BINARY '(' NUMBER ')'
  {
    $$ = &DataType{TypeName: "binary(" + string($3) + ")"}
  }
| VARBINARY '(' NUMBER ')'
  {
    $$ = &DataType{TypeName: "varbinary(" + string($3) + ")"}
  }
| TINYBLOB
  {
    $$ = &DataType{TypeName: "tinyblob"}
  }
| BLOB
  {
    $$ = &DataType{TypeName: "blob"}
  }
| MEDIUMBLOB
  {
    $$ = &DataType{TypeName: "mediumblob"}
  }
| LONGBLOB
  {
    $$ = &DataType{TypeName: "longblob"}
  }
| TINYTEXT data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{TypeName: "tinytext", IsBinary: $2, Charset: $3, Collate: $4 }
  }
| TEXT data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{TypeName: "text", IsBinary: $2, Charset: $3, Collate: $4 }
  }
| MEDIUMTEXT data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{TypeName: "mediumtext", IsBinary: $2, Charset: $3, Collate: $4 }
  }
| LONGTEXT data_type_binary_opt data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{TypeName: "longtext", IsBinary: $2, Charset: $3, Collate: $4 }
  }
| ENUM '(' value_list ')' data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{TypeName: "enum(" + String($3) + ")", Charset: $5, Collate: $6 }
  }
| SET '(' value_list ')' data_type_charset_opt data_type_collate_opt
  {
    $$ = &DataType{TypeName: "set(" + String($3) + ")", Charset: $5, Collate: $6 }
  }
  
not_null_opt:
  { $$ = false }
| NULL
  { $$ = false }
| NOT NULL
  { $$ = true }

default_value_opt:
  { $$ = nil}
| DEFAULT value
  { $$ = $2 }

auto_increment_opt:
  { $$ = false }
| AUTO_INCREMENT
  { $$ = true }

unique_or_primary_opt:
  { $$ = nil }
| UNIQUE
  { $$ = []byte("unique key") }
| UNIQUE KEY
  { $$ = []byte("unique key") }
| KEY
  { $$ = []byte("primary key") }
| PRIMARY KEY
  { $$ = []byte("primary key") }

column_comment_opt:
  { $$ = nil }
| COMMENT STRING
  {
    $$ = StrVal($2)
  }

column_format_opt:
  { $$ = nil }
| COLUMN_FORMAT FIXED
  { $$ = []byte("fixed") }
| COLUMN_FORMAT DYNAMIC
  { $$ = []byte("dynamic") }
| COLUMN_FORMAT DEFAULT
  { $$ = []byte("default") }

column_storage_opt:
  { $$ = nil }
| STORAGE DISK
  { $$ = []byte("disk") }
| STORAGE MEMORY
  { $$ = []byte("memory") }
| STORAGE DEFAULT
  { $$ = []byte("default") }

reference_definition_opt:
  { $$ = nil }
| reference_definition
  { $$ = $1 }

reference_definition:
  REFERENCES table_name '(' index_column_list ')' reference_match_opt reference_on_delete_or_update_opt
  { $$ = []byte("references " + String($2) + "(" + String($4) + ")" + string($6) + string($7) ) }

reference_match_opt:
  { $$ = nil}
| MATCH FULL
  { $$ = []byte("match full") }
| MATCH PARTIAL
  { $$ = []byte("match partial") }
| MATCH SIMPLE
  { $$ = []byte("match simple") }

reference_on_delete_or_update_opt:
  { $$ = nil }
| ON DELETE reference_option
  { $$ = []byte("on delete " + string($3)) }
| ON UPDATE reference_option
  { $$ = []byte("on update " + string($3)) }
| ON DELETE reference_option ON UPDATE reference_option
  { $$ = []byte("on delete " + string($3) + " on update " + string($6)) }

reference_option_opt:
  { $$ = nil }
| reference_option
  { $$ = $1 }

reference_option:
  RESTRICT
  { $$ = []byte("restrict") }
| CASCADE
  { $$ = []byte("cascade") }
| SET NULL
  { $$ = []byte("set null") }
| NO ACTION
  { $$ = []byte("no action") }

data_type_unsigned_opt:
  { $$ = false}
| UNSIGNED
  { $$ = true }

data_type_zerofill_opt:
  { $$ = false }
| ZEROFILL
  { $$ = true }

data_type_binary_opt:
  { $$ = false }
| BINARY
  { $$ = true }

data_type_charset_opt:
  { $$ = nil }
| CHARACTER SET charset_words
  { $$ = $3 }

data_type_collate_opt:
  { $$ = nil }
| COLLATE collate_words
  { $$ = $2 }

constraint_opt:
  { $$ = nil }
| CONSTRAINT sql_id
  { $$ = $2 }

table_option_list_opt:
  { $$ = nil }
| table_option_list
  { $$ = $1 }

table_option_list:
  table_option
  { $$ = OptionKeyValues{$1} }
| table_option_list table_option
  { $$ = append($1, $2) }

table_option:
  ENGINE '=' sql_id
  {
    $$ = &OptionKeyValue{Key: "engine", Value: string($3)}
  }
| AUTO_INCREMENT '=' NUMBER
  {
    $$ = &OptionKeyValue{Key: "auto_increment", Value: string($3)}
  }
| DEFAULT CHARACTER SET '=' charset_words
  {
    $$ = &OptionKeyValue{Key: "charset", Value: string($5)}
  }
| CHARACTER SET '=' charset_words
  {
    $$ = &OptionKeyValue{Key: "charset", Value: string($4)}
  }
| DEFAULT CHARSET '=' charset_words
  {
    $$ = &OptionKeyValue{Key: "charset", Value: string($4)}
  }
| CHARSET '=' charset_words
  {
    $$ = &OptionKeyValue{Key: "charset", Value: string($3)}
  }
| DEFAULT COLLATE '=' collate_words
  {
    $$ = &OptionKeyValue{Key: "collate", Value: string($4)}
  }
| COLLATE '=' collate_words
  {
    $$ = &OptionKeyValue{Key: "collate", Value: string($3)}
  }
| COMMENT '=' STRING
  {
    $$ = &OptionKeyValue{Key: "comment", Value: String(StrVal($3))}
  }

alter_specification_list_opt:
  { $$ = nil }
| alter_specification_list
  { $$ = $1 }

alter_specification_list:
  alter_specification
  { $$ = AlterSpecifications{$1} }
| alter_specification_list ',' alter_specification
  { $$ = append($1, $3) }

alter_specification:
  table_option_list
  {
    $$ = $1
  }
| ADD COLUMN column_name column_definition first_or_after_column_opt
  {
    $$ = &AddOrModifyColumnSpec{Action: "add", ColumnName: $3, ColumnDef: $4, FirstOrAfterColumn: $5 }
  }
| ADD INDEX sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &AddIndexSpec{Name: $3, IndexType: $4, IndexColumns: $6}
  }
| ADD KEY sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &AddIndexSpec{Name: $3, IndexType: $4, IndexColumns: $6}
  }
| ADD constraint_opt PRIMARY KEY index_type_opt '(' index_column_list ')'
  {
    $$ = &AddPrimaryKeySpec{Symbol: $2, IndexType: $5, IndexColumns: $7}
  }
| ADD constraint_opt UNIQUE sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &AddUniqueIndexSpec{Symbol: $2, Name: $4, IndexType: $5, IndexColumns: $7}
  }
| ADD constraint_opt UNIQUE INDEX sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &AddUniqueIndexSpec{Symbol: $2, Name: $5, IndexType: $6, IndexColumns: $8}
  }
| ADD constraint_opt UNIQUE KEY sql_id index_type_opt '(' index_column_list ')'
  {
    $$ = &AddUniqueIndexSpec{Symbol: $2, Name: $5, IndexType: $6, IndexColumns: $8}
  }
| ADD constraint_opt FOREIGN KEY  '(' index_column_list ')' reference_definition 
  {
    $$ = &AddForeignKeySpec{Symbol: $2, IndexColumns: $6, ReferenceDef: $8}
  }
| CHANGE COLUMN column_name column_name column_definition first_or_after_column_opt
  {
    $$ = &ChangeColumnSpec{OldColumnName: $3, ColumnName: $4, ColumnDef: $5, FirstOrAfterColumn: $6 }
  }
| MODIFY COLUMN column_name column_definition first_or_after_column_opt
  {
    $$ = &AddOrModifyColumnSpec{Action: "modify", ColumnName: $3, ColumnDef: $4, FirstOrAfterColumn: $5 }
  }
| DROP COLUMN column_name
  {
    $$ = &DropColumnSpec{ColumnName: $3}
  }
| DROP PRIMARY KEY
  {
    $$ = &DropPrimaryKeySpec{}
  }
| DROP INDEX sql_id
  {
    $$ = &DropIndexSpec{Name: $3}
  }
| DROP KEY sql_id
  {
    $$ = &DropIndexSpec{Name: $3}
  }
| DROP FOREIGN KEY sql_id
  {
    $$ = &DropForeignKeySpec{Name: $4}
  }
| DISABLE KEYS
  {
    $$ = &DisableKeysSpec{}
  }
| ENABLE KEYS
  {
    $$ = &EnableKeysSpec{}
  }

first_or_after_column_opt:
  { $$ = nil }
| first_or_after_column
  { $$ = $1 }

first_or_after_column:
  FIRST column_name
  {
    $$ = &FirstOrAfterColumn{FirstOrAfter: "first", ColumnName: $2}
  }
| AFTER column_name
  {
    $$ = &FirstOrAfterColumn{FirstOrAfter: "after", ColumnName: $2}
  }