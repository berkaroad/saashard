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
const PROCEDURE = 57481
const FUNCTION = 57482
const INDEXES = 57483
const KEYS = 57484
const TRIGGER = 57485
const TRIGGERS = 57486
const PLUGINS = 57487
const PROCESSLIST = 57488
const SLAVE = 57489
const PROFILES = 57490
const REPLACE = 57491
const ADMIN = 57492
const HELP = 57493
const OFFSET = 57494
const COLLATE = 57495
const CREATE = 57496
const ALTER = 57497
const DROP = 57498
const RENAME = 57499
const TABLE = 57500
const INDEX = 57501
const VIEW = 57502
const TO = 57503
const IGNORE = 57504
const IF = 57505
const UNIQUE = 57506
const USING = 57507

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

const yyNprod = 311
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 860

var yyAct = [...]int{

	145, 500, 285, 160, 532, 287, 142, 253, 388, 405,
	245, 495, 153, 408, 349, 143, 131, 331, 288, 3,
	162, 262, 261, 541, 329, 132, 337, 541, 128, 136,
	121, 541, 422, 423, 424, 425, 426, 255, 427, 428,
	80, 463, 42, 43, 44, 45, 73, 74, 74, 60,
	255, 255, 380, 68, 418, 70, 513, 248, 343, 71,
	84, 76, 77, 78, 385, 512, 148, 152, 478, 480,
	158, 123, 511, 341, 125, 122, 124, 75, 129, 344,
	135, 149, 150, 151, 303, 140, 156, 83, 81, 137,
	164, 210, 52, 51, 230, 259, 163, 148, 152, 233,
	234, 158, 53, 235, 167, 54, 139, 113, 159, 494,
	81, 135, 149, 150, 151, 391, 140, 156, 231, 236,
	232, 220, 221, 244, 154, 155, 133, 225, 299, 81,
	222, 252, 165, 224, 357, 258, 247, 139, 79, 159,
	271, 165, 433, 289, 378, 214, 215, 290, 489, 22,
	450, 301, 260, 451, 452, 154, 155, 133, 237, 543,
	127, 294, 297, 542, 148, 152, 482, 540, 158, 284,
	286, 546, 300, 488, 92, 91, 90, 462, 165, 149,
	150, 151, 381, 140, 156, 389, 444, 442, 379, 213,
	479, 216, 217, 218, 243, 82, 161, 208, 434, 492,
	493, 212, 368, 82, 139, 261, 159, 269, 272, 273,
	274, 275, 276, 271, 305, 340, 342, 339, 93, 94,
	311, 212, 154, 155, 157, 82, 82, 383, 508, 270,
	269, 272, 273, 274, 275, 276, 271, 262, 261, 130,
	164, 369, 306, 328, 82, 298, 163, 82, 274, 275,
	276, 271, 389, 334, 447, 157, 82, 496, 137, 355,
	356, 358, 496, 347, 414, 211, 361, 353, 510, 366,
	367, 302, 370, 371, 372, 373, 374, 375, 376, 377,
	362, 354, 219, 212, 365, 211, 251, 22, 26, 27,
	28, 359, 360, 82, 382, 137, 137, 364, 363, 164,
	509, 304, 358, 392, 72, 163, 307, 308, 486, 384,
	386, 23, 310, 24, 89, 25, 314, 315, 472, 390,
	435, 476, 157, 473, 114, 270, 269, 272, 273, 274,
	275, 276, 271, 164, 164, 21, 411, 475, 415, 163,
	413, 398, 399, 400, 474, 402, 470, 211, 298, 416,
	410, 471, 330, 404, 272, 273, 274, 275, 276, 271,
	407, 116, 380, 517, 353, 98, 432, 437, 438, 262,
	261, 330, 254, 419, 502, 491, 528, 527, 256, 436,
	352, 46, 117, 441, 88, 351, 22, 420, 137, 42,
	43, 44, 45, 403, 526, 291, 446, 519, 520, 332,
	393, 394, 443, 333, 164, 397, 298, 456, 255, 333,
	163, 401, 295, 458, 457, 352, 455, 293, 292, 448,
	351, 410, 483, 481, 431, 461, 465, 464, 115, 346,
	345, 326, 249, 353, 353, 468, 469, 484, 485, 430,
	246, 242, 487, 240, 241, 126, 29, 538, 529, 239,
	490, 270, 269, 272, 273, 274, 275, 276, 271, 439,
	10, 539, 501, 209, 164, 498, 166, 516, 497, 238,
	503, 119, 9, 8, 440, 504, 270, 269, 272, 273,
	274, 275, 276, 271, 335, 250, 7, 453, 454, 514,
	6, 63, 5, 4, 515, 270, 269, 272, 273, 274,
	275, 276, 271, 64, 62, 422, 423, 424, 425, 426,
	382, 427, 428, 524, 87, 522, 102, 61, 521, 530,
	501, 67, 22, 66, 65, 85, 256, 533, 533, 533,
	531, 406, 534, 535, 507, 459, 523, 164, 525, 506,
	544, 409, 547, 163, 309, 467, 330, 548, 313, 549,
	545, 316, 317, 318, 319, 320, 321, 322, 323, 324,
	325, 148, 152, 327, 312, 158, 229, 228, 227, 226,
	96, 95, 97, 223, 118, 165, 149, 150, 151, 22,
	140, 156, 536, 48, 449, 336, 69, 417, 338, 120,
	412, 537, 518, 152, 499, 505, 158, 466, 445, 296,
	387, 139, 147, 159, 144, 146, 165, 149, 150, 151,
	460, 291, 156, 93, 94, 141, 263, 99, 100, 154,
	155, 138, 101, 103, 104, 105, 107, 108, 477, 109,
	22, 111, 112, 350, 159, 421, 395, 396, 348, 110,
	134, 429, 47, 257, 106, 86, 152, 41, 20, 158,
	154, 155, 22, 26, 27, 28, 11, 19, 18, 165,
	149, 150, 151, 17, 291, 156, 49, 50, 55, 56,
	57, 58, 59, 16, 15, 14, 23, 13, 24, 30,
	25, 12, 2, 1, 0, 0, 0, 159, 0, 0,
	82, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 39, 154, 155, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 157,
	0, 82, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 0, 37, 38, 0,
	0, 0, 0, 0, 0, 265, 267, 0, 0, 0,
	157, 277, 278, 279, 280, 281, 282, 283, 268, 266,
	264, 270, 269, 272, 273, 274, 275, 276, 271, 0,
	0, 0, 0, 0, 82, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 157, 0, 0, 0, 0, 0, 0,
	0, 29, 40, 0, 0, 0, 31, 32, 34, 33,
	168, 169, 170, 171, 172, 173, 174, 175, 176, 177,
	178, 179, 180, 181, 182, 183, 184, 185, 186, 187,
	188, 189, 190, 191, 192, 193, 194, 195, 196, 197,
	198, 199, 200, 201, 202, 203, 204, 205, 206, 207,
}
var yyPact = [...]int{

	647, -1000, -1000, 348, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 343, -1000, -1000, -58, -1000, -1000, -1000, -1000, -1000,
	282, -120, -129, -96, -112, -1000, 49, -1000, -1000, 95,
	-79, 574, 508, -1000, -1000, -1000, -1000, 496, -1000, 75,
	470, -1000, -44, -1000, -1000, 394, -130, 394, 565, 446,
	348, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -148, -99,
	95, -1000, -97, 95, -1000, 411, -150, 95, -150, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 77, -1000, 343, 107,
	437, 717, 717, -1000, -1000, 434, 191, 191, 0, 191,
	191, 273, -31, 564, -13, -19, 560, 559, 558, 557,
	-55, -1000, -27, -1000, -1000, 74, 444, 420, 394, 394,
	407, 134, 95, -1000, 406, -1000, -119, 398, 465, 229,
	95, 363, -1000, -1000, 76, 68, 179, 685, -1000, 541,
	144, -1000, -1000, -1000, 572, 379, 378, -1000, 373, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 572,
	303, 36, -1000, 106, -1000, 67, 717, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -84, 191,
	-1000, 572, 541, -1000, 191, 191, -1000, -1000, -1000, 394,
	211, 555, -1000, 394, 191, 191, 394, 394, 394, 394,
	394, 394, 394, 394, 394, 394, -1000, 397, 394, 98,
	536, 370, -1000, 464, -154, -1000, 45, -1000, 396, -1000,
	-1000, 395, -1000, -1000, 346, 77, 572, -1000, -1000, 95,
	54, 541, 541, 572, 356, 223, 572, 572, 181, 572,
	572, 572, 572, 572, 572, 572, 572, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 685, -37, 7, 1, 685,
	-1000, 625, 46, 77, -1000, 574, 122, 419, 98, 22,
	572, 95, -1000, 365, -1000, 419, 179, -1000, -1000, 191,
	-1000, 394, 394, 191, -1000, -1000, 536, 536, 536, 191,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 364, 361, 518,
	541, 517, 98, 98, -1000, -1000, 207, 95, -1000, -122,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 342, 458,
	405, 381, 58, -1000, -1000, 153, -1000, -1000, -1000, -1000,
	146, 419, -1000, 356, 572, 572, 419, 400, -1000, 453,
	276, 130, -1000, 168, 168, 57, 57, 57, -1000, -1000,
	572, -1000, 419, -1000, 6, 77, 5, 189, -1000, 541,
	-1000, 56, 419, -1000, -1000, 191, 191, -1000, -1000, -1000,
	-1000, -1000, 517, 98, 518, 511, 521, 179, -1000, 356,
	348, 303, -4, -1000, 393, -1000, -1000, 392, -1000, 534,
	346, 346, -1000, -1000, 299, 271, 297, 290, 274, 13,
	-1000, 389, -15, 388, 572, 572, -1000, 419, 249, 572,
	-1000, 419, -1000, -8, -1000, 62, -1000, 572, 311, -1000,
	104, 15, -1000, -1000, -1000, 205, 200, 511, -1000, 572,
	329, -1000, -1000, 98, -1000, -1000, 527, 520, 458, 171,
	-1000, 253, -1000, 221, -1000, -1000, -1000, -1000, -102, -109,
	-118, -1000, -1000, -1000, 419, 419, 572, 419, -1000, -1000,
	419, 572, -1000, -1000, -1000, -1000, 441, -1000, -1000, 318,
	-1000, 375, 356, -1000, -1000, 518, 541, 572, 541, -1000,
	-1000, 355, 338, 337, 419, 419, 421, 572, -1000, -1000,
	-1000, -1000, 511, 179, 317, 179, 95, 95, 95, 575,
	-1000, 431, -14, -1000, -18, -22, 98, -1000, 543, 96,
	-1000, 95, -1000, -1000, 303, -1000, 95, -1000, 95, -1000,
}
var yyPgo = [...]int{

	0, 683, 682, 18, 493, 492, 490, 486, 473, 472,
	460, 681, 677, 675, 674, 673, 663, 658, 657, 656,
	648, 642, 335, 647, 645, 304, 16, 25, 643, 641,
	640, 638, 14, 635, 633, 324, 628, 4, 24, 29,
	621, 616, 13, 615, 2, 15, 5, 610, 605, 12,
	604, 6, 602, 600, 8, 599, 598, 597, 595, 9,
	594, 1, 592, 7, 591, 17, 590, 11, 3, 20,
	91, 160, 589, 588, 587, 586, 585, 0, 10, 104,
	584, 314, 583,
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
	5, 6, 20, 20, 82, 21, 22, 22, 23, 23,
	23, 23, 23, 24, 24, 26, 26, 27, 27, 27,
	30, 30, 28, 28, 28, 31, 31, 32, 32, 32,
	32, 29, 29, 29, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 34, 34, 34, 35, 35, 36, 36,
	36, 36, 37, 37, 38, 38, 39, 39, 39, 39,
	39, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 41, 41, 41, 41, 41, 41, 41, 42, 42,
	47, 47, 45, 45, 49, 46, 46, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 48, 48, 50, 50, 50, 52,
	55, 55, 53, 53, 54, 56, 56, 51, 51, 43,
	43, 43, 43, 57, 57, 58, 58, 59, 59, 60,
	60, 61, 62, 62, 62, 63, 63, 63, 63, 64,
	64, 64, 65, 65, 66, 66, 67, 67, 68, 68,
	69, 71, 71, 72, 72, 25, 25, 73, 73, 73,
	73, 73, 74, 74, 75, 75, 76, 76, 77, 77,
	78, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 80, 80, 80, 80, 81, 81, 81, 70, 70,
	70,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 5, 12, 3, 8, 8, 6, 6, 8, 7,
	2, 2, 2, 2, 2, 2, 2, 2, 4, 5,
	4, 4, 6, 7, 1, 2, 1, 1, 2, 5,
	8, 4, 6, 7, 4, 5, 4, 5, 5, 5,
	4, 4, 5, 5, 4, 4, 4, 6, 5, 7,
	6, 7, 5, 5, 6, 6, 6, 6, 5, 5,
	5, 5, 5, 5, 3, 4, 4, 2, 3, 2,
	2, 3, 1, 2, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 1, 3, 3, 2,
	3, 3, 3, 4, 3, 4, 5, 6, 3, 4,
	2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 3, 3, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	4, 5, 4, 1, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 1, 3,
	3, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 1, 0, 2, 1, 1,
	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 1, 0, 1, 1, 0, 2,
	2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -19, -11, -12, -13, -14, -15, -16, -17, -18,
	-20, -22, 5, 29, 31, 33, 6, 7, 8, 164,
	32, 169, 170, 172, 171, 87, 88, 90, 91, 55,
	165, -23, 41, 42, 43, 44, 38, -21, -82, -21,
	-21, 151, 150, 160, 163, -21, -21, -21, -21, -21,
	-3, -7, -8, -10, -9, -4, -5, -6, 173, -75,
	175, 179, -25, 175, 177, 173, 173, 174, 175, 89,
	-77, 34, 149, 166, -3, 17, -24, 18, -22, -81,
	101, 100, 99, 143, 144, 101, 100, 102, -81, 147,
	148, 152, 46, 153, 154, 155, 174, 156, 157, 159,
	169, 161, 162, 151, -35, 34, -25, -35, 9, 25,
	-72, 178, 174, -77, 173, -77, 34, -71, 178, -77,
	-71, -26, -27, 80, -30, 34, -39, -44, -40, 60,
	39, -43, -51, -45, -50, -77, -48, -52, 20, 35,
	36, 37, 21, -49, 78, 79, 40, 178, 24, 62,
	-68, 89, -69, -51, -77, 34, 29, -79, 103, 104,
	105, 106, 107, 108, 109, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 131, 132, 133, 134,
	135, 136, 137, 138, 139, 140, 141, 142, -79, 29,
	-70, 74, 10, -70, 145, 146, -70, -70, -70, 9,
	152, 153, 161, 9, 146, 146, 9, 9, 9, 9,
	149, 173, 175, 154, 155, 158, 146, 84, 25, 29,
	-35, -35, 34, 60, -77, -78, 34, -78, 176, 34,
	20, 57, -77, -63, 9, 45, 15, -28, -77, 19,
	84, 59, 58, -41, 75, 60, 74, 61, 73, 77,
	76, 83, 78, 79, 80, 81, 82, 66, 67, 68,
	69, 70, 71, 72, -39, -44, -39, -46, -3, -44,
	-44, 39, 39, 39, -49, 39, -55, -44, 45, 92,
	66, 84, -79, 168, -70, -44, -39, -70, -70, -35,
	-70, 9, 9, -35, -70, -70, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, 34, -35, -68, -38,
	10, -65, 29, 39, -78, 20, -76, 180, -73, 172,
	170, 28, 171, 13, 34, 34, 34, -78, -31, -32,
	-34, 39, 34, -49, -27, -44, -77, 80, -77, -39,
	-39, -44, -45, 75, 74, 61, -44, -44, 21, 60,
	-44, -44, -44, -44, -44, -44, -44, -44, 181, 181,
	45, 181, -44, 181, -26, 18, -26, -53, -54, 63,
	-69, 93, -44, 35, -70, -35, -35, -70, -38, -38,
	-38, -70, -65, 29, -38, -59, 13, -39, -42, 24,
	-3, -68, -66, -51, 57, -77, -78, -74, 176, -38,
	45, -33, 47, 48, 49, 50, 51, 53, 54, -29,
	34, 19, -32, 84, 45, 167, -45, -44, -44, 59,
	21, -44, 181, -26, 181, -56, -54, 65, -39, -80,
	94, 97, 98, -70, -70, -42, -68, -59, -63, 14,
	-47, -45, 181, 45, 34, 34, -57, 11, -32, -32,
	47, 52, 47, 52, 47, 47, 47, -36, 55, 177,
	56, 34, 181, 34, -44, -44, 59, -44, 181, 86,
	-44, 64, 95, 96, 94, -67, 57, -67, -63, -60,
	-61, -44, 45, -51, -78, -58, 12, 14, 57, 47,
	47, 174, 174, 174, -44, -44, 26, 45, -62, 22,
	23, -45, -59, -39, -46, -39, 39, 39, 39, 27,
	-61, -63, -37, -77, -37, -37, 7, -64, 16, 30,
	181, 45, 181, 181, -68, 7, 75, -77, -77, -77,
}
var yyDef = [...]int{

	96, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 94, 94, 94, 94, 94, 94, 94, 94,
	0, 254, 245, 0, 0, 44, 0, 46, 47, 0,
	92, 0, 98, 100, 101, 102, 97, 103, 96, 305,
	305, 87, 0, 89, 90, 0, 245, 0, 0, 0,
	30, 31, 32, 33, 34, 35, 36, 37, 243, 0,
	0, 255, 0, 0, 246, 0, 241, 0, 241, 45,
	48, 258, 259, 93, 23, 99, 0, 104, 95, 0,
	0, 0, 0, 306, 307, 0, 308, 308, 0, 308,
	308, 308, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 84, 0, 88, 91, 136, 0, 0, 0, 0,
	0, 0, 0, 260, 0, 260, 0, 0, 0, 0,
	0, 225, 105, 107, 112, 258, 110, 111, 146, 0,
	0, 177, 178, 179, 0, 207, 0, 193, 0, 209,
	210, 211, 212, 173, 196, 197, 198, 194, 195, 200,
	38, 0, 238, 0, 207, 258, 0, 40, 261, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 282,
	283, 284, 285, 286, 287, 288, 289, 290, 291, 292,
	293, 294, 295, 296, 297, 298, 299, 300, 41, 308,
	60, 0, 0, 61, 308, 308, 64, 65, 66, 0,
	308, 0, 85, 0, 308, 308, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 86, 0, 0, 0,
	144, 232, 260, 0, 256, 51, 0, 54, 0, 56,
	242, 0, 260, 21, 0, 0, 0, 108, 113, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 161, 162, 163,
	164, 165, 166, 167, 149, 0, 0, 0, 0, 175,
	188, 0, 0, 0, 160, 0, 0, 201, 0, 0,
	0, 0, 39, 0, 59, 309, 310, 62, 63, 308,
	68, 0, 0, 308, 72, 73, 144, 144, 144, 308,
	78, 79, 80, 81, 82, 83, 137, 232, 144, 217,
	0, 0, 0, 0, 49, 244, 0, 0, 260, 252,
	247, 248, 249, 250, 251, 55, 57, 58, 144, 115,
	121, 0, 133, 135, 106, 226, 114, 109, 208, 147,
	148, 151, 152, 0, 0, 0, 154, 0, 158, 0,
	180, 181, 182, 183, 184, 185, 186, 187, 150, 172,
	0, 174, 175, 189, 0, 0, 0, 205, 202, 0,
	239, 0, 240, 42, 67, 308, 308, 70, 74, 75,
	76, 77, 0, 0, 217, 225, 0, 145, 26, 0,
	169, 27, 0, 234, 0, 257, 52, 0, 253, 213,
	0, 0, 124, 125, 0, 0, 0, 0, 0, 138,
	122, 0, 0, 0, 0, 0, 153, 155, 0, 0,
	159, 176, 190, 0, 192, 0, 203, 0, 0, 43,
	0, 0, 304, 69, 71, 236, 236, 225, 29, 0,
	168, 170, 233, 0, 260, 53, 215, 0, 116, 119,
	126, 0, 128, 0, 130, 131, 132, 117, 0, 0,
	0, 123, 118, 134, 227, 228, 0, 156, 191, 199,
	206, 0, 301, 302, 303, 24, 0, 25, 28, 218,
	219, 222, 0, 235, 50, 217, 0, 0, 0, 127,
	129, 0, 0, 0, 157, 204, 0, 0, 221, 223,
	224, 171, 225, 216, 214, 120, 0, 0, 0, 0,
	220, 229, 0, 142, 0, 0, 0, 22, 0, 0,
	139, 0, 140, 141, 237, 230, 0, 143, 0, 231,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 82, 77, 3,
	39, 181, 80, 78, 45, 79, 84, 81, 3, 3,
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
	177, 178, 179, 180,
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
		//line ./yacc.y:233
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:239
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:241
		{
			yyVAL.statement = yyDollar[1].setStmt
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:243
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:245
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:261
		{
			yyVAL.statement = nil
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:265
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 22:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:269
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:273
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:279
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:283
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
		//line ./yacc.y:295
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:299
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
		//line ./yacc.y:311
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:317
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:323
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].selStmt}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:327
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:331
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:335
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:339
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:343
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].setStmt}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:347
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:351
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:357
		{
			yyVAL.setStmt = &SetVariable{
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				Exprs:    yyDollar[4].updateExprs,
			}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:365
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[5].bytes),
			}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:372
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[4].bytes),
			}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:379
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
			}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:386
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
				Collate:  string(yyDollar[6].bytes),
			}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:394
		{
			yyVAL.setStmt = &SetTransactionIsolationLevel{
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				IsolationLevel: string(yyDollar[7].bytes),
			}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:404
		{
			yyVAL.statement = &Begin{}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:408
		{
			yyVAL.statement = &Begin{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:414
		{
			yyVAL.statement = &Commit{}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:420
		{
			yyVAL.statement = &Rollback{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:426
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:432
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:436
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:441
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:447
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:451
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:456
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:462
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:468
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:472
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:477
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:483
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:487
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:491
		{
			yyVAL.showStmt = &ShowCollation{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:495
		{
			yyVAL.showStmt = &ShowVariables{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:499
		{
			yyVAL.showStmt = &ShowStatus{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:503
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:507
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:511
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:515
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:519
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:523
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 70:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:527
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:531
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:535
		{
			yyVAL.showStmt = &ShowProcedureStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:539
		{
			yyVAL.showStmt = &ShowFunctionStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:543
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:547
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:551
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:555
		{
			yyVAL.showStmt = &ShowTriggers{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:559
		{
			yyVAL.showStmt = &ShowCreateDatabase{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:563
		{
			yyVAL.showStmt = &ShowCreateTable{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:567
		{
			yyVAL.showStmt = &ShowCreateView{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:571
		{
			yyVAL.showStmt = &ShowCreateProcedure{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:575
		{
			yyVAL.showStmt = &ShowCreateFunction{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:579
		{
			yyVAL.showStmt = &ShowCreateTrigger{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:583
		{
			yyVAL.showStmt = &ShowProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:587
		{
			yyVAL.showStmt = &ShowFullProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:591
		{
			yyVAL.showStmt = &ShowSlaveStatus{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:595
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:599
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:603
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:607
		{
			yyVAL.showStmt = &ShowProfiles{}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:613
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[3].tableName}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:620
		{
			yyVAL.statement = nil
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:622
		{
			yyVAL.statement = nil
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:625
		{
			SetAllowComments(yylex, true)
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:629
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:635
		{
			yyVAL.bytes2 = nil
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:639
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:645
		{
			yyVAL.str = AST_UNION
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:649
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:653
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:657
		{
			yyVAL.str = AST_EXCEPT
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:661
		{
			yyVAL.str = AST_INTERSECT
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:666
		{
			yyVAL.str = ""
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:670
		{
			yyVAL.str = AST_DISTINCT
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:676
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:680
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:686
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:690
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:694
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:700
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:704
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 112:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:709
		{
			yyVAL.bytes = nil
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:713
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:717
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:723
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:727
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:733
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:737
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:741
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:745
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 121:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:750
		{
			yyVAL.bytes = nil
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:754
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:758
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:764
		{
			yyVAL.str = AST_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:768
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:772
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:776
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:780
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:784
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:788
		{
			yyVAL.str = AST_JOIN
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:792
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:796
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:802
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:806
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:810
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:816
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:820
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 138:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:825
		{
			yyVAL.indexHints = nil
		}
	case 139:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:829
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:833
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 141:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:837
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:843
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:847
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 144:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:852
		{
			yyVAL.boolExpr = nil
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:856
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:863
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:867
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:871
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:875
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:881
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:885
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:889
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:893
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:897
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:901
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:905
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:909
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:913
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:917
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:923
		{
			yyVAL.str = AST_EQ
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:927
		{
			yyVAL.str = AST_LT
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:931
		{
			yyVAL.str = AST_GT
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:935
		{
			yyVAL.str = AST_LE
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:939
		{
			yyVAL.str = AST_GE
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:943
		{
			yyVAL.str = AST_NE
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:947
		{
			yyVAL.str = AST_NSE
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:953
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:957
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:963
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:967
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:973
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:977
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:983
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:989
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:993
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:999
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1003
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1007
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1011
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1015
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1019
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1023
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1027
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1031
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1035
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1039
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1043
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
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1058
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 190:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1062
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 191:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1066
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 192:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1070
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1074
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1080
		{
			yyVAL.bytes = IF_BYTES
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1084
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1090
		{
			yyVAL.byt = AST_UPLUS
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1094
		{
			yyVAL.byt = AST_UMINUS
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1098
		{
			yyVAL.byt = AST_TILDA
		}
	case 199:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1104
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1109
		{
			yyVAL.valExpr = nil
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1113
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1119
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1123
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 204:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1129
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1134
		{
			yyVAL.valExpr = nil
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1138
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1144
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1148
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1154
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1158
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1162
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1166
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1171
		{
			yyVAL.valExprs = nil
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1175
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1180
		{
			yyVAL.boolExpr = nil
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1184
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1189
		{
			yyVAL.orderBy = nil
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1193
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1199
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1203
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1209
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1214
		{
			yyVAL.str = AST_ASC
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1218
		{
			yyVAL.str = AST_ASC
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1222
		{
			yyVAL.str = AST_DESC
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1227
		{
			yyVAL.limit = nil
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1231
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1235
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 228:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1239
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 229:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1244
		{
			yyVAL.str = ""
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1248
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 231:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1252
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
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1265
		{
			yyVAL.columns = nil
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1269
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1275
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1279
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1284
		{
			yyVAL.updateExprs = nil
		}
	case 237:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1288
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1294
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1298
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1304
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 241:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1309
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1311
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1314
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1316
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1319
		{
			yyVAL.str = ""
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1321
		{
			yyVAL.str = AST_IGNORE
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1325
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1327
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1329
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1331
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1333
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1336
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1338
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1341
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1343
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1346
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1348
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1352
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1356
		{
			yyVAL.bytes = []byte("database")
		}
	case 260:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1361
		{
			ForceEOF(yylex)
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1367
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1369
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1371
		{
			yyVAL.bytes = []byte("big5")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1373
		{
			yyVAL.bytes = []byte("binary")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1375
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1377
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1379
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1381
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1383
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1385
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1387
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1389
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1391
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1393
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1395
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1397
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1399
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1401
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1403
		{
			yyVAL.bytes = []byte("greek")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1405
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1407
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1409
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1411
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1413
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1415
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1417
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1419
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1421
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1423
		{
			yyVAL.bytes = []byte("macce")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1425
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1427
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1429
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1431
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1433
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1435
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1437
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1439
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 298:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1441
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1443
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 300:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1445
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 301:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1449
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 302:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1451
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 303:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1453
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 304:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1455
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 305:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1458
		{
			yyVAL.bytes = nil
		}
	case 306:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1460
		{
			yyVAL.bytes = []byte("session")
		}
	case 307:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1462
		{
			yyVAL.bytes = []byte("global")
		}
	case 308:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1465
		{
			yyVAL.expr = nil
		}
	case 309:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1467
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 310:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1471
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
