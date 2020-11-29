// CMSC 430
// Duane J. Jarc

// This file contains the bodies of the evaluation functions

#include <string>
#include <vector>
#include <cmath>

using namespace std;

#include "values.h"
#include "listing.h"

int evaluateReduction(Operators operator_, int head, int tail)
{
	if (operator_ == ADD)
		return head + tail;
	return head * tail;
}


int evaluateRelational(int left, Operators operator_, int right)
{
	int result;
	switch (operator_)
	{
		case LESS:
			result = left < right;
			break;
	}
	return result;
}

double evaluateArithmetic(double left, Operators operator_, double right)
{
	int result;
	switch (operator_)
	{
		case ADD:
			result = left + right;
			break;
		case MULTIPLY:
			result = left * right;
			break;
	}
	return result;
}

