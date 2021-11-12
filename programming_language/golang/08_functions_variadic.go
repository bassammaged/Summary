package main

import "fmt"

/*
	# ----- Variadic functions
		- can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
*/
// functions recieve integer arguments as a range
func withNoLimitArgument(nums ...int) {
	for _, v := range nums {
		print(v)
	}
}

// practice on empty interface
func print(v interface{}) {
	fmt.Println("Value:", v)
}

func VariadicFunc() {
	fmt.Println("===== Start Variadic =====")
	withNoLimitArgument(1, 2, 3, 4, 5, 6, 7, 8)
}
