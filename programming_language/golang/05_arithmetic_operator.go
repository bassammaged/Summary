package main

import (
	"fmt"
)

/*
	# -------------------------------- start arithmetic operators -------------------------------- #
		==> Quick overview: Following table shows all the arithmetic operators supported by Go language.
							Assume variable A holds 10 and variable B holds 20, then:
							------------------------------------------------------------------------------------
							|     Operator		|			Description					|		Example			|
							------------------------------------------------------------------------------------
							|       +			|			Adds two operands			|   A + B gives 30	   |
							-----------------------------------------------------------------------------------
							|       -			| Subtracts 2nd operand from the 1st    |   A + B gives 30	  |
							-----------------------------------------------------------------------------------
							|       *			| 		Multiplies both operands	    |   A * B gives 200	  |
							-----------------------------------------------------------------------------------
							|       /			| Dividesnumerator by the denominator.  |   A + B gives 30	  |
							-----------------------------------------------------------------------------------
							|       %			| Modulus operator; gives the 			|					  |
							| 					| remainder after an integer division.  |   A + B gives 30	  |
							-----------------------------------------------------------------------------------
							|       ++			| Increment operator. It increases 		|					  |
							|					|      the integer value by one.   		|   A + B gives 30	  |
							-----------------------------------------------------------------------------------
							|       --			| Decrement operator. It decreases 		|					  |
							|					|      the integer value by one   		|   A + B gives 30	  |
							-----------------------------------------------------------------------------------
	# --------------------------------   end arithmetic operators -------------------------------- #
*/

func ArithmeticOperator() {
	fmt.Println("Addition: 5+5=", 5+5)
}
