// Code generated by goyacc -o y.go parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

//line parser.y:6
type yySymType struct {
	yys int
	val any
}

const String = 57346
const Number = 57347
const Print = 57348
const Label = 57349
const Type = 57350
const Comment = 57351
const Get = 57352
const Put = 57353
const Post = 57354
const Delete = 57355
const Head = 57356
const Header = 57357
const Body = 57358
const Json = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"String",
	"Number",
	"Print",
	"Label",
	"Type",
	"Comment",
	"Get",
	"Put",
	"Post",
	"Delete",
	"Head",
	"Header",
	"Body",
	"Json",
	"'='",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"' '",
	"'('",
	"')'",
	"'.'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 5,
	1, 11,
	9, 11,
	-2, 34,
	-1, 9,
	1, 8,
	9, 8,
	-2, 35,
}

const yyPrivate = 57344

const yyLast = 105

var yyAct = [...]int8{
	40, 70, 11, 58, 14, 51, 52, 59, 47, 85,
	24, 15, 16, 17, 18, 19, 84, 23, 68, 24,
	67, 23, 22, 10, 23, 39, 48, 37, 49, 24,
	23, 83, 53, 55, 56, 57, 61, 62, 63, 64,
	69, 66, 27, 28, 29, 30, 38, 36, 20, 50,
	41, 65, 25, 15, 16, 17, 18, 19, 29, 30,
	43, 71, 60, 74, 75, 76, 77, 72, 80, 21,
	81, 45, 82, 20, 12, 5, 13, 4, 15, 16,
	17, 18, 19, 43, 78, 44, 20, 79, 54, 26,
	73, 35, 34, 33, 32, 31, 42, 9, 46, 8,
	7, 6, 3, 2, 1,
}

var yyPact = [...]int16{
	68, 60, -1000, -1000, -1000, 4, -1000, -1000, -1000, -1000,
	-1000, 3, 28, 82, 23, 91, 90, 89, 88, 87,
	-1000, -1000, 43, 78, 64, 1, -1000, 81, 81, 81,
	81, 47, 47, 47, 47, 47, -1000, 23, -1000, 3,
	-1000, -2, -1000, 27, -1000, -1000, 16, -5, -1000, -7,
	15, -1000, -1000, 37, -1000, 37, -1000, -1000, 45, 47,
	86, 45, 45, 45, 45, 80, -1000, -1000, -1000, -1000,
	-1000, 66, -1000, -1000, -1000, -1000, -1000, -1000, 6, -9,
	-16, -1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 104, 103, 102, 101, 100, 99, 96, 23, 2,
	98, 0, 4, 7, 3, 1,
}

var yyR1 = [...]int8{
	0, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 5, 10, 10, 4, 4, 4, 4, 4, 6,
	11, 11, 11, 2, 2, 2, 2, 2, 2, 7,
	7, 7, 7, 12, 12, 12, 13, 14, 14, 14,
	15, 15, 15, 8, 8, 8, 8, 8, 9, 9,
}

var yyR2 = [...]int8{
	0, 1, 1, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 4, 4, 3, 4, 4, 2,
	4, 4, 4, 3, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 1, 1, 1, 2, 1, 0, 2,
	0, 2, 2, 4, 4, 4, 4, 4, 3, 3,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, 9, 7, -4, -5, -6, -7,
	-8, -9, 6, 8, -12, 10, 11, 12, 13, 14,
	5, 9, 18, 26, 26, 24, 7, 19, 20, 21,
	22, 4, 4, 4, 4, 4, 4, -12, -8, -9,
	-11, 7, -7, 17, 7, 7, -10, 7, 25, -9,
	-8, 4, 5, -12, 7, -12, -12, -12, -14, -13,
	15, -14, -14, -14, -14, 24, 25, 25, 25, 25,
	-15, 16, -13, 4, -15, -15, -15, -15, 4, 7,
	-9, 4, -11, 25, 25, 25,
}

var yyDef = [...]int8{
	0, -2, 1, 2, 3, -2, 5, 6, 7, -2,
	9, 10, 0, 0, 0, 0, 0, 0, 0, 0,
	33, 4, 25, 0, 0, 0, 19, 0, 0, 0,
	0, 38, 38, 38, 38, 38, 23, 24, 26, 27,
	28, 34, 35, 0, 48, 49, 0, 0, 16, 0,
	0, 12, 13, 29, 34, 30, 31, 32, 40, 37,
	0, 40, 40, 40, 40, 0, 14, 15, 17, 18,
	43, 0, 39, 36, 44, 45, 46, 47, 0, 0,
	0, 41, 42, 20, 21, 22,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 23, 3, 3, 3, 3, 3, 3, 3,
	24, 25, 21, 19, 3, 20, 26, 22, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 18,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
}

var yyTok3 = [...]int8{
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
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
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
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
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
	yyn = int(yyPact[yystate])
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
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
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
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
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
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:43
		{
			PrintYySymDebug(yyDollar[1])
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:49
		{
			PrintYySym(yyDollar[3])
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:50
		{
			PrintGlobalYySym(yyDollar[3])
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:51
		{
			println("")
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:52
		{
			PrintYySym(yyDollar[3])
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:53
		{
			PrintYySym(yyDollar[3])
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:56
		{
			PrintType(yyDollar[2])
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:59
		{
			yyVAL = yySymType{val: StrToMap(yyDollar[3].val)}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:60
		{
			yyVAL = yySymType{val: StrToMap(GlobalVarRead(yyDollar[3]))}
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:61
		{
			yyVAL = yySymType{val: StrToMap(yyDollar[3].val)}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:64
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:65
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:66
		{
			SyntaxError()
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:67
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:68
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:69
		{
			GlobalVarWrite(yyDollar[1], yyDollar[3].val)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:72
		{
			yyVAL = yySymType{val: ValAdd(yyDollar[1].val, yyDollar[3].val)}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:73
		{
			yyVAL = yySymType{val: ValMinus(yyDollar[1].val, yyDollar[3].val)}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:74
		{
			yyVAL = yySymType{val: ValMul(yyDollar[1].val, yyDollar[3].val)}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:75
		{
			yyVAL = yySymType{val: ValDiv(yyDollar[1].val, yyDollar[3].val)}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:78
		{
			yyVAL = yyDollar[1]
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:79
		{
			yyVAL = yySymType{val: GlobalVarRead(yyDollar[1])}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:80
		{
			yyVAL = yyDollar[1]
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:83
		{
			yyVAL = yyDollar[2]
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:86
		{
			yyVAL = yyDollar[1]
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:88
		{
			yyVAL = joinHeaders(yyDollar[1], yyDollar[2])
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:92
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:93
		{
			yyVAL = yySymType{val: yyDollar[2].val}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:96
		{
			yyVAL = HttpSend("GET", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:97
		{
			yyVAL = HttpSend("PUT", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:98
		{
			yyVAL = HttpSend("POST", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:99
		{
			yyVAL = HttpSend("DELETE", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:100
		{
			yyVAL = HttpSend("HEAD", yyDollar[2].val.(string), yyDollar[3], yyDollar[4])
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:103
		{
			v := GlobalVarRead(yyDollar[1])
			yyVAL = yySymType{val: WalkField(v, yyDollar[3].val)}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:107
		{
			yyVAL = yySymType{val: WalkField(yyDollar[1].val, yyDollar[3].val)}
		}
	}
	goto yystack /* stack new state and value */
}
