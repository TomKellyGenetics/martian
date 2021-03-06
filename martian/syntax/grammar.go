// Code generated by goyacc -p mm -o grammar.go grammar.y. DO NOT EDIT.

//line grammar.y:2
//
// Copyright (c) 2014 10X Genomics, Inc. All rights reserved.
//
// MRO grammar.
//

package syntax

import __yyfmt__ "fmt"

//line grammar.y:8

import (
	"strings"
)

//line grammar.y:16
type mmSymType struct {
	yys       int
	global    *Ast
	srcfile   *SourceFile
	arr       int16
	loc       int
	val       []byte
	modifiers *Modifiers
	dec       Dec
	decs      []Dec
	inparam   *InParam
	outparam  *OutParam
	retains   []*RetainParam
	stretains *RetainParams
	i_params  *InParams
	o_params  *OutParams
	res       *Resources
	par_tuple paramsTuple
	src       *SrcParam
	exp       Exp
	exps      []Exp
	rexp      *RefExp
	vexp      *ValExp
	kvpairs   map[string]Exp
	call      *CallStm
	calls     []*CallStm
	binding   *BindStm
	bindings  *BindStms
	retstm    *ReturnStm
	plretains *PipelineRetains
	reflist   []*RefExp
	includes  []*Include
	intern    *stringIntern
}

const SKIP = 57346
const COMMENT = 57347
const INVALID = 57348
const SEMICOLON = 57349
const COLON = 57350
const COMMA = 57351
const EQUALS = 57352
const LBRACKET = 57353
const RBRACKET = 57354
const LPAREN = 57355
const RPAREN = 57356
const LBRACE = 57357
const RBRACE = 57358
const SWEEP = 57359
const RETURN = 57360
const SELF = 57361
const FILETYPE = 57362
const STAGE = 57363
const PIPELINE = 57364
const CALL = 57365
const SPLIT = 57366
const USING = 57367
const RETAIN = 57368
const LOCAL = 57369
const PREFLIGHT = 57370
const VOLATILE = 57371
const DISABLED = 57372
const STRICT = 57373
const IN = 57374
const OUT = 57375
const SRC = 57376
const AS = 57377
const THREADS = 57378
const MEM_GB = 57379
const VMEM_GB = 57380
const SPECIAL = 57381
const ID = 57382
const LITSTRING = 57383
const NUM_FLOAT = 57384
const NUM_INT = 57385
const DOT = 57386
const PY = 57387
const EXEC = 57388
const COMPILED = 57389
const MAP = 57390
const INT = 57391
const STRING = 57392
const FLOAT = 57393
const PATH = 57394
const BOOL = 57395
const TRUE = 57396
const FALSE = 57397
const NULL = 57398
const DEFAULT = 57399
const INCLUDE_DIRECTIVE = 57400

var mmToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SKIP",
	"COMMENT",
	"INVALID",
	"SEMICOLON",
	"COLON",
	"COMMA",
	"EQUALS",
	"LBRACKET",
	"RBRACKET",
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"SWEEP",
	"RETURN",
	"SELF",
	"FILETYPE",
	"STAGE",
	"PIPELINE",
	"CALL",
	"SPLIT",
	"USING",
	"RETAIN",
	"LOCAL",
	"PREFLIGHT",
	"VOLATILE",
	"DISABLED",
	"STRICT",
	"IN",
	"OUT",
	"SRC",
	"AS",
	"THREADS",
	"MEM_GB",
	"VMEM_GB",
	"SPECIAL",
	"ID",
	"LITSTRING",
	"NUM_FLOAT",
	"NUM_INT",
	"DOT",
	"PY",
	"EXEC",
	"COMPILED",
	"MAP",
	"INT",
	"STRING",
	"FLOAT",
	"PATH",
	"BOOL",
	"TRUE",
	"FALSE",
	"NULL",
	"DEFAULT",
	"INCLUDE_DIRECTIVE",
}
var mmStatenames = [...]string{}

const mmEofCode = 1
const mmErrCode = 2
const mmInitialStackSize = 16

//line grammar.y:742

//line yacctab:1
var mmExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 65,
	13, 113,
	35, 113,
	-2, 72,
	-1, 66,
	13, 116,
	35, 116,
	-2, 73,
	-1, 67,
	13, 123,
	35, 123,
	-2, 74,
}

const mmPrivate = 57344

const mmLast = 642

var mmAct = [...]int{

	36, 35, 101, 98, 33, 180, 90, 164, 4, 31,
	118, 24, 26, 14, 139, 19, 20, 58, 135, 136,
	137, 9, 59, 60, 200, 10, 75, 76, 71, 37,
	42, 64, 61, 70, 49, 52, 47, 43, 46, 53,
	40, 50, 16, 21, 22, 8, 51, 44, 45, 48,
	38, 13, 11, 12, 228, 227, 41, 39, 226, 87,
	103, 54, 229, 76, 19, 20, 15, 150, 203, 195,
	182, 84, 85, 82, 179, 169, 62, 89, 29, 88,
	25, 80, 99, 208, 86, 91, 56, 230, 117, 116,
	204, 205, 206, 207, 104, 222, 210, 186, 175, 111,
	161, 181, 166, 81, 58, 113, 181, 166, 126, 28,
	130, 159, 160, 9, 117, 117, 132, 10, 58, 127,
	128, 129, 16, 21, 22, 8, 8, 117, 138, 16,
	21, 22, 8, 111, 149, 7, 145, 172, 125, 27,
	153, 72, 8, 13, 11, 12, 197, 133, 73, 155,
	140, 198, 165, 151, 190, 57, 19, 20, 15, 27,
	6, 168, 92, 171, 188, 187, 177, 173, 178, 189,
	176, 157, 112, 183, 78, 94, 95, 96, 97, 77,
	68, 193, 63, 69, 192, 167, 221, 220, 219, 196,
	218, 199, 217, 82, 131, 209, 108, 107, 106, 105,
	111, 236, 9, 216, 235, 234, 10, 233, 146, 232,
	37, 42, 231, 225, 224, 49, 52, 47, 43, 46,
	53, 40, 50, 214, 211, 201, 194, 51, 44, 45,
	48, 38, 13, 11, 12, 9, 83, 41, 39, 10,
	184, 162, 156, 37, 42, 19, 20, 15, 49, 52,
	47, 43, 46, 53, 40, 50, 144, 143, 142, 141,
	51, 44, 45, 48, 38, 13, 11, 12, 9, 32,
	41, 39, 10, 100, 74, 1, 37, 42, 19, 20,
	15, 49, 52, 47, 43, 46, 53, 40, 50, 3,
	34, 5, 23, 51, 44, 45, 48, 38, 13, 11,
	12, 9, 191, 41, 39, 10, 158, 170, 79, 37,
	42, 19, 20, 15, 49, 52, 47, 43, 46, 53,
	40, 50, 93, 110, 154, 55, 51, 44, 45, 48,
	38, 13, 11, 12, 114, 148, 41, 39, 185, 212,
	174, 202, 115, 42, 19, 20, 15, 49, 52, 47,
	43, 46, 53, 40, 50, 102, 18, 17, 30, 51,
	44, 45, 48, 38, 134, 2, 163, 0, 152, 41,
	39, 124, 119, 120, 122, 121, 123, 42, 0, 0,
	0, 49, 52, 47, 43, 46, 53, 40, 50, 0,
	0, 0, 0, 51, 44, 45, 48, 38, 166, 0,
	223, 0, 0, 41, 39, 37, 42, 0, 0, 0,
	49, 52, 47, 43, 46, 53, 40, 50, 0, 0,
	0, 0, 51, 44, 45, 48, 38, 0, 215, 0,
	0, 0, 41, 39, 42, 0, 0, 0, 49, 52,
	47, 43, 46, 53, 40, 50, 0, 0, 0, 0,
	51, 44, 45, 48, 38, 0, 213, 0, 0, 0,
	41, 39, 42, 0, 0, 0, 49, 52, 47, 43,
	46, 53, 40, 50, 0, 0, 0, 152, 51, 44,
	45, 48, 38, 0, 0, 0, 42, 0, 41, 39,
	49, 52, 47, 43, 46, 53, 40, 50, 0, 0,
	0, 0, 51, 44, 45, 48, 38, 0, 147, 0,
	0, 0, 41, 39, 42, 0, 0, 0, 49, 52,
	47, 43, 46, 53, 40, 50, 0, 0, 0, 0,
	51, 44, 45, 48, 38, 0, 0, 37, 42, 0,
	41, 39, 49, 52, 47, 43, 46, 53, 40, 50,
	0, 0, 0, 0, 51, 44, 45, 48, 38, 0,
	109, 0, 0, 0, 41, 39, 42, 0, 0, 0,
	49, 52, 47, 43, 46, 53, 40, 50, 0, 0,
	0, 0, 51, 44, 45, 48, 38, 0, 0, 0,
	42, 0, 41, 39, 49, 52, 47, 43, 46, 53,
	40, 50, 0, 0, 0, 0, 51, 44, 45, 48,
	38, 0, 0, 0, 42, 0, 41, 39, 49, 52,
	47, 65, 66, 67, 40, 50, 0, 0, 0, 0,
	51, 44, 45, 48, 38, 0, 0, 0, 0, 0,
	41, 39,
}
var mmPact = [...]int{

	102, -1000, 22, 109, 84, -1000, 37, -1000, -1000, 257,
	45, -1000, -1000, -1000, -1000, -1000, 570, -1000, -1000, -1000,
	-1000, 570, 570, 109, 84, 35, 84, -1000, 169, -1000,
	594, 171, -1000, -1000, -1000, -1000, -11, -16, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 132, 266, 19, -1000, 166,
	161, 84, -1000, -1000, 68, -1000, -1000, -1000, 224, -1000,
	570, 570, 43, -1000, 290, -1000, 570, -1000, -1000, 148,
	-1000, 570, -1000, -1000, -1000, -1000, 265, -1000, -1000, -1000,
	28, 28, -1000, -1000, 189, 188, 187, 186, 546, 159,
	290, 55, -1000, 323, 94, -39, -39, -39, 518, -1000,
	-1000, 184, -1000, -1000, 133, -1000, -27, 323, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -17, 135, 250, 249, 248,
	247, 191, 494, 110, 26, -1000, -1000, -1000, -1000, 466,
	103, -1000, -1000, -1000, -1000, 233, 158, -1000, 86, 87,
	232, 357, 173, 66, 119, 84, -1000, 290, 72, 157,
	153, -1000, -1000, -1000, 65, 61, -1000, -1000, 231, -1000,
	71, 84, 152, 155, -1000, 141, -1000, -1000, 28, -1000,
	217, -1000, -1000, 60, -1000, 130, 138, -1000, 10, 216,
	-1000, 54, 28, 82, -1000, -1000, 215, -1000, -1000, 442,
	214, -1000, 414, -1000, 182, 180, 178, 177, 176, 81,
	-1000, -1000, 386, -1000, -1000, -1000, 204, 15, 12, 11,
	21, 56, -1000, -1000, 203, -1000, 200, 198, 196, 195,
	192, -1000, -1000, -1000, -1000, -1000, -1000,
}
var mmPgo = [...]int{

	0, 365, 0, 138, 10, 7, 364, 5, 358, 14,
	135, 357, 356, 289, 355, 342, 341, 340, 339, 338,
	6, 2, 335, 334, 4, 1, 290, 13, 9, 325,
	8, 324, 323, 322, 3, 308, 307, 306, 302, 275,
}
var mmR1 = [...]int{

	0, 39, 39, 39, 39, 39, 39, 39, 1, 1,
	13, 13, 10, 10, 10, 12, 11, 37, 37, 38,
	38, 38, 38, 38, 38, 17, 17, 16, 16, 3,
	3, 9, 9, 20, 20, 14, 14, 21, 21, 15,
	15, 15, 15, 15, 15, 23, 5, 7, 4, 4,
	4, 4, 4, 4, 4, 6, 6, 6, 22, 22,
	22, 36, 19, 19, 18, 18, 31, 31, 30, 30,
	30, 8, 8, 8, 8, 35, 35, 33, 33, 33,
	33, 34, 34, 32, 32, 32, 28, 28, 29, 29,
	24, 24, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 27, 27, 25, 25, 25, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2,
}
var mmR2 = [...]int{

	0, 2, 3, 2, 1, 2, 1, 1, 3, 2,
	2, 1, 3, 1, 1, 11, 10, 0, 4, 0,
	5, 5, 5, 5, 5, 0, 4, 0, 3, 3,
	1, 0, 3, 0, 2, 6, 5, 0, 2, 4,
	5, 6, 5, 6, 7, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 0, 6,
	5, 4, 0, 4, 0, 3, 2, 1, 6, 8,
	5, 0, 2, 2, 2, 0, 2, 4, 4, 4,
	4, 0, 2, 4, 8, 7, 3, 1, 5, 3,
	1, 1, 3, 4, 2, 2, 3, 4, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1,
}
var mmChk = [...]int{

	-1000, -39, -1, -13, -30, -26, 58, -10, 23, 11,
	15, 42, 43, 41, -27, 56, 20, -11, -12, 54,
	55, 21, 22, -13, -30, 58, -30, -10, 25, 41,
	-8, -28, 12, -24, -26, -25, -2, 19, 40, 47,
	30, 46, 20, 27, 37, 38, 28, 26, 39, 24,
	31, 36, 25, 29, 16, -29, 41, -3, -2, -2,
	-2, -30, 41, 13, -2, 27, 28, 29, 9, 12,
	44, 44, 9, 16, 8, 7, 44, 13, 13, -35,
	13, 35, -24, 12, -2, -2, 41, 16, -24, -2,
	-20, -20, 14, -33, 27, 28, 29, 30, -34, -2,
	8, -21, -14, 32, -21, 10, 10, 10, 10, 14,
	-32, -2, 13, -24, -23, -15, 34, 33, -4, 49,
	50, 52, 51, 53, 48, -3, 14, -27, -27, -27,
	-25, 10, -34, 14, -6, 45, 46, 47, -4, -9,
	15, 9, 9, 9, 9, -24, 17, 14, -22, 24,
	41, -9, 11, -2, -31, -30, 9, 13, -37, 25,
	25, 13, 9, 9, -5, -2, 41, 12, -5, 9,
	-36, -30, 18, -28, -17, 26, 13, 13, -20, 9,
	-7, 41, 9, -5, 9, -19, 26, 13, 9, 14,
	13, -38, -20, -21, 9, 9, -7, 16, 13, -34,
	14, 9, -16, 14, 36, 37, 38, 39, 29, -21,
	14, 9, -18, 14, 9, 14, -2, 10, 10, 10,
	10, 10, 14, 14, -25, 9, 43, 43, 43, 41,
	31, 9, 9, 9, 9, 9, 9,
}
var mmDef = [...]int{

	0, -2, 0, 4, 6, 7, 0, 11, 71, 0,
	0, 98, 99, 100, 101, 102, 0, 13, 14, 103,
	104, 0, 0, 1, 3, 0, 5, 10, 0, 9,
	0, 0, 94, 87, 90, 91, 106, 0, 108, 109,
	110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
	120, 121, 122, 123, 95, 0, 0, 0, 30, 0,
	0, 2, 8, 75, 0, -2, -2, -2, 0, 92,
	0, 0, 0, 96, 0, 12, 0, 33, 33, 0,
	81, 0, 86, 93, 105, 107, 0, 97, 89, 29,
	37, 37, 70, 76, 0, 0, 0, 0, 0, 0,
	0, 0, 34, 0, 0, 0, 0, 0, 0, 68,
	82, 0, 81, 88, 0, 38, 0, 0, 31, 48,
	49, 50, 51, 52, 53, 54, 0, 0, 0, 0,
	0, 0, 0, 58, 0, 55, 56, 57, 31, 0,
	0, 77, 78, 79, 80, 0, 0, 69, 17, 0,
	0, 0, 0, 0, 0, 67, 83, 0, 25, 0,
	0, 33, 45, 39, 0, 0, 46, 32, 0, 36,
	62, 66, 0, 0, 16, 0, 19, 33, 37, 40,
	0, 47, 42, 0, 35, 0, 0, 81, 0, 0,
	27, 0, 37, 0, 41, 43, 0, 15, 64, 0,
	0, 85, 0, 18, 0, 0, 0, 0, 0, 0,
	60, 44, 0, 61, 84, 26, 0, 0, 0, 0,
	0, 0, 59, 63, 0, 28, 0, 0, 0, 0,
	0, 65, 20, 21, 22, 23, 24,
}
var mmTok1 = [...]int{

	1,
}
var mmTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58,
}
var mmTok3 = [...]int{
	0,
}

var mmErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	mmDebug        = 0
	mmErrorVerbose = false
)

type mmLexer interface {
	Lex(lval *mmSymType) int
	Error(s string)
}

type mmParser interface {
	Parse(mmLexer) int
	Lookahead() int
}

type mmParserImpl struct {
	lval  mmSymType
	stack [mmInitialStackSize]mmSymType
	char  int
}

func (p *mmParserImpl) Lookahead() int {
	return p.char
}

func mmNewParser() mmParser {
	return &mmParserImpl{}
}

const mmFlag = -1000

func mmTokname(c int) string {
	if c >= 1 && c-1 < len(mmToknames) {
		if mmToknames[c-1] != "" {
			return mmToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func mmStatname(s int) string {
	if s >= 0 && s < len(mmStatenames) {
		if mmStatenames[s] != "" {
			return mmStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func mmErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !mmErrorVerbose {
		return "syntax error"
	}

	for _, e := range mmErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + mmTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := mmPact[state]
	for tok := TOKSTART; tok-1 < len(mmToknames); tok++ {
		if n := base + tok; n >= 0 && n < mmLast && mmChk[mmAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if mmDef[state] == -2 {
		i := 0
		for mmExca[i] != -1 || mmExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; mmExca[i] >= 0; i += 2 {
			tok := mmExca[i]
			if tok < TOKSTART || mmExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if mmExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += mmTokname(tok)
	}
	return res
}

func mmlex1(lex mmLexer, lval *mmSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = mmTok1[0]
		goto out
	}
	if char < len(mmTok1) {
		token = mmTok1[char]
		goto out
	}
	if char >= mmPrivate {
		if char < mmPrivate+len(mmTok2) {
			token = mmTok2[char-mmPrivate]
			goto out
		}
	}
	for i := 0; i < len(mmTok3); i += 2 {
		token = mmTok3[i+0]
		if token == char {
			token = mmTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = mmTok2[1] /* unknown char */
	}
	if mmDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", mmTokname(token), uint(char))
	}
	return char, token
}

func mmParse(mmlex mmLexer) int {
	return mmNewParser().Parse(mmlex)
}

func (mmrcvr *mmParserImpl) Parse(mmlex mmLexer) int {
	var mmn int
	var mmVAL mmSymType
	var mmDollar []mmSymType
	_ = mmDollar // silence set and not used
	mmS := mmrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	mmstate := 0
	mmrcvr.char = -1
	mmtoken := -1 // mmrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		mmstate = -1
		mmrcvr.char = -1
		mmtoken = -1
	}()
	mmp := -1
	goto mmstack

ret0:
	return 0

ret1:
	return 1

mmstack:
	/* put a state and value onto the stack */
	if mmDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", mmTokname(mmtoken), mmStatname(mmstate))
	}

	mmp++
	if mmp >= len(mmS) {
		nyys := make([]mmSymType, len(mmS)*2)
		copy(nyys, mmS)
		mmS = nyys
	}
	mmS[mmp] = mmVAL
	mmS[mmp].yys = mmstate

mmnewstate:
	mmn = mmPact[mmstate]
	if mmn <= mmFlag {
		goto mmdefault /* simple state */
	}
	if mmrcvr.char < 0 {
		mmrcvr.char, mmtoken = mmlex1(mmlex, &mmrcvr.lval)
	}
	mmn += mmtoken
	if mmn < 0 || mmn >= mmLast {
		goto mmdefault
	}
	mmn = mmAct[mmn]
	if mmChk[mmn] == mmtoken { /* valid shift */
		mmrcvr.char = -1
		mmtoken = -1
		mmVAL = mmrcvr.lval
		mmstate = mmn
		if Errflag > 0 {
			Errflag--
		}
		goto mmstack
	}

mmdefault:
	/* default state action */
	mmn = mmDef[mmstate]
	if mmn == -2 {
		if mmrcvr.char < 0 {
			mmrcvr.char, mmtoken = mmlex1(mmlex, &mmrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if mmExca[xi+0] == -1 && mmExca[xi+1] == mmstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			mmn = mmExca[xi+0]
			if mmn < 0 || mmn == mmtoken {
				break
			}
		}
		mmn = mmExca[xi+1]
		if mmn < 0 {
			goto ret0
		}
	}
	if mmn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			mmlex.Error(mmErrorMessage(mmstate, mmtoken))
			Nerrs++
			if mmDebug >= 1 {
				__yyfmt__.Printf("%s", mmStatname(mmstate))
				__yyfmt__.Printf(" saw %s\n", mmTokname(mmtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for mmp >= 0 {
				mmn = mmPact[mmS[mmp].yys] + mmErrCode
				if mmn >= 0 && mmn < mmLast {
					mmstate = mmAct[mmn] /* simulate a shift of "error" */
					if mmChk[mmstate] == mmErrCode {
						goto mmstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if mmDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", mmS[mmp].yys)
				}
				mmp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if mmDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", mmTokname(mmtoken))
			}
			if mmtoken == mmEofCode {
				goto ret1
			}
			mmrcvr.char = -1
			mmtoken = -1
			goto mmnewstate /* try again in the same state */
		}
	}

	/* reduction by production mmn */
	if mmDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", mmn, mmStatname(mmstate))
	}

	mmnt := mmn
	mmpt := mmp
	_ = mmpt // guard against "declared and not used"

	mmp -= mmR2[mmn]
	// mmp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if mmp+1 >= len(mmS) {
		nyys := make([]mmSymType, len(mmS)*2)
		copy(nyys, mmS)
		mmS = nyys
	}
	mmVAL = mmS[mmp+1]

	/* consult goto table to find next state */
	mmn = mmR1[mmn]
	mmg := mmPgo[mmn]
	mmj := mmg + mmS[mmp].yys + 1

	if mmj >= mmLast {
		mmstate = mmAct[mmg]
	} else {
		mmstate = mmAct[mmj]
		if mmChk[mmstate] != -mmn {
			mmstate = mmAct[mmg]
		}
	}
	// dummy call; replaced with literal code
	switch mmnt {

	case 1:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:94
		{
			{
				global := NewAst(mmDollar[2].decs, nil, mmDollar[2].srcfile)
				global.Includes = mmDollar[1].includes
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 2:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:100
		{
			{
				global := NewAst(mmDollar[2].decs, mmDollar[3].call, mmDollar[2].srcfile)
				global.Includes = mmDollar[1].includes
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 3:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:106
		{
			{
				global := NewAst(nil, mmDollar[2].call, mmDollar[2].srcfile)
				global.Includes = mmDollar[1].includes
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 4:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:112
		{
			{
				global := NewAst(mmDollar[1].decs, nil, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 5:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:117
		{
			{
				global := NewAst(mmDollar[1].decs, mmDollar[2].call, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 6:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:122
		{
			{
				global := NewAst(nil, mmDollar[1].call, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
			}
		}
	case 7:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:127
		{
			{
				global := NewAst(nil, nil, mmDollar[1].srcfile)
				mmlex.(*mmLexInfo).global = global
				mmlex.(*mmLexInfo).exp = mmDollar[1].vexp
			}
		}
	case 8:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:136
		{
			{
				mmVAL.includes = append(mmDollar[1].includes, &Include{
					Node:  NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Value: mmDollar[3].intern.unquote(mmDollar[3].val),
				})
			}
		}
	case 9:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:142
		{
			{
				mmVAL.includes = []*Include{
					{
						Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
						Value: mmDollar[2].intern.unquote(mmDollar[2].val),
					},
				}
			}
		}
	case 10:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:152
		{
			{
				mmVAL.decs = append(mmDollar[1].decs, mmDollar[2].dec)
			}
		}
	case 11:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:154
		{
			{
				mmVAL.decs = []Dec{mmDollar[1].dec}
			}
		}
	case 12:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:159
		{
			{
				mmVAL.dec = &UserType{
					Node: NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:   mmDollar[2].intern.Get(mmDollar[2].val),
				}
			}
		}
	case 15:
		mmDollar = mmS[mmpt-11 : mmpt+1]
//line grammar.y:169
		{
			{
				mmVAL.dec = &Pipeline{
					Node:      NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:        mmDollar[2].intern.Get(mmDollar[2].val),
					InParams:  mmDollar[4].i_params,
					OutParams: mmDollar[5].o_params,
					Calls:     mmDollar[8].calls,
					Callables: &Callables{Table: make(map[string]Callable)},
					Ret:       mmDollar[9].retstm,
					Retain:    mmDollar[10].plretains,
				}
			}
		}
	case 16:
		mmDollar = mmS[mmpt-10 : mmpt+1]
//line grammar.y:183
		{
			{
				mmVAL.dec = &Stage{
					Node:      NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:        mmDollar[2].intern.Get(mmDollar[2].val),
					InParams:  mmDollar[4].i_params,
					OutParams: mmDollar[5].o_params,
					Src:       mmDollar[6].src,
					ChunkIns:  mmDollar[8].par_tuple.Ins,
					ChunkOuts: mmDollar[8].par_tuple.Outs,
					Split:     mmDollar[8].par_tuple.Present,
					Resources: mmDollar[9].res,
					Retain:    mmDollar[10].stretains,
				}
			}
		}
	case 17:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:200
		{
			{
				mmVAL.res = nil
			}
		}
	case 18:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:202
		{
			{
				mmDollar[3].res.Node = NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile)
				mmVAL.res = mmDollar[3].res
			}
		}
	case 19:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:210
		{
			{
				mmVAL.res = new(Resources)
			}
		}
	case 20:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:212
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.ThreadNode = &n
				i := parseInt(mmDollar[4].val)
				mmDollar[1].res.Threads = int16(i)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 21:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:220
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.MemNode = &n
				i := parseInt(mmDollar[4].val)
				mmDollar[1].res.MemGB = int16(i)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 22:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:228
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.VMemNode = &n
				i := parseInt(mmDollar[4].val)
				mmDollar[1].res.VMemGB = int16(i)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 23:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:236
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.SpecialNode = &n
				mmDollar[1].res.Special = mmDollar[4].intern.unquote(mmDollar[4].val)
				mmVAL.res = mmDollar[1].res
			}
		}
	case 24:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:243
		{
			{
				n := NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile)
				mmDollar[1].res.VolatileNode = &n
				mmDollar[1].res.StrictVolatile = true
				mmVAL.res = mmDollar[1].res
			}
		}
	case 25:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:253
		{
			{
				mmVAL.stretains = nil
			}
		}
	case 26:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:255
		{
			{
				mmVAL.stretains = &RetainParams{
					Node:   NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Params: mmDollar[3].retains,
				}
			}
		}
	case 27:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:265
		{
			{
				mmVAL.retains = nil
			}
		}
	case 28:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:267
		{
			{
				mmVAL.retains = append(mmDollar[1].retains, &RetainParam{
					Node: NewAstNode(mmDollar[2].loc, mmDollar[2].srcfile),
					Id:   mmDollar[2].intern.Get(mmDollar[2].val),
				})
			}
		}
	case 29:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:278
		{
			{
				idd := append(mmDollar[1].val, '.')
				mmVAL.val = append(idd, mmDollar[3].val...)
			}
		}
	case 30:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:283
		{
			{
				// set capacity == length so append doesn't overwrite
				// other parts of the buffer later.
				mmVAL.val = mmDollar[1].val[:len(mmDollar[1].val):len(mmDollar[1].val)]
			}
		}
	case 31:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:292
		{
			{
				mmVAL.arr = 0
			}
		}
	case 32:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:294
		{
			{
				mmVAL.arr++
			}
		}
	case 33:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:299
		{
			{
				mmVAL.i_params = &InParams{Table: make(map[string]*InParam)}
			}
		}
	case 34:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:301
		{
			{
				mmDollar[1].i_params.List = append(mmDollar[1].i_params.List, mmDollar[2].inparam)
				mmVAL.i_params = mmDollar[1].i_params
			}
		}
	case 35:
		mmDollar = mmS[mmpt-6 : mmpt+1]
//line grammar.y:309
		{
			{
				mmVAL.inparam = &InParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
					Help:     unquote(mmDollar[5].val),
				}
			}
		}
	case 36:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:317
		{
			{
				mmVAL.inparam = &InParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
				}
			}
		}
	case 37:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:327
		{
			{
				mmVAL.o_params = &OutParams{Table: make(map[string]*OutParam)}
			}
		}
	case 38:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:329
		{
			{
				mmDollar[1].o_params.List = append(mmDollar[1].o_params.List, mmDollar[2].outparam)
				mmVAL.o_params = mmDollar[1].o_params
			}
		}
	case 39:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:337
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       default_out_name,
				}
			}
		}
	case 40:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:344
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       default_out_name,
					Help:     unquote(mmDollar[4].val),
				}
			}
		}
	case 41:
		mmDollar = mmS[mmpt-6 : mmpt+1]
//line grammar.y:352
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       default_out_name,
					Help:     unquote(mmDollar[4].val),
					OutName:  mmDollar[5].intern.unquote(mmDollar[5].val),
				}
			}
		}
	case 42:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:361
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
				}
			}
		}
	case 43:
		mmDollar = mmS[mmpt-6 : mmpt+1]
//line grammar.y:368
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
					Help:     unquote(mmDollar[5].val),
				}
			}
		}
	case 44:
		mmDollar = mmS[mmpt-7 : mmpt+1]
//line grammar.y:376
		{
			{
				mmVAL.outparam = &OutParam{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Tname:    mmDollar[2].intern.Get(mmDollar[2].val),
					ArrayDim: mmDollar[3].arr,
					Id:       mmDollar[4].intern.Get(mmDollar[4].val),
					Help:     unquote(mmDollar[5].val),
					OutName:  mmDollar[6].intern.unquote(mmDollar[6].val),
				}
			}
		}
	case 45:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:388
		{
			{
				cmd := strings.TrimSpace(mmDollar[3].intern.unquote(mmDollar[3].val))
				stagecodeParts := strings.Fields(mmDollar[3].intern.unquote(mmDollar[3].val))
				mmVAL.src = &SrcParam{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Lang: StageLanguage(mmDollar[2].intern.Get(mmDollar[2].val)),
					cmd:  cmd,
					Path: stagecodeParts[0],
					Args: stagecodeParts[1:],
				}
			}
		}
	case 58:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:425
		{
			{
				mmVAL.par_tuple = paramsTuple{
					Present: false,
					Ins:     &InParams{Table: make(map[string]*InParam)},
					Outs:    &OutParams{Table: make(map[string]*OutParam)},
				}
			}
		}
	case 59:
		mmDollar = mmS[mmpt-6 : mmpt+1]
//line grammar.y:433
		{
			{
				mmVAL.par_tuple = paramsTuple{
					Present: true,
					Ins:     mmDollar[4].i_params,
					Outs:    mmDollar[5].o_params,
				}
			}
		}
	case 60:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:439
		{
			{
				mmVAL.par_tuple = paramsTuple{
					Present: true,
					Ins:     mmDollar[3].i_params,
					Outs:    mmDollar[4].o_params,
				}
			}
		}
	case 61:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:448
		{
			{
				mmVAL.retstm = &ReturnStm{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Bindings: mmDollar[3].bindings,
				}
			}
		}
	case 62:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:456
		{
			{
				mmVAL.plretains = nil
			}
		}
	case 63:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:458
		{
			{
				mmVAL.plretains = &PipelineRetains{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Refs: mmDollar[3].reflist,
				}
			}
		}
	case 64:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:465
		{
			{
				mmVAL.reflist = nil
			}
		}
	case 65:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:467
		{
			{
				mmVAL.reflist = append(mmDollar[1].reflist, mmDollar[2].rexp)
			}
		}
	case 66:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:471
		{
			{
				mmVAL.calls = append(mmDollar[1].calls, mmDollar[2].call)
			}
		}
	case 67:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:473
		{
			{
				mmVAL.calls = []*CallStm{mmDollar[1].call}
			}
		}
	case 68:
		mmDollar = mmS[mmpt-6 : mmpt+1]
//line grammar.y:478
		{
			{
				id := mmDollar[3].intern.Get(mmDollar[3].val)
				mmVAL.call = &CallStm{
					Node:      NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Modifiers: mmDollar[2].modifiers,
					Id:        id,
					DecId:     id,
					Bindings:  mmDollar[5].bindings,
				}
			}
		}
	case 69:
		mmDollar = mmS[mmpt-8 : mmpt+1]
//line grammar.y:487
		{
			{
				mmVAL.call = &CallStm{
					Node:      NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Modifiers: mmDollar[2].modifiers,
					Id:        mmDollar[5].intern.Get(mmDollar[5].val),
					DecId:     mmDollar[3].intern.Get(mmDollar[3].val),
					Bindings:  mmDollar[7].bindings,
				}
			}
		}
	case 70:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:495
		{
			{
				mmDollar[1].call.Modifiers.Bindings = mmDollar[4].bindings
				mmVAL.call = mmDollar[1].call
			}
		}
	case 71:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:503
		{
			{
				mmVAL.modifiers = new(Modifiers)
			}
		}
	case 72:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:505
		{
			{
				mmVAL.modifiers.Local = true
			}
		}
	case 73:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:507
		{
			{
				mmVAL.modifiers.Preflight = true
			}
		}
	case 74:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:509
		{
			{
				mmVAL.modifiers.Volatile = true
			}
		}
	case 75:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:514
		{
			{
				mmVAL.bindings = &BindStms{
					Node:  NewAstNode(mmDollar[0].loc, mmDollar[0].srcfile),
					Table: make(map[string]*BindStm),
				}
			}
		}
	case 76:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:519
		{
			{
				mmDollar[1].bindings.List = append(mmDollar[1].bindings.List, mmDollar[2].binding)
				mmVAL.bindings = mmDollar[1].bindings
			}
		}
	case 77:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:527
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   local,
					Exp:  mmDollar[3].vexp,
				}
			}
		}
	case 78:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:533
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   preflight,
					Exp:  mmDollar[3].vexp,
				}
			}
		}
	case 79:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:539
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   volatile,
					Exp:  mmDollar[3].vexp,
				}
			}
		}
	case 80:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:545
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   disabled,
					Exp:  mmDollar[3].rexp,
				}
			}
		}
	case 81:
		mmDollar = mmS[mmpt-0 : mmpt+1]
//line grammar.y:553
		{
			{
				mmVAL.bindings = &BindStms{
					Node:  NewAstNode(mmDollar[0].loc, mmDollar[0].srcfile),
					Table: make(map[string]*BindStm),
				}
			}
		}
	case 82:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:558
		{
			{
				mmDollar[1].bindings.List = append(mmDollar[1].bindings.List, mmDollar[2].binding)
				mmVAL.bindings = mmDollar[1].bindings
			}
		}
	case 83:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:566
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   mmDollar[1].intern.Get(mmDollar[1].val),
					Exp:  mmDollar[3].exp,
				}
			}
		}
	case 84:
		mmDollar = mmS[mmpt-8 : mmpt+1]
//line grammar.y:572
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   mmDollar[1].intern.Get(mmDollar[1].val),
					Exp: &ValExp{
						Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
						Kind:  KindArray,
						Value: mmDollar[5].exps,
					},
					Sweep: true,
				}
			}
		}
	case 85:
		mmDollar = mmS[mmpt-7 : mmpt+1]
//line grammar.y:583
		{
			{
				mmVAL.binding = &BindStm{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Id:   mmDollar[1].intern.Get(mmDollar[1].val),
					Exp: &ValExp{
						Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
						Kind:  KindArray,
						Value: mmDollar[5].exps,
					},
					Sweep: true,
				}
			}
		}
	case 86:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:597
		{
			{
				mmVAL.exps = append(mmDollar[1].exps, mmDollar[3].exp)
			}
		}
	case 87:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:599
		{
			{
				mmVAL.exps = []Exp{mmDollar[1].exp}
			}
		}
	case 88:
		mmDollar = mmS[mmpt-5 : mmpt+1]
//line grammar.y:604
		{
			{
				mmDollar[1].kvpairs[unquote(mmDollar[3].val)] = mmDollar[5].exp
				mmVAL.kvpairs = mmDollar[1].kvpairs
			}
		}
	case 89:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:609
		{
			{
				mmVAL.kvpairs = map[string]Exp{unquote(mmDollar[1].val): mmDollar[3].exp}
			}
		}
	case 90:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:614
		{
			{
				mmVAL.exp = mmDollar[1].vexp
			}
		}
	case 91:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:616
		{
			{
				mmVAL.exp = mmDollar[1].rexp
			}
		}
	case 92:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:620
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindArray,
					Value: mmDollar[2].exps,
				}
			}
		}
	case 93:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:626
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindArray,
					Value: mmDollar[2].exps,
				}
			}
		}
	case 94:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:632
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindArray,
					Value: make([]Exp, 0),
				}
			}
		}
	case 95:
		mmDollar = mmS[mmpt-2 : mmpt+1]
//line grammar.y:638
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindMap,
					Value: make(map[string]interface{}, 0),
				}
			}
		}
	case 96:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:644
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindMap,
					Value: mmDollar[2].kvpairs,
				}
			}
		}
	case 97:
		mmDollar = mmS[mmpt-4 : mmpt+1]
//line grammar.y:650
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindMap,
					Value: mmDollar[2].kvpairs,
				}
			}
		}
	case 98:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:656
		{
			{ // Lexer guarantees parseable float strings.
				f := parseFloat(mmDollar[1].val)
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindFloat,
					Value: f,
				}
			}
		}
	case 99:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:665
		{
			{ // Lexer guarantees parseable int strings.
				i := parseInt(mmDollar[1].val)
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindInt,
					Value: i,
				}
			}
		}
	case 100:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:674
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindString,
					Value: unquote(mmDollar[1].val),
				}
			}
		}
	case 102:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:681
		{
			{
				mmVAL.vexp = &ValExp{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind: KindNull,
				}
			}
		}
	case 103:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:689
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindBool,
					Value: true,
				}
			}
		}
	case 104:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:695
		{
			{
				mmVAL.vexp = &ValExp{
					Node:  NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:  KindBool,
					Value: false,
				}
			}
		}
	case 105:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:703
		{
			{
				mmVAL.rexp = &RefExp{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:     KindCall,
					Id:       mmDollar[1].intern.Get(mmDollar[1].val),
					OutputId: mmDollar[3].intern.Get(mmDollar[3].val),
				}
			}
		}
	case 106:
		mmDollar = mmS[mmpt-1 : mmpt+1]
//line grammar.y:710
		{
			{
				mmVAL.rexp = &RefExp{
					Node:     NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind:     KindCall,
					Id:       mmDollar[1].intern.Get(mmDollar[1].val),
					OutputId: default_out_name,
				}
			}
		}
	case 107:
		mmDollar = mmS[mmpt-3 : mmpt+1]
//line grammar.y:717
		{
			{
				mmVAL.rexp = &RefExp{
					Node: NewAstNode(mmDollar[1].loc, mmDollar[1].srcfile),
					Kind: KindSelf,
					Id:   mmDollar[3].intern.Get(mmDollar[3].val),
				}
			}
		}
	}
	goto mmstack /* stack new state and value */
}
