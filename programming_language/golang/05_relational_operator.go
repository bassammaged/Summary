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
							|       ==			| It checks if the values of two 		|					  	|
							| 					| operands are equal or not; if yes,  	|  	  					|
							| 					| the condition becomes true.  			| (A == B) is not true. |
							-------------------------------------------------------------------------------------
							|       !=			| It checks if the values of two 		|					  	|
							| 					| operands are equal or not; if the  	|   					|
							| 					| values are not equal, then the   		|   				  	|
							| 					| condition becomes true.   			|   (A != B) is true.	|
							-----------------------------------------------------------------------------------
							|       >			| It checks if the value of left  		|					  	|
							| 					| operand is greater than the value  	|   					|
							| 					| of right operand; if yes, the   		|   				  	|
							| 					| condition becomes true.   			|   (A > B) is true.	|
							-----------------------------------------------------------------------------------
							|       <			| It checks if the value of left  		|					  	|
							| 					| operand is less than the value  		|   					|
							| 					| of right operand; if yes, the   		|   				  	|
							| 					| condition becomes true.   			|   (A < B) is true.	|
							-----------------------------------------------------------------------------------
							|       >=			| It checks if the value of left  		|					  	|
							| 					|  operand is greater than or equal to  |   					|
							| 					| the value of the right operand; if    |   				  	|
							| 					| yes, the condition becomes true.   	|   (A >= B) is true.	|
							-----------------------------------------------------------------------------------
							|       <=			| It checks if the value of left  		|					  	|
							| 					|  operand is less than or equal to  	|   					|
							| 					| the value of the right operand; if    |   				  	|
							| 					| yes, the condition becomes true.   	|   (A <= B) is true.	|
							-----------------------------------------------------------------------------------

	# --------------------------------   end relational operators -------------------------------- #
*/

func RelationalOperater() {
	var num1 int32 = 109
	num2 := 100.5
	fmt.Printf("===== Relational Operation =====")
	fmt.Println("[+] The result of <= relational operator: ", float64(num1) <= num2)
}
