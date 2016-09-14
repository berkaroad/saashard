//line ./yacc.y:2

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

import __yyfmt__ "fmt"

//line ./yacc.y:42
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
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line ./yacc.y:70
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	setStmt     SetStatement
	showStmt    ShowStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
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
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const NULL = 57363
const ASC = 57364
const DESC = 57365
const VALUES = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const LOCK = 57372
const SHOW = 57373
const EXPLAIN = 57374
const DESCRIBE = 57375
const ID = 57376
const STRING = 57377
const NUMBER = 57378
const VALUE_ARG = 57379
const COMMENT = 57380
const UNION = 57381
const MINUS = 57382
const EXCEPT = 57383
const INTERSECT = 57384
const FULL = 57385
const JOIN = 57386
const STRAIGHT_JOIN = 57387
const LEFT = 57388
const RIGHT = 57389
const INNER = 57390
const OUTER = 57391
const CROSS = 57392
const NATURAL = 57393
const USE = 57394
const FORCE = 57395
const ON = 57396
const OR = 57397
const AND = 57398
const NOT = 57399
const BETWEEN = 57400
const CASE = 57401
const WHEN = 57402
const THEN = 57403
const ELSE = 57404
const LE = 57405
const GE = 57406
const NE = 57407
const NULL_SAFE_EQUAL = 57408
const IS = 57409
const LIKE = 57410
const IN = 57411
const UNARY = 57412
const END = 57413
const BEGIN = 57414
const START = 57415
const TRANSACTION = 57416
const COMMIT = 57417
const ROLLBACK = 57418
const ISOLATION = 57419
const LEVEL = 57420
const READ = 57421
const COMMITTED = 57422
const UNCOMMITTED = 57423
const REPEATABLE = 57424
const SERIALIZABLE = 57425
const NAMES = 57426
const CHARSET = 57427
const CHARACTER = 57428
const COLLATION = 57429
const ARMSCII8 = 57430
const ASCII = 57431
const BIG5 = 57432
const BINARY = 57433
const CP1250 = 57434
const CP1251 = 57435
const CP1256 = 57436
const CP1257 = 57437
const CP850 = 57438
const CP852 = 57439
const CP866 = 57440
const CP932 = 57441
const DEC8 = 57442
const EUCJPMS = 57443
const EUCKR = 57444
const GB2312 = 57445
const GBK = 57446
const GEOSTD8 = 57447
const GREEK = 57448
const HEBREW = 57449
const HP8 = 57450
const KEYBCS2 = 57451
const KOI8R = 57452
const KOI8U = 57453
const LATIN1 = 57454
const LATIN2 = 57455
const LATIN5 = 57456
const LATIN7 = 57457
const MACCE = 57458
const MACROMAN = 57459
const SJIS = 57460
const SWE7 = 57461
const TIS620 = 57462
const UCS2 = 57463
const EJIS = 57464
const UTF16 = 57465
const UTF16LE = 57466
const UTF32 = 57467
const UTF8 = 57468
const UTF8MB4 = 57469
const SESSION = 57470
const GLOBAL = 57471
const VARIABLES = 57472
const STATUS = 57473
const DATABASES = 57474
const SCHEMAS = 57475
const DATABASE = 57476
const STORAGE = 57477
const ENGINES = 57478
const TABLES = 57479
const COLUMNS = 57480
const FIELDS = 57481
const PROCEDURE = 57482
const FUNCTION = 57483
const INDEXES = 57484
const KEYS = 57485
const TRIGGER = 57486
const TRIGGERS = 57487
const PLUGINS = 57488
const PROCESSLIST = 57489
const SLAVE = 57490
const PROFILES = 57491
const REPLACE = 57492
const ADMIN = 57493
const HELP = 57494
const OFFSET = 57495
const COLLATE = 57496
const CREATE = 57497
const ALTER = 57498
const DROP = 57499
const RENAME = 57500
const TABLE = 57501
const INDEX = 57502
const VIEW = 57503
const TO = 57504
const IGNORE = 57505
const IF = 57506
const UNIQUE = 57507
const USING = 57508
const POSITION = 57509

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"SHOW",
	"EXPLAIN",
	"DESCRIBE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"'('",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"FULL",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"IN",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'.'",
	"UNARY",
	"END",
	"BEGIN",
	"START",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"ISOLATION",
	"LEVEL",
	"READ",
	"COMMITTED",
	"UNCOMMITTED",
	"REPEATABLE",
	"SERIALIZABLE",
	"NAMES",
	"CHARSET",
	"CHARACTER",
	"COLLATION",
	"ARMSCII8",
	"ASCII",
	"BIG5",
	"BINARY",
	"CP1250",
	"CP1251",
	"CP1256",
	"CP1257",
	"CP850",
	"CP852",
	"CP866",
	"CP932",
	"DEC8",
	"EUCJPMS",
	"EUCKR",
	"GB2312",
	"GBK",
	"GEOSTD8",
	"GREEK",
	"HEBREW",
	"HP8",
	"KEYBCS2",
	"KOI8R",
	"KOI8U",
	"LATIN1",
	"LATIN2",
	"LATIN5",
	"LATIN7",
	"MACCE",
	"MACROMAN",
	"SJIS",
	"SWE7",
	"TIS620",
	"UCS2",
	"EJIS",
	"UTF16",
	"UTF16LE",
	"UTF32",
	"UTF8",
	"UTF8MB4",
	"SESSION",
	"GLOBAL",
	"VARIABLES",
	"STATUS",
	"DATABASES",
	"SCHEMAS",
	"DATABASE",
	"STORAGE",
	"ENGINES",
	"TABLES",
	"COLUMNS",
	"FIELDS",
	"PROCEDURE",
	"FUNCTION",
	"INDEXES",
	"KEYS",
	"TRIGGER",
	"TRIGGERS",
	"PLUGINS",
	"PROCESSLIST",
	"SLAVE",
	"PROFILES",
	"REPLACE",
	"ADMIN",
	"HELP",
	"OFFSET",
	"COLLATE",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"POSITION",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 314
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 875

var yyAct = [...]int{

	153, 514, 290, 163, 547, 357, 143, 258, 397, 416,
	292, 250, 509, 154, 556, 293, 3, 419, 137, 345,
	144, 165, 129, 339, 337, 122, 133, 74, 148, 152,
	267, 266, 162, 393, 556, 429, 152, 556, 388, 162,
	80, 476, 136, 149, 150, 151, 60, 141, 157, 168,
	149, 150, 151, 253, 296, 157, 527, 84, 275, 274,
	277, 278, 279, 280, 281, 276, 526, 388, 140, 525,
	160, 124, 388, 123, 126, 388, 125, 160, 130, 42,
	43, 44, 45, 75, 170, 235, 155, 156, 134, 138,
	167, 238, 239, 155, 156, 240, 166, 433, 434, 435,
	436, 437, 351, 438, 439, 68, 213, 70, 491, 493,
	236, 71, 237, 73, 309, 74, 83, 349, 76, 77,
	78, 81, 114, 352, 249, 168, 52, 51, 264, 223,
	224, 225, 257, 241, 81, 168, 263, 53, 252, 226,
	54, 230, 508, 81, 294, 217, 218, 229, 295, 506,
	507, 400, 558, 305, 22, 386, 79, 82, 503, 289,
	291, 444, 297, 303, 82, 529, 307, 365, 276, 148,
	152, 561, 557, 162, 265, 555, 501, 211, 462, 475,
	164, 463, 464, 168, 149, 150, 151, 161, 141, 157,
	158, 242, 215, 128, 161, 306, 398, 158, 391, 279,
	280, 281, 276, 248, 216, 456, 219, 220, 221, 140,
	453, 160, 266, 387, 267, 266, 373, 311, 376, 510,
	505, 389, 425, 92, 91, 90, 256, 155, 156, 372,
	371, 492, 304, 495, 312, 398, 82, 459, 267, 266,
	82, 317, 215, 72, 510, 167, 89, 524, 336, 82,
	82, 166, 222, 215, 308, 523, 214, 377, 82, 342,
	348, 350, 347, 138, 363, 364, 366, 93, 94, 355,
	489, 369, 131, 361, 374, 375, 488, 378, 379, 380,
	381, 382, 383, 384, 385, 367, 368, 362, 487, 21,
	370, 485, 22, 26, 27, 28, 486, 98, 82, 390,
	117, 304, 390, 394, 390, 167, 214, 259, 366, 401,
	392, 166, 395, 261, 388, 532, 23, 214, 24, 310,
	25, 338, 483, 516, 313, 314, 399, 484, 161, 338,
	316, 158, 42, 43, 44, 45, 322, 323, 88, 543,
	115, 167, 167, 260, 422, 22, 426, 166, 424, 409,
	410, 411, 534, 535, 542, 421, 431, 418, 427, 413,
	541, 415, 296, 301, 304, 443, 499, 277, 278, 279,
	280, 281, 276, 361, 360, 448, 449, 360, 300, 359,
	299, 430, 359, 275, 274, 277, 278, 279, 280, 281,
	276, 452, 447, 414, 298, 46, 390, 402, 118, 442,
	496, 494, 478, 341, 454, 458, 275, 274, 277, 278,
	279, 280, 281, 276, 441, 167, 340, 460, 469, 477,
	116, 166, 403, 354, 471, 470, 341, 407, 408, 421,
	353, 468, 334, 445, 412, 254, 251, 481, 482, 247,
	127, 474, 553, 10, 544, 361, 361, 244, 497, 498,
	9, 212, 29, 500, 8, 7, 554, 6, 502, 169,
	245, 246, 504, 531, 275, 274, 277, 278, 279, 280,
	281, 276, 243, 120, 63, 515, 451, 167, 512, 343,
	255, 64, 511, 517, 5, 62, 61, 87, 67, 518,
	455, 275, 274, 277, 278, 279, 280, 281, 276, 450,
	4, 85, 528, 433, 434, 435, 436, 437, 530, 438,
	439, 465, 466, 467, 22, 66, 275, 274, 277, 278,
	279, 280, 281, 276, 390, 261, 480, 521, 472, 537,
	417, 65, 539, 420, 545, 515, 520, 536, 338, 538,
	319, 540, 548, 548, 548, 546, 318, 549, 550, 234,
	233, 22, 167, 232, 231, 559, 446, 562, 166, 228,
	227, 119, 563, 315, 564, 560, 551, 48, 320, 321,
	461, 344, 324, 325, 326, 327, 328, 329, 330, 331,
	332, 333, 148, 152, 335, 69, 162, 428, 102, 274,
	277, 278, 279, 280, 281, 276, 168, 149, 150, 151,
	346, 141, 157, 121, 22, 423, 552, 433, 434, 435,
	436, 437, 152, 438, 439, 162, 533, 522, 513, 519,
	152, 479, 140, 162, 160, 168, 149, 150, 151, 457,
	296, 157, 302, 168, 149, 150, 151, 146, 296, 157,
	155, 156, 96, 95, 97, 275, 274, 277, 278, 279,
	280, 281, 276, 160, 396, 147, 145, 159, 404, 405,
	406, 160, 473, 142, 268, 139, 490, 358, 432, 155,
	156, 356, 135, 440, 262, 132, 86, 155, 156, 41,
	20, 11, 19, 18, 17, 93, 94, 16, 15, 99,
	100, 14, 13, 12, 101, 103, 104, 105, 106, 108,
	109, 2, 110, 1, 112, 113, 22, 26, 27, 28,
	0, 82, 111, 0, 0, 0, 0, 107, 0, 0,
	0, 0, 0, 0, 0, 47, 0, 0, 0, 0,
	23, 0, 24, 30, 25, 0, 0, 0, 0, 0,
	82, 161, 0, 0, 158, 0, 0, 0, 82, 49,
	50, 55, 56, 57, 58, 59, 39, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	161, 0, 0, 158, 0, 0, 0, 0, 161, 0,
	0, 158, 0, 0, 0, 0, 0, 0, 35, 36,
	0, 37, 38, 171, 172, 173, 174, 175, 176, 177,
	178, 179, 180, 181, 182, 183, 184, 185, 186, 187,
	188, 189, 190, 191, 192, 193, 194, 195, 196, 197,
	198, 199, 200, 201, 202, 203, 204, 205, 206, 207,
	208, 209, 210, 270, 272, 0, 0, 0, 0, 282,
	283, 284, 285, 286, 287, 288, 273, 271, 269, 275,
	274, 277, 278, 279, 280, 281, 276, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 29, 40, 0, 0,
	0, 31, 32, 34, 33,
}
var yyPact = [...]int{

	701, -1000, -1000, 291, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 357, -1000, -1000, -24, -1000, -1000, -1000, -1000, -1000,
	287, -69, -63, -91, -56, -1000, 67, -1000, -1000, 100,
	-51, 546, 484, -1000, -1000, -1000, -1000, 469, -1000, 124,
	542, -1000, -29, -1000, -1000, 386, -151, 386, 552, 448,
	291, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -154, -102,
	100, -1000, -98, 100, -1000, 406, -157, 100, -157, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 8, -1000, 357, 91,
	430, 690, 690, -1000, -1000, 422, 182, 182, 0, 182,
	182, 243, -23, 551, 550, 1, -5, 545, 544, 541,
	540, -64, -1000, -13, -1000, -1000, 107, 447, 418, 386,
	386, 405, 143, 100, -1000, 402, -1000, -124, 401, 460,
	169, 100, 298, -1000, -1000, 109, 90, 180, 773, -1000,
	562, 149, -1000, -1000, -1000, 591, -1000, -1000, 355, -1000,
	-1000, -1000, -1000, 341, -1000, -1000, -1000, -1000, 339, 324,
	591, -1000, -1000, 256, 61, -1000, 129, -1000, 82, 690,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -55, 182, -1000, 591, 562, -1000, 182, 182, -1000,
	-1000, -1000, 386, 232, 537, 531, -1000, 386, 386, 182,
	182, 386, 386, 386, 386, 386, 386, 386, 386, 386,
	386, -1000, 398, 386, 101, 528, 387, -1000, 459, -162,
	-1000, 89, -1000, 396, -1000, -1000, 389, -1000, -1000, 343,
	8, 591, -1000, -1000, 100, 87, 562, 562, 591, 323,
	155, 591, 591, 197, 591, 591, 591, 591, 591, 591,
	591, 591, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	773, -28, 30, 38, 773, -1000, 599, -1000, 546, 15,
	591, 591, 133, 569, 101, 58, 591, 100, -1000, 362,
	-1000, 569, 180, -1000, -1000, 182, -1000, 386, 386, 386,
	182, 182, -1000, -1000, 528, 528, 528, 182, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 364, 319, 517, 562, 509,
	101, 101, -1000, -1000, 165, 100, -1000, -142, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 311, 456, 380, 340,
	77, -1000, -1000, 388, -1000, -1000, -1000, -1000, 153, 569,
	-1000, 323, 591, 591, 569, 440, -1000, 455, 289, 512,
	-1000, 119, 119, 85, 85, 85, -1000, -1000, 591, -1000,
	569, -1000, 27, 591, 415, 22, 172, -1000, 562, -1000,
	84, 569, -1000, -1000, 182, 182, 182, -1000, -1000, -1000,
	-1000, -1000, -1000, 509, 101, 517, 510, 514, 180, -1000,
	323, 291, 256, -4, -1000, 385, -1000, -1000, 368, -1000,
	515, 343, 343, -1000, -1000, 275, 244, 241, 229, 223,
	53, -1000, 367, 50, 366, 591, 591, -1000, 569, 307,
	591, -1000, 569, -1000, -7, 591, -1000, 72, -1000, 591,
	156, -1000, 54, 48, -1000, -1000, -1000, -1000, 162, 187,
	510, -1000, 591, 278, -1000, -1000, 101, -1000, -1000, 524,
	513, 456, 560, -1000, 208, -1000, 200, -1000, -1000, -1000,
	-1000, -106, -109, -119, -1000, -1000, -1000, 569, 569, 591,
	569, -1000, -18, -1000, 569, 591, -1000, -1000, -1000, -1000,
	437, -1000, -1000, 270, -1000, 330, 323, -1000, -1000, 517,
	562, 591, 562, -1000, -1000, 321, 315, 300, 569, -1000,
	569, 417, 591, -1000, -1000, -1000, -1000, 510, 180, 269,
	180, 100, 100, 100, 559, -1000, 426, -8, -1000, -11,
	-31, 101, -1000, 558, 96, -1000, 100, -1000, -1000, 256,
	-1000, 100, -1000, 100, -1000,
}
var yyPgo = [...]int{

	0, 703, 701, 15, 500, 484, 457, 455, 454, 450,
	443, 693, 692, 691, 688, 687, 684, 683, 682, 681,
	680, 725, 289, 679, 676, 243, 675, 26, 674, 673,
	672, 671, 5, 668, 667, 340, 666, 4, 24, 18,
	665, 664, 17, 663, 2, 20, 10, 662, 657, 13,
	656, 6, 655, 654, 8, 637, 632, 629, 621, 619,
	9, 618, 1, 616, 7, 606, 23, 605, 12, 3,
	21, 106, 193, 603, 600, 587, 585, 571, 0, 11,
	84, 570, 246, 567,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 7, 7, 10, 10, 8, 9,
	19, 19, 19, 19, 19, 19, 19, 19, 4, 4,
	4, 4, 4, 4, 15, 15, 16, 17, 18, 11,
	11, 11, 12, 12, 12, 13, 14, 14, 14, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 6, 20, 20, 83, 21, 22, 22,
	23, 23, 23, 23, 23, 24, 24, 26, 26, 27,
	27, 27, 30, 30, 28, 28, 28, 31, 31, 32,
	32, 32, 29, 29, 29, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 34, 34, 34, 35, 35, 36,
	36, 36, 36, 37, 37, 38, 38, 39, 39, 39,
	39, 39, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 41, 41, 41, 41, 41, 41, 41, 42,
	42, 47, 47, 45, 45, 49, 46, 46, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 55, 55, 55, 55, 55, 48, 48, 50,
	50, 50, 52, 56, 56, 53, 53, 54, 57, 57,
	51, 51, 43, 43, 43, 43, 58, 58, 59, 59,
	60, 60, 61, 61, 62, 63, 63, 63, 64, 64,
	64, 64, 65, 65, 65, 66, 66, 67, 67, 68,
	68, 69, 69, 70, 72, 72, 73, 73, 25, 25,
	74, 74, 74, 74, 74, 75, 75, 76, 76, 77,
	77, 78, 78, 79, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 81, 81, 81, 81, 82, 82,
	82, 71, 71, 71,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 5, 12, 3, 8, 8, 6, 6, 8, 7,
	2, 2, 2, 2, 2, 2, 2, 2, 4, 5,
	4, 4, 6, 7, 1, 2, 1, 1, 2, 5,
	8, 4, 6, 7, 4, 5, 4, 5, 5, 5,
	4, 4, 5, 5, 4, 4, 4, 6, 5, 7,
	6, 6, 7, 7, 5, 5, 6, 6, 6, 6,
	5, 5, 5, 5, 5, 5, 3, 4, 4, 2,
	3, 2, 2, 3, 1, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 1, 3, 3,
	2, 3, 3, 3, 4, 3, 4, 5, 6, 3,
	4, 2, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 3, 3, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	1, 1, 3, 4, 5, 6, 4, 1, 1, 1,
	1, 1, 5, 0, 1, 1, 2, 4, 0, 2,
	1, 3, 1, 1, 1, 1, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 1, 0, 1,
	1, 0, 2, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -19, -11, -12, -13, -14, -15, -16, -17, -18,
	-20, -22, 5, 29, 31, 33, 6, 7, 8, 165,
	32, 170, 171, 173, 172, 87, 88, 90, 91, 55,
	166, -23, 41, 42, 43, 44, 38, -21, -83, -21,
	-21, 151, 150, 161, 164, -21, -21, -21, -21, -21,
	-3, -7, -8, -10, -9, -4, -5, -6, 174, -76,
	176, 180, -25, 176, 178, 174, 174, 175, 176, 89,
	-78, 34, 149, 167, -3, 17, -24, 18, -22, -82,
	101, 100, 99, 143, 144, 101, 100, 102, -82, 147,
	148, 152, 46, 153, 154, 155, 156, 175, 157, 158,
	160, 170, 162, 163, 151, -35, 34, -25, -35, 9,
	25, -73, 179, 175, -78, 174, -78, 34, -72, 179,
	-78, -72, -26, -27, 80, -30, 34, -39, -44, -40,
	60, 39, -43, -51, -45, -50, -55, -52, 20, 35,
	36, 37, 21, -78, -49, 78, 79, 40, 182, -48,
	62, 179, 24, -69, 89, -70, -51, -78, 34, 29,
	-80, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	142, -80, 29, -71, 74, 10, -71, 145, 146, -71,
	-71, -71, 9, 152, 153, 154, 162, 9, 9, 146,
	146, 9, 9, 9, 9, 149, 174, 176, 155, 156,
	159, 146, 84, 25, 29, -35, -35, 34, 60, -78,
	-79, 34, -79, 177, 34, 20, 57, -78, -64, 9,
	45, 15, -28, -78, 19, 84, 59, 58, -41, 75,
	60, 74, 61, 73, 77, 76, 83, 78, 79, 80,
	81, 82, 66, 67, 68, 69, 70, 71, 72, -39,
	-44, -39, -46, -3, -44, -44, 39, -49, 39, 39,
	39, 39, -56, -44, 45, 92, 66, 84, -80, 169,
	-71, -44, -39, -71, -71, -35, -71, 9, 9, 9,
	-35, -35, -71, -71, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, 34, -35, -69, -38, 10, -66,
	29, 39, -79, 20, -77, 181, -74, 173, 171, 28,
	172, 13, 34, 34, 34, -79, -31, -32, -34, 39,
	34, -49, -27, -44, -78, 80, -78, -39, -39, -44,
	-45, 75, 74, 61, -44, -44, 21, 60, -44, -44,
	-44, -44, -44, -44, -44, -44, 183, 183, 45, 183,
	-44, 183, -46, 18, -44, -46, -53, -54, 63, -70,
	93, -44, 35, -71, -35, -35, -35, -71, -71, -38,
	-38, -38, -71, -66, 29, -38, -60, 13, -39, -42,
	24, -3, -69, -67, -51, 57, -78, -79, -75, 177,
	-38, 45, -33, 47, 48, 49, 50, 51, 53, 54,
	-29, 34, 19, -32, 84, 45, 168, -45, -44, -44,
	59, 21, -44, 183, -46, 75, 183, -57, -54, 65,
	-39, -81, 94, 97, 98, -71, -71, -71, -42, -69,
	-60, -64, 14, -47, -45, 183, 45, 34, 34, -58,
	11, -32, -32, 47, 52, 47, 52, 47, 47, 47,
	-36, 55, 178, 56, 34, 183, 34, -44, -44, 59,
	-44, 183, -44, 86, -44, 64, 95, 96, 94, -68,
	57, -68, -64, -61, -62, -44, 45, -51, -79, -59,
	12, 14, 57, 47, 47, 175, 175, 175, -44, 183,
	-44, 26, 45, -63, 22, 23, -45, -60, -39, -46,
	-39, 39, 39, 39, 27, -62, -64, -37, -78, -37,
	-37, 7, -65, 16, 30, 183, 45, 183, 183, -69,
	7, 75, -78, -78, -78,
}
var yyDef = [...]int{

	98, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 96, 96, 96, 96, 96, 96, 96, 96,
	0, 257, 248, 0, 0, 44, 0, 46, 47, 0,
	94, 0, 100, 102, 103, 104, 99, 105, 98, 308,
	308, 89, 0, 91, 92, 0, 248, 0, 0, 0,
	30, 31, 32, 33, 34, 35, 36, 37, 246, 0,
	0, 258, 0, 0, 249, 0, 244, 0, 244, 45,
	48, 261, 262, 95, 23, 101, 0, 106, 97, 0,
	0, 0, 0, 309, 310, 0, 311, 311, 0, 311,
	311, 311, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 90, 93, 137, 0, 0, 0,
	0, 0, 0, 0, 263, 0, 263, 0, 0, 0,
	0, 0, 228, 107, 109, 114, 261, 112, 113, 147,
	0, 0, 178, 179, 180, 0, 190, 191, 0, 212,
	213, 214, 215, 210, 174, 199, 200, 201, 0, 0,
	203, 197, 198, 38, 0, 241, 0, 210, 261, 0,
	40, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 282,
	283, 284, 285, 286, 287, 288, 289, 290, 291, 292,
	293, 294, 295, 296, 297, 298, 299, 300, 301, 302,
	303, 41, 311, 60, 0, 0, 61, 311, 311, 64,
	65, 66, 0, 311, 0, 0, 87, 0, 0, 311,
	311, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 88, 0, 0, 0, 145, 235, 263, 0, 259,
	51, 0, 54, 0, 56, 245, 0, 263, 21, 0,
	0, 0, 110, 115, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 162, 163, 164, 165, 166, 167, 168, 150,
	0, 0, 0, 0, 176, 189, 0, 161, 0, 0,
	0, 0, 0, 204, 0, 0, 0, 0, 39, 0,
	59, 312, 313, 62, 63, 311, 68, 0, 0, 0,
	311, 311, 74, 75, 145, 145, 145, 311, 80, 81,
	82, 83, 84, 85, 138, 235, 145, 220, 0, 0,
	0, 0, 49, 247, 0, 0, 263, 255, 250, 251,
	252, 253, 254, 55, 57, 58, 145, 117, 122, 0,
	134, 136, 108, 229, 116, 111, 211, 148, 149, 152,
	153, 0, 0, 0, 155, 0, 159, 0, 181, 182,
	183, 184, 185, 186, 187, 188, 151, 173, 0, 175,
	176, 192, 0, 0, 0, 0, 208, 205, 0, 242,
	0, 243, 42, 67, 311, 311, 311, 70, 71, 76,
	77, 78, 79, 0, 0, 220, 228, 0, 146, 26,
	0, 170, 27, 0, 237, 0, 260, 52, 0, 256,
	216, 0, 0, 125, 126, 0, 0, 0, 0, 0,
	139, 123, 0, 0, 0, 0, 0, 154, 156, 0,
	0, 160, 177, 193, 0, 0, 196, 0, 206, 0,
	0, 43, 0, 0, 307, 69, 72, 73, 239, 239,
	228, 29, 0, 169, 171, 236, 0, 263, 53, 218,
	0, 118, 0, 127, 0, 129, 0, 131, 132, 133,
	119, 0, 0, 0, 124, 120, 135, 230, 231, 0,
	157, 194, 0, 202, 209, 0, 304, 305, 306, 24,
	0, 25, 28, 221, 222, 225, 0, 238, 50, 220,
	0, 0, 0, 128, 130, 0, 0, 0, 158, 195,
	207, 0, 0, 224, 226, 227, 172, 228, 219, 217,
	121, 0, 0, 0, 0, 223, 232, 0, 143, 0,
	0, 0, 22, 0, 0, 140, 0, 141, 142, 240,
	233, 0, 144, 0, 234,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 82, 77, 3,
	39, 183, 80, 78, 45, 79, 84, 81, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	67, 66, 68, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 83, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 76, 3, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 41, 42, 43,
	44, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 69, 70, 71, 72, 73, 74, 75, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165, 166,
	167, 168, 169, 170, 171, 172, 173, 174, 175, 176,
	177, 178, 179, 180, 181, 182,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:238
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:244
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:246
		{
			yyVAL.statement = yyDollar[1].setStmt
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:248
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:250
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:266
		{
			yyVAL.statement = nil
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:270
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 22:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:274
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:278
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:284
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:288
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:300
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:304
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 28:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:316
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:322
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:328
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].selStmt}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:332
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:336
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:340
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:344
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:348
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].setStmt}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:352
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:356
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:362
		{
			yyVAL.setStmt = &SetVariable{
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				Exprs:    yyDollar[4].updateExprs,
			}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:370
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[5].bytes),
			}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:377
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[4].bytes),
			}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:384
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
			}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:391
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
				Collate:  string(yyDollar[6].bytes),
			}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:399
		{
			yyVAL.setStmt = &SetTransactionIsolationLevel{
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				IsolationLevel: string(yyDollar[7].bytes),
			}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:409
		{
			yyVAL.statement = &Begin{}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:413
		{
			yyVAL.statement = &Begin{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:419
		{
			yyVAL.statement = &Commit{}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:425
		{
			yyVAL.statement = &Rollback{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:431
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:437
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:441
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:446
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:452
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:456
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:461
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:467
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:473
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:477
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:482
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:488
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:492
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:496
		{
			yyVAL.showStmt = &ShowCollation{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:500
		{
			yyVAL.showStmt = &ShowVariables{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:504
		{
			yyVAL.showStmt = &ShowStatus{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:508
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:512
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:516
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:520
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:524
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:528
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 70:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:532
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:536
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 72:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:540
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 73:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:544
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:548
		{
			yyVAL.showStmt = &ShowProcedureStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:552
		{
			yyVAL.showStmt = &ShowFunctionStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:556
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:560
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:564
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 79:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:568
		{
			yyVAL.showStmt = &ShowTriggers{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:572
		{
			yyVAL.showStmt = &ShowCreateDatabase{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:576
		{
			yyVAL.showStmt = &ShowCreateTable{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:580
		{
			yyVAL.showStmt = &ShowCreateView{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:584
		{
			yyVAL.showStmt = &ShowCreateProcedure{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:588
		{
			yyVAL.showStmt = &ShowCreateFunction{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:592
		{
			yyVAL.showStmt = &ShowCreateTrigger{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:596
		{
			yyVAL.showStmt = &ShowProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:600
		{
			yyVAL.showStmt = &ShowFullProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:604
		{
			yyVAL.showStmt = &ShowSlaveStatus{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:608
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:612
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:616
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:620
		{
			yyVAL.showStmt = &ShowProfiles{}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:626
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[3].tableName}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:633
		{
			yyVAL.statement = nil
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:635
		{
			yyVAL.statement = nil
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:638
		{
			SetAllowComments(yylex, true)
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:642
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:648
		{
			yyVAL.bytes2 = nil
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:652
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:658
		{
			yyVAL.str = AST_UNION
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:662
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:666
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:670
		{
			yyVAL.str = AST_EXCEPT
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:674
		{
			yyVAL.str = AST_INTERSECT
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:679
		{
			yyVAL.str = ""
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:683
		{
			yyVAL.str = AST_DISTINCT
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:689
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:693
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:699
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:703
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:707
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:713
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:717
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:722
		{
			yyVAL.bytes = nil
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:726
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:730
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:736
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:740
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:746
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:750
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:754
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:759
		{
			yyVAL.bytes = nil
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:763
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:767
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:773
		{
			yyVAL.str = AST_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:777
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:781
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:785
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:789
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:793
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:797
		{
			yyVAL.str = AST_JOIN
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:801
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:805
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:811
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:815
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:819
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:825
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:829
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 139:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:834
		{
			yyVAL.indexHints = nil
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:838
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 141:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:842
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 142:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:846
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:852
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:856
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 145:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:861
		{
			yyVAL.boolExpr = nil
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:865
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:872
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:876
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:880
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:884
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:890
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:894
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:898
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:902
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:906
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:910
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:914
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:918
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:922
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:926
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:932
		{
			yyVAL.str = AST_EQ
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:936
		{
			yyVAL.str = AST_LT
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:940
		{
			yyVAL.str = AST_GT
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:944
		{
			yyVAL.str = AST_LE
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:948
		{
			yyVAL.str = AST_GE
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:952
		{
			yyVAL.str = AST_NE
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:956
		{
			yyVAL.str = AST_NSE
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:962
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:966
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:972
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:976
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:982
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:986
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:992
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:998
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1002
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1008
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1012
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1016
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1020
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1024
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1028
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1032
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1036
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1040
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1044
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1048
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1052
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1067
		{
			yyVAL.valExpr = yyDollar[1].funcExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1071
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1077
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1081
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].valExprs}
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1085
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].valExprs}
		}
	case 195:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:1089
		{
			yyVAL.funcExpr = &FuncExpr{Name: []byte("locate"), Exprs: ValExprs{yyDollar[3].valExpr, yyDollar[5].valExpr}}
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1093
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].valExprs}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1099
		{
			yyVAL.bytes = IF_BYTES
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1103
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1109
		{
			yyVAL.byt = AST_UPLUS
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1113
		{
			yyVAL.byt = AST_UMINUS
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1117
		{
			yyVAL.byt = AST_TILDA
		}
	case 202:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1123
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1128
		{
			yyVAL.valExpr = nil
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1132
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1138
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1142
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 207:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1148
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1153
		{
			yyVAL.valExpr = nil
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1157
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1163
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1167
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1173
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1177
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1181
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1185
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1190
		{
			yyVAL.valExprs = nil
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1194
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1199
		{
			yyVAL.boolExpr = nil
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1203
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1208
		{
			yyVAL.orderBy = nil
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1212
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1218
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1222
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 224:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1228
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1233
		{
			yyVAL.str = AST_ASC
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1237
		{
			yyVAL.str = AST_ASC
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1241
		{
			yyVAL.str = AST_DESC
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1246
		{
			yyVAL.limit = nil
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1250
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 230:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1254
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1258
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1263
		{
			yyVAL.str = ""
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1267
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 234:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1271
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1284
		{
			yyVAL.columns = nil
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1288
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1294
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1298
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1303
		{
			yyVAL.updateExprs = nil
		}
	case 240:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1307
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1313
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1317
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 243:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1323
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1328
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1330
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1333
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1335
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1338
		{
			yyVAL.str = ""
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1340
		{
			yyVAL.str = AST_IGNORE
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1344
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1346
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1348
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1350
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1352
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1355
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1357
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1360
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1362
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1365
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1367
		{
			yyVAL.empty = struct{}{}
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1371
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1375
		{
			yyVAL.bytes = []byte("database")
		}
	case 263:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1380
		{
			ForceEOF(yylex)
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1386
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1388
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1390
		{
			yyVAL.bytes = []byte("big5")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1392
		{
			yyVAL.bytes = []byte("binary")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1394
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1396
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1398
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1400
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1402
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1404
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1406
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1408
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1410
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1412
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1414
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1416
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1418
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1420
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1422
		{
			yyVAL.bytes = []byte("greek")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1424
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1426
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1428
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1430
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1432
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1434
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1436
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1438
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1440
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1442
		{
			yyVAL.bytes = []byte("macce")
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1444
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1446
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1448
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1450
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1452
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 298:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1454
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1456
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 300:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1458
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 301:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1460
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1462
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 303:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1464
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 304:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1468
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 305:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1470
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 306:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1472
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 307:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1474
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 308:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1477
		{
			yyVAL.bytes = nil
		}
	case 309:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1479
		{
			yyVAL.bytes = []byte("session")
		}
	case 310:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1481
		{
			yyVAL.bytes = []byte("global")
		}
	case 311:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1484
		{
			yyVAL.expr = nil
		}
	case 312:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1486
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 313:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1490
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
