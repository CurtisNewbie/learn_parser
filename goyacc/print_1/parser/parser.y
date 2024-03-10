%{
package parser

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

%left '+' '-'
%left '*' '/'

%start expression

%%

expression:
    assignment
    | statement

statement:
    print_st
    | label_st
    | type_st
    | arith_st

label_st:
    Label { PrintYySymDebug($1, yylex) }

Value:
    String | Number

print_st:
    Print Value { println($2.val) }
    | Print Label { PrintYySym($2, yylex) }

type_st:
    Type Label { PrintType($2, yylex) }

assignment:
    Label '=' String { GlobalVarWrite(yylex, $1, $3.val) }
    | Label '=' Number { GlobalVarWrite(yylex, $1, $3.val) }
    | Label '=' Label { GlobalVarCopy(yylex, $1, $3) }
    | Label '=' arith_st { GlobalVarWrite(yylex, $1, $3.val) }
    | Label '=' { SyntaxError() }

arith_st:
    nval_expr '+' nval_expr { $$ = yySymType{ val: ValAdd($1.val, $3.val) } }
    | nval_expr '-' nval_expr { $$ = yySymType{ val: ValMinus($1.val, $3.val) } }
    | nval_expr '*' nval_expr { $$ = yySymType{ val: ValMul($1.val, $3.val) } }
    | nval_expr '/' nval_expr { $$ = yySymType{ val: ValDiv($1.val, $3.val) } }

nval_expr :
    Number { $$ = $1 }
    | Label { $$ = yySymType { val: GlobalVarRead(yylex, $1) } }
    | arith_st { $$ = $1 }