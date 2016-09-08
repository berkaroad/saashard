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

const yyNprod = 307
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 841

var yyAct = [...]int{

	141, 496, 281, 156, 528, 283, 138, 249, 384, 401,
	241, 491, 149, 404, 345, 139, 127, 327, 284, 3,
	158, 537, 258, 257, 325, 128, 333, 537, 124, 132,
	42, 43, 44, 45, 117, 537, 381, 70, 144, 148,
	76, 251, 154, 459, 418, 419, 420, 421, 422, 59,
	423, 424, 131, 145, 146, 147, 251, 136, 152, 251,
	80, 376, 64, 414, 66, 474, 476, 119, 67, 69,
	121, 70, 244, 509, 125, 72, 73, 74, 135, 508,
	155, 507, 118, 120, 226, 133, 160, 71, 299, 229,
	230, 206, 159, 231, 79, 339, 150, 151, 129, 109,
	22, 26, 27, 28, 163, 52, 51, 227, 77, 228,
	337, 216, 217, 255, 430, 53, 340, 77, 98, 240,
	218, 144, 148, 210, 211, 154, 161, 248, 77, 232,
	221, 254, 243, 220, 490, 131, 145, 146, 147, 285,
	136, 152, 387, 286, 374, 266, 265, 268, 269, 270,
	271, 272, 267, 295, 353, 75, 539, 290, 293, 488,
	489, 135, 538, 155, 485, 280, 282, 78, 429, 377,
	536, 297, 92, 91, 93, 256, 484, 478, 458, 150,
	151, 129, 233, 267, 542, 209, 475, 212, 213, 214,
	296, 440, 161, 204, 438, 153, 375, 446, 379, 385,
	447, 448, 266, 265, 268, 269, 270, 271, 272, 267,
	301, 270, 271, 272, 267, 89, 90, 258, 257, 95,
	96, 307, 208, 78, 97, 99, 100, 101, 103, 104,
	239, 105, 78, 107, 108, 431, 160, 257, 302, 324,
	106, 78, 159, 78, 385, 102, 443, 157, 123, 330,
	78, 336, 338, 335, 133, 351, 352, 354, 29, 343,
	504, 364, 357, 349, 208, 362, 363, 298, 366, 367,
	368, 369, 370, 371, 372, 373, 358, 350, 153, 492,
	361, 88, 87, 86, 410, 247, 207, 355, 356, 85,
	378, 133, 133, 360, 359, 160, 506, 300, 354, 388,
	365, 159, 303, 304, 505, 380, 382, 78, 306, 215,
	208, 68, 310, 311, 472, 386, 265, 268, 269, 270,
	271, 272, 267, 126, 110, 89, 90, 21, 207, 160,
	160, 524, 407, 471, 411, 159, 409, 394, 395, 396,
	94, 398, 258, 257, 470, 412, 406, 294, 487, 400,
	268, 269, 270, 271, 272, 267, 403, 294, 376, 492,
	349, 468, 428, 433, 434, 466, 469, 112, 513, 415,
	467, 498, 250, 523, 207, 432, 84, 522, 252, 437,
	287, 113, 22, 326, 133, 42, 43, 44, 45, 326,
	348, 291, 442, 515, 516, 347, 289, 390, 439, 288,
	160, 393, 46, 452, 399, 427, 159, 397, 251, 454,
	453, 348, 451, 328, 329, 444, 347, 406, 416, 479,
	426, 457, 389, 329, 294, 477, 461, 460, 111, 349,
	349, 464, 465, 480, 481, 342, 341, 322, 483, 236,
	237, 245, 242, 238, 122, 534, 486, 266, 265, 268,
	269, 270, 271, 272, 267, 10, 9, 8, 497, 535,
	160, 494, 148, 235, 493, 154, 499, 205, 7, 162,
	525, 500, 512, 234, 482, 161, 145, 146, 147, 115,
	287, 152, 22, 449, 450, 510, 62, 63, 61, 436,
	511, 266, 265, 268, 269, 270, 271, 272, 267, 60,
	331, 405, 246, 155, 83, 81, 378, 252, 503, 520,
	455, 518, 402, 502, 517, 526, 497, 463, 326, 150,
	151, 308, 225, 529, 529, 529, 527, 224, 530, 531,
	223, 222, 519, 160, 521, 219, 540, 114, 543, 159,
	305, 541, 22, 544, 309, 545, 532, 312, 313, 314,
	315, 316, 317, 318, 319, 320, 321, 144, 148, 323,
	22, 154, 48, 445, 144, 148, 332, 65, 154, 413,
	334, 161, 145, 146, 147, 116, 136, 152, 161, 145,
	146, 147, 408, 136, 152, 418, 419, 420, 421, 422,
	78, 423, 424, 533, 514, 22, 495, 135, 501, 155,
	462, 441, 292, 383, 135, 143, 155, 140, 142, 456,
	137, 148, 259, 134, 154, 150, 151, 473, 153, 346,
	417, 435, 150, 151, 161, 145, 146, 147, 344, 287,
	152, 130, 391, 392, 22, 26, 27, 28, 266, 265,
	268, 269, 270, 271, 272, 267, 47, 425, 253, 82,
	41, 20, 155, 11, 19, 18, 17, 16, 23, 15,
	24, 30, 25, 14, 13, 12, 6, 5, 150, 151,
	49, 50, 54, 55, 56, 57, 58, 4, 2, 1,
	0, 0, 0, 0, 39, 0, 78, 0, 261, 263,
	0, 0, 0, 78, 273, 274, 275, 276, 277, 278,
	279, 264, 262, 260, 266, 265, 268, 269, 270, 271,
	272, 267, 0, 0, 153, 0, 35, 36, 0, 37,
	38, 153, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 153, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 29, 40, 0, 0, 0, 31, 32, 34,
	33, 164, 165, 166, 167, 168, 169, 170, 171, 172,
	173, 174, 175, 176, 177, 178, 179, 180, 181, 182,
	183, 184, 185, 186, 187, 188, 189, 190, 191, 192,
	193, 194, 195, 196, 197, 198, 199, 200, 201, 202,
	203,
}
var yyPact = [...]int{

	629, -1000, -1000, 344, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 364, -1000, -1000, -45, -1000, -1000, -1000, -1000, -1000,
	95, -110, -105, -85, -97, -1000, 66, -1000, -1000, 83,
	-71, 555, 488, -1000, -1000, -1000, -1000, 486, -1000, 182,
	72, -1000, -52, -1000, 394, -139, 394, 528, 454, 344,
	-1000, -1000, -1000, -1000, -143, -91, 83, -1000, -89, 83,
	-1000, 410, -149, 83, -149, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 101, -1000, 364, 158, 440, 698, 698, -1000,
	-1000, 438, 254, 254, -22, 254, 254, 300, -41, 526,
	-13, -16, 522, 521, 518, 513, -65, -1000, -17, -1000,
	-1000, 98, 448, 434, 394, 394, 409, 170, 83, -1000,
	408, -1000, -103, 407, 482, 228, 83, 363, -1000, -1000,
	94, 91, 159, 628, -1000, 544, 537, -1000, -1000, -1000,
	441, 360, 357, -1000, 352, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 441, 312, 61, -1000, 124,
	-1000, 87, 698, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -79, 254, -1000, 441, 544, -1000,
	254, 254, -1000, -1000, -1000, 394, 212, 512, -1000, 394,
	254, 254, 394, 394, 394, 394, 394, 394, 394, 394,
	394, 394, -1000, 403, 394, 92, 508, 384, -1000, 480,
	-153, -1000, 82, -1000, 402, -1000, -1000, 401, -1000, -1000,
	356, 101, 441, -1000, -1000, 83, 74, 544, 544, 441,
	341, 219, 441, 441, 240, 441, 441, 441, 441, 441,
	441, 441, 441, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 628, -36, 16, -11, 628, -1000, 590, 18, 101,
	-1000, 555, 136, 126, 92, 49, 441, 83, -1000, 387,
	-1000, 126, 159, -1000, -1000, 254, -1000, 394, 394, 254,
	-1000, -1000, 508, 508, 508, 254, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 375, 379, 499, 544, 477, 92, 92,
	-1000, -1000, 227, 83, -1000, -112, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 373, 538, 386, 377, 84, -1000,
	-1000, 69, -1000, -1000, -1000, -1000, 178, 126, -1000, 341,
	441, 441, 126, 562, -1000, 468, 272, 239, -1000, 131,
	131, 100, 100, 100, -1000, -1000, 441, -1000, 126, -1000,
	14, 101, 11, 181, -1000, 544, -1000, 103, 126, -1000,
	-1000, 254, 254, -1000, -1000, -1000, -1000, -1000, 477, 92,
	499, 492, 496, 159, -1000, 341, 344, 312, -2, -1000,
	393, -1000, -1000, 392, -1000, 506, 356, 356, -1000, -1000,
	318, 314, 297, 286, 267, 10, -1000, 391, -3, 385,
	441, 441, -1000, 126, 415, 441, -1000, 126, -1000, -4,
	-1000, 78, -1000, 441, 284, -1000, 64, 40, -1000, -1000,
	-1000, 222, 302, 492, -1000, 441, 326, -1000, -1000, 92,
	-1000, -1000, 501, 494, 538, 203, -1000, 257, -1000, 249,
	-1000, -1000, -1000, -1000, -92, -94, -100, -1000, -1000, -1000,
	126, 126, 441, 126, -1000, -1000, 126, 441, -1000, -1000,
	-1000, -1000, 446, -1000, -1000, 323, -1000, 371, 341, -1000,
	-1000, 499, 544, 441, 544, -1000, -1000, 338, 334, 292,
	126, 126, 443, 441, -1000, -1000, -1000, -1000, 492, 159,
	313, 159, 83, 83, 83, 539, -1000, 429, -10, -1000,
	-18, -24, 92, -1000, 534, 109, -1000, 83, -1000, -1000,
	312, -1000, 83, -1000, 83, -1000,
}
var yyPgo = [...]int{

	0, 679, 678, 18, 677, 667, 666, 468, 457, 456,
	455, 665, 664, 663, 659, 657, 656, 655, 654, 653,
	651, 646, 327, 650, 649, 311, 16, 25, 648, 647,
	631, 628, 14, 620, 619, 324, 617, 4, 24, 29,
	613, 612, 13, 610, 2, 15, 5, 609, 608, 12,
	607, 6, 605, 603, 8, 602, 601, 600, 598, 9,
	596, 1, 594, 7, 593, 17, 582, 11, 3, 20,
	91, 248, 575, 570, 569, 567, 566, 0, 10, 104,
	563, 289, 562,
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
	5, 5, 5, 5, 5, 5, 5, 6, 20, 20,
	82, 21, 22, 22, 23, 23, 23, 23, 23, 24,
	24, 26, 26, 27, 27, 27, 30, 30, 28, 28,
	28, 31, 31, 32, 32, 32, 32, 29, 29, 29,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 34,
	34, 34, 35, 35, 36, 36, 36, 36, 37, 37,
	38, 38, 39, 39, 39, 39, 39, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 41, 41, 41,
	41, 41, 41, 41, 42, 42, 47, 47, 45, 45,
	49, 46, 46, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	48, 48, 50, 50, 50, 52, 55, 55, 53, 53,
	54, 56, 56, 51, 51, 43, 43, 43, 43, 57,
	57, 58, 58, 59, 59, 60, 60, 61, 62, 62,
	62, 63, 63, 63, 63, 64, 64, 64, 65, 65,
	66, 66, 67, 67, 68, 68, 69, 71, 71, 72,
	72, 25, 25, 73, 73, 73, 73, 73, 74, 74,
	75, 75, 76, 76, 77, 77, 78, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 79, 79, 79,
	79, 79, 79, 79, 79, 79, 79, 80, 80, 80,
	80, 81, 81, 81, 70, 70, 70,
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
	5, 3, 4, 4, 2, 3, 2, 3, 1, 2,
	0, 2, 0, 2, 1, 2, 1, 1, 1, 0,
	1, 1, 3, 1, 2, 3, 1, 1, 0, 1,
	2, 1, 3, 3, 3, 3, 5, 0, 1, 2,
	1, 1, 2, 3, 2, 3, 2, 2, 2, 1,
	3, 1, 1, 3, 0, 5, 5, 5, 1, 3,
	0, 2, 1, 3, 3, 2, 3, 3, 3, 4,
	3, 4, 5, 6, 3, 4, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 1, 1, 3, 3, 1,
	3, 1, 3, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 3, 4, 5, 4, 1,
	1, 1, 1, 1, 1, 5, 0, 1, 1, 2,
	4, 0, 2, 1, 3, 1, 1, 1, 1, 0,
	3, 0, 2, 0, 3, 1, 3, 2, 0, 1,
	1, 0, 2, 4, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 1, 0, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 2, 2,
	1, 0, 1, 1, 0, 2, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -19, -11, -12, -13, -14, -15, -16, -17, -18,
	-20, -22, 5, 29, 31, 33, 6, 7, 8, 163,
	32, 168, 169, 171, 170, 87, 88, 90, 91, 55,
	164, -23, 41, 42, 43, 44, 38, -21, -82, -21,
	-21, 151, 150, 160, -21, -21, -21, -21, -21, -3,
	-7, -8, -10, -9, 172, -75, 174, 178, -25, 174,
	176, 172, 172, 173, 174, 89, -77, 34, 149, 165,
	-3, 17, -24, 18, -22, -81, 101, 100, 99, 143,
	144, 101, 100, 102, -81, 147, 148, 152, 46, 153,
	154, 155, 173, 156, 157, 159, 168, 161, 162, 151,
	-35, 34, -25, -35, 9, 25, -72, 177, 173, -77,
	172, -77, 34, -71, 177, -77, -71, -26, -27, 80,
	-30, 34, -39, -44, -40, 60, 39, -43, -51, -45,
	-50, -77, -48, -52, 20, 35, 36, 37, 21, -49,
	78, 79, 40, 177, 24, 62, -68, 89, -69, -51,
	-77, 34, 29, -79, 103, 104, 105, 106, 107, 108,
	109, 110, 111, 112, 113, 114, 115, 116, 117, 118,
	119, 120, 121, 122, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	139, 140, 141, 142, -79, 29, -70, 74, 10, -70,
	145, 146, -70, -70, -70, 9, 152, 153, 161, 9,
	146, 146, 9, 9, 9, 9, 149, 172, 174, 154,
	155, 158, 146, 84, 25, 29, -35, -35, 34, 60,
	-77, -78, 34, -78, 175, 34, 20, 57, -77, -63,
	9, 45, 15, -28, -77, 19, 84, 59, 58, -41,
	75, 60, 74, 61, 73, 77, 76, 83, 78, 79,
	80, 81, 82, 66, 67, 68, 69, 70, 71, 72,
	-39, -44, -39, -46, -3, -44, -44, 39, 39, 39,
	-49, 39, -55, -44, 45, 92, 66, 84, -79, 167,
	-70, -44, -39, -70, -70, -35, -70, 9, 9, -35,
	-70, -70, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, 34, -35, -68, -38, 10, -65, 29, 39,
	-78, 20, -76, 179, -73, 171, 169, 28, 170, 13,
	34, 34, 34, -78, -31, -32, -34, 39, 34, -49,
	-27, -44, -77, 80, -77, -39, -39, -44, -45, 75,
	74, 61, -44, -44, 21, 60, -44, -44, -44, -44,
	-44, -44, -44, -44, 180, 180, 45, 180, -44, 180,
	-26, 18, -26, -53, -54, 63, -69, 93, -44, 35,
	-70, -35, -35, -70, -38, -38, -38, -70, -65, 29,
	-38, -59, 13, -39, -42, 24, -3, -68, -66, -51,
	57, -77, -78, -74, 175, -38, 45, -33, 47, 48,
	49, 50, 51, 53, 54, -29, 34, 19, -32, 84,
	45, 166, -45, -44, -44, 59, 21, -44, 180, -26,
	180, -56, -54, 65, -39, -80, 94, 97, 98, -70,
	-70, -42, -68, -59, -63, 14, -47, -45, 180, 45,
	34, 34, -57, 11, -32, -32, 47, 52, 47, 52,
	47, 47, 47, -36, 55, 176, 56, 34, 180, 34,
	-44, -44, 59, -44, 180, 86, -44, 64, 95, 96,
	94, -67, 57, -67, -63, -60, -61, -44, 45, -51,
	-78, -58, 12, 14, 57, 47, 47, 173, 173, 173,
	-44, -44, 26, 45, -62, 22, 23, -45, -59, -39,
	-46, -39, 39, 39, 39, 27, -61, -63, -37, -77,
	-37, -37, 7, -64, 16, 30, 180, 45, 180, 180,
	-68, 7, 75, -77, -77, -77,
}
var yyDef = [...]int{

	92, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 90, 90, 90, 90, 90, 90, 90, 90,
	0, 250, 241, 0, 0, 41, 0, 43, 44, 0,
	88, 0, 94, 96, 97, 98, 93, 99, 92, 301,
	301, 84, 0, 86, 0, 241, 0, 0, 0, 30,
	31, 32, 33, 34, 239, 0, 0, 251, 0, 0,
	242, 0, 237, 0, 237, 42, 45, 254, 255, 89,
	23, 95, 0, 100, 91, 0, 0, 0, 0, 302,
	303, 0, 304, 304, 0, 304, 304, 304, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 81, 0, 85,
	87, 132, 0, 0, 0, 0, 0, 0, 0, 256,
	0, 256, 0, 0, 0, 0, 0, 221, 101, 103,
	108, 254, 106, 107, 142, 0, 0, 173, 174, 175,
	0, 203, 0, 189, 0, 205, 206, 207, 208, 169,
	192, 193, 194, 190, 191, 196, 35, 0, 234, 0,
	203, 254, 0, 37, 257, 258, 259, 260, 261, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 282,
	283, 284, 285, 286, 287, 288, 289, 290, 291, 292,
	293, 294, 295, 296, 38, 304, 57, 0, 0, 58,
	304, 304, 61, 62, 63, 0, 304, 0, 82, 0,
	304, 304, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 83, 0, 0, 0, 140, 228, 256, 0,
	252, 48, 0, 51, 0, 53, 238, 0, 256, 21,
	0, 0, 0, 104, 109, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 157, 158, 159, 160, 161, 162, 163,
	145, 0, 0, 0, 0, 171, 184, 0, 0, 0,
	156, 0, 0, 197, 0, 0, 0, 0, 36, 0,
	56, 305, 306, 59, 60, 304, 65, 0, 0, 304,
	69, 70, 140, 140, 140, 304, 75, 76, 77, 78,
	79, 80, 133, 228, 140, 213, 0, 0, 0, 0,
	46, 240, 0, 0, 256, 248, 243, 244, 245, 246,
	247, 52, 54, 55, 140, 111, 117, 0, 129, 131,
	102, 222, 110, 105, 204, 143, 144, 147, 148, 0,
	0, 0, 150, 0, 154, 0, 176, 177, 178, 179,
	180, 181, 182, 183, 146, 168, 0, 170, 171, 185,
	0, 0, 0, 201, 198, 0, 235, 0, 236, 39,
	64, 304, 304, 67, 71, 72, 73, 74, 0, 0,
	213, 221, 0, 141, 26, 0, 165, 27, 0, 230,
	0, 253, 49, 0, 249, 209, 0, 0, 120, 121,
	0, 0, 0, 0, 0, 134, 118, 0, 0, 0,
	0, 0, 149, 151, 0, 0, 155, 172, 186, 0,
	188, 0, 199, 0, 0, 40, 0, 0, 300, 66,
	68, 232, 232, 221, 29, 0, 164, 166, 229, 0,
	256, 50, 211, 0, 112, 115, 122, 0, 124, 0,
	126, 127, 128, 113, 0, 0, 0, 119, 114, 130,
	223, 224, 0, 152, 187, 195, 202, 0, 297, 298,
	299, 24, 0, 25, 28, 214, 215, 218, 0, 231,
	47, 213, 0, 0, 0, 123, 125, 0, 0, 0,
	153, 200, 0, 0, 217, 219, 220, 167, 221, 212,
	210, 116, 0, 0, 0, 0, 216, 225, 0, 138,
	0, 0, 0, 22, 0, 0, 135, 0, 136, 137,
	233, 226, 0, 139, 0, 227,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:535
		{
			yyVAL.showStmt = &ShowSlaveStatus{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:539
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:543
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:547
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:553
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[3].tableName}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:560
		{
			yyVAL.statement = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:562
		{
			yyVAL.statement = nil
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:565
		{
			SetAllowComments(yylex, true)
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:569
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:575
		{
			yyVAL.bytes2 = nil
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:579
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:585
		{
			yyVAL.str = AST_UNION
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:589
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:593
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:597
		{
			yyVAL.str = AST_EXCEPT
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:601
		{
			yyVAL.str = AST_INTERSECT
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:606
		{
			yyVAL.str = ""
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:610
		{
			yyVAL.str = AST_DISTINCT
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:616
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:620
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:626
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:630
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:634
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:640
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:644
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:649
		{
			yyVAL.bytes = nil
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:653
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:657
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:663
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:667
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:673
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:677
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:681
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:685
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:690
		{
			yyVAL.bytes = nil
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:694
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:698
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:704
		{
			yyVAL.str = AST_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:708
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:712
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:716
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:720
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:724
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:728
		{
			yyVAL.str = AST_JOIN
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:732
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:736
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:742
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:746
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:750
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:756
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:760
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 134:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:765
		{
			yyVAL.indexHints = nil
		}
	case 135:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:769
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 136:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:773
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 137:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:777
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:783
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:787
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 140:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:792
		{
			yyVAL.boolExpr = nil
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:796
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:803
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:807
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:811
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:815
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:821
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:825
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 149:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:829
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:833
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:837
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:841
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:845
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:849
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:853
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:857
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:863
		{
			yyVAL.str = AST_EQ
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:867
		{
			yyVAL.str = AST_LT
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:871
		{
			yyVAL.str = AST_GT
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:875
		{
			yyVAL.str = AST_LE
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:879
		{
			yyVAL.str = AST_GE
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:883
		{
			yyVAL.str = AST_NE
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:887
		{
			yyVAL.str = AST_NSE
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:893
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:897
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:903
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:907
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 168:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:913
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:917
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:923
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:929
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:933
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:939
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:943
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:947
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:951
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:955
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:959
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:963
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:967
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:971
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:975
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:979
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:983
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
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:998
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1002
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 187:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1006
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1010
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1014
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1020
		{
			yyVAL.bytes = IF_BYTES
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1024
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1030
		{
			yyVAL.byt = AST_UPLUS
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1034
		{
			yyVAL.byt = AST_UMINUS
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1038
		{
			yyVAL.byt = AST_TILDA
		}
	case 195:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1044
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1049
		{
			yyVAL.valExpr = nil
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1053
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1059
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1063
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 200:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1069
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1074
		{
			yyVAL.valExpr = nil
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1078
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1084
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 204:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1088
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1094
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1098
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1102
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1106
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1111
		{
			yyVAL.valExprs = nil
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1115
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1120
		{
			yyVAL.boolExpr = nil
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1124
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1129
		{
			yyVAL.orderBy = nil
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1133
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1139
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1143
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1149
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1154
		{
			yyVAL.str = AST_ASC
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1158
		{
			yyVAL.str = AST_ASC
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1162
		{
			yyVAL.str = AST_DESC
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1167
		{
			yyVAL.limit = nil
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1171
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 223:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1175
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 224:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1179
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1184
		{
			yyVAL.str = ""
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1188
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 227:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1192
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
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1205
		{
			yyVAL.columns = nil
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1209
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1215
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 231:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1219
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1224
		{
			yyVAL.updateExprs = nil
		}
	case 233:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1228
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1234
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1238
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 236:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1244
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1249
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1251
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1254
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1256
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1259
		{
			yyVAL.str = ""
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1261
		{
			yyVAL.str = AST_IGNORE
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1269
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1271
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1276
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1278
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1281
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1283
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1286
		{
			yyVAL.empty = struct{}{}
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1288
		{
			yyVAL.empty = struct{}{}
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1292
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1296
		{
			yyVAL.bytes = []byte("database")
		}
	case 256:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1301
		{
			ForceEOF(yylex)
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1307
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1309
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1311
		{
			yyVAL.bytes = []byte("big5")
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1313
		{
			yyVAL.bytes = []byte("binary")
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1315
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1317
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1319
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1321
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1323
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1325
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1327
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1329
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1331
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1333
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1335
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1337
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1339
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1341
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1343
		{
			yyVAL.bytes = []byte("greek")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1345
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1347
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1349
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1351
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1353
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1355
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1357
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1359
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1361
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1363
		{
			yyVAL.bytes = []byte("macce")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1365
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1367
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1369
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1371
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1373
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1375
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1377
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1379
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1381
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1383
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1385
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 297:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1389
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 298:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1391
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 299:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1393
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 300:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1395
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 301:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1398
		{
			yyVAL.bytes = nil
		}
	case 302:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1400
		{
			yyVAL.bytes = []byte("session")
		}
	case 303:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1402
		{
			yyVAL.bytes = []byte("global")
		}
	case 304:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1405
		{
			yyVAL.expr = nil
		}
	case 305:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1407
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 306:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1411
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
