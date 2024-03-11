package parser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

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
	Debugf("GlobalVarWrite: %v, %v\n", y.val, val)
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
	fmt.Printf("%v: %#v <%v>\n", n, v, reflect.TypeOf(v))
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

func HttpSend(method string, url string, header yySymType, body yySymType) yySymType {
	fmt.Printf("Sending '%v %v', %#v, %#v\n", method, url, header, body)

	var br io.Reader = nil
	if body.val != nil {
		br = bytes.NewReader([]byte(body.val.(string)))
	}

	r, err := http.NewRequest(method, url, br)
	if err != nil {
		fmt.Printf("ERROR: failed to send http request, %v", err)
		return yySymType{}
	}

	if br != nil {
		if body.hint == "json" {
			r.Header.Set("Content-Type", "application/json")
		} else if body.hint == "text" {
			r.Header.Set("Content-Type", "text/plain")
		}
	}
	if harr, ok := header.val.([]string); ok {
		for _, v := range harr {
			tk := strings.SplitN(v, ":", 1)
			if len(tk) > 1 {
				r.Header.Add(tk[0], tk[1])
			}
		}
	} else if hs, ok := header.val.(string); ok {
		tk := strings.SplitN(hs, ":", 1)
		if len(tk) > 1 {
			r.Header.Add(tk[0], tk[1])
		}
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Printf("ERROR: failed to send http request, %v\n", err)
		return yySymType{}
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("ERROR: failed to read http response, %v\n", err)
		return yySymType{}
	}

	Debugf("Returned response: %d, body: %s", res.StatusCode, buf)
	return yySymType{val: map[string]any{"code": res.StatusCode, "body": buf}}
}

func joinHeaders(h1 yySymType, h2 yySymType) yySymType {
	v1 := h1.val
	v2 := h2.val

	arr := []string{}
	if varr, ok := v1.([]string); ok {
		arr = append(arr, varr...)
	} else {
		arr = append(arr, cast.ToString(v1))
	}

	if varr, ok := v2.([]string); ok {
		arr = append(arr, varr...)
	} else {
		arr = append(arr, cast.ToString(v2))
	}
	return yySymType{
		val: arr,
	}
}

func WalkField(obj yySymType, field any) any {
	fieldStr := cast.ToString(field)
	v := GlobalVarRead(obj)
	if v == nil {
		return nil
	}
	if m, ok := v.(map[string]any); ok {
		return m[fieldStr]
	}
	return nil
}
