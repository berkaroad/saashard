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
const REPLACE = 57490
const ADMIN = 57491
const HELP = 57492
const OFFSET = 57493
const COLLATE = 57494
const CREATE = 57495
const ALTER = 57496
const DROP = 57497
const RENAME = 57498
const TABLE = 57499
const INDEX = 57500
const VIEW = 57501
const TO = 57502
const IGNORE = 57503
const IF = 57504
const UNIQUE = 57505
const USING = 57506

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

const yyNprod = 310
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 829

var yyAct = [...]int{

	144, 499, 284, 159, 531, 286, 141, 252, 387, 404,
	244, 494, 152, 407, 348, 142, 130, 330, 287, 3,
	161, 261, 260, 336, 328, 131, 127, 540, 540, 135,
	120, 42, 43, 44, 45, 73, 384, 540, 147, 151,
	79, 254, 157, 462, 421, 422, 423, 424, 425, 59,
	426, 427, 134, 148, 149, 150, 417, 139, 155, 254,
	83, 247, 254, 67, 379, 69, 512, 147, 151, 70,
	122, 157, 72, 124, 73, 511, 510, 128, 138, 121,
	158, 134, 148, 149, 150, 123, 139, 155, 136, 163,
	342, 209, 74, 477, 479, 162, 153, 154, 132, 147,
	151, 302, 82, 157, 112, 340, 235, 138, 258, 158,
	101, 343, 224, 164, 148, 149, 150, 166, 139, 155,
	219, 220, 243, 80, 223, 153, 154, 132, 80, 221,
	251, 75, 76, 77, 257, 246, 80, 52, 51, 138,
	164, 158, 288, 377, 213, 214, 289, 53, 449, 433,
	229, 450, 451, 491, 492, 232, 233, 153, 154, 234,
	293, 296, 542, 541, 95, 94, 96, 81, 283, 285,
	380, 390, 539, 230, 356, 231, 487, 481, 461, 493,
	269, 268, 271, 272, 273, 274, 275, 270, 212, 298,
	215, 216, 217, 488, 443, 156, 81, 441, 382, 378,
	78, 432, 300, 22, 26, 27, 28, 92, 93, 207,
	259, 98, 99, 304, 478, 236, 100, 102, 103, 104,
	106, 107, 270, 108, 156, 110, 111, 23, 81, 24,
	164, 25, 109, 273, 274, 275, 270, 105, 81, 163,
	364, 305, 327, 81, 126, 162, 339, 341, 338, 310,
	211, 81, 333, 363, 362, 81, 156, 136, 354, 355,
	357, 545, 346, 218, 211, 360, 352, 299, 365, 366,
	434, 369, 370, 371, 372, 373, 374, 375, 376, 361,
	353, 388, 211, 301, 367, 160, 91, 90, 89, 242,
	358, 359, 260, 381, 136, 136, 507, 388, 163, 446,
	303, 357, 391, 495, 162, 306, 307, 413, 383, 385,
	250, 309, 71, 509, 210, 313, 314, 508, 389, 261,
	260, 21, 129, 368, 113, 490, 261, 260, 210, 471,
	92, 93, 163, 163, 472, 410, 475, 414, 162, 412,
	397, 398, 399, 474, 401, 81, 210, 88, 415, 409,
	297, 473, 403, 271, 272, 273, 274, 275, 270, 406,
	297, 29, 495, 352, 469, 431, 436, 437, 115, 470,
	87, 527, 418, 421, 422, 423, 424, 425, 435, 426,
	427, 116, 440, 42, 43, 44, 45, 136, 268, 271,
	272, 273, 274, 275, 270, 445, 518, 519, 97, 379,
	393, 442, 253, 163, 396, 329, 455, 516, 255, 162,
	400, 329, 457, 456, 501, 454, 351, 402, 447, 331,
	409, 350, 526, 525, 460, 290, 294, 332, 292, 332,
	291, 392, 352, 352, 467, 468, 483, 484, 254, 482,
	419, 486, 239, 240, 46, 430, 297, 480, 464, 489,
	269, 268, 271, 272, 273, 274, 275, 270, 463, 114,
	429, 500, 345, 163, 497, 151, 344, 496, 157, 502,
	325, 248, 245, 22, 503, 241, 125, 485, 164, 148,
	149, 150, 10, 290, 155, 9, 452, 453, 513, 537,
	8, 238, 208, 514, 269, 268, 271, 272, 273, 274,
	275, 270, 351, 538, 165, 528, 158, 350, 515, 381,
	237, 22, 523, 62, 521, 439, 63, 520, 529, 500,
	118, 61, 153, 154, 334, 249, 532, 532, 532, 530,
	408, 533, 534, 86, 7, 522, 163, 524, 84, 543,
	255, 546, 162, 308, 506, 22, 547, 312, 548, 438,
	315, 316, 317, 318, 319, 320, 321, 322, 323, 324,
	147, 151, 326, 6, 157, 60, 269, 268, 271, 272,
	273, 274, 275, 270, 164, 148, 149, 150, 5, 139,
	155, 458, 22, 269, 268, 271, 272, 273, 274, 275,
	270, 405, 505, 81, 66, 4, 466, 329, 151, 311,
	138, 157, 158, 47, 228, 227, 226, 225, 222, 65,
	117, 164, 148, 149, 150, 544, 290, 155, 153, 154,
	535, 156, 22, 26, 27, 28, 64, 49, 50, 54,
	55, 56, 57, 58, 22, 394, 395, 48, 448, 158,
	335, 68, 416, 337, 119, 411, 23, 536, 24, 30,
	25, 517, 498, 504, 465, 153, 154, 444, 295, 386,
	146, 143, 145, 459, 140, 262, 137, 476, 349, 420,
	347, 133, 39, 428, 256, 85, 41, 20, 11, 19,
	18, 17, 16, 15, 14, 13, 12, 2, 1, 81,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 0, 37, 38, 0,
	0, 0, 0, 0, 0, 0, 0, 156, 0, 0,
	0, 264, 266, 0, 0, 0, 81, 276, 277, 278,
	279, 280, 281, 282, 267, 265, 263, 269, 268, 271,
	272, 273, 274, 275, 270, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 156, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	29, 40, 0, 0, 0, 31, 32, 34, 33, 167,
	168, 169, 170, 171, 172, 173, 174, 175, 176, 177,
	178, 179, 180, 181, 182, 183, 184, 185, 186, 187,
	188, 189, 190, 191, 192, 193, 194, 195, 196, 197,
	198, 199, 200, 201, 202, 203, 204, 205, 206,
}
var yyPact = [...]int{

	617, -1000, -1000, 342, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 406, -1000, -1000, -13, -1000, -1000, -1000, -1000, -1000,
	198, -109, -102, -80, -41, -1000, 111, -1000, -1000, 102,
	-63, 629, 521, -1000, -1000, -1000, -1000, 515, -1000, 187,
	64, -1000, -47, -1000, 425, -141, 425, 601, 495, 342,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -147, -94, 102,
	-1000, -87, 102, -1000, 442, -151, 102, -151, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 47, -1000, 406, 196, 475,
	686, 686, -1000, -1000, 463, 272, 272, -1, 272, 272,
	254, -32, 599, -22, -34, 598, 597, 596, 595, 1,
	-1000, -40, -1000, -1000, 131, 485, 462, 425, 425, 441,
	229, 102, -1000, 438, -1000, -114, 437, 505, 253, 102,
	393, -1000, -1000, 89, 126, 268, 661, -1000, 79, 540,
	-1000, -1000, -1000, 444, 391, 389, -1000, 387, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 444, 315,
	97, -1000, 201, -1000, 118, 686, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -66, 272, -1000,
	444, 79, -1000, 272, 272, -1000, -1000, -1000, 425, 240,
	590, -1000, 425, 272, 272, 425, 425, 425, 425, 425,
	425, 425, 425, 425, 425, -1000, 436, 425, 106, 587,
	390, -1000, 504, -156, -1000, 77, -1000, 432, -1000, -1000,
	428, -1000, -1000, 382, 47, 444, -1000, -1000, 102, 94,
	79, 79, 444, 386, 179, 444, 444, 263, 444, 444,
	444, 444, 444, 444, 444, 444, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 661, -37, 19, -10, 661, -1000,
	577, 18, 47, -1000, 629, 218, 507, 106, 78, 444,
	102, -1000, 396, -1000, 507, 268, -1000, -1000, 272, -1000,
	425, 425, 272, -1000, -1000, 587, 587, 587, 272, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 388, 401, 578, 79,
	506, 106, 106, -1000, -1000, 250, 102, -1000, -119, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 395, 326, 426,
	468, 117, -1000, -1000, 104, -1000, -1000, -1000, -1000, 233,
	507, -1000, 386, 444, 444, 507, 490, -1000, 494, 275,
	311, -1000, 153, 153, 139, 139, 139, -1000, -1000, 444,
	-1000, 507, -1000, 17, 47, 14, 234, -1000, 79, -1000,
	54, 507, -1000, -1000, 272, 272, -1000, -1000, -1000, -1000,
	-1000, 506, 106, 578, 525, 567, 268, -1000, 386, 342,
	315, -2, -1000, 424, -1000, -1000, 414, -1000, 585, 382,
	382, -1000, -1000, 317, 282, 304, 296, 289, 38, -1000,
	413, -3, 405, 444, 444, -1000, 507, 418, 444, -1000,
	507, -1000, -4, -1000, 107, -1000, 444, 261, -1000, 58,
	85, -1000, -1000, -1000, 246, 305, 525, -1000, 444, 369,
	-1000, -1000, 106, -1000, -1000, 580, 530, 326, 239, -1000,
	270, -1000, 266, -1000, -1000, -1000, -1000, -97, -98, -107,
	-1000, -1000, -1000, 507, 507, 444, 507, -1000, -1000, 507,
	444, -1000, -1000, -1000, -1000, 482, -1000, -1000, 362, -1000,
	374, 386, -1000, -1000, 578, 79, 444, 79, -1000, -1000,
	384, 383, 332, 507, 507, 478, 444, -1000, -1000, -1000,
	-1000, 525, 268, 354, 268, 102, 102, 102, 613, -1000,
	473, -8, -1000, -17, -18, 106, -1000, 608, 186, -1000,
	102, -1000, -1000, 315, -1000, 102, -1000, 102, -1000,
}
var yyPgo = [...]int{

	0, 688, 687, 18, 595, 578, 563, 534, 490, 485,
	482, 686, 685, 684, 683, 682, 681, 680, 679, 678,
	677, 603, 321, 676, 675, 312, 16, 25, 674, 673,
	671, 670, 14, 669, 668, 324, 667, 4, 24, 29,
	666, 665, 13, 664, 2, 15, 5, 663, 662, 12,
	661, 6, 660, 659, 8, 658, 657, 654, 653, 9,
	652, 1, 651, 7, 647, 17, 645, 11, 3, 20,
	91, 244, 644, 643, 642, 641, 640, 0, 10, 117,
	638, 347, 637,
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
	6, 20, 20, 82, 21, 22, 22, 23, 23, 23,
	23, 23, 24, 24, 26, 26, 27, 27, 27, 30,
	30, 28, 28, 28, 31, 31, 32, 32, 32, 32,
	29, 29, 29, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 34, 34, 34, 35, 35, 36, 36, 36,
	36, 37, 37, 38, 38, 39, 39, 39, 39, 39,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	41, 41, 41, 41, 41, 41, 41, 42, 42, 47,
	47, 45, 45, 49, 46, 46, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 48, 48, 50, 50, 50, 52, 55,
	55, 53, 53, 54, 56, 56, 51, 51, 43, 43,
	43, 43, 57, 57, 58, 58, 59, 59, 60, 60,
	61, 62, 62, 62, 63, 63, 63, 63, 64, 64,
	64, 65, 65, 66, 66, 67, 67, 68, 68, 69,
	71, 71, 72, 72, 25, 25, 73, 73, 73, 73,
	73, 74, 74, 75, 75, 76, 76, 77, 77, 78,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	80, 80, 80, 80, 81, 81, 81, 70, 70, 70,
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
	3, 1, 2, 0, 2, 0, 2, 1, 2, 1,
	1, 1, 0, 1, 1, 3, 1, 2, 3, 1,
	1, 0, 1, 2, 1, 3, 3, 3, 3, 5,
	0, 1, 2, 1, 1, 2, 3, 2, 3, 2,
	2, 2, 1, 3, 1, 1, 3, 0, 5, 5,
	5, 1, 3, 0, 2, 1, 3, 3, 2, 3,
	3, 3, 4, 3, 4, 5, 6, 3, 4, 2,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	3, 3, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 1, 3, 3,
	0, 2, 0, 3, 0, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 1, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 1, 0, 1, 1, 0, 2, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -19, -11, -12, -13, -14, -15, -16, -17, -18,
	-20, -22, 5, 29, 31, 33, 6, 7, 8, 163,
	32, 168, 169, 171, 170, 87, 88, 90, 91, 55,
	164, -23, 41, 42, 43, 44, 38, -21, -82, -21,
	-21, 151, 150, 160, -21, -21, -21, -21, -21, -3,
	-7, -8, -10, -9, -4, -5, -6, 172, -75, 174,
	178, -25, 174, 176, 172, 172, 173, 174, 89, -77,
	34, 149, 165, -3, 17, -24, 18, -22, -81, 101,
	100, 99, 143, 144, 101, 100, 102, -81, 147, 148,
	152, 46, 153, 154, 155, 173, 156, 157, 159, 168,
	161, 162, 151, -35, 34, -25, -35, 9, 25, -72,
	177, 173, -77, 172, -77, 34, -71, 177, -77, -71,
	-26, -27, 80, -30, 34, -39, -44, -40, 60, 39,
	-43, -51, -45, -50, -77, -48, -52, 20, 35, 36,
	37, 21, -49, 78, 79, 40, 177, 24, 62, -68,
	89, -69, -51, -77, 34, 29, -79, 103, 104, 105,
	106, 107, 108, 109, 110, 111, 112, 113, 114, 115,
	116, 117, 118, 119, 120, 121, 122, 123, 124, 125,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 142, -79, 29, -70,
	74, 10, -70, 145, 146, -70, -70, -70, 9, 152,
	153, 161, 9, 146, 146, 9, 9, 9, 9, 149,
	172, 174, 154, 155, 158, 146, 84, 25, 29, -35,
	-35, 34, 60, -77, -78, 34, -78, 175, 34, 20,
	57, -77, -63, 9, 45, 15, -28, -77, 19, 84,
	59, 58, -41, 75, 60, 74, 61, 73, 77, 76,
	83, 78, 79, 80, 81, 82, 66, 67, 68, 69,
	70, 71, 72, -39, -44, -39, -46, -3, -44, -44,
	39, 39, 39, -49, 39, -55, -44, 45, 92, 66,
	84, -79, 167, -70, -44, -39, -70, -70, -35, -70,
	9, 9, -35, -70, -70, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, 34, -35, -68, -38, 10,
	-65, 29, 39, -78, 20, -76, 179, -73, 171, 169,
	28, 170, 13, 34, 34, 34, -78, -31, -32, -34,
	39, 34, -49, -27, -44, -77, 80, -77, -39, -39,
	-44, -45, 75, 74, 61, -44, -44, 21, 60, -44,
	-44, -44, -44, -44, -44, -44, -44, 180, 180, 45,
	180, -44, 180, -26, 18, -26, -53, -54, 63, -69,
	93, -44, 35, -70, -35, -35, -70, -38, -38, -38,
	-70, -65, 29, -38, -59, 13, -39, -42, 24, -3,
	-68, -66, -51, 57, -77, -78, -74, 175, -38, 45,
	-33, 47, 48, 49, 50, 51, 53, 54, -29, 34,
	19, -32, 84, 45, 166, -45, -44, -44, 59, 21,
	-44, 180, -26, 180, -56, -54, 65, -39, -80, 94,
	97, 98, -70, -70, -42, -68, -59, -63, 14, -47,
	-45, 180, 45, 34, 34, -57, 11, -32, -32, 47,
	52, 47, 52, 47, 47, 47, -36, 55, 176, 56,
	34, 180, 34, -44, -44, 59, -44, 180, 86, -44,
	64, 95, 96, 94, -67, 57, -67, -63, -60, -61,
	-44, 45, -51, -78, -58, 12, 14, 57, 47, 47,
	173, 173, 173, -44, -44, 26, 45, -62, 22, 23,
	-45, -59, -39, -46, -39, 39, 39, 39, 27, -61,
	-63, -37, -77, -37, -37, 7, -64, 16, 30, 180,
	45, 180, 180, -68, 7, 75, -77, -77, -77,
}
var yyDef = [...]int{

	95, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 93, 93, 93, 93, 93, 93, 93, 93,
	0, 253, 244, 0, 0, 44, 0, 46, 47, 0,
	91, 0, 97, 99, 100, 101, 96, 102, 95, 304,
	304, 87, 0, 89, 0, 244, 0, 0, 0, 30,
	31, 32, 33, 34, 35, 36, 37, 242, 0, 0,
	254, 0, 0, 245, 0, 240, 0, 240, 45, 48,
	257, 258, 92, 23, 98, 0, 103, 94, 0, 0,
	0, 0, 305, 306, 0, 307, 307, 0, 307, 307,
	307, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	84, 0, 88, 90, 135, 0, 0, 0, 0, 0,
	0, 0, 259, 0, 259, 0, 0, 0, 0, 0,
	224, 104, 106, 111, 257, 109, 110, 145, 0, 0,
	176, 177, 178, 0, 206, 0, 192, 0, 208, 209,
	210, 211, 172, 195, 196, 197, 193, 194, 199, 38,
	0, 237, 0, 206, 257, 0, 40, 260, 261, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 282,
	283, 284, 285, 286, 287, 288, 289, 290, 291, 292,
	293, 294, 295, 296, 297, 298, 299, 41, 307, 60,
	0, 0, 61, 307, 307, 64, 65, 66, 0, 307,
	0, 85, 0, 307, 307, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 86, 0, 0, 0, 143,
	231, 259, 0, 255, 51, 0, 54, 0, 56, 241,
	0, 259, 21, 0, 0, 0, 107, 112, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 160, 161, 162, 163,
	164, 165, 166, 148, 0, 0, 0, 0, 174, 187,
	0, 0, 0, 159, 0, 0, 200, 0, 0, 0,
	0, 39, 0, 59, 308, 309, 62, 63, 307, 68,
	0, 0, 307, 72, 73, 143, 143, 143, 307, 78,
	79, 80, 81, 82, 83, 136, 231, 143, 216, 0,
	0, 0, 0, 49, 243, 0, 0, 259, 251, 246,
	247, 248, 249, 250, 55, 57, 58, 143, 114, 120,
	0, 132, 134, 105, 225, 113, 108, 207, 146, 147,
	150, 151, 0, 0, 0, 153, 0, 157, 0, 179,
	180, 181, 182, 183, 184, 185, 186, 149, 171, 0,
	173, 174, 188, 0, 0, 0, 204, 201, 0, 238,
	0, 239, 42, 67, 307, 307, 70, 74, 75, 76,
	77, 0, 0, 216, 224, 0, 144, 26, 0, 168,
	27, 0, 233, 0, 256, 52, 0, 252, 212, 0,
	0, 123, 124, 0, 0, 0, 0, 0, 137, 121,
	0, 0, 0, 0, 0, 152, 154, 0, 0, 158,
	175, 189, 0, 191, 0, 202, 0, 0, 43, 0,
	0, 303, 69, 71, 235, 235, 224, 29, 0, 167,
	169, 232, 0, 259, 53, 214, 0, 115, 118, 125,
	0, 127, 0, 129, 130, 131, 116, 0, 0, 0,
	122, 117, 133, 226, 227, 0, 155, 190, 198, 205,
	0, 300, 301, 302, 24, 0, 25, 28, 217, 218,
	221, 0, 234, 50, 216, 0, 0, 0, 126, 128,
	0, 0, 0, 156, 203, 0, 0, 220, 222, 223,
	170, 224, 215, 213, 119, 0, 0, 0, 0, 219,
	228, 0, 141, 0, 0, 0, 22, 0, 0, 138,
	0, 139, 140, 236, 229, 0, 142, 0, 230,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 82, 77, 3,
	39, 180, 80, 78, 45, 79, 84, 81, 3, 3,
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
	177, 178, 179,
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
		//line ./yacc.y:232
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:238
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:240
		{
			yyVAL.statement = yyDollar[1].setStmt
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:242
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:244
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:260
		{
			yyVAL.statement = nil
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:264
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 22:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:268
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:272
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:278
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:282
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
		//line ./yacc.y:294
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:298
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
		//line ./yacc.y:310
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:316
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:322
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].selStmt}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:326
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:330
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:334
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:338
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:342
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].setStmt}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:346
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:350
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].showStmt}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:356
		{
			yyVAL.setStmt = &SetVariable{
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				Exprs:    yyDollar[4].updateExprs,
			}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:364
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[5].bytes),
			}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:371
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[4].bytes),
			}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:378
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
			}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:385
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
				Collate:  string(yyDollar[6].bytes),
			}
		}
	case 43:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:393
		{
			yyVAL.setStmt = &SetTransactionIsolationLevel{
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				IsolationLevel: string(yyDollar[7].bytes),
			}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:403
		{
			yyVAL.statement = &Begin{}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:407
		{
			yyVAL.statement = &Begin{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:413
		{
			yyVAL.statement = &Commit{}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:419
		{
			yyVAL.statement = &Rollback{}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:425
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:431
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:435
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:440
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:446
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:450
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:455
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:461
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:467
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:471
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:476
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:482
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:486
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:490
		{
			yyVAL.showStmt = &ShowCollation{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:494
		{
			yyVAL.showStmt = &ShowVariables{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:498
		{
			yyVAL.showStmt = &ShowStatus{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:502
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:506
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:510
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:514
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:518
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:522
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 70:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:526
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:530
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:534
		{
			yyVAL.showStmt = &ShowProcedureStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:538
		{
			yyVAL.showStmt = &ShowFunctionStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:542
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:546
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:550
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:554
		{
			yyVAL.showStmt = &ShowTriggers{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:558
		{
			yyVAL.showStmt = &ShowCreateDatabase{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:562
		{
			yyVAL.showStmt = &ShowCreateTable{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:566
		{
			yyVAL.showStmt = &ShowCreateView{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:570
		{
			yyVAL.showStmt = &ShowCreateProcedure{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:574
		{
			yyVAL.showStmt = &ShowCreateFunction{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:578
		{
			yyVAL.showStmt = &ShowCreateTrigger{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:582
		{
			yyVAL.showStmt = &ShowProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:586
		{
			yyVAL.showStmt = &ShowFullProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:590
		{
			yyVAL.showStmt = &ShowSlaveStatus{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:594
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:598
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:602
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:608
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[3].tableName}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:615
		{
			yyVAL.statement = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:617
		{
			yyVAL.statement = nil
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:620
		{
			SetAllowComments(yylex, true)
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:624
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:630
		{
			yyVAL.bytes2 = nil
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:634
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:640
		{
			yyVAL.str = AST_UNION
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:644
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:648
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:652
		{
			yyVAL.str = AST_EXCEPT
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:656
		{
			yyVAL.str = AST_INTERSECT
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:661
		{
			yyVAL.str = ""
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:665
		{
			yyVAL.str = AST_DISTINCT
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:671
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:675
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:681
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:685
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:689
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:695
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:699
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:704
		{
			yyVAL.bytes = nil
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:708
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:712
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:718
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:722
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:728
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:732
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:736
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 119:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:740
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 120:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:745
		{
			yyVAL.bytes = nil
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:749
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:753
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:759
		{
			yyVAL.str = AST_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:763
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:767
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:771
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:775
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:779
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:783
		{
			yyVAL.str = AST_JOIN
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:787
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:791
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:797
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:801
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:805
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:811
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:815
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 137:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:820
		{
			yyVAL.indexHints = nil
		}
	case 138:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:824
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 139:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:828
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 140:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:832
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:838
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:842
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 143:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:847
		{
			yyVAL.boolExpr = nil
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:851
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:858
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:862
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:866
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:870
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:876
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:880
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:884
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:888
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:892
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:896
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:900
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:904
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:908
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:912
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:918
		{
			yyVAL.str = AST_EQ
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:922
		{
			yyVAL.str = AST_LT
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:926
		{
			yyVAL.str = AST_GT
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:930
		{
			yyVAL.str = AST_LE
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:934
		{
			yyVAL.str = AST_GE
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:938
		{
			yyVAL.str = AST_NE
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:942
		{
			yyVAL.str = AST_NSE
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:948
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:952
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:958
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:962
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:968
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:972
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:978
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:984
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:988
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:994
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:998
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1002
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1006
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1010
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1014
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1018
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1022
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1026
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1030
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1034
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1038
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
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1053
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1057
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 190:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1061
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1065
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1069
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1075
		{
			yyVAL.bytes = IF_BYTES
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1079
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1085
		{
			yyVAL.byt = AST_UPLUS
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1089
		{
			yyVAL.byt = AST_UMINUS
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1093
		{
			yyVAL.byt = AST_TILDA
		}
	case 198:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1099
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1104
		{
			yyVAL.valExpr = nil
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1108
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1114
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1118
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 203:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1124
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1129
		{
			yyVAL.valExpr = nil
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1133
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1139
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1143
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1149
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1153
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1157
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1161
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1166
		{
			yyVAL.valExprs = nil
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1170
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1175
		{
			yyVAL.boolExpr = nil
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1179
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1184
		{
			yyVAL.orderBy = nil
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1188
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1194
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 219:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1198
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1204
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1209
		{
			yyVAL.str = AST_ASC
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1213
		{
			yyVAL.str = AST_ASC
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1217
		{
			yyVAL.str = AST_DESC
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1222
		{
			yyVAL.limit = nil
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1226
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 226:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1230
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1234
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1239
		{
			yyVAL.str = ""
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1243
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 230:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1247
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
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1260
		{
			yyVAL.columns = nil
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1264
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1270
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1274
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1279
		{
			yyVAL.updateExprs = nil
		}
	case 236:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1283
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1289
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1293
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1299
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 240:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1304
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1306
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1309
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1311
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1314
		{
			yyVAL.str = ""
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1316
		{
			yyVAL.str = AST_IGNORE
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1320
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1322
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1324
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1326
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1328
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1331
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1333
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1336
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1338
		{
			yyVAL.empty = struct{}{}
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1341
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1343
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1347
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1351
		{
			yyVAL.bytes = []byte("database")
		}
	case 259:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1356
		{
			ForceEOF(yylex)
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1362
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1364
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1366
		{
			yyVAL.bytes = []byte("big5")
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1368
		{
			yyVAL.bytes = []byte("binary")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1370
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1372
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1374
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1376
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1378
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1380
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1382
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1384
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1386
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1388
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1390
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1392
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1394
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1396
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1398
		{
			yyVAL.bytes = []byte("greek")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1400
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1402
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1404
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1406
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1408
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1410
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1412
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1414
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1416
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1418
		{
			yyVAL.bytes = []byte("macce")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1420
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1422
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1424
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1426
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1428
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1430
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1432
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1434
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1436
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 298:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1438
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1440
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 300:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1444
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 301:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1446
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 302:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1448
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 303:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1450
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 304:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1453
		{
			yyVAL.bytes = nil
		}
	case 305:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1455
		{
			yyVAL.bytes = []byte("session")
		}
	case 306:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1457
		{
			yyVAL.bytes = []byte("global")
		}
	case 307:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1460
		{
			yyVAL.expr = nil
		}
	case 308:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1462
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 309:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1466
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
