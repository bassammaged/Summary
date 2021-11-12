package main

import "fmt"

/*
	# ----- What is anonymous function in golang?
		- Go language provides a special feature known as an anonymous function. An anonymous function is a function which doesn’t contain any name. It is useful
		 when you want to create an inline function. In Go language, an anonymous function can form a closure. An anonymous function is also known as function literal.

		 - In Go language, you are allowed to assign an anonymous function to a variable. When you assign a function to a variable, then the type of the variable
		  is of function type and you can call that variable like a function call.

*/

// Declare Anonymous function with a little of dynamic
var whoAreYou func() = func() {
	// Nothing Here
}

func AnonymousFunction() {
	fmt.Println("[======= Start anonymous function =======")

	// ----- Declare anonymous function inline as a variable
	// Declare the function
	var doStuff func() = func() {
		fmt.Println("[+] Doing Stuff ....")
	}
	// Calling the function
	doStuff()

	// ----- Using the dynamic anonymous function
	whoAreYou = func() {
		fmt.Println("[+] What's the hell? Who are you?")
	}
	whoAreYou()

	// Make it with another body
	whoAreYou = func() {
		fmt.Println("[+] Get your hand away. Who are you?")
	}
	whoAreYou()

}
