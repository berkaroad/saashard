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
const ID = 57374
const STRING = 57375
const NUMBER = 57376
const VALUE_ARG = 57377
const COMMENT = 57378
const UNION = 57379
const MINUS = 57380
const EXCEPT = 57381
const INTERSECT = 57382
const FULL = 57383
const JOIN = 57384
const STRAIGHT_JOIN = 57385
const LEFT = 57386
const RIGHT = 57387
const INNER = 57388
const OUTER = 57389
const CROSS = 57390
const NATURAL = 57391
const USE = 57392
const FORCE = 57393
const ON = 57394
const OR = 57395
const AND = 57396
const NOT = 57397
const BETWEEN = 57398
const CASE = 57399
const WHEN = 57400
const THEN = 57401
const ELSE = 57402
const LE = 57403
const GE = 57404
const NE = 57405
const NULL_SAFE_EQUAL = 57406
const IS = 57407
const LIKE = 57408
const IN = 57409
const UNARY = 57410
const END = 57411
const BEGIN = 57412
const START = 57413
const TRANSACTION = 57414
const COMMIT = 57415
const ROLLBACK = 57416
const ISOLATION = 57417
const LEVEL = 57418
const READ = 57419
const COMMITTED = 57420
const UNCOMMITTED = 57421
const REPEATABLE = 57422
const NAMES = 57423
const CHARSET = 57424
const CHARACTER = 57425
const ARMSCII8 = 57426
const ASCII = 57427
const BIG5 = 57428
const BINARY = 57429
const CP1250 = 57430
const CP1251 = 57431
const CP1256 = 57432
const CP1257 = 57433
const CP850 = 57434
const CP852 = 57435
const CP866 = 57436
const CP932 = 57437
const DEC8 = 57438
const EUCJPMS = 57439
const EUCKR = 57440
const GB2312 = 57441
const GBK = 57442
const GEOSTD8 = 57443
const GREEK = 57444
const HEBREW = 57445
const HP8 = 57446
const KEYBCS2 = 57447
const KOI8R = 57448
const KOI8U = 57449
const LATIN1 = 57450
const LATIN2 = 57451
const LATIN5 = 57452
const LATIN7 = 57453
const MACCE = 57454
const MACROMAN = 57455
const SJIS = 57456
const SWE7 = 57457
const TIS620 = 57458
const UCS2 = 57459
const EJIS = 57460
const UTF16 = 57461
const UTF16LE = 57462
const UTF32 = 57463
const UTF8 = 57464
const UTF8MB4 = 57465
const SESSION = 57466
const GLOBAL = 57467
const VARIABLES = 57468
const STATUS = 57469
const DATABASES = 57470
const STORAGE = 57471
const ENGINES = 57472
const TABLES = 57473
const COLUMNS = 57474
const PROCEDURE = 57475
const FUNCTION = 57476
const INDEXES = 57477
const KEYS = 57478
const REPLACE = 57479
const ADMIN = 57480
const HELP = 57481
const OFFSET = 57482
const COLLATE = 57483
const CREATE = 57484
const ALTER = 57485
const DROP = 57486
const RENAME = 57487
const TABLE = 57488
const INDEX = 57489
const VIEW = 57490
const TO = 57491
const IGNORE = 57492
const IF = 57493
const UNIQUE = 57494
const USING = 57495

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
	"NAMES",
	"CHARSET",
	"CHARACTER",
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
	"STORAGE",
	"ENGINES",
	"TABLES",
	"COLUMNS",
	"PROCEDURE",
	"FUNCTION",
	"INDEXES",
	"KEYS",
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

const yyNprod = 281
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 738

var yyAct = [...]int{

	123, 120, 258, 142, 489, 457, 260, 226, 194, 360,
	452, 355, 131, 220, 204, 261, 3, 109, 106, 316,
	273, 284, 352, 144, 126, 130, 88, 110, 136, 81,
	235, 234, 498, 121, 60, 498, 113, 127, 128, 129,
	19, 118, 134, 63, 53, 498, 114, 52, 228, 53,
	83, 228, 66, 85, 228, 371, 130, 89, 197, 136,
	347, 314, 117, 471, 137, 470, 469, 147, 127, 128,
	129, 115, 264, 134, 82, 146, 145, 438, 440, 84,
	132, 133, 111, 193, 37, 38, 39, 40, 72, 126,
	130, 201, 54, 136, 196, 137, 55, 56, 57, 279,
	219, 113, 127, 128, 129, 65, 118, 134, 146, 224,
	209, 132, 133, 231, 150, 211, 212, 141, 215, 126,
	130, 262, 222, 136, 47, 263, 49, 117, 214, 137,
	50, 147, 127, 128, 129, 202, 118, 134, 461, 267,
	270, 146, 145, 345, 272, 132, 133, 111, 386, 387,
	388, 389, 390, 419, 391, 392, 420, 117, 500, 137,
	276, 499, 207, 208, 257, 259, 364, 459, 460, 147,
	135, 497, 58, 350, 448, 132, 133, 408, 449, 397,
	406, 130, 278, 271, 136, 244, 346, 313, 233, 439,
	108, 503, 147, 127, 128, 129, 290, 264, 134, 277,
	281, 135, 245, 246, 247, 248, 249, 244, 298, 356,
	294, 288, 61, 206, 348, 291, 192, 356, 297, 411,
	137, 295, 300, 301, 143, 234, 74, 64, 466, 306,
	307, 115, 322, 323, 325, 135, 132, 133, 468, 328,
	320, 453, 333, 334, 367, 337, 338, 339, 340, 341,
	342, 343, 344, 299, 302, 312, 321, 305, 324, 335,
	308, 309, 310, 200, 280, 135, 467, 349, 115, 115,
	398, 329, 442, 235, 234, 205, 146, 145, 332, 325,
	365, 326, 327, 351, 353, 368, 75, 77, 76, 94,
	357, 331, 330, 359, 369, 436, 336, 146, 145, 363,
	373, 243, 242, 245, 246, 247, 248, 249, 244, 372,
	275, 222, 235, 234, 87, 146, 382, 374, 451, 435,
	377, 362, 453, 378, 379, 380, 135, 432, 430, 78,
	79, 320, 433, 431, 401, 402, 383, 434, 396, 275,
	347, 287, 289, 286, 51, 475, 274, 375, 376, 381,
	405, 247, 248, 249, 244, 115, 485, 358, 319, 146,
	145, 274, 414, 318, 400, 107, 410, 484, 416, 415,
	407, 413, 90, 222, 37, 38, 39, 40, 296, 384,
	19, 483, 399, 227, 477, 478, 107, 264, 71, 229,
	423, 424, 19, 268, 275, 126, 130, 320, 320, 136,
	266, 444, 445, 412, 428, 429, 447, 147, 127, 128,
	129, 366, 118, 134, 450, 425, 265, 228, 138, 319,
	458, 97, 107, 455, 318, 454, 148, 61, 147, 443,
	462, 395, 232, 117, 441, 137, 243, 242, 245, 246,
	247, 248, 249, 244, 394, 61, 422, 421, 65, 472,
	486, 132, 133, 293, 473, 243, 242, 245, 246, 247,
	248, 249, 244, 386, 387, 388, 389, 390, 349, 391,
	392, 292, 481, 479, 93, 92, 225, 198, 458, 195,
	191, 487, 86, 203, 490, 490, 490, 488, 495, 491,
	492, 149, 140, 474, 146, 145, 19, 501, 139, 504,
	91, 404, 496, 282, 505, 199, 506, 69, 229, 67,
	465, 480, 446, 482, 417, 221, 78, 79, 361, 464,
	95, 105, 104, 96, 98, 99, 100, 102, 103, 243,
	242, 245, 246, 247, 248, 249, 244, 427, 274, 101,
	304, 135, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 163, 164, 165, 166, 167, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 187, 188,
	189, 190, 19, 20, 21, 22, 242, 245, 246, 247,
	248, 249, 244, 303, 218, 217, 216, 213, 210, 73,
	502, 493, 19, 238, 240, 42, 23, 418, 34, 250,
	251, 252, 253, 254, 255, 256, 241, 239, 237, 243,
	242, 245, 246, 247, 248, 249, 244, 403, 283, 48,
	33, 370, 285, 80, 41, 223, 494, 476, 456, 463,
	426, 409, 269, 354, 243, 242, 245, 246, 247, 248,
	249, 244, 125, 122, 124, 43, 44, 45, 46, 311,
	119, 236, 28, 29, 116, 30, 31, 59, 437, 62,
	317, 385, 315, 112, 393, 230, 68, 36, 70, 18,
	17, 16, 15, 14, 13, 12, 11, 10, 9, 8,
	7, 6, 5, 4, 2, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 32,
	35, 0, 0, 0, 24, 25, 27, 26,
}
var yyPact = [...]int{

	577, -1000, -1000, 335, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -37, -116, -69, -65, -1000, 85,
	-1000, -1000, -1000, 395, -1000, 73, 597, 492, -1000, -1000,
	-1000, 489, -1000, -121, 416, 590, 190, -137, -88, 395,
	-1000, -82, 395, -1000, 450, -140, 395, -140, -1000, 475,
	-1000, -1000, 377, 385, -1000, 108, -1000, -1000, 69, -1000,
	382, 473, 463, 416, 137, 393, 462, 443, -1000, -1000,
	448, 158, 395, -1000, 447, -1000, -106, 445, 485, 208,
	395, 416, 454, 203, 21, 203, 589, -31, 588, -14,
	-24, 587, 586, 585, -1000, -45, 491, 396, 444, 374,
	-1000, -1000, 413, 106, 217, 545, -1000, 99, 375, -1000,
	-1000, -1000, 160, 379, 363, -1000, 356, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 160, -1000, 416,
	396, 528, 296, 70, -1000, 135, -1000, 100, -57, 443,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 483, -147, -1000, 183, -1000, 439, -1000, -1000,
	421, -1000, 349, 203, -1000, 160, 99, 203, 203, -1000,
	416, 584, 531, 416, 203, 203, 416, 416, 416, -1000,
	-1000, 350, 335, 18, -1000, -1000, -1000, 326, 69, 160,
	-1000, -1000, 395, 180, 99, 99, 160, 350, 219, 160,
	160, 238, 160, 160, 160, 160, 160, 160, 160, 160,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 545, -26,
	17, 45, 545, -1000, 35, 4, 69, -1000, 597, 148,
	381, 328, 351, 505, 99, 396, 75, 160, 395, 378,
	-1000, -1000, -1000, 189, 395, -1000, -109, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 491, 396, -1000, 381, 217,
	-1000, -1000, 203, 416, 416, 203, -1000, -1000, 203, 203,
	203, 306, -1000, -1000, 396, 336, 418, 412, 387, 97,
	-1000, -1000, 227, -1000, -1000, -1000, -1000, 168, 381, -1000,
	350, 160, 160, 381, 570, -1000, 480, 126, 511, -1000,
	273, 273, 104, 104, 104, -1000, -1000, 160, -1000, 381,
	-1000, 11, 69, 8, 156, -1000, 99, 491, 396, 505,
	493, 500, 217, -1000, 61, 381, -1000, 415, -1000, -1000,
	414, -1000, -1000, 296, -1000, 203, 203, -1000, -1000, -1000,
	-1000, 350, -1000, 526, 326, 326, -1000, -1000, 283, 282,
	292, 274, 250, 24, -1000, 402, 103, 397, 160, 160,
	-1000, 381, 455, 160, -1000, 381, -1000, 5, -1000, 94,
	-1000, 160, 256, 186, 267, 493, -1000, 160, -1000, 74,
	46, -1000, -1000, -1000, -1000, -1000, 507, 496, 418, 173,
	-1000, 221, -1000, 193, -1000, -1000, -1000, -1000, -96, -97,
	-99, -1000, -1000, -1000, 381, 381, 160, 381, -1000, -1000,
	381, 160, -1000, 467, -1000, -1000, 302, -1000, 362, -1000,
	-1000, -1000, -1000, 505, 99, 160, 99, -1000, -1000, 344,
	330, 319, 381, 381, 423, 160, -1000, -1000, -1000, 493,
	217, 297, 217, 395, 395, 395, 594, -1000, 472, 2,
	-1000, -8, -11, 396, -1000, 593, 118, -1000, 395, -1000,
	-1000, 296, -1000, 395, -1000, 395, -1000,
}
var yyPgo = [...]int{

	0, 695, 694, 15, 693, 692, 691, 690, 689, 688,
	687, 686, 685, 684, 683, 682, 681, 680, 679, 634,
	678, 677, 676, 344, 17, 27, 675, 674, 673, 672,
	19, 671, 670, 43, 668, 4, 20, 46, 664, 661,
	13, 660, 2, 33, 6, 659, 654, 12, 653, 1,
	652, 643, 11, 642, 641, 640, 639, 9, 638, 5,
	637, 7, 636, 18, 635, 10, 3, 23, 14, 314,
	633, 632, 631, 629, 628, 0, 8, 114, 607, 226,
	605,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 4, 4, 15, 15, 5, 6, 7, 7, 7,
	7, 7, 7, 12, 12, 13, 14, 16, 8, 8,
	8, 9, 9, 9, 10, 11, 11, 11, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 18, 18, 80, 19, 20, 20,
	21, 21, 21, 21, 21, 22, 22, 24, 24, 25,
	25, 25, 28, 28, 26, 26, 26, 29, 29, 30,
	30, 30, 30, 27, 27, 27, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 32, 32, 32, 33, 33,
	34, 34, 34, 34, 35, 35, 36, 36, 37, 37,
	37, 37, 37, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 39, 39, 39, 39, 39, 39, 39,
	40, 40, 45, 45, 43, 43, 47, 44, 44, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 46, 46, 48, 48,
	48, 50, 53, 53, 51, 51, 52, 54, 54, 49,
	49, 41, 41, 41, 41, 55, 55, 56, 56, 57,
	57, 58, 58, 59, 60, 60, 60, 61, 61, 61,
	61, 62, 62, 62, 63, 63, 64, 64, 65, 65,
	66, 66, 67, 69, 69, 70, 70, 23, 23, 71,
	71, 71, 71, 71, 72, 72, 73, 73, 74, 74,
	75, 76, 77, 77, 77, 77, 77, 77, 77, 77,
	77, 77, 77, 77, 77, 77, 77, 77, 77, 77,
	77, 77, 77, 77, 77, 77, 77, 77, 77, 77,
	77, 77, 77, 77, 77, 77, 77, 77, 77, 77,
	77, 77, 78, 78, 78, 79, 79, 79, 68, 68,
	68,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 5, 12,
	3, 8, 8, 6, 6, 8, 7, 4, 4, 6,
	5, 4, 7, 1, 2, 1, 1, 2, 5, 8,
	4, 6, 7, 4, 5, 4, 5, 5, 5, 4,
	5, 5, 4, 6, 7, 6, 7, 5, 5, 6,
	6, 6, 3, 4, 4, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 1, 3, 3, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 3, 4, 5, 4, 1, 1, 1, 1, 1,
	1, 5, 0, 1, 1, 2, 4, 0, 2, 1,
	3, 1, 1, 1, 1, 0, 3, 0, 2, 0,
	3, 1, 3, 2, 0, 1, 1, 0, 2, 4,
	4, 0, 2, 4, 0, 3, 1, 3, 0, 5,
	1, 3, 3, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 0, 1, 1, 0, 2,
	2,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, -14, -15, -16, -17, -18, 5,
	6, 7, 8, 29, 157, 158, 160, 159, 85, 86,
	88, 89, 152, 53, 31, 153, -21, 39, 40, 41,
	42, -19, -80, -19, -19, -19, -19, 161, -73, 163,
	167, -23, 163, 165, 161, 161, 162, 163, 87, -19,
	-75, 32, -19, -33, 154, 32, -3, 17, -22, 18,
	-20, -23, -33, 9, -79, 96, 98, 97, 139, 140,
	-70, 166, 162, -75, 161, -75, 32, -69, 166, -75,
	-69, 25, 98, 97, -79, 143, 146, 44, 147, 148,
	149, 162, 150, 151, 145, 144, -63, 37, 82, -24,
	-25, 78, -28, 32, -37, -42, -38, 58, 37, -41,
	-49, -43, -48, -75, -46, -50, 20, 33, 34, 35,
	21, -47, 76, 77, 38, 166, 24, 60, 36, 25,
	29, -33, -66, 87, -67, -49, -75, 32, 33, 29,
	-77, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 32, 58, -75, -76, 32, -76, 164, 32, 20,
	55, -75, -33, 29, -68, 72, 10, 141, 142, -68,
	9, 146, 147, 9, 142, 142, 9, 9, 9, 145,
	-40, 24, -3, -64, -49, 32, -61, 9, 43, 15,
	-26, -75, 19, 82, 57, 56, -39, 73, 58, 72,
	59, 71, 75, 74, 81, 76, 77, 78, 79, 80,
	64, 65, 66, 67, 68, 69, 70, -37, -42, -37,
	-44, -3, -42, -42, 37, 37, 37, -47, 37, -53,
	-42, -33, -66, -36, 10, 43, 90, 64, 82, 156,
	-77, -76, 20, -74, 168, -71, 160, 158, 28, 159,
	13, 32, 32, 32, -76, -63, 29, -68, -42, -37,
	-68, -68, -33, 9, 9, -33, -68, -68, -33, -33,
	-33, -45, -43, 169, 43, -29, -30, -32, 37, 32,
	-47, -25, -42, -75, 78, -75, -37, -37, -42, -43,
	73, 72, 59, -42, -42, 21, 58, -42, -42, -42,
	-42, -42, -42, -42, -42, 169, 169, 43, 169, -42,
	169, -24, 18, -24, -51, -52, 61, -63, 29, -36,
	-57, 13, -37, -67, 91, -42, 33, 55, -75, -76,
	-72, 164, -40, -66, -68, -33, -33, -68, -68, -68,
	-68, 43, -49, -36, 43, -31, 45, 46, 47, 48,
	49, 51, 52, -27, 32, 19, -30, 82, 43, 155,
	-43, -42, -42, 57, 21, -42, 169, -24, 169, -54,
	-52, 63, -37, -40, -66, -57, -61, 14, -78, 92,
	95, 32, 32, -68, -68, -43, -55, 11, -30, -30,
	45, 50, 45, 50, 45, 45, 45, -34, 53, 165,
	54, 32, 169, 32, -42, -42, 57, -42, 169, 84,
	-42, 62, -65, 55, -65, -61, -58, -59, -42, 93,
	94, 92, -76, -56, 12, 14, 55, 45, 45, 162,
	162, 162, -42, -42, 26, 43, -60, 22, 23, -57,
	-37, -44, -37, 37, 37, 37, 27, -59, -61, -35,
	-75, -35, -35, 7, -62, 16, 30, 169, 43, 169,
	169, -66, 7, 73, -75, -75, -75,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 66,
	66, 66, 66, 66, 226, 217, 0, 0, 33, 0,
	35, 36, 66, 0, 66, 0, 0, 70, 72, 73,
	74, 75, 68, 217, 0, 0, 275, 215, 0, 0,
	227, 0, 0, 218, 0, 213, 0, 213, 34, 0,
	37, 230, 275, 204, 65, 108, 20, 71, 0, 76,
	67, 0, 0, 0, 0, 0, 0, 0, 276, 277,
	0, 0, 0, 231, 0, 231, 0, 0, 0, 0,
	0, 0, 0, 278, 0, 278, 0, 0, 0, 0,
	0, 0, 0, 0, 62, 0, 0, 0, 0, 197,
	77, 79, 84, 230, 82, 83, 118, 0, 0, 149,
	150, 151, 0, 179, 0, 165, 0, 181, 182, 183,
	184, 145, 168, 169, 170, 166, 167, 172, 69, 0,
	0, 116, 27, 0, 210, 0, 179, 230, 28, 0,
	31, 232, 233, 234, 235, 236, 237, 238, 239, 240,
	241, 242, 243, 244, 245, 246, 247, 248, 249, 250,
	251, 252, 253, 254, 255, 256, 257, 258, 259, 260,
	261, 262, 263, 264, 265, 266, 267, 268, 269, 270,
	271, 231, 0, 228, 40, 0, 43, 0, 45, 214,
	0, 231, 204, 278, 49, 0, 0, 278, 278, 52,
	0, 0, 0, 0, 278, 278, 0, 0, 0, 63,
	64, 0, 141, 0, 206, 109, 18, 0, 0, 0,
	80, 85, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	133, 134, 135, 136, 137, 138, 139, 121, 0, 0,
	0, 0, 147, 160, 0, 0, 0, 132, 0, 0,
	173, 204, 116, 189, 0, 0, 0, 0, 0, 0,
	30, 38, 216, 0, 0, 231, 224, 219, 220, 221,
	222, 223, 44, 46, 47, 0, 0, 48, 279, 280,
	50, 51, 278, 0, 0, 278, 57, 58, 278, 278,
	278, 140, 142, 205, 0, 116, 87, 93, 0, 105,
	107, 78, 198, 86, 81, 180, 119, 120, 123, 124,
	0, 0, 0, 126, 0, 130, 0, 152, 153, 154,
	155, 156, 157, 158, 159, 122, 144, 0, 146, 147,
	161, 0, 0, 0, 177, 174, 0, 0, 0, 189,
	197, 0, 117, 211, 0, 212, 29, 0, 229, 41,
	0, 225, 23, 24, 53, 278, 278, 55, 59, 60,
	61, 0, 207, 185, 0, 0, 96, 97, 0, 0,
	0, 0, 0, 110, 94, 0, 0, 0, 0, 0,
	125, 127, 0, 0, 131, 148, 162, 0, 164, 0,
	175, 0, 0, 208, 208, 197, 26, 0, 32, 0,
	0, 231, 42, 54, 56, 143, 187, 0, 88, 91,
	98, 0, 100, 0, 102, 103, 104, 89, 0, 0,
	0, 95, 90, 106, 199, 200, 0, 128, 163, 171,
	178, 0, 21, 0, 22, 25, 190, 191, 194, 272,
	273, 274, 39, 189, 0, 0, 0, 99, 101, 0,
	0, 0, 129, 176, 0, 0, 193, 195, 196, 197,
	188, 186, 92, 0, 0, 0, 0, 192, 201, 0,
	114, 0, 0, 0, 19, 0, 0, 111, 0, 112,
	113, 209, 202, 0, 115, 0, 203,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 80, 75, 3,
	37, 169, 78, 76, 43, 77, 82, 79, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	65, 64, 66, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 81, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 74, 3, 38,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 39, 40, 41, 42, 44,
	45, 46, 47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 67,
	68, 69, 70, 71, 72, 73, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165, 166,
	167, 168,
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
		//line ./yacc.y:185
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:191
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:212
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, Limit: yyDollar[5].limit}
		}
	case 19:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line ./yacc.y:216
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:220
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:227
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:231
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:243
		{
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:247
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:260
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:266
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:272
		{
			yyVAL.statement = &Set{
				SetType:        SetTypeNameValue,
				Comments:       Comments(yyDollar[2].bytes2),
				Scope:          string(yyDollar[3].bytes),
				NameValueExprs: yyDollar[4].updateExprs,
			}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:281
		{
			yyVAL.statement = &Set{
				SetType:  SetTypeNames,
				Comments: Comments(yyDollar[2].bytes2),
				SpaceSplitExprs: SpaceSplitExprs{
					&SpaceSplitExpr{
						Name: "names", Expr: yyDollar[4].bytes,
					},
				},
			}
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:293
		{
			yyVAL.statement = &Set{
				SetType:  SetTypeNames,
				Comments: Comments(yyDollar[2].bytes2),
				SpaceSplitExprs: SpaceSplitExprs{
					&SpaceSplitExpr{
						Name: "names", Expr: yyDollar[4].bytes,
					},
					&SpaceSplitExpr{
						Name: "collate", Expr: yyDollar[6].bytes,
					},
				},
			}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:308
		{
			yyVAL.statement = &Set{
				SetType:  SetTypeCharset,
				Comments: Comments(yyDollar[2].bytes2),
				SpaceSplitExprs: SpaceSplitExprs{
					&SpaceSplitExpr{
						Name: "character set", Expr: yyDollar[5].bytes,
					},
				},
			}
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:320
		{
			yyVAL.statement = &Set{
				SetType:  SetTypeCharset,
				Comments: Comments(yyDollar[2].bytes2),
				SpaceSplitExprs: SpaceSplitExprs{
					&SpaceSplitExpr{
						Name: "charset", Expr: yyDollar[4].bytes,
					},
				},
			}
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:332
		{
			yyVAL.statement = &Set{
				SetType:  SetTypeTransactionIsolationLevel,
				Comments: Comments(yyDollar[2].bytes2),
				Scope:    string(yyDollar[3].bytes),
				SpaceSplitExprs: SpaceSplitExprs{
					&SpaceSplitExpr{
						Name: "transaction isolation level", Expr: yyDollar[7].bytes,
					},
				},
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:347
		{
			yyVAL.statement = &Begin{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:351
		{
			yyVAL.statement = &Begin{}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:358
		{
			yyVAL.statement = &Commit{}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:364
		{
			yyVAL.statement = &Rollback{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:370
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:376
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[4].bytes}
		}
	case 39:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./yacc.y:380
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:385
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:391
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 42:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:395
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Ignore: yyDollar[2].str, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:400
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:406
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:412
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:416
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:421
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:427
		{
			yyVAL.statement = &Show{Type: ShowTypeCharset, Comments: Comments(yyDollar[2].bytes2), Section: "character set", LikeOrWhere: yyDollar[5].expr}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:431
		{
			yyVAL.statement = &Show{Type: ShowTypeCharset, Comments: Comments(yyDollar[2].bytes2), Section: "charset", LikeOrWhere: yyDollar[4].expr}
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:435
		{
			yyVAL.statement = &Show{Type: ShowTypeVariables, Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), Section: "VARIABLES", LikeOrWhere: yyDollar[5].expr}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:439
		{
			yyVAL.statement = &Show{Type: ShowTypeStatus, Comments: Comments(yyDollar[2].bytes2), Scope: string(yyDollar[3].bytes), Section: "status", LikeOrWhere: yyDollar[5].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:443
		{
			yyVAL.statement = &Show{Type: ShowTypeDatabases, Comments: Comments(yyDollar[2].bytes2), Section: "databases", LikeOrWhere: yyDollar[4].expr}
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:447
		{
			yyVAL.statement = &Show{Type: ShowTypeTables, Comments: Comments(yyDollar[2].bytes2), Section: "tables", From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 54:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:451
		{
			yyVAL.statement = &Show{Type: ShowTypeTables, Comments: Comments(yyDollar[2].bytes2), Section: "full tables", From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:455
		{
			yyVAL.statement = &Show{Type: ShowTypeColumns, Comments: Comments(yyDollar[2].bytes2), Section: "columns", From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 56:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./yacc.y:459
		{
			yyVAL.statement = &Show{Type: ShowTypeColumns, Comments: Comments(yyDollar[2].bytes2), Section: "full columns", From: yyDollar[6].tableName, LikeOrWhere: yyDollar[7].expr}
		}
	case 57:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:463
		{
			yyVAL.statement = &Show{Type: ShowTypeProcedureStatus, Comments: Comments(yyDollar[2].bytes2), Section: "full columns", LikeOrWhere: yyDollar[5].expr}
		}
	case 58:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:467
		{
			yyVAL.statement = &Show{Type: ShowTypeFunctionStatus, Comments: Comments(yyDollar[2].bytes2), Section: "full columns", LikeOrWhere: yyDollar[5].expr}
		}
	case 59:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:471
		{
			yyVAL.statement = &Show{Type: ShowTypeIndexes, Comments: Comments(yyDollar[2].bytes2), Section: "index", From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 60:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:475
		{
			yyVAL.statement = &Show{Type: ShowTypeIndexes, Comments: Comments(yyDollar[2].bytes2), Section: "indexes", From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:479
		{
			yyVAL.statement = &Show{Type: ShowTypeIndexes, Comments: Comments(yyDollar[2].bytes2), Section: "keys", From: yyDollar[5].tableName, LikeOrWhere: yyDollar[6].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:483
		{
			yyVAL.statement = &Show{Type: ShowTypeEngines, Comments: Comments(yyDollar[2].bytes2), Section: "engines"}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:487
		{
			yyVAL.statement = &Show{Type: ShowTypeEngines, Comments: Comments(yyDollar[2].bytes2), Section: "storage engines"}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:493
		{
			yyVAL.statement = &Admin{Region: yyDollar[2].tableName, Columns: yyDollar[3].columns, Rows: yyDollar[4].insRows}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:497
		{
			yyVAL.statement = &AdminHelp{}
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:502
		{
			SetAllowComments(yylex, true)
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:506
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:512
		{
			yyVAL.bytes2 = nil
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:516
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:522
		{
			yyVAL.str = AST_UNION
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:526
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:530
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:534
		{
			yyVAL.str = AST_EXCEPT
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:538
		{
			yyVAL.str = AST_INTERSECT
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:543
		{
			yyVAL.str = ""
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:547
		{
			yyVAL.str = AST_DISTINCT
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:553
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:557
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:563
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:567
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:571
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:577
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:581
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 84:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:586
		{
			yyVAL.bytes = nil
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:590
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:594
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:600
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:604
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:610
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:614
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:618
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:622
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:627
		{
			yyVAL.bytes = nil
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:631
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:635
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:641
		{
			yyVAL.str = AST_JOIN
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:645
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:649
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:653
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:657
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:661
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:665
		{
			yyVAL.str = AST_JOIN
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:669
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:673
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:679
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:683
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:687
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:693
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:697
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 110:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:702
		{
			yyVAL.indexHints = nil
		}
	case 111:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:706
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:710
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:714
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:720
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:724
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 116:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:729
		{
			yyVAL.boolExpr = nil
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:733
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:740
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 120:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:744
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:748
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:752
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:758
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:762
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].tuple}
		}
	case 125:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:766
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].tuple}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:770
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 127:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:774
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 128:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:778
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 129:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./yacc.y:782
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:786
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:790
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:794
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:800
		{
			yyVAL.str = AST_EQ
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:804
		{
			yyVAL.str = AST_LT
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:808
		{
			yyVAL.str = AST_GT
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:812
		{
			yyVAL.str = AST_LE
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:816
		{
			yyVAL.str = AST_GE
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:820
		{
			yyVAL.str = AST_NE
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:824
		{
			yyVAL.str = AST_NSE
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:830
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:834
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:840
		{
			yyVAL.values = Values{yyDollar[1].tuple}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:844
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].tuple)
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:850
		{
			yyVAL.tuple = ValTuple(yyDollar[2].valExprs)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:854
		{
			yyVAL.tuple = yyDollar[1].subquery
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:860
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:866
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:870
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:876
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:880
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:884
		{
			yyVAL.valExpr = yyDollar[1].tuple
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:888
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:892
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:896
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:900
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:904
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:908
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:912
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:916
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:920
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
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:935
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 162:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:939
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:943
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:947
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:951
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:957
		{
			yyVAL.bytes = IF_BYTES
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:961
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:967
		{
			yyVAL.byt = AST_UPLUS
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:971
		{
			yyVAL.byt = AST_UMINUS
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:975
		{
			yyVAL.byt = AST_TILDA
		}
	case 171:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:981
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:986
		{
			yyVAL.valExpr = nil
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:990
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:996
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1000
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 176:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1006
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1011
		{
			yyVAL.valExpr = nil
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1015
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1021
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 180:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1025
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1031
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1035
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1039
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1043
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1048
		{
			yyVAL.valExprs = nil
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1052
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1057
		{
			yyVAL.boolExpr = nil
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1061
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1066
		{
			yyVAL.orderBy = nil
		}
	case 190:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1070
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1076
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1080
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1086
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1091
		{
			yyVAL.str = AST_ASC
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1095
		{
			yyVAL.str = AST_ASC
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1099
		{
			yyVAL.str = AST_DESC
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1104
		{
			yyVAL.limit = nil
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1108
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1112
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 200:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1116
		{
			yyVAL.limit = &Limit{Offset: yyDollar[4].valExpr, Rowcount: yyDollar[2].valExpr}
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1121
		{
			yyVAL.str = ""
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1125
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 203:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./yacc.y:1129
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
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1142
		{
			yyVAL.columns = nil
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1146
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1152
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1156
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1161
		{
			yyVAL.updateExprs = nil
		}
	case 209:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./yacc.y:1165
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1171
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1175
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1181
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1186
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1188
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1191
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./yacc.y:1193
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1196
		{
			yyVAL.str = ""
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1198
		{
			yyVAL.str = AST_IGNORE
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1202
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1204
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1206
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1208
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1210
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1213
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1215
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1218
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1220
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1223
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1225
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1229
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 231:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1234
		{
			ForceEOF(yylex)
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1240
		{
			yyVAL.bytes = []byte("armscii8")
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1242
		{
			yyVAL.bytes = []byte("ascii")
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1244
		{
			yyVAL.bytes = []byte("big5")
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1246
		{
			yyVAL.bytes = []byte("binary")
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1248
		{
			yyVAL.bytes = []byte("cp1250")
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1250
		{
			yyVAL.bytes = []byte("cp1251")
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1252
		{
			yyVAL.bytes = []byte("cp1256")
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1254
		{
			yyVAL.bytes = []byte("cp1257")
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1256
		{
			yyVAL.bytes = []byte("cp850")
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1258
		{
			yyVAL.bytes = []byte("cp852")
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1260
		{
			yyVAL.bytes = []byte("cp866")
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1262
		{
			yyVAL.bytes = []byte("cp932")
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1264
		{
			yyVAL.bytes = []byte("dec8")
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1266
		{
			yyVAL.bytes = []byte("eucjpms")
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1268
		{
			yyVAL.bytes = []byte("euckr")
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1270
		{
			yyVAL.bytes = []byte("gb2312")
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1272
		{
			yyVAL.bytes = []byte("gbk")
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1274
		{
			yyVAL.bytes = []byte("geostd8")
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1276
		{
			yyVAL.bytes = []byte("greek")
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1278
		{
			yyVAL.bytes = []byte("hebrew")
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1280
		{
			yyVAL.bytes = []byte("hp8")
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1282
		{
			yyVAL.bytes = []byte("keybcs2")
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1284
		{
			yyVAL.bytes = []byte("koi8r")
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1286
		{
			yyVAL.bytes = []byte("koi8u")
		}
	case 256:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1288
		{
			yyVAL.bytes = []byte("latin1")
		}
	case 257:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1290
		{
			yyVAL.bytes = []byte("latin2")
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1292
		{
			yyVAL.bytes = []byte("latin5")
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1294
		{
			yyVAL.bytes = []byte("latin7")
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1296
		{
			yyVAL.bytes = []byte("macce")
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1298
		{
			yyVAL.bytes = []byte("macroman")
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1300
		{
			yyVAL.bytes = []byte("sjis")
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1302
		{
			yyVAL.bytes = []byte("swe7")
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1304
		{
			yyVAL.bytes = []byte("tis620")
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1306
		{
			yyVAL.bytes = []byte("ucs2")
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1308
		{
			yyVAL.bytes = []byte("ujis")
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1310
		{
			yyVAL.bytes = []byte("utf16")
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1312
		{
			yyVAL.bytes = []byte("utf16le")
		}
	case 269:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1314
		{
			yyVAL.bytes = []byte("utf32")
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1316
		{
			yyVAL.bytes = []byte("utf8")
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1318
		{
			yyVAL.bytes = []byte("utf8mb4")
		}
	case 272:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1323
		{
			yyVAL.bytes = []byte("read committed")
		}
	case 273:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1325
		{
			yyVAL.bytes = []byte("read uncommitted")
		}
	case 274:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1327
		{
			yyVAL.bytes = []byte("repeatable read")
		}
	case 275:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1330
		{
			yyVAL.bytes = nil
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1332
		{
			yyVAL.bytes = []byte("session")
		}
	case 277:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./yacc.y:1334
		{
			yyVAL.bytes = []byte("global")
		}
	case 278:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./yacc.y:1337
		{
			yyVAL.expr = nil
		}
	case 279:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1339
		{
			yyVAL.expr = &LikeExpr{Expr: yyDollar[2].valExpr}
		}
	case 280:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./yacc.y:1343
		{
			yyVAL.expr = &WhereExpr{Expr: yyDollar[2].boolExpr}
		}
	}
	goto yystack /* stack new state and value */
}
