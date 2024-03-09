%{
package parser

import (
    "fmt"
)

%}

%union{
    intv int
    strv string
}

%token String
%token Number
%token Print
%token Label

%start expression

%%

expression:
    assignment
    | print_st

print_st: Print String
{
    println($2.strv)
}
| Print Number
{
    println($2.intv)
}
| Print Label
{
    v := yylex.(*vm).globalvar[$2.strv]
    fmt.Printf("%#v\n", v)
}

assignment: Label '=' String
{
    yylex.(*vm).globalvar[$1.strv] = $3.strv
}
| Label '=' Number
{
    yylex.(*vm).globalvar[$1.strv] = $3.intv
}