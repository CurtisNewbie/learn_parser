package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/curtisnewbie/miso/miso"
	"github.com/spf13/cast"
)

var (
	vmrt    = newVm()
	printKw = []string{"print", "echo"}
	typeKw  = "type"
)

type vm struct {
	script    string
	offset    int
	globalvar map[string]any
}

func (v *vm) Lex(lval *yySymType) int {
	for {
		miso.Debug("run for")
		if c, ok := v.next(); ok {
			switch {
			case c == '\'':
				return v.parseString(lval)
			case c == '\n' || c == ' ' || c == '\t':
				v.move(1)
				continue
			case unicode.IsLetter(c):
				// TODO: refactor to a table like matching
				if d, ok := v.parseKeywords(lval, printKw, Print); ok {
					return d
				}
				if d, ok := v.parseKeyword(lval, typeKw, Type); ok {
					return d
				}
				return v.parseLabel(lval)
			case unicode.IsDigit(c):
				return v.parseNumber(lval)
			case c == '/' && v.lookAheadIs(1, '/'):
				v.move(len(v.script) - v.offset)
				return 0
			default:
				miso.Debugf("default %v, %v", c, string(c))
				v.move(1)
				return int(c)
			}
		} else {
			break
		}
	}
	return 0
}

func (v *vm) Error(s string) {
	panic(fmt.Sprintf("%v - %v, %#v", s, v.offset, v))
}

func (v *vm) next() (rune, bool) {
	miso.Debug("next")
	if v.offset >= len(v.script) {
		return 0, false
	}
	c := v.script[v.offset]
	return rune(c), true
}

func (v *vm) lookAheadAt(n int) (rune, bool) {
	miso.Debugf("lookAheadAt %v", n)
	if v.offset+n >= len(v.script) {
		return 0, false
	}
	c := v.script[v.offset+n]
	return rune(c), true
}

func (v *vm) lookAheadIs(n int, c rune) bool {
	if ch, ok := v.lookAheadAt(1); ok && ch == c {
		return true
	}
	return false
}

func (v *vm) move(gap int) {
	v.offset = v.offset + gap
	miso.Debugf("move %v to %v", gap, v.offset)
}

func (v *vm) parseNumber(lval *yySymType) int {
	miso.Debugf("parseNumber, remaining=%v", v.remaining())
	i := 1
	isFloat := false
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c != '.' && !unicode.IsDigit(c) {
				break
			} else {
				if c == '.' {
					isFloat = true
				}
				i += 1
			}
		} else {
			break
		}
	}
	if isFloat {
		lval.val = cast.ToFloat64(v.script[v.offset : v.offset+i])
	} else {
		lval.val = cast.ToInt64(v.script[v.offset : v.offset+i])
	}
	miso.Debugf("label.val: %v", lval.val)
	v.move(i)
	miso.Debugf("offset: %v", v.remaining())
	return Number
}

func (v *vm) parseLabel(lval *yySymType) int {
	miso.Debugf("parseLabel, remaining=%v", v.remaining())
	i := 1
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if !unicode.IsLetter(c) {
				break
			} else {
				i += 1
			}
		} else {
			break
		}
	}
	lval.val = v.script[v.offset : v.offset+i]
	miso.Debugf("label.val: %v", lval.val)
	v.move(i)
	miso.Debugf("offset: %v", v.script[v.offset:])
	return Label
}

func (v *vm) remaining() string {
	return v.script[v.offset:]
}

func (v *vm) parseString(lval *yySymType) int {
	miso.Debugf("parsestring, starting at: %v", v.offset)
	i := 2
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c == '\'' {
				lval.val = v.script[v.offset+1 : v.offset+i]
				miso.Debugf("lval.val : %v, %v, %v", v.script[v.offset+1:v.offset+i], v.offset, v.offset+i)
				v.move(i + 1)
				return String
			}
			i += 1
		} else {
			break
		}
	}
	panic(fmt.Sprintf("illegal syntax for string, %#v", v))
}

func (v *vm) parseKeywords(lval *yySymType, keywords []string, keywordType int) (int, bool) {
	for _, kw := range keywords {
		if i, ok := v.parseKeyword(lval, kw, keywordType); ok {
			return i, true
		}
	}
	return 0, false
}

func (v *vm) parseKeyword(lval *yySymType, keyword string, keywordType int) (int, bool) {
	miso.Debugf("parseKeyword, %v, %v", keyword, keywordType)

	kwl := len(keyword)
	if v.offset+kwl >= len(v.script) {
		return 0, false
	}

	pre := v.script[v.offset : v.offset+kwl]
	miso.Debugf("pre: %v", pre)
	if pre == keyword {
		v.move(kwl)
		return keywordType, true
	}
	return 0, false
}

func newVm() *vm {
	return &vm{
		globalvar: make(map[string]any),
	}
}

func Parse(s string) {
	lines := strings.Split(s, "\n")
	miso.SetLogLevel("info")
	yyErrorVerbose = true
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		vmrt.script = l
		vmrt.offset = 0
		yyParse(vmrt)
		miso.Debugf("vars: %#v\n", vmrt.globalvar)
	}
}
