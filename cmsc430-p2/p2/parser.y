// Michael DeAngelo
// CMSC 430 - project 2
// November 17, 2020
// Updated to project 2 specifications from skeleton code provided

/* Compiler Theory and Design
   Dr. Duane J. Jarc */



%{

#include <string>

using namespace std;

#include "listing.h"

int yylex();
void yyerror(const char* message);

%}

%define parse.error verbose

%token IDENTIFIER
%token INT_LITERAL

%token ADDOP MULOP RELOP ANDOP

%token BEGIN_ BOOLEAN END ENDREDUCE FUNCTION INTEGER IS REDUCE RETURNS

%token ARROW CASE ELSE ENDCASE ENDIF IF OTHERS REAL THEN WHEN OROP NOTOP REMOP EXPOP REAL_LITERAL BOOL_LITERAL

%%
function:
	function_header var1ormore body ;

function_header:
	FUNCTION IDENTIFIER parameters RETURNS type ';' | error_ ;

var1ormore:
    var1ormore variable | ;

parameters:
    parameter | parameters ',' parameter | ;

parameter:
    IDENTIFIER ':' type

variable:
	IDENTIFIER ':' type IS statement_ ;

type:
	INTEGER |
	REAL |
	BOOLEAN ;

body:
	BEGIN_ statement_ END ';' ;

statement_:
	statement ';' |
	IF expression THEN statement ';' ELSE statement ';' ENDIF |
	CASE expression IS case OTHERS ARROW statement ';' ENDCASE ;

statement:
	expression |
	statement_ |
	REDUCE operator reductions ENDREDUCE ;

operator:
    NOTOP |
    EXPOP |
	ADDOP |
	MULOP |
	REMOP |
	OROP ;

reductions:
	reductions statement_ |
	;

expression:
	expression ANDOP relation |
	expression OROP relation |
	relation ;

relation:
	relation RELOP term |
	term;

term:
	term ADDOP factor |
	factor ;

factor:
    factor EXPOP primary |
	factor MULOP primary |
	factor REMOP primary |
	primary ;

primary:
	'(' expression ')' |
	NOTOP factor |
	INT_LITERAL | REAL_LITERAL | BOOL_LITERAL |
	IDENTIFIER ;

error_:
    error_list | error_ error_list

error_list:
    error ';' ;

case:
    case_item | case case_item

case_item:
    WHEN INT_LITERAL ARROW statement ';' ;

%%

void yyerror(const char* message)
{
	appendError(SYNTAX, message);
}

int main(int argc, char *argv[])
{
	firstLine();
	yyparse();
	lastLine();
	return 0;
}
