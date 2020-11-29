/* Compiler Theory and Design
   Duane J. Jarc */

/* This file contains flex input file */

%{
#include <cstdio>
#include <string>

using namespace std;

#include "values.h"
#include "listing.h"
#include "tokens.h"

%}

%option noyywrap

ws		[ \t\r]+
comment		(\-\-.*|\/\/.*)\n
line		[\n]
id		 [a-zA-Z](_?([a-zA-Z]|[0-9]))*
digit		[0-9]
int		{digit}+
punc		[\(\),:;]

real    [0-9]+\.[0-9]*(?:[eE]{1}(?:[\+-])?[0-9]+)?
bool    [true|false]
%%

{ws}		{ ECHO; }
{comment}	{ ECHO; nextLine();}
{line}		{ ECHO; nextLine();}
"<"		{ ECHO; yylval.oper = LESS; return(RELOP); }
"+"		{ ECHO; yylval.oper = ADD; return(ADDOP); }
"*"		{ ECHO; yylval.oper = MULTIPLY; return(MULOP); }
begin		{ ECHO; return(BEGIN_); }
boolean		{ ECHO; return(BOOLEAN); }
end		{ ECHO; return(END); }
endreduce	{ ECHO; return(ENDREDUCE); }
function	{ ECHO; return(FUNCTION); }
integer		{ ECHO; return(INTEGER); }
is		{ ECHO; return(IS); }
reduce		{ ECHO; return(REDUCE); }
returns		{ ECHO; return(RETURNS); }
and		{ ECHO; return(ANDOP); }

"=>"   { ECHO; return(ARROW); }
case    { ECHO; return(CASE); }
else    { ECHO; return(ELSE); }
endcase { ECHO; return(ENDCASE); }
endif   { ECHO; return(ENDIF); }
if      { ECHO; return(IF); }
others  { ECHO; return(OTHERS); }
real    { ECHO; return(REAL); }
then    { ECHO; return(THEN); }
when    { ECHO; return(WHEN); }

or      { ECHO; return(OROP); }
not     { ECHO; return(NOTOP); }

"="     { ECHO; return(RELOP); }
"/="    { ECHO; return(RELOP); }
">"     { ECHO; return(RELOP); }
">="    { ECHO; return(RELOP); }
"<="    { ECHO; return(RELOP); }

"-"     { ECHO; return(ADDOP); }
"/"     { ECHO; return(MULOP); }

rem     { ECHO; return(REMOP); }
"**"    { ECHO; return(EXPOP); }

{real}  { ECHO; return(REAL_LITERAL); }
{bool}  { ECHO; return(BOOL_LITERAL); }

{id}		{ ECHO; yylval.iden = (CharPtr)malloc(yyleng + 1); 
			strcpy(yylval.iden, yytext); return(IDENTIFIER);}
{int}		{ ECHO; yylval.value = atoi(yytext); return(INT_LITERAL); }
{punc}		{ ECHO; return(yytext[0]); }
.		{ ECHO; appendError(LEXICAL, yytext); }

%%
