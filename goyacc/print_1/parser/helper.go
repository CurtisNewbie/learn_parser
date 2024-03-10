package parser

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

func ValAdd(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) + cast.ToFloat64(b)
	} else {
		return a.(int64) + b.(int64)
	}
}

func ValMinus(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) - cast.ToFloat64(b)
	} else {
		return a.(int64) - b.(int64)
	}
}

func ValMul(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) * cast.ToFloat64(b)
	} else {
		return a.(int64) * b.(int64)
	}
}

func ValDiv(a any, b any) any {
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta.Kind() == reflect.Float64 || tb.Kind() == reflect.Float64 {
		return cast.ToFloat64(a) / cast.ToFloat64(b)
	} else {
		return a.(int64) / b.(int64)
	}
}

func GlobalVar(yylex yyLexer) map[string]any {
	return yylex.(*vm).globalvar
}

func GlobalVarRead(yylex yyLexer, y yySymType) any {
	n := y.val.(string)
	return GlobalVar(yylex)[n]
}

func GlobalVarWrite(yylex yyLexer, y yySymType, val any) {
	n := y.val.(string)
	GlobalVar(yylex)[n] = val
}

func GlobalVarCopy(yylex yyLexer, to yySymType, from yySymType) {
	v := GlobalVarRead(yylex, from)
	GlobalVarWrite(yylex, to, v)
}

func PrintYySymDebug(y yySymType, yylex yyLexer) {
	n := y.val.(string)
	glob := GlobalVar(yylex)
	v := glob[n]
	fmt.Printf("%#v <%v>\n", v, reflect.TypeOf(v))
}

func PrintYySym(y yySymType, yylex yyLexer) {
	n := y.val.(string)
	glob := GlobalVar(yylex)
	v := glob[n]
	fmt.Printf("%v\n", v)
}

func PrintType(y yySymType, yylex yyLexer) {
	n := y.val.(string)
	val := yylex.(*vm).globalvar[n]
	typ := reflect.TypeOf(val)
	fmt.Printf("%s: <%v>\n", n, typ)
}

func SyntaxError() {
	println("syntax error")
}
