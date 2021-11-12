package main

import "fmt"

/*
	# ----- What is functions
		- A function is a group of statements that exist within a program for the purpose of performing a specific
		task. At a high level, a function takes an input and returns an output. Function allows you to extract
		commonly used block of code into a single component.

		# -- How to create a function
			- A declaration begins with the func keyword, followed by the name you want the function to have,
			a pair of parentheses (), and then a block containing the function's code.
				- Simple function:
					func SimpleFunction() {
						fmt.Println("Hello World")
					}
			- Simple function with parameters:
					func add(x int, y int) {
						total := 0
						total = x + y
						fmt.Println(total)
					}
			- Simple function with return value:
					func add(x int, y int) int {
						total := 0
						total = x + y
						return total
					}
			- Functions Returning Multiple Values
					fun rectangle(l int, w int) (area int, parameter int) {
						parameter 	= 2 * (l + w)
						area 		= l * w
						return
					}

					// Another example
					fun rectangle(l,w int) (int, int){
						parameter 	:= 2 * (l+w)
						area 		:= l*w
						return parameter,area
					}

*/
func parameter(l, w int) {
	fmt.Printf("Parameter of rectangle = length * width: %v\n", l*w)
}

func calcMultiReturn(l, w int) (int, int) {
	area := l * w
	parameter := 2 * (l + w)
	return area, parameter
}

func calcMultiReturn2(l, w int) (area, parameter int) {
	area = l * w
	parameter = 2 * (l + w)
	return
}

func Run() {
	fmt.Println("===== Functions =====")
	parameter(5, 3)
	area, parameter := calcMultiReturn(12, 11)
	fmt.Printf("The area of rect.: %v, the parameter of rect.: %v\n", area, parameter)
	area, _ = calcMultiReturn2(10, 6)
	fmt.Printf("The area of rect.: %v\n", area)

}
