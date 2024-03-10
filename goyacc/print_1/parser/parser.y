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

arith_st: plus_st | minus_st | mul_st | div_st

label_st: Label
{
    PrintYySymDebug($1, yylex)
}

Value: String | Number

print_st: Print Value {
    println($2.val)
}
| Print Label
{
    PrintYySym($2, yylex)
}

assignment: Label '=' String
{
    GlobalVarWrite(yylex, $1, $3.val)
}
| Label '=' Number
{
    GlobalVarWrite(yylex, $1, $3.val)
}
| Label '=' Label
{
    GlobalVarCopy(yylex, $1, $3)
}
| Label '=' arith_st {
    GlobalVarWrite(yylex, $1, $3.val)
}
| Label '='


type_st: Type Label {
    PrintType($1, yylex)
}

plus_st: Number '+' Number
{
	val := ValAdd($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| Label '+' Number
{
    v1 := GlobalVarRead(yylex, $1)
    res := ValAdd(v1, $3.val)
    $$ = yySymType {
        val: res,
    }
}
| Number '+' Label
{
    v3 := GlobalVarRead(yylex, $3)
    res := ValAdd(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}
| Label '+' Label
{
    v1 := GlobalVarRead(yylex, $1)
    v3 := GlobalVarRead(yylex, $3)
    res := ValAdd(v1, v3)
    $$ = yySymType {
        val: res,
    }
}
| arith_st '+' Number {
	val := ValAdd($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| arith_st '+' Label {
    v3 := GlobalVarRead(yylex, $3)
    res := ValAdd(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}

minus_st: Number '-' Number
{
	val := ValMinus($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| Label '-' Number
{
    v1 := GlobalVarRead(yylex, $1)
    res := ValMinus(v1, $3.val)
    $$ = yySymType {
        val: res,
    }
}
| Number '-' Label
{
    v3 := GlobalVarRead(yylex, $3)
    res := ValMinus(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}
| Label '-' Label
{
    v1 := GlobalVarRead(yylex, $1)
    v3 := GlobalVarRead(yylex, $3)
    res := ValMinus(v1, v3)
    $$ = yySymType {
        val: res,
    }
}
| arith_st '-' Number {
	val := ValMinus($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| arith_st '-' Label {
    v3 := GlobalVarRead(yylex, $3)
    res := ValMinus(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}

mul_st: Number '*' Number
{
	val := ValMul($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| Label '*' Number
{
    v1 := GlobalVarRead(yylex, $1)
    res := ValMul(v1, $3.val)
    $$ = yySymType {
        val: res,
    }
}
| Number '*' Label
{
    v3 := GlobalVarRead(yylex, $3)
    res := ValMul(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}
| Label '*' Label
{
    v1 := GlobalVarRead(yylex, $1)
    v3 := GlobalVarRead(yylex, $3)
    res := ValMul(v1, v3)
    $$ = yySymType {
        val: res,
    }
}
| arith_st '*' Number {
	val := ValMul($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| arith_st '*' Label {
    v3 := GlobalVarRead(yylex, $3)
    res := ValMul(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}

div_st: Number '/' Number
{
	val := ValMul($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| Label '/' Number
{
    v1 := GlobalVarRead(yylex, $1)
    res := ValMul(v1, $3.val)
    $$ = yySymType {
        val: res,
    }
}
| Number '/' Label
{
    v3 := GlobalVarRead(yylex, $3)
    res := ValMul(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}
| Label '/' Label
{
    v1 := GlobalVarRead(yylex, $1)
    v3 := GlobalVarRead(yylex, $3)
    res := ValMul(v1, v3)
    $$ = yySymType {
        val: res,
    }
}
| arith_st '/' Number {
	val := ValMul($1.val, $3.val)
	sy := yySymType{
		val: val,
	}
    $$ = sy
}
| arith_st '/' Label {
    v3 := GlobalVarRead(yylex, $3)
    res := ValMul(v3, $1.val)
    $$ = yySymType {
        val: res,
    }
}