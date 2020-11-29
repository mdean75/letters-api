// Michael DeAngelo
// CMSC 430 - project 2
// November 17, 2020
// Updated to project 2 specifications from skeleton code provided

/* Compiler Theory and Design
   Dr. Duane J. Jarc */



%{

#include <string>

using namespace std;

#include "values.h"
#include "listing.h"
#include "symbols.h"

int yylex();
void yyerror(const char* message);

Symbols<int> symbols;

int result;

%}

%define parse.error verbose

%union
{
	CharPtr iden;
	Operators oper;
	int value;
}

%token <iden> IDENTIFIER
%token <value> INT_LITERAL

%token <oper> ADDOP MULOP RELOP
%token ANDOP

%token BEGIN_ BOOLEAN END ENDREDUCE FUNCTION INTEGER IS REDUCE RETURNS

%token ARROW CASE ELSE ENDCASE ENDIF IF OTHERS REAL THEN WHEN OROP NOTOP REMOP EXPOP REAL_LITERAL BOOL_LITERAL

%type <value> body statement_ statement reductions expression relation term
	factor primary
%type <oper> operator

%%
function:
	function_header var1ormore body {result = $3;} ;

function_header:
	FUNCTION IDENTIFIER parameters RETURNS type ';' | error_ ;

var1ormore:
    var1ormore variable | ;

parameters:
    parameter | parameters ',' parameter | ;

parameter:
    IDENTIFIER ':' type

variable:
	IDENTIFIER ':' type IS statement_ {symbols.insert($1, $5);} ;

type:
	INTEGER |
	REAL |
	BOOLEAN ;

body:
	BEGIN_ statement_ END ';' {$$ = $2;} ;

statement_:
	statement ';' |
	IF expression THEN statement ';' ELSE statement ';' ENDIF |
	CASE expression IS case OTHERS ARROW statement ';' ENDCASE |
	error ';' {$$ = 0;} ;

statement:
	expression |
	statement_ |
	REDUCE operator reductions ENDREDUCE {$$ = $3;} ;

operator:
    NOTOP |
    EXPOP |
	ADDOP |
	MULOP |
	REMOP |
	OROP ;

reductions:
	reductions statement_ {$$ = evaluateReduction($<oper>0, $1, $2);} |
	{$$ = $<oper>0 == ADD ? 0 : 1;} ;

expression:
	expression ANDOP relation {$$ = $1 && $3;} |
	expression OROP relation {$$ = $1 && $3;} |
	relation ;

relation:
	relation RELOP term {$$ = evaluateRelational($1, $2, $3);} |
	term;

term:
	term ADDOP factor {$$ = evaluateArithmetic($1, $2, $3);} |
	factor ;

factor:
    factor EXPOP primary {$$ = evaluateArithmetic($1, $2, $3);} |
	factor MULOP primary {$$ = evaluateArithmetic($1, $2, $3);} |
	factor REMOP primary {$$ = evaluateArithmetic($1, $2, $3);} |
	primary ;

primary:
	'(' expression ')' {$$ = $2;} |
	NOTOP factor |
	INT_LITERAL | REAL_LITERAL | BOOL_LITERAL |
	IDENTIFIER {if (!symbols.find($1, $$)) appendError(UNDECLARED, $1);} ;

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
	if (lastLine() == 0)
    	cout << "Result = " << result << endl;
    return 0;
}
