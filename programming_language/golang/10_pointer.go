package main

import (
	"fmt"
	f "fmt"
)

/*
	# ----- What is a pointer?
		- Pointers in Go programming language or Golang is a variable that is used to store the memory
		address of another variable. Pointers in Golang is also termed as the special variables. The variables
		are used to store some data at a particular memory address in the system. The memory address is always
		found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

	# ----- What is the need for the pointers?
		- To understand this need, first, we have to understand the concept of variables. Variables are the names given
		to a memory location where the actual data is stored. To access the stored data we need the address of that
		particular memory location. To remember all the memory addresses(Hexadecimal Format) manually is an overhead
		that’s why we use variables to store data and variables can be accessed just by using their name.
		Golang also allows saving a hexadecimal number into a variable using the literal expression i.e. number
		starting from 0x is a hexadecimal number.

	# ----- Pointer
		- A pointer is a special kind of variable that is not only used to store the memory addresses of
		other variables but also points where the memory is located and provides ways to find out the value stored
		at that memory location. It is generally termed as a Special kind of Variable because it is almost declared
		as a variable but with *(dereferencing operator).

		# -- Declaration and Initialization of Pointers
			- Before we start there are two important operators which we will use in pointers i.e.
				- * Operator also termed as the dereferencing operator used to declare pointer variable and
				access the value stored in the address.
				- & operator termed as address operator used to returns the address of a variable or to
				access the address of a variable to a pointer.

*/

func Pointer() {
	f.Println("===== Start pointer =====")
	var x int = 0xff // Assign hexdecimal value at x

	// return the memory location
	fmt.Printf("[+] The memory location: %v\n", &x)

	// assign the value of memory location of x
	var y *int = &x
	fmt.Printf("[+] The value of memory location of x assigned to y: %v\n", y)

}
