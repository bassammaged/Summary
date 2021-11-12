package main

import "fmt"

/*
        Bitwise operator works on bits and performs bit by bit operation. Assume if a = 60; and b = 13; Now in binary
		format they will be as follows:
            a = 0011 1100
            b = 0000 1101

            a&b 	= 0000 1100			# Binary and operator
            a|b 	= 0011 1101			# Binary or operator
            a^b 	= 0011 0001			# Binary XOR operator
			~a  	= 1100 0011			# Binary not operator
			a << 2	= 1111 0000			# Binary Left shift operator
			a >> 2	= 0000 1111			# Binary right shift operator
*/

func BitwiseOperator() {
	var a int = 60
	b := 13

	fmt.Println("===== Bitwise Operator =====")
	fmt.Printf("[+] The result of a|b: %v, where a=%v and b=%v\n", a|b, a, b)
}
