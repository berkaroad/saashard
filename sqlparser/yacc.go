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

const yyNprod = 313
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 843

var yyAct = [...]int{

	153, 509, 289, 162, 541, 291, 143, 257, 355, 413,
	249, 504, 154, 292, 3, 144, 416, 394, 337, 266,
	265, 550, 132, 550, 335, 133, 550, 137, 259, 343,
	164, 430, 431, 432, 433, 434, 472, 435, 436, 259,
	80, 259, 148, 152, 60, 386, 161, 487, 489, 42,
	43, 44, 45, 129, 122, 84, 136, 149, 150, 151,
	74, 141, 157, 68, 73, 70, 74, 426, 252, 71,
	522, 124, 521, 520, 126, 76, 77, 78, 130, 123,
	125, 75, 140, 307, 159, 83, 114, 212, 240, 138,
	166, 391, 115, 148, 152, 229, 165, 161, 228, 169,
	155, 156, 134, 52, 51, 263, 503, 136, 149, 150,
	151, 81, 141, 157, 53, 234, 397, 54, 303, 167,
	81, 237, 238, 79, 248, 239, 81, 498, 167, 222,
	223, 224, 256, 140, 441, 159, 262, 251, 305, 225,
	235, 264, 236, 384, 293, 241, 216, 217, 294, 275,
	118, 155, 156, 134, 501, 502, 555, 363, 552, 304,
	551, 296, 301, 549, 22, 497, 491, 128, 288, 290,
	488, 82, 442, 471, 163, 214, 452, 395, 450, 148,
	152, 374, 385, 161, 395, 215, 455, 218, 219, 220,
	387, 247, 210, 167, 149, 150, 151, 265, 141, 157,
	505, 160, 349, 274, 273, 276, 277, 278, 279, 280,
	275, 458, 244, 245, 459, 460, 309, 347, 422, 140,
	375, 159, 82, 350, 266, 265, 82, 255, 302, 519,
	22, 26, 27, 28, 82, 82, 302, 155, 156, 213,
	505, 82, 310, 82, 166, 518, 131, 334, 315, 214,
	165, 21, 160, 485, 23, 389, 24, 340, 25, 484,
	266, 265, 138, 361, 362, 364, 500, 353, 306, 483,
	367, 359, 386, 372, 373, 526, 376, 377, 378, 379,
	380, 381, 382, 383, 368, 360, 273, 276, 277, 278,
	279, 280, 275, 365, 366, 443, 72, 511, 388, 308,
	88, 138, 138, 166, 311, 312, 364, 398, 82, 165,
	314, 537, 536, 213, 313, 336, 320, 321, 89, 318,
	319, 390, 392, 322, 323, 324, 325, 326, 327, 328,
	329, 330, 331, 396, 535, 333, 481, 336, 160, 166,
	166, 482, 419, 371, 423, 165, 421, 406, 407, 408,
	428, 418, 410, 117, 295, 424, 370, 369, 479, 412,
	346, 348, 345, 480, 415, 299, 440, 221, 214, 98,
	359, 298, 302, 445, 446, 92, 91, 90, 411, 427,
	278, 279, 280, 275, 297, 444, 358, 492, 339, 449,
	29, 357, 46, 490, 138, 276, 277, 278, 279, 280,
	275, 400, 42, 43, 44, 45, 404, 405, 401, 402,
	403, 454, 166, 409, 451, 465, 439, 399, 165, 93,
	94, 467, 466, 456, 418, 338, 474, 464, 473, 116,
	243, 438, 213, 470, 352, 339, 351, 477, 478, 332,
	211, 359, 359, 253, 250, 493, 494, 22, 246, 127,
	496, 528, 529, 430, 431, 432, 433, 434, 499, 435,
	436, 547, 102, 274, 273, 276, 277, 278, 279, 280,
	275, 510, 168, 166, 507, 548, 358, 506, 538, 512,
	525, 357, 10, 9, 513, 22, 242, 258, 120, 461,
	462, 463, 448, 260, 8, 7, 341, 254, 523, 6,
	5, 4, 87, 524, 417, 274, 273, 276, 277, 278,
	279, 280, 275, 63, 64, 85, 96, 95, 97, 388,
	260, 516, 533, 259, 531, 62, 61, 530, 539, 510,
	67, 66, 65, 468, 414, 515, 542, 542, 542, 540,
	476, 543, 544, 532, 336, 534, 166, 317, 316, 553,
	233, 556, 165, 232, 231, 230, 557, 227, 558, 93,
	94, 226, 119, 99, 100, 554, 545, 22, 101, 103,
	104, 105, 106, 108, 109, 48, 110, 457, 112, 113,
	148, 152, 342, 69, 161, 425, 111, 344, 121, 420,
	152, 107, 546, 161, 167, 149, 150, 151, 22, 141,
	157, 527, 508, 167, 149, 150, 151, 514, 295, 157,
	475, 453, 300, 146, 152, 393, 147, 161, 145, 158,
	140, 469, 159, 142, 267, 139, 486, 167, 149, 150,
	151, 159, 295, 157, 22, 26, 27, 28, 155, 156,
	356, 429, 354, 135, 437, 261, 495, 155, 156, 86,
	41, 20, 11, 19, 18, 159, 17, 16, 23, 15,
	24, 30, 25, 274, 273, 276, 277, 278, 279, 280,
	275, 155, 156, 14, 13, 12, 2, 1, 0, 269,
	271, 0, 0, 0, 39, 281, 282, 283, 284, 285,
	286, 287, 272, 270, 268, 274, 273, 276, 277, 278,
	279, 280, 275, 0, 447, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 0, 0, 35, 36, 82, 37,
	38, 274, 273, 276, 277, 278, 279, 280, 275, 47,
	430, 431, 432, 433, 434, 0, 435, 436, 0, 160,
	517, 0, 82, 0, 0, 0, 0, 0, 160, 0,
	0, 0, 0, 49, 50, 55, 56, 57, 58, 59,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 160, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 29, 40, 0, 0, 0, 31,
	32, 34, 33, 170, 171, 172, 173, 174, 175, 176,
	177, 178, 179, 180, 181, 182, 183, 184, 185, 186,
	187, 188, 189, 190, 191, 192, 193, 194, 195, 196,
	197, 198, 199, 200, 201, 202, 203, 204, 205, 206,
	207, 208, 209,
}
var yyPact = [...]int{

	629, -1000, -1000, 361, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 354, -1000, -1000, -47, -1000, -1000, -1000, -1000, -1000,
	225, -111, -112, -93, -99, -1000, 34, -1000, -1000, 92,
	-82, 562, 498, -1000, -1000, -1000, -1000, 484, -1000, 276,
	416, -1000, -65, -1000, -1000, 395, -118, 395, 553, 463,
	361, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -125, -96,
	92, -1000, -94, 92, -1000, 415, -126, 92, -126, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 22, -1000, 354, 85,
	443, 700, 700, -1000, -1000, 411, 165, 165, 1, 165,
	165, 358, -23, 552, 548, -48, -51, 546, 545, 544,
	541, -34, -1000, -58, -1000, -1000, 61, 461, 401, 395,
	395, 414, 131, 92, -1000, 410, -1000, -109, 409, 477,
	170, 92, 478, -1000, -1000, 86, 57, 166, 619, -1000,
	560, 159, -1000, -1000, -1000, 569, -1000, -1000, 345, -1000,
	-1000, -1000, -1000, 332, -1000, -1000, -1000, -1000, 326, 569,
	-1000, -1000, 191, 26, -1000, 93, -1000, 54, 700, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-86, 165, -1000, 569, 560, -1000, 165, 165, -1000, -1000,
	-1000, 395, 239, 539, 538, -1000, 395, 395, 165, 165,
	395, 395, 395, 395, 395, 395, 395, 395, 395, 395,
	-1000, 405, 395, 94, 534, 396, -1000, 476, -152, -1000,
	189, -1000, 402, -1000, -1000, 400, -1000, -1000, 352, 22,
	569, -1000, -1000, 92, 77, 560, 560, 569, 315, 282,
	569, 569, 160, 569, 569, 569, 569, 569, 569, 569,
	569, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 619,
	-39, 0, 8, 619, -1000, 593, -1000, 562, 73, 22,
	114, 387, 94, 23, 569, 92, -1000, 382, -1000, 387,
	166, -1000, -1000, 165, -1000, 395, 395, 395, 165, 165,
	-1000, -1000, 534, 534, 534, 165, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 349, 327, 521, 560, 480, 94, 94,
	-1000, -1000, 161, 92, -1000, -110, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 305, 406, 397, 442, 50, -1000,
	-1000, 127, -1000, -1000, -1000, -1000, 138, 387, -1000, 315,
	569, 569, 387, 645, -1000, 471, 317, 209, -1000, 300,
	300, 66, 66, 66, -1000, -1000, 569, -1000, 387, -1000,
	-4, 22, -6, 121, -1000, 560, -1000, 117, 387, -1000,
	-1000, 165, 165, 165, -1000, -1000, -1000, -1000, -1000, -1000,
	480, 94, 521, 505, 519, 166, -1000, 315, 361, 191,
	-9, -1000, 394, -1000, -1000, 392, -1000, 529, 352, 352,
	-1000, -1000, 311, 289, 222, 212, 206, -8, -1000, 359,
	-16, 353, 569, 569, -1000, 387, 587, 569, -1000, 387,
	-1000, -17, -1000, 41, -1000, 569, 202, -1000, 59, 12,
	-1000, -1000, -1000, -1000, 143, 183, 505, -1000, 569, 252,
	-1000, -1000, 94, -1000, -1000, 523, 507, 406, 683, -1000,
	198, -1000, 182, -1000, -1000, -1000, -1000, -102, -103, -105,
	-1000, -1000, -1000, 387, 387, 569, 387, -1000, -1000, 387,
	569, -1000, -1000, -1000, -1000, 454, -1000, -1000, 230, -1000,
	429, 315, -1000, -1000, 521, 560, 569, 560, -1000, -1000,
	295, 273, 272, 387, 387, 451, 569, -1000, -1000, -1000,
	-1000, 505, 166, 227, 166, 92, 92, 92, 559, -1000,
	445, -19, -1000, -22, -24, 94, -1000, 558, 81, -1000,
	92, -1000, -1000, 191, -1000, 92, -1000, 92, -1000,
}
var yyPgo = [...]int{

	0, 677, 676, 13, 501, 500, 499, 495, 494, 483,
	482, 675, 674, 673, 659, 657, 656, 654, 653, 652,
	651, 729, 251, 650, 649, 296, 22, 25, 645, 644,
	643, 642, 8, 641, 640, 92, 626, 4, 24, 27,
	625, 624, 16, 623, 2, 15, 5, 621, 619, 12,
	618, 6, 616, 615, 17, 613, 612, 611, 610, 607,
	9, 602, 1, 601, 7, 592, 18, 589, 11, 3,
	30, 87, 167, 588, 587, 585, 583, 582, 0, 10,
	99, 577, 318, 575,
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
	44, 44, 55, 55, 55, 55, 48, 48, 50, 50,
	50, 52, 56, 56, 53, 53, 54, 57, 57, 51,
	51, 43, 43, 43, 43, 58, 58, 59, 59, 60,
	60, 61, 61, 62, 63, 63, 63, 64, 64, 64,
	64, 65, 65, 65, 66, 66, 67, 67, 68, 68,
	69, 69, 70, 72, 72, 73, 73, 25, 25, 74,
	74, 74, 74, 74, 75, 75, 76, 76, 77, 77,
	78, 78, 79, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 80, 80, 80, 80, 80, 80, 80,
	80, 80, 80, 81, 81, 81, 81, 82, 82, 82,
	71, 71, 71,
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
	1, 1, 3, 4, 5, 4, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	4, 0, 2, 4, 0, 3, 1, 3, 0, 5,
	1, 3, 3, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 1, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 1, 0, 1, 1,
	0, 2, 2,
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
	36, 37, 21, -78, -49, 78, 79, 40, -48, 62,
	179, 24, -69, 89, -70, -51, -78, 34, 29, -80,
	103, 104, 105, 106, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 129, 130, 131, 132,
	133, 134, 135, 136, 137, 138, 139, 140, 141, 142,
	-80, 29, -71, 74, 10, -71, 145, 146, -71, -71,
	-71, 9, 152, 153, 154, 162, 9, 9, 146, 146,
	9, 9, 9, 9, 149, 174, 176, 155, 156, 159,
	146, 84, 25, 29, -35, -35, 34, 60, -78, -79,
	34, -79, 177, 34, 20, 57, -78, -64, 9, 45,
	15, -28, -78, 19, 84, 59, 58, -41, 75, 60,
	74, 61, 73, 77, 76, 83, 78, 79, 80, 81,
	82, 66, 67, 68, 69, 70, 71, 72, -39, -44,
	-39, -46, -3, -44, -44, 39, -49, 39, 39, 39,
	-56, -44, 45, 92, 66, 84, -80, 169, -71, -44,
	-39, -71, -71, -35, -71, 9, 9, 9, -35, -35,
	-71, -71, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, 34, -35, -69, -38, 10, -66, 29, 39,
	-79, 20, -77, 181, -74, 173, 171, 28, 172, 13,
	34, 34, 34, -79, -31, -32, -34, 39, 34, -49,
	-27, -44, -78, 80, -78, -39, -39, -44, -45, 75,
	74, 61, -44, -44, 21, 60, -44, -44, -44, -44,
	-44, -44, -44, -44, 182, 182, 45, 182, -44, 182,
	-26, 18, -26, -53, -54, 63, -70, 93, -44, 35,
	-71, -35, -35, -35, -71, -71, -38, -38, -38, -71,
	-66, 29, -38, -60, 13, -39, -42, 24, -3, -69,
	-67, -51, 57, -78, -79, -75, 177, -38, 45, -33,
	47, 48, 49, 50, 51, 53, 54, -29, 34, 19,
	-32, 84, 45, 168, -45, -44, -44, 59, 21, -44,
	182, -26, 182, -57, -54, 65, -39, -81, 94, 97,
	98, -71, -71, -71, -42, -69, -60, -64, 14, -47,
	-45, 182, 45, 34, 34, -58, 11, -32, -32, 47,
	52, 47, 52, 47, 47, 47, -36, 55, 178, 56,
	34, 182, 34, -44, -44, 59, -44, 182, 86, -44,
	64, 95, 96, 94, -68, 57, -68, -64, -61, -62,
	-44, 45, -51, -79, -59, 12, 14, 57, 47, 47,
	175, 175, 175, -44, -44, 26, 45, -63, 22, 23,
	-45, -60, -39, -46, -39, 39, 39, 39, 27, -62,
	-64, -37, -78, -37, -37, 7, -65, 16, 30, 182,
	45, 182, 182, -69, 7, 75, -78, -78, -78,
}
var yyDef = [...]int{

	98, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 96, 96, 96, 96, 96, 96, 96, 96,
	0, 256, 247, 0, 0, 44, 0, 46, 47, 0,
	94, 0, 100, 102, 103, 104, 99, 105, 98, 307,
	307, 89, 0, 91, 92, 0, 247, 0, 0, 0,
	30, 31, 32, 33, 34, 35, 36, 37, 245, 0,
	0, 257, 0, 0, 248, 0, 243, 0, 243, 45,
	48, 260, 261, 95, 23, 101, 0, 106, 97, 0,
	0, 0, 0, 308, 309, 0, 310, 310, 0, 310,
	310, 310, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 86, 0, 90, 93, 137, 0, 0, 0,
	0, 0, 0, 0, 262, 0, 262, 0, 0, 0,
	0, 0, 227, 107, 109, 114, 260, 112, 113, 147,
	0, 0, 178, 179, 180, 0, 190, 191, 0, 211,
	212, 213, 214, 209, 174, 198, 199, 200, 0, 202,
	196, 197, 38, 0, 240, 0, 209, 260, 0, 40,
	263, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 282,
	283, 284, 285, 286, 287, 288, 289, 290, 291, 292,
	293, 294, 295, 296, 297, 298, 299, 300, 301, 302,
	41, 310, 60, 0, 0, 61, 310, 310, 64, 65,
	66, 0, 310, 0, 0, 87, 0, 0, 310, 310,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	88, 0, 0, 0, 145, 234, 262, 0, 258, 51,
	0, 54, 0, 56, 244, 0, 262, 21, 0, 0,
	0, 110, 115, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 162, 163, 164, 165, 166, 167, 168, 150, 0,
	0, 0, 0, 176, 189, 0, 161, 0, 0, 0,
	0, 203, 0, 0, 0, 0, 39, 0, 59, 311,
	312, 62, 63, 310, 68, 0, 0, 0, 310, 310,
	74, 75, 145, 145, 145, 310, 80, 81, 82, 83,
	84, 85, 138, 234, 145, 219, 0, 0, 0, 0,
	49, 246, 0, 0, 262, 254, 249, 250, 251, 252,
	253, 55, 57, 58, 145, 117, 122, 0, 134, 136,
	108, 228, 116, 111, 210, 148, 149, 152, 153, 0,
	0, 0, 155, 0, 159, 0, 181, 182, 183, 184,
	185, 186, 187, 188, 151, 173, 0, 175, 176, 192,
	0, 0, 0, 207, 204, 0, 241, 0, 242, 42,
	67, 310, 310, 310, 70, 71, 76, 77, 78, 79,
	0, 0, 219, 227, 0, 146, 26, 0, 170, 27,
	0, 236, 0, 259, 52, 0, 255, 215, 0, 0,
	125, 126, 0, 0, 0, 0, 0, 139, 123, 0,
	0, 0, 0, 0, 154, 156, 0, 0, 160, 177,
	193, 0, 195, 0, 205, 0, 0, 43, 0, 0,
	306, 69, 72, 73, 238, 238, 227, 29, 0, 169,
	171, 235, 0, 262, 53, 217, 0, 118, 0, 127,
	0, 129, 0, 131, 132, 133, 119, 0, 0, 0,
	124, 120, 135, 229, 230, 0, 157, 194, 201, 208,
	0, 303, 304, 305, 24, 0, 25, 28, 220, 221,
	224, 0, 237, 50, 219, 0, 0, 0, 128, 130,
	0, 0, 0, 158, 206, 0, 0, 223, 225, 226,
	172, 227, 218, 216, 121, 0, 0, 0, 0, 222,
	231, 0, 143, 0, 0, 0, 22, 0, 0, 140,
	0, 141, 142, 239, 232, 0, 144, 0, 233,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 82, 77, 3,
	39, 182, 80, 78, 45, 79, 84, 81, 3, 3,
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
	177, 178, 179, 180, 181,
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
		//line ./yacc.y:235
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:241
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:243
		{
			yyVAL.statement = yyDollar[1].setStmt
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:245
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:247
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:263
		{
			yyVAL.statement = nil
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:267
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 22:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:271
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:275
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:281
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:285
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
		//line ./yacc.y:297
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:301
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
		//line ./yacc.y:313
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:319
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:325
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].selStmt}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:329
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:333
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:337
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:341
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:345
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].setStmt}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:349
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:353
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:359
		{
			yyVAL.setStmt = &SetVariable{
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				Exprs:    yyDollar[4].updateExprs,
			}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:367
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[5].bytes),
			}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:374
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[4].bytes),
			}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:381
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
			}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:388
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
				Collate:  string(yyDollar[6].bytes),
			}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:396
		{
			yyVAL.setStmt = &SetTransactionIsolationLevel{
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				IsolationLevel: string(yyDollar[7].bytes),
			}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:406
		{
			yyVAL.statement = &Begin{}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:410
		{
			yyVAL.statement = &Begin{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:416
		{
			yyVAL.statement = &Commit{}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:422
		{
			yyVAL.statement = &Rollback{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:428
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:434
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:438
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:443
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:449
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:453
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:458
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:464
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:470
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:474
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:479
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:485
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:489
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:493
		{
			yyVAL.showStmt = &ShowCollation{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:497
		{
			yyVAL.showStmt = &ShowVariables{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:501
		{
			yyVAL.showStmt = &ShowStatus{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:505
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:509
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:513
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:517
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:521
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:525
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 70:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:529
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:533
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 72:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:537
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 73:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:541
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:545
		{
			yyVAL.showStmt = &ShowProcedureStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:549
		{
			yyVAL.showStmt = &ShowFunctionStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:553
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:557
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:561
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 79:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:565
		{
			yyVAL.showStmt = &ShowTriggers{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:569
		{
			yyVAL.showStmt = &ShowCreateDatabase{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:573
		{
			yyVAL.showStmt = &ShowCreateTable{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:577
		{
			yyVAL.showStmt = &ShowCreateView{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:581
		{
			yyVAL.showStmt = &ShowCreateProcedure{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:585
		{
			yyVAL.showStmt = &ShowCreateFunction{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:589
		{
			yyVAL.showStmt = &ShowCreateTrigger{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:593
		{
			yyVAL.showStmt = &ShowProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:597
		{
			yyVAL.showStmt = &ShowFullProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:601
		{
			yyVAL.showStmt = &ShowSlaveStatus{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:605
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:609
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:613
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:617
		{
			yyVAL.showStmt = &ShowProfiles{}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:623
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[3].tableName}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:630
		{
			yyVAL.statement = nil
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:632
		{
			yyVAL.statement = nil
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:635
		{
			SetAllowComments(yylex, true)
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:639
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:645
		{
			yyVAL.bytes2 = nil
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:649
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:655
		{
			yyVAL.str = AST_UNION
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:659
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:663
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:667
		{
			yyVAL.str = AST_EXCEPT
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:671
		{
			yyVAL.str = AST_INTERSECT
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:676
		{
			yyVAL.str = ""
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:680
		{
			yyVAL.str = AST_DISTINCT
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:686
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:690
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:696
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:700
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:704
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:710
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:714
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:719
		{
			yyVAL.bytes = nil
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:723
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:727
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:733
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:737
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:743
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:747
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:751
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:756
		{
			yyVAL.bytes = nil
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:760
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:764
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:770
		{
			yyVAL.str = AST_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:774
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:778
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:782
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:786
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:790
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:794
		{
			yyVAL.str = AST_JOIN
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:798
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:802
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:808
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:812
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:816
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:822
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:826
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 139:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:831
		{
			yyVAL.indexHints = nil
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:835
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 141:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:839
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 142:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:843
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:849
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:853
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 145:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:858
		{
			yyVAL.boolExpr = nil
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:862
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:869
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:873
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:877
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:881
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:887
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:891
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:895
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:899
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:903
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:907
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:911
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:915
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:919
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:923
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:929
		{
			yyVAL.str = AST_EQ
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:933
		{
			yyVAL.str = AST_LT
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:937
		{
			yyVAL.str = AST_GT
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:941
		{
			yyVAL.str = AST_LE
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:945
		{
			yyVAL.str = AST_GE
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:949
		{
			yyVAL.str = AST_NE
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:953
		{
			yyVAL.str = AST_NSE
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:959
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:963
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:969
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:973
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:979
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:983
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:989
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:995
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:999
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1005
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1009
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1013
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1017
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1021
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1025
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1029
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1033
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1037
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1041
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1045
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1049
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
		//line ./yacc.y:1064
		{
			yyVAL.valExpr = yyDollar[1].funcExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1068
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1074
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1078
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1082
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1086
		{
			yyVAL.funcExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1092
		{
			yyVAL.bytes = IF_BYTES
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1096
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1102
		{
			yyVAL.byt = AST_UPLUS
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1106
		{
			yyVAL.byt = AST_UMINUS
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1110
		{
			yyVAL.byt = AST_TILDA
		}
	case 201:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1116
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1121
		{
			yyVAL.valExpr = nil
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1125
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1131
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1135
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 206:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1141
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1146
		{
			yyVAL.valExpr = nil
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1150
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1156
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1160
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1166
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1170
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1174
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1178
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1183
		{
			yyVAL.valExprs = nil
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1187
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1192
		{
			yyVAL.boolExpr = nil
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1196
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1201
		{
			yyVAL.orderBy = nil
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1205
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1211
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1215
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1221
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1226
		{
			yyVAL.str = AST_ASC
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1230
		{
			yyVAL.str = AST_ASC
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1234
		{
			yyVAL.str = AST_DESC
		}
	case 227:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1239
		{
			yyVAL.limit = nil
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1243
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 229:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1247
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 230:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1251
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1256
		{
			yyVAL.str = ""
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1260
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 233:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1264
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
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1277
		{
			yyVAL.columns = nil
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1281
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1287
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1291
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1296
		{
			yyVAL.updateExprs = nil
		}
	case 239:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1300
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1306
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1310
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 242:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1316
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 243:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1321
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1323
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1326
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1328
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1331
		{
			yyVAL.str = ""
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1333
		{
			yyVAL.str = AST_IGNORE
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1337
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1339
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1341
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1343
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1345
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1348
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1350
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1353
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1355
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1358
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1360
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1364
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1368
		{
			yyVAL.bytes = []byte("database")
		}
	case 262:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1373
		{
			ForceEOF(yylex)
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1379
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1381
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1383
		{
			yyVAL.bytes = []byte("big5")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1385
		{
			yyVAL.bytes = []byte("binary")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1387
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1389
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1391
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1393
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1395
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1397
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1399
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1401
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1403
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1405
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1407
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1409
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1411
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1413
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1415
		{
			yyVAL.bytes = []byte("greek")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1417
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1419
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1421
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1423
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1425
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1427
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1429
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1431
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1433
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1435
		{
			yyVAL.bytes = []byte("macce")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1437
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1439
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1441
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1443
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1445
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1447
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 298:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1449
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1451
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 300:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1453
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 301:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1455
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1457
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 303:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1461
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 304:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1463
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 305:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1465
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 306:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1467
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 307:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1470
		{
			yyVAL.bytes = nil
		}
	case 308:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1472
		{
			yyVAL.bytes = []byte("session")
		}
	case 309:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1474
		{
			yyVAL.bytes = []byte("global")
		}
	case 310:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1477
		{
			yyVAL.expr = nil
		}
	case 311:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1479
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 312:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1483
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
