%{
package parser

import (
    "fmt"
    "reflect"
)

%}

%union{
	typ string
	val any
}

%token String
%token Number
%token Print
%token Label
%token Type

%start expression

%%

expression:
    assignment
    | statement

statement:
    print_st
    | label_st
    | type_st

label_st: Label
{
    fmt.Printf("Syntax error")
}

Value: String | Number

print_st: Print Value {
    println($2.val)
}
| Print Label
{
    n := $2.val.(string)
    v := yylex.(*vm).globalvar[n]
    fmt.Printf("%#v\n", v)
}

assignment: Label '=' String
{
    n := $1.val.(string)
    yylex.(*vm).globalvar[n] = $3.val
}
| Label '=' Number
{
    n := $1.val.(string)
    yylex.(*vm).globalvar[n] = $3.val
}
| Label '=' Label
{
    n1 := $1.val.(string)
    n2 := $3.val.(string)
    yylex.(*vm).globalvar[n1] = yylex.(*vm).globalvar[n2]
}

type_st: Type Label {
    n := $2.val.(string)
    val := yylex.(*vm).globalvar[n]
    typ := reflect.TypeOf(val)
    fmt.Printf("%s: <%v>\n", n, typ)
}