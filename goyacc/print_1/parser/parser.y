%{
package parser

%}

%union{
    hint string
	val any
}

%token String
%token Number
%token Print
%token Label
%token Type
%token Comment
%token Get Put Post Delete Head
%token Header Body Json

%left '='
%left '+' '-'
%left '*' '/'
%left Header ' '

%start expression

%%

expression:
    assignment
    | statement
    | Comment
    | expression Comment

statement:
    print_st
    | label_st
    | type_st
    | arith_st
    | network_st

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

header_sg:
    Header String { $$ = $2 }

header_st:
    header_sg { $$ = $1 }
    | /* empty */
    | header_sg header_sg { $$ = joinHeaders($1, $2) }

body_st:
    /* empty */
    | Body String { $$ = yySymType{ val: $2.val, hint: "text" } }
    | Body Json '(' String ')' { $$ = yySymType{ val: $4.val, hint: "json" } }

network_st:
    Get String header_st body_st { HttpSend("GET", $2.val.(string), $3, $4) }
    | Put String header_st body_st { HttpSend("PUT", $2.val.(string), $3, $4) }
    | Post String header_st body_st { HttpSend("POST", $2.val.(string), $3, $4) }
    | Delete String header_st body_st { HttpSend("DELETE", $2.val.(string), $3, $4) }
    | Head String header_st body_st { HttpSend("HEAD", $2.val.(string), $3, $4) }