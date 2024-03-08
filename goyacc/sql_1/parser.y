%{
package main

%}

%union{
    val interface{}
}

%token LexError
%token <val> String Number

%type <val> value

%%

select:
    'SELECT' delimited_value+ 'FROM' value

delimited_value: value ',' value | value ',' | value

value: String | Number
