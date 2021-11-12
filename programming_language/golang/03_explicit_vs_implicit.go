package main

import "fmt"

/*
	# -------------------------------- start Explicit vs Implicit -------------------------------- #
		// ------- Explicit
			Describes something that is very clear and without vagueness or ambiguity.

		// ------- Implicit
			Often functions as the opposite, referring to something that is understood, but not described clearly
			or directly
	# --------------------------------   end Explicit vs Implicit -------------------------------- #
*/

func ExplicitVsImplicit() {
	// ------- Implicit
	var Name = "Moscu"
	Age := 25

	// ------- Explicit
	var CSP string = "default 'self' *.test.com"

	fmt.Println("===== Explicit vs Implicit =====")
	fmt.Printf("Expilict: %v vs Implicit: %v, %v", CSP, Name, Age)
}
