%{
package parser

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
    print($2.strv)
}
| Print Number
{
    print($2.intv)
}

value: String | Number
{
    $$ = $1
}

assignment: Label '=' String
{
    yylex.(*vm).globalvar[$1.strv] = $3
}
| Label '=' Number
{
    yylex.(*vm).globalvar[$1.intv] = $3
}
