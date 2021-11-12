package main

import (
	"fmt"
)

/*
	# -------------------------------- start relational operators -------------------------------- #
		==> Quick overview: The following table lists all the relational operators supported by Go language.
							Assume variable A holds 10 and variable B holds 20, then
							-------------------------------------------------------------------------------------
							|     Operator		|			Description					|		Example			|
							-------------------------------------------------------------------------------------
							|       &&			| Called Logical AND operator. If 		|					  	|
							| 					| both the operands are false,  		|  	  					|
							| 					| then the condition becomes false.  	|  (A && B) is false.	|
							-------------------------------------------------------------------------------------
							|       ||			| Called Logical OR Operator. 			|					  	|
							| 					| If any of the two operands is true,  	|   					|
							| 					| then the condition becomes true.   	|  (A || B) is false.	|
							-----------------------------------------------------------------------------------
							|       !			| Called Logical NOT Operator. Use to  	|					  	|
							| 					| reverses the logical state of its.  	|   					|
							| 					| operand. If a condition is true, then |   				  	|
							| 					| Logical NOT operator makes it false.	|   !(A && B) is true.	|
							-----------------------------------------------------------------------------------
	# --------------------------------   end relational operators -------------------------------- #
*/

func LogicalOperator() {
	fmt.Println("You can test it with Conditional statement for example.")
}
