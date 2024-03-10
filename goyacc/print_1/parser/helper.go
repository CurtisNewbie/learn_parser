package parser

import (
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

var (
	debug = false
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

func GlobalVar() map[string]any {
	return vmrt.globalvar
}

func GlobalVarRead(y yySymType) any {
	n := y.val.(string)
	return GlobalVar()[n]
}

func GlobalVarWrite(y yySymType, val any) {
	n := y.val.(string)
	GlobalVar()[n] = val
}

func GlobalVarCopy(to yySymType, from yySymType) {
	v := GlobalVarRead(from)
	GlobalVarWrite(to, v)
}

func PrintYySymDebug(y yySymType) {
	n := y.val.(string)
	glob := GlobalVar()
	v := glob[n]
	fmt.Printf("%#v <%v>\n", v, reflect.TypeOf(v))
}

func PrintYySym(y yySymType) {
	fmt.Printf("%v\n", y.val)
}

func PrintGlobalYySym(y yySymType) {
	n := y.val.(string)
	glob := GlobalVar()
	v := glob[n]
	fmt.Printf("%v\n", v)
}

func PrintType(y yySymType) {
	n := y.val.(string)
	val := vmrt.globalvar[n]
	typ := reflect.TypeOf(val)
	fmt.Printf("%s: <%v>\n", n, typ)
}

func SyntaxError() {
	println("syntax error")
}

func EnableDebug() {
	debug = true
}

func Debug(args ...any) {
	if debug {
		j := make([]any, len(args)+1)
		j[0] = "[DEBUG]"
		for i := range args {
			j[i+1] = args[i]
		}
		fmt.Println(j...)
	}
}

func Debugf(s string, args ...any) {
	if debug {
		fmt.Printf("[DEBUG] "+s+"\n", args...)
	}
}
