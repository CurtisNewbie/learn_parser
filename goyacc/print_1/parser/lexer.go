package parser

import (
	"fmt"
	"unicode"
)

//go:generate goyacc parser.y

type vm struct {
	script    string
	offset    int
	globalvar map[string]any
}

func (v *vm) Lex(lval *yySymType) int {
	for {
		// println("debug: for")
		if c, ok := v.next(); ok {
			switch {
			case c == '\'':
				return v.parseString(lval)
			case c == '\n' || c == ' ':
				v.move(1)
				continue
			case c == 'p': // print or assignment
				if d, ok := v.parsePrint(lval); ok {
					return d
				}
				return v.parseLabel(lval)
			case unicode.IsLetter(c):
				return v.parseLabel(lval)
			default:
				// println("debug: default", c, string(c))
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
	// println("debug: next")
	if v.offset >= len(v.script) {
		return 0, false
	}
	c := v.script[v.offset]
	return rune(c), true
}

func (v *vm) lookAheadAt(n int) (rune, bool) {
	// println("debug: lookAheadAt ", n)
	if v.offset+n >= len(v.script) {
		return 0, false
	}
	c := v.script[v.offset+n]
	return rune(c), true
}

func (v *vm) move(gap int) {
	v.offset = v.offset + gap
	// println("debug: move", gap, "to", v.offset)
}

func (v *vm) parseLabel(lval *yySymType) int {
	// println("debug: parseLabel, v.script[v.offset:]", v.script[v.offset:])
	i := 1
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c == ' ' {
				break
			} else {
				i += 1
			}
		} else {
			break
		}
	}
	lval.strv = v.script[v.offset : v.offset+i]
	// println("debug: label.strv ", lval.strv)
	v.move(i)
	// println("debug: offset: ", v.script[v.offset:])
	return Label
}

func (v *vm) parseString(lval *yySymType) int {
	// println("debug: parsestring, starting at", v.offset)
	i := 2
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c == '\'' {
				lval.strv = v.script[v.offset : v.offset+i]
				// println("debug: lval.strv ", v.script[v.offset+1:v.offset+i], " ", v.offset, " ", v.offset+i)
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

func (v *vm) parsePrint(lval *yySymType) (int, bool) {
	if v.offset+5 >= len(v.script) {
		return 0, false
	}

	i := 1
	for {
		if c, ok := v.lookAheadAt(i); ok {
			if c == ' ' {
				i += 1
				continue
			} else {
				break
			}
		} else {
			break
		}
	}

	pre := v.script[v.offset+i : v.offset+i+5]
	if pre == "print" {
		v.move(i + 5)
		return Print, true
	}
	return 0, false
}

func newVm(s string) *vm {
	return &vm{
		script:    s,
		globalvar: make(map[string]any),
	}
}

func Parse(s string) {
	yyErrorVerbose = true
	vm := newVm(s)
	yyParse(vm)
	fmt.Printf("vars: %#v\n", vm.globalvar)
}
