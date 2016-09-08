//line ./yacc.y:2
package sqlparser

import __yyfmt__ "fmt"

//line ./yacc.y:2
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line ./yacc.y:27
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
const REPLACE = 57489
const ADMIN = 57490
const HELP = 57491
const OFFSET = 57492
const COLLATE = 57493
const CREATE = 57494
const ALTER = 57495
const DROP = 57496
const RENAME = 57497
const TABLE = 57498
const INDEX = 57499
const VIEW = 57500
const TO = 57501
const IGNORE = 57502
const IF = 57503
const UNIQUE = 57504
const USING = 57505

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

const yyNprod = 306
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 783

var yyAct = [...]int{

	140, 494, 279, 155, 526, 281, 137, 247, 343, 399,
	239, 489, 148, 282, 3, 138, 402, 331, 382, 325,
	157, 535, 126, 123, 323, 256, 255, 131, 127, 379,
	535, 143, 147, 535, 249, 153, 42, 43, 44, 45,
	76, 457, 249, 249, 59, 130, 144, 145, 146, 374,
	135, 151, 116, 70, 412, 80, 416, 417, 418, 419,
	420, 242, 421, 422, 507, 506, 64, 118, 66, 505,
	120, 134, 67, 154, 124, 472, 474, 69, 225, 70,
	72, 73, 74, 228, 229, 132, 159, 230, 117, 149,
	150, 128, 158, 143, 147, 205, 119, 153, 71, 297,
	226, 79, 227, 108, 220, 215, 216, 130, 144, 145,
	146, 488, 135, 151, 217, 52, 51, 219, 238, 385,
	22, 26, 27, 28, 337, 53, 246, 293, 77, 160,
	252, 241, 253, 134, 75, 154, 209, 210, 283, 335,
	444, 160, 284, 445, 446, 338, 372, 77, 486, 487,
	22, 149, 150, 128, 483, 537, 288, 291, 427, 295,
	78, 109, 278, 280, 536, 143, 147, 534, 482, 153,
	268, 269, 270, 265, 375, 456, 438, 436, 540, 160,
	144, 145, 146, 373, 135, 151, 254, 152, 476, 208,
	377, 211, 212, 213, 428, 473, 156, 162, 231, 265,
	294, 147, 77, 383, 153, 134, 383, 154, 441, 299,
	237, 362, 122, 255, 160, 144, 145, 146, 112, 285,
	151, 207, 78, 149, 150, 264, 263, 266, 267, 268,
	269, 270, 265, 502, 159, 300, 490, 322, 408, 245,
	158, 504, 154, 78, 78, 256, 255, 328, 351, 152,
	363, 485, 132, 349, 350, 352, 78, 341, 149, 150,
	355, 347, 78, 360, 361, 503, 364, 365, 366, 367,
	368, 369, 370, 371, 356, 234, 235, 29, 348, 334,
	336, 333, 68, 353, 354, 206, 203, 125, 376, 132,
	132, 256, 255, 159, 78, 21, 352, 386, 292, 158,
	298, 88, 87, 86, 85, 301, 302, 305, 207, 378,
	380, 304, 359, 384, 429, 308, 309, 78, 466, 292,
	470, 152, 469, 467, 468, 358, 357, 159, 159, 78,
	405, 490, 409, 158, 407, 392, 393, 394, 111, 404,
	374, 396, 511, 410, 84, 89, 90, 398, 464, 214,
	207, 496, 401, 465, 426, 94, 152, 522, 347, 296,
	324, 431, 432, 42, 43, 44, 45, 413, 397, 248,
	521, 324, 206, 430, 520, 250, 303, 435, 327, 46,
	307, 285, 132, 310, 311, 312, 313, 314, 315, 316,
	317, 318, 319, 289, 321, 414, 22, 287, 159, 388,
	440, 450, 437, 391, 158, 249, 292, 452, 451, 395,
	404, 442, 346, 449, 206, 286, 326, 345, 387, 455,
	425, 233, 477, 462, 463, 346, 327, 347, 347, 475,
	345, 478, 479, 143, 147, 424, 481, 153, 459, 513,
	514, 458, 110, 340, 484, 339, 320, 160, 144, 145,
	146, 532, 135, 151, 243, 240, 495, 236, 159, 492,
	121, 204, 491, 161, 497, 533, 523, 389, 390, 498,
	510, 10, 9, 134, 8, 154, 263, 266, 267, 268,
	269, 270, 265, 508, 7, 447, 448, 22, 509, 232,
	114, 149, 150, 264, 263, 266, 267, 268, 269, 270,
	265, 98, 62, 63, 376, 61, 403, 518, 329, 516,
	434, 244, 515, 524, 495, 60, 83, 81, 250, 22,
	501, 527, 527, 527, 525, 453, 528, 529, 517, 400,
	519, 159, 500, 461, 538, 147, 541, 158, 153, 324,
	306, 542, 224, 543, 223, 222, 480, 221, 160, 144,
	145, 146, 218, 285, 151, 92, 91, 93, 113, 48,
	433, 539, 78, 264, 263, 266, 267, 268, 269, 270,
	265, 530, 22, 443, 330, 65, 154, 264, 263, 266,
	267, 268, 269, 270, 265, 411, 332, 115, 406, 152,
	531, 512, 149, 150, 493, 499, 460, 439, 89, 90,
	290, 381, 95, 96, 142, 139, 141, 97, 99, 100,
	101, 103, 104, 454, 105, 136, 107, 22, 26, 27,
	28, 257, 106, 133, 471, 344, 415, 102, 264, 263,
	266, 267, 268, 269, 270, 265, 342, 129, 423, 251,
	82, 23, 41, 24, 30, 25, 266, 267, 268, 269,
	270, 265, 416, 417, 418, 419, 420, 20, 421, 422,
	11, 19, 18, 78, 259, 261, 17, 39, 47, 16,
	271, 272, 273, 274, 275, 276, 277, 262, 260, 258,
	264, 263, 266, 267, 268, 269, 270, 265, 15, 14,
	152, 13, 49, 50, 54, 55, 56, 57, 58, 35,
	36, 12, 37, 38, 163, 164, 165, 166, 167, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 187, 188,
	189, 190, 191, 192, 193, 194, 195, 196, 197, 198,
	199, 200, 201, 202, 6, 5, 4, 2, 1, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 29, 40, 0, 0, 0, 31,
	32, 34, 33,
}
var yyPact = [...]int{

	612, -1000, -1000, 322, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 341, -1000, -1000, -35, -1000, -1000, -1000, -1000, -1000,
	115, -105, -96, -73, -91, -1000, 45, -1000, -1000, 94,
	-63, 567, 500, -1000, -1000, -1000, -1000, 498, -1000, 202,
	455, -1000, -48, -1000, 408, -122, 408, 549, 465, 322,
	-1000, -1000, -1000, -1000, -124, -84, 94, -1000, -75, 94,
	-1000, 426, -153, 94, -153, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 73, -1000, 341, 107, 434, 601, 601, -1000,
	-1000, 432, 211, 211, -9, 211, 211, 340, -47, 543,
	-29, -42, 538, 536, 535, 533, -71, -1000, -1000, -1000,
	114, 464, 392, 408, 408, 423, 150, 94, -1000, 421,
	-1000, -113, 420, 491, 182, 94, 360, -1000, -1000, 113,
	102, 233, 604, -1000, 413, 145, -1000, -1000, -1000, 180,
	376, 358, -1000, 354, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 180, 253, 35, -1000, 134, -1000,
	75, 601, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -67, 211, -1000, 180, 413, -1000, 211,
	211, -1000, -1000, -1000, 408, 298, 531, -1000, 408, 211,
	211, 408, 408, 408, 408, 408, 408, 408, 408, 408,
	408, 412, 408, 95, 529, 387, -1000, 488, -161, -1000,
	111, -1000, 411, -1000, -1000, 409, -1000, -1000, 378, 73,
	180, -1000, -1000, 94, 168, 413, 413, 180, 342, 251,
	180, 180, 190, 180, 180, 180, 180, 180, 180, 180,
	180, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 604,
	-33, 4, -5, 604, -1000, 514, 11, 73, -1000, 567,
	140, 552, 95, 26, 180, 94, -1000, 383, -1000, 552,
	233, -1000, -1000, 211, -1000, 408, 408, 211, -1000, -1000,
	529, 529, 529, 211, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 339, 361, 516, 413, 482, 95, 95, -1000, -1000,
	181, 94, -1000, -120, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 350, 605, 401, 391, 74, -1000, -1000, 149,
	-1000, -1000, -1000, -1000, 154, 552, -1000, 342, 180, 180,
	552, 501, -1000, 489, 568, 399, -1000, 90, 90, 116,
	116, 116, -1000, -1000, 180, -1000, 552, -1000, -2, 73,
	-3, 143, -1000, 413, -1000, 46, 552, -1000, -1000, 211,
	211, -1000, -1000, -1000, -1000, -1000, 482, 95, 516, 503,
	511, 233, -1000, 342, 322, 253, -4, -1000, 407, -1000,
	-1000, 404, -1000, 522, 378, 378, -1000, -1000, 301, 271,
	277, 275, 273, 20, -1000, 395, 9, 388, 180, 180,
	-1000, 552, 487, 180, -1000, 552, -1000, -11, -1000, 68,
	-1000, 180, 187, -1000, 53, 17, -1000, -1000, -1000, 179,
	274, 503, -1000, 180, 306, -1000, -1000, 95, -1000, -1000,
	520, 506, 605, 176, -1000, 218, -1000, 194, -1000, -1000,
	-1000, -1000, -103, -107, -108, -1000, -1000, -1000, 552, 552,
	180, 552, -1000, -1000, 552, 180, -1000, -1000, -1000, -1000,
	444, -1000, -1000, 297, -1000, 417, 342, -1000, -1000, 516,
	413, 180, 413, -1000, -1000, 335, 331, 318, 552, 552,
	439, 180, -1000, -1000, -1000, -1000, 503, 233, 295, 233,
	94, 94, 94, 564, -1000, 435, -12, -1000, -15, -24,
	95, -1000, 554, 103, -1000, 94, -1000, -1000, 253, -1000,
	94, -1000, 94, -1000,
}
var yyPgo = [...]int{

	0, 748, 747, 13, 746, 745, 744, 484, 474, 472,
	471, 701, 691, 689, 688, 669, 666, 662, 661, 660,
	657, 668, 295, 642, 640, 282, 22, 28, 639, 638,
	637, 636, 8, 626, 625, 161, 624, 4, 24, 27,
	623, 621, 16, 615, 2, 15, 5, 613, 606, 12,
	605, 6, 604, 601, 18, 600, 597, 596, 595, 9,
	594, 1, 591, 7, 590, 19, 588, 11, 3, 20,
	95, 212, 587, 586, 585, 575, 574, 0, 10, 197,
	573, 304, 559,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 7, 7, 10, 10, 8, 9,
	19, 19, 19, 19, 19, 4, 4, 4, 4, 4,
	4, 15, 15, 16, 17, 18, 11, 11, 11, 12,
	12, 12, 13, 14, 14, 14, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 6, 20, 20, 82,
	21, 22, 22, 23, 23, 23, 23, 23, 24, 24,
	26, 26, 27, 27, 27, 30, 30, 28, 28, 28,
	31, 31, 32, 32, 32, 32, 29, 29, 29, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 34, 34,
	34, 35, 35, 36, 36, 36, 36, 37, 37, 38,
	38, 39, 39, 39, 39, 39, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 41, 41, 41, 41,
	41, 41, 41, 42, 42, 47, 47, 45, 45, 49,
	46, 46, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 48,
	48, 50, 50, 50, 52, 55, 55, 53, 53, 54,
	56, 56, 51, 51, 43, 43, 43, 43, 57, 57,
	58, 58, 59, 59, 60, 60, 61, 62, 62, 62,
	63, 63, 63, 63, 64, 64, 64, 65, 65, 66,
	66, 67, 67, 68, 68, 69, 71, 71, 72, 72,
	25, 25, 73, 73, 73, 73, 73, 74, 74, 75,
	75, 76, 76, 77, 77, 78, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 80, 80, 80, 80,
	81, 81, 81, 70, 70, 70,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 5, 12, 3, 8, 8, 6, 6, 8, 7,
	2, 2, 2, 2, 2, 4, 5, 4, 4, 6,
	7, 1, 2, 1, 1, 2, 5, 8, 4, 6,
	7, 4, 5, 4, 5, 5, 5, 4, 4, 5,
	5, 4, 4, 4, 6, 5, 7, 6, 7, 5,
	5, 6, 6, 6, 6, 5, 5, 5, 5, 5,
	5, 3, 4, 2, 3, 2, 3, 1, 2, 0,
	2, 0, 2, 1, 2, 1, 1, 1, 0, 1,
	1, 3, 1, 2, 3, 1, 1, 0, 1, 2,
	1, 3, 3, 3, 3, 5, 0, 1, 2, 1,
	1, 2, 3, 2, 3, 2, 2, 2, 1, 3,
	1, 1, 3, 0, 5, 5, 5, 1, 3, 0,
	2, 1, 3, 3, 2, 3, 3, 3, 4, 3,
	4, 5, 6, 3, 4, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 3, 3, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 5, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 4, 0, 2, 4, 0, 3, 1,
	3, 0, 5, 1, 3, 3, 0, 2, 0, 3,
	0, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 2, 1,
	0, 1, 1, 0, 2, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -19, -11, -12, -13, -14, -15, -16, -17, -18,
	-20, -22, 5, 29, 31, 33, 6, 7, 8, 162,
	32, 167, 168, 170, 169, 87, 88, 90, 91, 55,
	163, -23, 41, 42, 43, 44, 38, -21, -82, -21,
	-21, 151, 150, 160, -21, -21, -21, -21, -21, -3,
	-7, -8, -10, -9, 171, -75, 173, 177, -25, 173,
	175, 171, 171, 172, 173, 89, -77, 34, 149, 164,
	-3, 17, -24, 18, -22, -81, 101, 100, 99, 143,
	144, 101, 100, 102, -81, 147, 148, 152, 46, 153,
	154, 155, 172, 156, 157, 159, 167, 161, 151, -35,
	34, -25, -35, 9, 25, -72, 176, 172, -77, 171,
	-77, 34, -71, 176, -77, -71, -26, -27, 80, -30,
	34, -39, -44, -40, 60, 39, -43, -51, -45, -50,
	-77, -48, -52, 20, 35, 36, 37, 21, -49, 78,
	79, 40, 176, 24, 62, -68, 89, -69, -51, -77,
	34, 29, -79, 103, 104, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
	120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	140, 141, 142, -79, 29, -70, 74, 10, -70, 145,
	146, -70, -70, -70, 9, 152, 153, 161, 9, 146,
	146, 9, 9, 9, 9, 149, 171, 173, 154, 155,
	158, 84, 25, 29, -35, -35, 34, 60, -77, -78,
	34, -78, 174, 34, 20, 57, -77, -63, 9, 45,
	15, -28, -77, 19, 84, 59, 58, -41, 75, 60,
	74, 61, 73, 77, 76, 83, 78, 79, 80, 81,
	82, 66, 67, 68, 69, 70, 71, 72, -39, -44,
	-39, -46, -3, -44, -44, 39, 39, 39, -49, 39,
	-55, -44, 45, 92, 66, 84, -79, 166, -70, -44,
	-39, -70, -70, -35, -70, 9, 9, -35, -70, -70,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	34, -35, -68, -38, 10, -65, 29, 39, -78, 20,
	-76, 178, -73, 170, 168, 28, 169, 13, 34, 34,
	34, -78, -31, -32, -34, 39, 34, -49, -27, -44,
	-77, 80, -77, -39, -39, -44, -45, 75, 74, 61,
	-44, -44, 21, 60, -44, -44, -44, -44, -44, -44,
	-44, -44, 179, 179, 45, 179, -44, 179, -26, 18,
	-26, -53, -54, 63, -69, 93, -44, 35, -70, -35,
	-35, -70, -38, -38, -38, -70, -65, 29, -38, -59,
	13, -39, -42, 24, -3, -68, -66, -51, 57, -77,
	-78, -74, 174, -38, 45, -33, 47, 48, 49, 50,
	51, 53, 54, -29, 34, 19, -32, 84, 45, 165,
	-45, -44, -44, 59, 21, -44, 179, -26, 179, -56,
	-54, 65, -39, -80, 94, 97, 98, -70, -70, -42,
	-68, -59, -63, 14, -47, -45, 179, 45, 34, 34,
	-57, 11, -32, -32, 47, 52, 47, 52, 47, 47,
	47, -36, 55, 175, 56, 34, 179, 34, -44, -44,
	59, -44, 179, 86, -44, 64, 95, 96, 94, -67,
	57, -67, -63, -60, -61, -44, 45, -51, -78, -58,
	12, 14, 57, 47, 47, 172, 172, 172, -44, -44,
	26, 45, -62, 22, 23, -45, -59, -39, -46, -39,
	39, 39, 39, 27, -61, -63, -37, -77, -37, -37,
	7, -64, 16, 30, 179, 45, 179, 179, -68, 7,
	75, -77, -77, -77,
}
var yyDef = [...]int{

	91, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 89, 89, 89, 89, 89, 89, 89, 89,
	0, 249, 240, 0, 0, 41, 0, 43, 44, 0,
	87, 0, 93, 95, 96, 97, 92, 98, 91, 300,
	300, 83, 0, 85, 0, 240, 0, 0, 0, 30,
	31, 32, 33, 34, 238, 0, 0, 250, 0, 0,
	241, 0, 236, 0, 236, 42, 45, 253, 254, 88,
	23, 94, 0, 99, 90, 0, 0, 0, 0, 301,
	302, 0, 303, 303, 0, 303, 303, 303, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 81, 84, 86,
	131, 0, 0, 0, 0, 0, 0, 0, 255, 0,
	255, 0, 0, 0, 0, 0, 220, 100, 102, 107,
	253, 105, 106, 141, 0, 0, 172, 173, 174, 0,
	202, 0, 188, 0, 204, 205, 206, 207, 168, 191,
	192, 193, 189, 190, 195, 35, 0, 233, 0, 202,
	253, 0, 37, 256, 257, 258, 259, 260, 261, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 282,
	283, 284, 285, 286, 287, 288, 289, 290, 291, 292,
	293, 294, 295, 38, 303, 57, 0, 0, 58, 303,
	303, 61, 62, 63, 0, 303, 0, 82, 0, 303,
	303, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 139, 227, 255, 0, 251, 48,
	0, 51, 0, 53, 237, 0, 255, 21, 0, 0,
	0, 103, 108, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 156, 157, 158, 159, 160, 161, 162, 144, 0,
	0, 0, 0, 170, 183, 0, 0, 0, 155, 0,
	0, 196, 0, 0, 0, 0, 36, 0, 56, 304,
	305, 59, 60, 303, 65, 0, 0, 303, 69, 70,
	139, 139, 139, 303, 75, 76, 77, 78, 79, 80,
	132, 227, 139, 212, 0, 0, 0, 0, 46, 239,
	0, 0, 255, 247, 242, 243, 244, 245, 246, 52,
	54, 55, 139, 110, 116, 0, 128, 130, 101, 221,
	109, 104, 203, 142, 143, 146, 147, 0, 0, 0,
	149, 0, 153, 0, 175, 176, 177, 178, 179, 180,
	181, 182, 145, 167, 0, 169, 170, 184, 0, 0,
	0, 200, 197, 0, 234, 0, 235, 39, 64, 303,
	303, 67, 71, 72, 73, 74, 0, 0, 212, 220,
	0, 140, 26, 0, 164, 27, 0, 229, 0, 252,
	49, 0, 248, 208, 0, 0, 119, 120, 0, 0,
	0, 0, 0, 133, 117, 0, 0, 0, 0, 0,
	148, 150, 0, 0, 154, 171, 185, 0, 187, 0,
	198, 0, 0, 40, 0, 0, 299, 66, 68, 231,
	231, 220, 29, 0, 163, 165, 228, 0, 255, 50,
	210, 0, 111, 114, 121, 0, 123, 0, 125, 126,
	127, 112, 0, 0, 0, 118, 113, 129, 222, 223,
	0, 151, 186, 194, 201, 0, 296, 297, 298, 24,
	0, 25, 28, 213, 214, 217, 0, 230, 47, 212,
	0, 0, 0, 122, 124, 0, 0, 0, 152, 199,
	0, 0, 216, 218, 219, 166, 220, 211, 209, 115,
	0, 0, 0, 0, 215, 224, 0, 137, 0, 0,
	0, 22, 0, 0, 134, 0, 135, 136, 232, 225,
	0, 138, 0, 226,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 82, 77, 3,
	39, 179, 80, 78, 45, 79, 84, 81, 3, 3,
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
	177, 178,
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
		//line ./yacc.y:189
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:195
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:197
		{
			yyVAL.statement = yyDollar[1].setStmt
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:199
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:201
		{
			yyVAL.statement = yyDollar[1].showStmt
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:217
		{
			yyVAL.statement = nil
		}
	case 21:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:221
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 22:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:225
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:229
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:235
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:239
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
		//line ./yacc.y:251
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 27:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:255
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
		//line ./yacc.y:267
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 29:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:273
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:279
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].selStmt}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:283
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:287
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:291
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:295
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:301
		{
			yyVAL.setStmt = &SetVariable{
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				Exprs:    yyDollar[4].updateExprs,
			}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:309
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[5].bytes),
			}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:316
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[4].bytes),
			}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:323
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
			}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:330
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
				Collate:  string(yyDollar[6].bytes),
			}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:338
		{
			yyVAL.setStmt = &SetTransactionIsolationLevel{
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				IsolationLevel: string(yyDollar[7].bytes),
			}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:348
		{
			yyVAL.statement = &Begin{}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:352
		{
			yyVAL.statement = &Begin{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:358
		{
			yyVAL.statement = &Commit{}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:364
		{
			yyVAL.statement = &Rollback{}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:370
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:376
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:380
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:385
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:391
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:395
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:400
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:406
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:412
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:416
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:421
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:427
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:431
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:435
		{
			yyVAL.showStmt = &ShowCollation{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:439
		{
			yyVAL.showStmt = &ShowVariables{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 60:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:443
		{
			yyVAL.showStmt = &ShowStatus{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:447
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:451
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:455
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:459
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:463
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 66:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:467
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 67:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:471
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 68:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:475
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:479
		{
			yyVAL.showStmt = &ShowProcedureStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:483
		{
			yyVAL.showStmt = &ShowFunctionStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:487
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:491
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:495
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 74:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:499
		{
			yyVAL.showStmt = &ShowTriggers{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:503
		{
			yyVAL.showStmt = &ShowCreateDatabase{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:507
		{
			yyVAL.showStmt = &ShowCreateTable{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:511
		{
			yyVAL.showStmt = &ShowCreateView{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:515
		{
			yyVAL.showStmt = &ShowCreateProcedure{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:519
		{
			yyVAL.showStmt = &ShowCreateFunction{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:523
		{
			yyVAL.showStmt = &ShowCreateTrigger{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:527
		{
			yyVAL.showStmt = &ShowProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:531
		{
			yyVAL.showStmt = &ShowFullProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:535
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:539
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:543
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:549
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[3].tableName}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:556
		{
			yyVAL.statement = nil
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:558
		{
			yyVAL.statement = nil
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:561
		{
			SetAllowComments(yylex, true)
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:565
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:571
		{
			yyVAL.bytes2 = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:575
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:581
		{
			yyVAL.str = AST_UNION
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:585
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:589
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:593
		{
			yyVAL.str = AST_EXCEPT
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:597
		{
			yyVAL.str = AST_INTERSECT
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:602
		{
			yyVAL.str = ""
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:606
		{
			yyVAL.str = AST_DISTINCT
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:612
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:616
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:622
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:626
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:630
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:636
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:640
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:645
		{
			yyVAL.bytes = nil
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:649
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:653
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:659
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:663
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:669
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:673
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:677
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 115:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:681
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:686
		{
			yyVAL.bytes = nil
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:690
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:694
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:700
		{
			yyVAL.str = AST_JOIN
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:704
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:708
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:712
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:716
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:720
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:724
		{
			yyVAL.str = AST_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:728
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:732
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:738
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:742
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:746
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:752
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:756
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 133:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:761
		{
			yyVAL.indexHints = nil
		}
	case 134:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:765
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 135:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:769
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 136:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:773
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:779
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:783
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 139:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:788
		{
			yyVAL.boolExpr = nil
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:792
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:799
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:803
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:807
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:811
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:817
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:821
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:825
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:829
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:833
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:837
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:841
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:845
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:849
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:853
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:859
		{
			yyVAL.str = AST_EQ
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:863
		{
			yyVAL.str = AST_LT
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:867
		{
			yyVAL.str = AST_GT
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:871
		{
			yyVAL.str = AST_LE
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:875
		{
			yyVAL.str = AST_GE
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:879
		{
			yyVAL.str = AST_NE
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:883
		{
			yyVAL.str = AST_NSE
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:889
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:893
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:899
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 166:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:903
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:909
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:913
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:919
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:925
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:929
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:935
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:939
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:943
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:947
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:951
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:955
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:959
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:963
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:967
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:971
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:975
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:979
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
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:994
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:998
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 186:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1002
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1006
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1010
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1016
		{
			yyVAL.bytes = IF_BYTES
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1020
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1026
		{
			yyVAL.byt = AST_UPLUS
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1030
		{
			yyVAL.byt = AST_UMINUS
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1034
		{
			yyVAL.byt = AST_TILDA
		}
	case 194:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1040
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1045
		{
			yyVAL.valExpr = nil
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1049
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1055
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1059
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1065
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1070
		{
			yyVAL.valExpr = nil
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1074
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1080
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1084
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1090
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1094
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1098
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1102
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1107
		{
			yyVAL.valExprs = nil
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1111
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1116
		{
			yyVAL.boolExpr = nil
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1120
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1125
		{
			yyVAL.orderBy = nil
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1129
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1135
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1139
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1145
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1150
		{
			yyVAL.str = AST_ASC
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1154
		{
			yyVAL.str = AST_ASC
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1158
		{
			yyVAL.str = AST_DESC
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1163
		{
			yyVAL.limit = nil
		}
	case 221:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1167
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 222:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1171
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 223:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1175
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1180
		{
			yyVAL.str = ""
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1184
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 226:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1188
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
	case 227:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1201
		{
			yyVAL.columns = nil
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1205
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1211
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1215
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1220
		{
			yyVAL.updateExprs = nil
		}
	case 232:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1224
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1230
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1234
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1240
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1245
		{
			yyVAL.empty = struct{}{}
		}
	case 237:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1247
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1250
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1252
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1255
		{
			yyVAL.str = ""
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1257
		{
			yyVAL.str = AST_IGNORE
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1272
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1274
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1279
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1284
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1288
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1292
		{
			yyVAL.bytes = []byte("database")
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1297
		{
			ForceEOF(yylex)
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1303
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1305
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1307
		{
			yyVAL.bytes = []byte("big5")
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1309
		{
			yyVAL.bytes = []byte("binary")
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1311
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1313
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1315
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1317
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1319
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1321
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1323
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1325
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1327
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1329
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1331
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1333
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1335
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1337
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1339
		{
			yyVAL.bytes = []byte("greek")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1341
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1343
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1345
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1347
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1349
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1351
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1353
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1355
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1357
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1359
		{
			yyVAL.bytes = []byte("macce")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1361
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1363
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1365
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1367
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1369
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1371
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1373
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1375
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1377
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1379
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1381
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 296:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1385
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 297:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1387
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 298:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1389
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1391
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 300:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1394
		{
			yyVAL.bytes = nil
		}
	case 301:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1396
		{
			yyVAL.bytes = []byte("session")
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1398
		{
			yyVAL.bytes = []byte("global")
		}
	case 303:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1401
		{
			yyVAL.expr = nil
		}
	case 304:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1403
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 305:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1407
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
