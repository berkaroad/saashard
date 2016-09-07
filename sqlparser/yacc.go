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
const ID = 57375
const STRING = 57376
const NUMBER = 57377
const VALUE_ARG = 57378
const COMMENT = 57379
const UNION = 57380
const MINUS = 57381
const EXCEPT = 57382
const INTERSECT = 57383
const FULL = 57384
const JOIN = 57385
const STRAIGHT_JOIN = 57386
const LEFT = 57387
const RIGHT = 57388
const INNER = 57389
const OUTER = 57390
const CROSS = 57391
const NATURAL = 57392
const USE = 57393
const FORCE = 57394
const ON = 57395
const OR = 57396
const AND = 57397
const NOT = 57398
const BETWEEN = 57399
const CASE = 57400
const WHEN = 57401
const THEN = 57402
const ELSE = 57403
const LE = 57404
const GE = 57405
const NE = 57406
const NULL_SAFE_EQUAL = 57407
const IS = 57408
const LIKE = 57409
const IN = 57410
const UNARY = 57411
const END = 57412
const BEGIN = 57413
const START = 57414
const TRANSACTION = 57415
const COMMIT = 57416
const ROLLBACK = 57417
const ISOLATION = 57418
const LEVEL = 57419
const READ = 57420
const COMMITTED = 57421
const UNCOMMITTED = 57422
const REPEATABLE = 57423
const SERIALIZABLE = 57424
const NAMES = 57425
const CHARSET = 57426
const CHARACTER = 57427
const COLLATION = 57428
const ARMSCII8 = 57429
const ASCII = 57430
const BIG5 = 57431
const BINARY = 57432
const CP1250 = 57433
const CP1251 = 57434
const CP1256 = 57435
const CP1257 = 57436
const CP850 = 57437
const CP852 = 57438
const CP866 = 57439
const CP932 = 57440
const DEC8 = 57441
const EUCJPMS = 57442
const EUCKR = 57443
const GB2312 = 57444
const GBK = 57445
const GEOSTD8 = 57446
const GREEK = 57447
const HEBREW = 57448
const HP8 = 57449
const KEYBCS2 = 57450
const KOI8R = 57451
const KOI8U = 57452
const LATIN1 = 57453
const LATIN2 = 57454
const LATIN5 = 57455
const LATIN7 = 57456
const MACCE = 57457
const MACROMAN = 57458
const SJIS = 57459
const SWE7 = 57460
const TIS620 = 57461
const UCS2 = 57462
const EJIS = 57463
const UTF16 = 57464
const UTF16LE = 57465
const UTF32 = 57466
const UTF8 = 57467
const UTF8MB4 = 57468
const SESSION = 57469
const GLOBAL = 57470
const VARIABLES = 57471
const STATUS = 57472
const DATABASES = 57473
const SCHEMAS = 57474
const DATABASE = 57475
const STORAGE = 57476
const ENGINES = 57477
const TABLES = 57478
const COLUMNS = 57479
const PROCEDURE = 57480
const FUNCTION = 57481
const INDEXES = 57482
const KEYS = 57483
const TRIGGER = 57484
const TRIGGERS = 57485
const PLUGINS = 57486
const PROCESSLIST = 57487
const REPLACE = 57488
const ADMIN = 57489
const HELP = 57490
const OFFSET = 57491
const COLLATE = 57492
const CREATE = 57493
const ALTER = 57494
const DROP = 57495
const RENAME = 57496
const TABLE = 57497
const INDEX = 57498
const VIEW = 57499
const TO = 57500
const IGNORE = 57501
const IF = 57502
const UNIQUE = 57503
const USING = 57504

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

const yyNprod = 303
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 750

var yyAct = [...]int{

	135, 489, 274, 150, 521, 276, 132, 242, 338, 394,
	234, 484, 143, 277, 3, 133, 397, 320, 377, 152,
	251, 250, 121, 326, 318, 530, 530, 126, 122, 530,
	411, 412, 413, 414, 415, 244, 416, 417, 73, 452,
	244, 244, 56, 40, 41, 42, 43, 467, 469, 61,
	369, 63, 118, 76, 67, 64, 111, 374, 407, 138,
	142, 237, 502, 148, 113, 114, 66, 115, 67, 501,
	500, 119, 125, 139, 140, 141, 112, 130, 146, 68,
	292, 127, 154, 69, 70, 71, 75, 104, 153, 138,
	142, 50, 49, 148, 215, 200, 204, 205, 129, 214,
	149, 51, 125, 139, 140, 141, 483, 130, 146, 84,
	83, 82, 380, 233, 220, 157, 144, 145, 123, 223,
	224, 241, 288, 225, 72, 247, 236, 478, 129, 422,
	149, 210, 211, 278, 481, 482, 221, 279, 222, 290,
	212, 367, 21, 24, 25, 26, 144, 145, 123, 21,
	249, 283, 286, 85, 86, 155, 228, 273, 275, 532,
	531, 106, 471, 529, 138, 142, 260, 468, 148, 477,
	535, 117, 74, 451, 433, 431, 289, 155, 139, 140,
	141, 370, 130, 146, 368, 203, 378, 206, 207, 208,
	439, 202, 250, 440, 441, 378, 332, 436, 251, 250,
	198, 232, 497, 129, 294, 149, 263, 264, 265, 260,
	151, 330, 354, 485, 147, 357, 333, 372, 346, 403,
	240, 144, 145, 300, 202, 353, 352, 65, 154, 20,
	295, 316, 209, 202, 153, 258, 261, 262, 263, 264,
	265, 260, 323, 120, 147, 251, 250, 127, 344, 345,
	347, 480, 336, 358, 201, 350, 342, 499, 355, 356,
	81, 359, 360, 361, 362, 363, 364, 365, 366, 351,
	229, 230, 291, 343, 287, 498, 80, 461, 348, 349,
	105, 287, 462, 371, 127, 127, 485, 201, 154, 465,
	464, 347, 381, 463, 153, 293, 201, 459, 27, 369,
	296, 297, 460, 319, 373, 375, 299, 379, 506, 90,
	303, 304, 261, 262, 263, 264, 265, 260, 491, 147,
	243, 517, 154, 154, 319, 400, 245, 404, 153, 402,
	387, 388, 389, 391, 399, 516, 515, 409, 405, 392,
	44, 393, 40, 41, 42, 43, 280, 396, 322, 421,
	329, 331, 328, 342, 341, 244, 426, 427, 287, 340,
	284, 282, 408, 411, 412, 413, 414, 415, 425, 416,
	417, 298, 430, 281, 382, 302, 420, 127, 305, 306,
	307, 308, 309, 310, 311, 312, 313, 314, 315, 74,
	419, 321, 21, 154, 383, 435, 445, 432, 386, 153,
	322, 248, 447, 446, 390, 399, 437, 155, 444, 472,
	470, 454, 453, 107, 450, 74, 335, 334, 457, 458,
	341, 317, 342, 342, 238, 340, 473, 474, 138, 142,
	235, 476, 148, 231, 116, 527, 227, 9, 199, 479,
	518, 155, 139, 140, 141, 156, 130, 146, 505, 528,
	8, 490, 226, 154, 487, 7, 94, 486, 423, 492,
	109, 429, 384, 385, 493, 324, 59, 129, 6, 149,
	259, 258, 261, 262, 263, 264, 265, 260, 503, 60,
	442, 443, 21, 504, 58, 144, 145, 239, 79, 259,
	258, 261, 262, 263, 264, 265, 260, 57, 77, 371,
	245, 398, 513, 496, 511, 448, 395, 510, 519, 490,
	88, 87, 89, 495, 21, 456, 522, 522, 522, 520,
	319, 523, 524, 512, 301, 514, 154, 219, 218, 533,
	142, 536, 153, 148, 217, 216, 537, 213, 538, 108,
	534, 525, 155, 139, 140, 141, 21, 280, 146, 508,
	509, 46, 438, 85, 86, 325, 62, 91, 92, 406,
	327, 110, 93, 95, 96, 97, 99, 100, 401, 101,
	149, 103, 21, 24, 25, 26, 142, 102, 424, 148,
	526, 507, 98, 147, 488, 494, 144, 145, 155, 139,
	140, 141, 455, 280, 146, 434, 22, 285, 23, 28,
	376, 137, 259, 258, 261, 262, 263, 264, 265, 260,
	134, 136, 449, 131, 252, 128, 149, 466, 339, 410,
	337, 37, 124, 418, 246, 78, 39, 19, 254, 256,
	10, 18, 144, 145, 266, 267, 268, 269, 270, 271,
	272, 257, 255, 253, 259, 258, 261, 262, 263, 264,
	265, 260, 475, 33, 34, 45, 35, 36, 17, 16,
	15, 14, 13, 12, 11, 5, 4, 2, 1, 259,
	258, 261, 262, 263, 264, 265, 260, 0, 47, 48,
	52, 53, 54, 55, 147, 158, 159, 160, 161, 162,
	163, 164, 165, 166, 167, 168, 169, 170, 171, 172,
	173, 174, 175, 176, 177, 178, 179, 180, 181, 182,
	183, 184, 185, 186, 187, 188, 189, 190, 191, 192,
	193, 194, 195, 196, 197, 428, 0, 0, 27, 38,
	147, 0, 0, 29, 30, 32, 31, 0, 0, 0,
	0, 0, 259, 258, 261, 262, 263, 264, 265, 260,
}
var yyPact = [...]int{

	567, -1000, -1000, 302, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	303, -1000, -1000, -58, -1000, -1000, -1000, -1000, 137, -121,
	-106, -91, -87, -1000, 36, -1000, -1000, 356, -77, 541,
	481, -1000, -1000, -1000, -1000, 470, -1000, 11, 411, -1000,
	-63, -1000, -120, 380, 530, 435, 302, -1000, -1000, -1000,
	-1000, -119, -95, 356, -1000, -105, 356, -1000, 401, -123,
	356, -123, -1000, -1000, -1000, -1000, -1000, -1000, 69, -1000,
	303, 122, 416, 583, 583, -1000, -1000, 409, 181, 181,
	-48, 181, 181, 223, -20, 528, -46, -51, 526, 525,
	519, 518, -34, -1000, -1000, 427, 407, 73, 380, 380,
	400, 142, 356, -1000, 397, -1000, -112, 391, 467, 164,
	356, 311, -1000, -1000, 382, 67, 141, 569, -1000, 408,
	144, -1000, -1000, -1000, 555, 335, 323, -1000, 322, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 555,
	237, 31, -1000, 111, -1000, 56, 583, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -85, 181,
	-1000, 555, 408, -1000, 181, 181, -1000, -1000, -1000, 380,
	214, 515, -1000, 380, 181, 181, 380, 380, 380, 380,
	380, 380, 380, 380, 380, 380, 380, 374, 388, 510,
	362, -1000, 445, -154, -1000, 183, -1000, 384, -1000, -1000,
	383, -1000, -1000, 321, 69, 555, -1000, -1000, 356, 139,
	408, 408, 555, 308, 152, 555, 555, 194, 555, 555,
	555, 555, 555, 555, 555, 555, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 569, -37, 6, 3, 569, -1000,
	509, 39, 69, -1000, 541, 124, 395, 374, 20, 555,
	356, -1000, 340, -1000, 395, 141, -1000, -1000, 181, -1000,
	380, 380, 181, -1000, -1000, 510, 510, 510, 181, -1000,
	-1000, -1000, -1000, -1000, -1000, 310, 314, -1000, 493, 408,
	477, 374, 374, -1000, -1000, 163, 356, -1000, -115, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 293, 317, 357,
	387, 46, -1000, -1000, 414, -1000, -1000, -1000, -1000, 134,
	395, -1000, 308, 555, 555, 395, 667, -1000, 440, 235,
	159, -1000, 127, 127, 84, 84, 84, -1000, -1000, 555,
	-1000, 395, -1000, -3, 69, -4, 133, -1000, 408, -1000,
	97, 395, -1000, -1000, 181, 181, -1000, -1000, -1000, -1000,
	-1000, 477, 374, 493, 485, 491, 141, -1000, 308, 302,
	237, -5, -1000, 379, -1000, -1000, 378, -1000, 504, 321,
	321, -1000, -1000, 251, 231, 247, 244, 243, -7, -1000,
	377, -16, 376, 555, 555, -1000, 395, 594, 555, -1000,
	395, -1000, -9, -1000, 42, -1000, 555, 188, -1000, 40,
	13, -1000, -1000, -1000, 157, 230, 485, -1000, 555, 274,
	-1000, -1000, 374, -1000, -1000, 501, 489, 317, 146, -1000,
	229, -1000, 211, -1000, -1000, -1000, -1000, -101, -102, -109,
	-1000, -1000, -1000, 395, 395, 555, 395, -1000, -1000, 395,
	555, -1000, -1000, -1000, -1000, 422, -1000, -1000, 264, -1000,
	527, 308, -1000, -1000, 493, 408, 555, 408, -1000, -1000,
	298, 297, 283, 395, 395, 413, 555, -1000, -1000, -1000,
	-1000, 485, 141, 255, 141, 356, 356, 356, 534, -1000,
	419, -15, -1000, -18, -19, 374, -1000, 533, 96, -1000,
	356, -1000, -1000, 237, -1000, 356, -1000, 356, -1000,
}
var yyPgo = [...]int{

	0, 668, 667, 13, 666, 665, 468, 455, 450, 437,
	664, 663, 662, 661, 660, 659, 658, 631, 630, 627,
	655, 229, 626, 625, 227, 22, 28, 624, 623, 622,
	620, 8, 619, 618, 161, 617, 4, 24, 27, 615,
	614, 16, 613, 2, 15, 5, 612, 611, 12, 610,
	6, 601, 600, 18, 597, 595, 592, 585, 9, 584,
	1, 581, 7, 580, 17, 568, 11, 3, 19, 95,
	171, 561, 560, 559, 556, 555, 0, 10, 115, 552,
	260, 551,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 6, 6, 9, 9, 7, 8, 18,
	18, 18, 18, 18, 4, 4, 4, 4, 4, 4,
	14, 14, 15, 16, 17, 10, 10, 10, 11, 11,
	11, 12, 13, 13, 13, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 5, 19, 19, 81, 20, 21,
	21, 22, 22, 22, 22, 22, 23, 23, 25, 25,
	26, 26, 26, 29, 29, 27, 27, 27, 30, 30,
	31, 31, 31, 31, 28, 28, 28, 32, 32, 32,
	32, 32, 32, 32, 32, 32, 33, 33, 33, 34,
	34, 35, 35, 35, 35, 36, 36, 37, 37, 38,
	38, 38, 38, 38, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 40, 40, 40, 40, 40, 40,
	40, 41, 41, 46, 46, 44, 44, 48, 45, 45,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 47, 47, 49,
	49, 49, 51, 54, 54, 52, 52, 53, 55, 55,
	50, 50, 42, 42, 42, 42, 56, 56, 57, 57,
	58, 58, 59, 59, 60, 61, 61, 61, 62, 62,
	62, 62, 63, 63, 63, 64, 64, 65, 65, 66,
	66, 67, 67, 68, 70, 70, 71, 71, 24, 24,
	72, 72, 72, 72, 72, 73, 73, 74, 74, 75,
	75, 76, 77, 78, 78, 78, 78, 78, 78, 78,
	78, 78, 78, 78, 78, 78, 78, 78, 78, 78,
	78, 78, 78, 78, 78, 78, 78, 78, 78, 78,
	78, 78, 78, 78, 78, 78, 78, 78, 78, 78,
	78, 78, 78, 79, 79, 79, 79, 80, 80, 80,
	69, 69, 69,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	5, 12, 3, 8, 8, 6, 6, 8, 7, 2,
	2, 2, 2, 2, 4, 5, 4, 4, 6, 7,
	1, 2, 1, 1, 2, 5, 8, 4, 6, 7,
	4, 5, 4, 5, 5, 5, 4, 4, 5, 5,
	4, 4, 4, 6, 5, 7, 6, 7, 5, 5,
	6, 6, 6, 6, 5, 5, 5, 5, 5, 5,
	3, 4, 2, 3, 2, 1, 2, 0, 2, 0,
	2, 1, 2, 1, 1, 1, 0, 1, 1, 3,
	1, 2, 3, 1, 1, 0, 1, 2, 1, 3,
	3, 3, 3, 5, 0, 1, 2, 1, 1, 2,
	3, 2, 3, 2, 2, 2, 1, 3, 1, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 1,
	3, 3, 2, 3, 3, 3, 4, 3, 4, 5,
	6, 3, 4, 2, 1, 1, 1, 1, 1, 1,
	1, 2, 1, 1, 3, 3, 1, 3, 1, 3,
	1, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 3, 4, 5, 4, 1, 1, 1, 1,
	1, 1, 5, 0, 1, 1, 2, 4, 0, 2,
	1, 3, 1, 1, 1, 1, 0, 3, 0, 2,
	0, 3, 1, 3, 2, 0, 1, 1, 0, 2,
	4, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 1, 0, 1, 1,
	0, 2, 2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-18, -10, -11, -12, -13, -14, -15, -16, -17, -19,
	-21, 5, 29, 31, 6, 7, 8, 161, 32, 166,
	167, 169, 168, 86, 87, 89, 90, 54, 162, -22,
	40, 41, 42, 43, 37, -20, -81, -20, -20, 150,
	149, 159, -20, -20, -20, -20, -3, -6, -7, -9,
	-8, 170, -74, 172, 176, -24, 172, 174, 170, 170,
	171, 172, 88, -76, 33, 163, -3, 17, -23, 18,
	-21, -80, 100, 99, 98, 142, 143, 100, 99, 101,
	-80, 146, 147, 151, 45, 152, 153, 154, 171, 155,
	156, 158, 166, 160, 150, -24, -34, 33, 9, 25,
	-71, 175, 171, -76, 170, -76, 33, -70, 175, -76,
	-70, -25, -26, 79, -29, 33, -38, -43, -39, 59,
	38, -42, -50, -44, -49, -76, -47, -51, 20, 34,
	35, 36, 21, -48, 77, 78, 39, 175, 24, 61,
	-67, 88, -68, -50, -76, 33, 29, -78, 102, 103,
	104, 105, 106, 107, 108, 109, 110, 111, 112, 113,
	114, 115, 116, 117, 118, 119, 120, 121, 122, 123,
	124, 125, 126, 127, 128, 129, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 141, -78, 29,
	-69, 73, 10, -69, 144, 145, -69, -69, -69, 9,
	151, 152, 160, 9, 145, 145, 9, 9, 9, 9,
	148, 170, 172, 153, 154, 157, 25, 29, 83, -34,
	-34, 33, 59, -76, -77, 33, -77, 173, 33, 20,
	56, -76, -62, 9, 44, 15, -27, -76, 19, 83,
	58, 57, -40, 74, 59, 73, 60, 72, 76, 75,
	82, 77, 78, 79, 80, 81, 65, 66, 67, 68,
	69, 70, 71, -38, -43, -38, -45, -3, -43, -43,
	38, 38, 38, -48, 38, -54, -43, 44, 91, 65,
	83, -78, 165, -69, -43, -38, -69, -69, -34, -69,
	9, 9, -34, -69, -69, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -67, 33, -37, 10,
	-64, 29, 38, -77, 20, -75, 177, -72, 169, 167,
	28, 168, 13, 33, 33, 33, -77, -30, -31, -33,
	38, 33, -48, -26, -43, -76, 79, -76, -38, -38,
	-43, -44, 74, 73, 60, -43, -43, 21, 59, -43,
	-43, -43, -43, -43, -43, -43, -43, 178, 178, 44,
	178, -43, 178, -25, 18, -25, -52, -53, 62, -68,
	92, -43, 34, -69, -34, -34, -69, -37, -37, -37,
	-69, -64, 29, -37, -58, 13, -38, -41, 24, -3,
	-67, -65, -50, 56, -76, -77, -73, 173, -37, 44,
	-32, 46, 47, 48, 49, 50, 52, 53, -28, 33,
	19, -31, 83, 44, 164, -44, -43, -43, 58, 21,
	-43, 178, -25, 178, -55, -53, 64, -38, -79, 93,
	96, 97, -69, -69, -41, -67, -58, -62, 14, -46,
	-44, 178, 44, 33, 33, -56, 11, -31, -31, 46,
	51, 46, 51, 46, 46, 46, -35, 54, 174, 55,
	33, 178, 33, -43, -43, 58, -43, 178, 85, -43,
	63, 94, 95, 93, -66, 56, -66, -62, -59, -60,
	-43, 44, -50, -77, -57, 12, 14, 56, 46, 46,
	171, 171, 171, -43, -43, 26, 44, -61, 22, 23,
	-44, -58, -38, -45, -38, 38, 38, 38, 27, -60,
	-62, -36, -76, -36, -36, 7, -63, 16, 30, 178,
	44, 178, 178, -67, 7, 74, -76, -76, -76,
}
var yyDef = [...]int{

	89, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 87, 87, 87, 87, 87, 87, 87, 0, 247,
	238, 0, 0, 40, 0, 42, 43, 0, 85, 0,
	91, 93, 94, 95, 90, 96, 89, 297, 297, 82,
	0, 84, 238, 0, 0, 0, 29, 30, 31, 32,
	33, 236, 0, 0, 248, 0, 0, 239, 0, 234,
	0, 234, 41, 44, 251, 86, 22, 92, 0, 97,
	88, 0, 0, 0, 0, 298, 299, 0, 300, 300,
	0, 300, 300, 300, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 80, 83, 0, 0, 129, 0, 0,
	0, 0, 0, 252, 0, 252, 0, 0, 0, 0,
	0, 218, 98, 100, 105, 251, 103, 104, 139, 0,
	0, 170, 171, 172, 0, 200, 0, 186, 0, 202,
	203, 204, 205, 166, 189, 190, 191, 187, 188, 193,
	34, 0, 231, 0, 200, 251, 0, 36, 253, 254,
	255, 256, 257, 258, 259, 260, 261, 262, 263, 264,
	265, 266, 267, 268, 269, 270, 271, 272, 273, 274,
	275, 276, 277, 278, 279, 280, 281, 282, 283, 284,
	285, 286, 287, 288, 289, 290, 291, 292, 37, 300,
	56, 0, 0, 57, 300, 300, 60, 61, 62, 0,
	300, 0, 81, 0, 300, 300, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 137,
	225, 252, 0, 249, 47, 0, 50, 0, 52, 235,
	0, 252, 20, 0, 0, 0, 101, 106, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 154, 155, 156, 157,
	158, 159, 160, 142, 0, 0, 0, 0, 168, 181,
	0, 0, 0, 153, 0, 0, 194, 0, 0, 0,
	0, 35, 0, 55, 301, 302, 58, 59, 300, 64,
	0, 0, 300, 68, 69, 137, 137, 137, 300, 74,
	75, 76, 77, 78, 79, 225, 137, 130, 210, 0,
	0, 0, 0, 45, 237, 0, 0, 252, 245, 240,
	241, 242, 243, 244, 51, 53, 54, 137, 108, 114,
	0, 126, 128, 99, 219, 107, 102, 201, 140, 141,
	144, 145, 0, 0, 0, 147, 0, 151, 0, 173,
	174, 175, 176, 177, 178, 179, 180, 143, 165, 0,
	167, 168, 182, 0, 0, 0, 198, 195, 0, 232,
	0, 233, 38, 63, 300, 300, 66, 70, 71, 72,
	73, 0, 0, 210, 218, 0, 138, 25, 0, 162,
	26, 0, 227, 0, 250, 48, 0, 246, 206, 0,
	0, 117, 118, 0, 0, 0, 0, 0, 131, 115,
	0, 0, 0, 0, 0, 146, 148, 0, 0, 152,
	169, 183, 0, 185, 0, 196, 0, 0, 39, 0,
	0, 296, 65, 67, 229, 229, 218, 28, 0, 161,
	163, 226, 0, 252, 49, 208, 0, 109, 112, 119,
	0, 121, 0, 123, 124, 125, 110, 0, 0, 0,
	116, 111, 127, 220, 221, 0, 149, 184, 192, 199,
	0, 293, 294, 295, 23, 0, 24, 27, 211, 212,
	215, 0, 228, 46, 210, 0, 0, 0, 120, 122,
	0, 0, 0, 150, 197, 0, 0, 214, 216, 217,
	164, 218, 209, 207, 113, 0, 0, 0, 0, 213,
	222, 0, 135, 0, 0, 0, 21, 0, 0, 132,
	0, 133, 134, 230, 223, 0, 136, 0, 224,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 81, 76, 3,
	38, 178, 79, 77, 44, 78, 83, 80, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 65, 67, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 82, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 75, 3, 39,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 40, 41, 42, 43,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	68, 69, 70, 71, 72, 73, 74, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165, 166,
	167, 168, 169, 170, 171, 172, 173, 174, 175, 176,
	177,
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
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:215
		{
			yyVAL.statement = nil
		}
	case 20:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:219
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 21:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:223
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:227
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:233
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:237
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:249
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:253
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:265
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:271
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:277
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].selStmt}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:281
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:285
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:289
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:293
		{
			yyVAL.statement = &Explain{Statement: yyDollar[2].statement}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:299
		{
			yyVAL.setStmt = &SetVariable{
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				Exprs:    yyDollar[4].updateExprs,
			}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:307
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[5].bytes),
			}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:314
		{
			yyVAL.setStmt = &SetCharset{
				Comments: Comments(yyDollar[2].bytes2),
				Charset:  string(yyDollar[4].bytes),
			}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:321
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
			}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:328
		{
			yyVAL.setStmt = &SetNames{
				Comments: Comments(yyDollar[2].bytes2),
				Names:    string(yyDollar[4].bytes),
				Collate:  string(yyDollar[6].bytes),
			}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:336
		{
			yyVAL.setStmt = &SetTransactionIsolationLevel{
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				IsolationLevel: string(yyDollar[7].bytes),
			}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:346
		{
			yyVAL.statement = &Begin{}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:350
		{
			yyVAL.statement = &Begin{}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:356
		{
			yyVAL.statement = &Commit{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:362
		{
			yyVAL.statement = &Rollback{}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:368
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:374
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:378
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:383
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:389
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:393
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:398
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:404
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:410
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:414
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:419
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:425
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:429
		{
			yyVAL.showStmt = &ShowCharset{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:433
		{
			yyVAL.showStmt = &ShowCollation{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:437
		{
			yyVAL.showStmt = &ShowVariables{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:441
		{
			yyVAL.showStmt = &ShowStatus{Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), LikeOrWhere: yyDollar[5].expr}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:445
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:449
		{
			yyVAL.showStmt = &ShowDatabases{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:453
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[4].expr}
		}
	case 63:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:457
		{
			yyVAL.showStmt = &ShowTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:461
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 65:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:465
		{
			yyVAL.showStmt = &ShowFullTables{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 66:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:469
		{
			yyVAL.showStmt = &ShowColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 67:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:473
		{
			yyVAL.showStmt = &ShowFullColumns{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:477
		{
			yyVAL.showStmt = &ShowProcedureStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:481
		{
			yyVAL.showStmt = &ShowFunctionStatus{Comments: Comments(yyDollar[2].bytes2), LikeOrWhere: yyDollar[5].expr}
		}
	case 70:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:485
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:489
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 72:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:493
		{
			yyVAL.showStmt = &ShowIndex{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, Where: yyDollar[6].boolExpr}
		}
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:497
		{
			yyVAL.showStmt = &ShowTriggers{Comments: Comments(yyDollar[2].bytes2), From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 74:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:501
		{
			yyVAL.showStmt = &ShowCreateDatabase{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:505
		{
			yyVAL.showStmt = &ShowCreateTable{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:509
		{
			yyVAL.showStmt = &ShowCreateView{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:513
		{
			yyVAL.showStmt = &ShowCreateProcedure{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:517
		{
			yyVAL.showStmt = &ShowCreateFunction{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:521
		{
			yyVAL.showStmt = &ShowCreateTrigger{Comments: Comments(yyDollar[2].bytes2), Name: yyDollar[5].tableName}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:525
		{
			yyVAL.showStmt = &ShowProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:529
		{
			yyVAL.showStmt = &ShowFullProcessList{Comments: Comments(yyDollar[2].bytes2)}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:533
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:537
		{
			yyVAL.showStmt = &ShowEngines{}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:541
		{
			yyVAL.showStmt = &ShowPlugins{}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:547
		{
			yyVAL.statement = nil
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:549
		{
			yyVAL.statement = nil
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:552
		{
			SetAllowComments(yylex, true)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:556
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:562
		{
			yyVAL.bytes2 = nil
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:566
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:572
		{
			yyVAL.str = AST_UNION
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:576
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:580
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:584
		{
			yyVAL.str = AST_EXCEPT
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:588
		{
			yyVAL.str = AST_INTERSECT
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:593
		{
			yyVAL.str = ""
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:597
		{
			yyVAL.str = AST_DISTINCT
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:603
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:607
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:613
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:617
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:621
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:627
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:631
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:636
		{
			yyVAL.bytes = nil
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:640
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:644
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:650
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:654
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:660
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:664
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:668
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:672
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:677
		{
			yyVAL.bytes = nil
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:681
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 116:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:685
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:691
		{
			yyVAL.str = AST_JOIN
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:695
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:699
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:703
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:707
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:711
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:715
		{
			yyVAL.str = AST_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:719
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:723
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:729
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:733
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:737
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:743
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:747
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 131:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:752
		{
			yyVAL.indexHints = nil
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:756
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:760
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 134:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:764
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:770
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:774
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 137:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:779
		{
			yyVAL.boolExpr = nil
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:783
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:790
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:794
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:798
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:802
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:808
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:812
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 146:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:816
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:820
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:824
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:828
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:832
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:836
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:840
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:844
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:850
		{
			yyVAL.str = AST_EQ
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:854
		{
			yyVAL.str = AST_LT
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:858
		{
			yyVAL.str = AST_GT
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:862
		{
			yyVAL.str = AST_LE
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:866
		{
			yyVAL.str = AST_GE
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:870
		{
			yyVAL.str = AST_NE
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:874
		{
			yyVAL.str = AST_NSE
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:880
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:884
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:890
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:894
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:900
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:904
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:910
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:916
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:920
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:926
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:930
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:934
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:938
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:942
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:946
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:950
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:954
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:958
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:962
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:966
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:970
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
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:985
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 183:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:989
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 184:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:993
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:997
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1001
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1007
		{
			yyVAL.bytes = IF_BYTES
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1011
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1017
		{
			yyVAL.byt = AST_UPLUS
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1021
		{
			yyVAL.byt = AST_UMINUS
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1025
		{
			yyVAL.byt = AST_TILDA
		}
	case 192:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1031
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1036
		{
			yyVAL.valExpr = nil
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1040
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1046
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1050
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1056
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1061
		{
			yyVAL.valExpr = nil
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1065
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1071
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1075
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1081
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1085
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1089
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1093
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1098
		{
			yyVAL.valExprs = nil
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1102
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1107
		{
			yyVAL.boolExpr = nil
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1111
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 210:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1116
		{
			yyVAL.orderBy = nil
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1120
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1126
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1130
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1136
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1141
		{
			yyVAL.str = AST_ASC
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1145
		{
			yyVAL.str = AST_ASC
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1149
		{
			yyVAL.str = AST_DESC
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1154
		{
			yyVAL.limit = nil
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1158
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 220:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1162
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 221:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1166
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1171
		{
			yyVAL.str = ""
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1175
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 224:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1179
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
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1192
		{
			yyVAL.columns = nil
		}
	case 226:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1196
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1202
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1206
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 229:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1211
		{
			yyVAL.updateExprs = nil
		}
	case 230:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1215
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1221
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1225
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1231
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1236
		{
			yyVAL.empty = struct{}{}
		}
	case 235:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1238
		{
			yyVAL.empty = struct{}{}
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1241
		{
			yyVAL.empty = struct{}{}
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1243
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1246
		{
			yyVAL.str = ""
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1248
		{
			yyVAL.str = AST_IGNORE
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1252
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1254
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1256
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1258
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1260
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1268
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1270
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1273
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1279
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 252:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1284
		{
			ForceEOF(yylex)
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1290
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1292
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1294
		{
			yyVAL.bytes = []byte("big5")
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1296
		{
			yyVAL.bytes = []byte("binary")
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1298
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1300
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1302
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1304
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1306
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1308
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1310
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1312
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1314
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1316
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1318
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1320
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1322
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1324
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1326
		{
			yyVAL.bytes = []byte("greek")
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1328
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1330
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1332
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1334
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1336
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1338
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 278:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1340
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 279:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1342
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 280:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1344
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1346
		{
			yyVAL.bytes = []byte("macce")
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1348
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1350
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1352
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 285:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1354
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 286:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1356
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 287:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1358
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 288:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1360
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 289:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1362
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1364
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1366
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1368
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 293:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1372
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 294:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1374
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 295:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1376
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1378
		{
			yyVAL.bytes = []byte("serializable")
		}
	case 297:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1381
		{
			yyVAL.bytes = nil
		}
	case 298:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1383
		{
			yyVAL.bytes = []byte("session")
		}
	case 299:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1385
		{
			yyVAL.bytes = []byte("global")
		}
	case 300:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1388
		{
			yyVAL.expr = nil
		}
	case 301:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1390
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 302:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1394
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
