// Michael DeAngelo
// CMSC 430 - project 2
// November 17, 2020
// Updated to project 2 specifications from skeleton code provided

// Compiler Theory and Design
// Dr. Duane J. Jarc

// This file contains the bodies of the functions that produces the compilation
// listing

#include <cstdio>
#include <string>
#include <queue>

using namespace std;

#include "listing.h"

static int lineNumber;
static string error = "";
static int totalErrors = 0;
int lexicalErrors = 0;
int syntaxErrors = 0;
int semanticErrors = 0;
queue <string> errorQueue;

static void displayErrors();

void firstLine()
{
    lineNumber = 1;
    printf("\n%4d  ",lineNumber);
}

void nextLine()
{
    displayErrors();
    lineNumber++;
    printf("%4d  ",lineNumber);
}

int lastLine()
{
    printf("\r");
    displayErrors();
    if (totalErrors != 0) {
        printf("Lexical Errors: %d\n", lexicalErrors);
        printf("Syntax Errors: %d\n", syntaxErrors);
        printf("Semantic Errors: %d\n", semanticErrors);
    } else {
        printf("Compiled Successfully\n");
    }
    printf("     \n");
    return totalErrors;
}

void appendError(ErrorCategories errorCategory, string message)
{
    string messages[] = { "Lexical Error, Invalid Character ", "",
                          "Semantic Error, ", "Semantic Error, Duplicate Identifier: ",
                          "Semantic Error, Undeclared " };

    switch (errorCategory) {
        case LEXICAL:
            lexicalErrors++;
            break;
        case SYNTAX:
            syntaxErrors++;
            break;
        case GENERAL_SEMANTIC:
            semanticErrors++;
            break;

    }
    errorQueue.push(messages[errorCategory] + message);
    totalErrors++;
}

void displayErrors()
{
    while (!errorQueue.empty()) {
        printf("%s\n", errorQueue.front().c_str());
        errorQueue.pop();
    }
}
