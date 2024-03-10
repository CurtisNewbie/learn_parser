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
%token Comment

%left '='
%left '+' '-'
%left '*' '/'

%start expression

%%

expression:
    assignment
    | statement
    | Comment

statement:
    print_st
    | label_st
    | type_st
    | arith_st

label_st:
    Label { PrintYySymDebug($1) }

Value:
    String | Number

print_st:
    Print Value { PrintYySym($2) }
    | Print Label { PrintGlobalYySym($2) }
    | Print { println("") }

type_st:
    Type Label { PrintType($2) }

assignment:
    Label '=' String { GlobalVarWrite($1, $3.val) }
    | Label '=' eval_expr { GlobalVarWrite($1, $3.val) }
    | Label '=' { SyntaxError() }

arith_st:
    eval_expr '+' eval_expr { $$ = yySymType{ val: ValAdd($1.val, $3.val) } }
    | eval_expr '-' eval_expr { $$ = yySymType{ val: ValMinus($1.val, $3.val) } }
    | eval_expr '*' eval_expr { $$ = yySymType{ val: ValMul($1.val, $3.val) } }
    | eval_expr '/' eval_expr { $$ = yySymType{ val: ValDiv($1.val, $3.val) } }

eval_expr:
    Number { $$ = $1 }
    | Label { $$ = yySymType { val: GlobalVarRead($1) } }
    | arith_st { $$ = $1 }