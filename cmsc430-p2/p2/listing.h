// Michael DeAngelo
// CMSC 430 - project 2
// November 17, 2020
// Updated to project 2 specifications from skeleton code provided

// CMSC 430
// Dr. Duane J. Jarc

// This file contains the function prototypes for the functions that produce the // compilation listing

enum ErrorCategories {LEXICAL, SYNTAX, GENERAL_SEMANTIC, DUPLICATE_IDENTIFIER,
	UNDECLARED};

void firstLine();
void nextLine();
int lastLine();
void appendError(ErrorCategories errorCategory, string message);

