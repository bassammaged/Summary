package main

import "fmt"

/*
	# ---- Closure function
		- Go functions may be closures. A closure is a function value that references variables from outside its
		body. The function may access and assign to the referenced variables; in this sense the function is "bound" to
		the variables.

		- A closure is a special type of anonymous function that references variables declared outside of the function itself. Wwe will be creating a function dynamically,
		but in this case we will be using variables that weren’t passed into the function as a parameter, but instead were available when the function was declared.
		- This is very similar to how a regular function can reference global variables. You aren’t directly passing these variables into the function as a parameter,
		but the function has access to them when it is called.

			# -- Closures provide data isolation
				package main
				import "fmt"
				func main() {
					n := 0
					counter := func() int {
						n += 1
						return n
					}
					fmt.Println(counter())
					fmt.Println(counter())
				}

				- One problem with the previous example is a problem that can also pop up when using global variables. Any code inside of the main() function has
				access to n, so it is possible to increment the counter without actually calling counter(). That isn’t what we want; Instead we would rather have n isolated
				so that no other code has access to it.
				To do this we need to look at another interesting aspect of closures, which is the fact that they can still reference variables that they had access to
				during creation even if those variables are no longer referenced elsewhere.

*/

// Normal function include anonymously function call itself.
func inSeq() func() int {
	// It just calls for the first time
	fmt.Println("inSeq started...")
	var n int = 0
	// the func() calls everytime and affect the value of n
	return func() int {
		fmt.Println("Anonymously started...")
		n++
		return n
	}
}

func ClosureFunc() {
	fmt.Println("==== Start Closure =====")
	func1 := inSeq()
	fmt.Println(func1())
	fmt.Println(func1())
}
