package main

import "fmt"

/*
	# ----- What is iota
		- Iota is a useful concept for creating incrementing constants in Go. However, there are several areas
		where iota may not be appropriate to use.
*/
func FuncIota() {
	// Basic Iota Usage

	// const (
	// 	Red int = iota
	// 	Orange
	// 	Yellow
	// 	Green
	// 	Blue
	// 	Indigo
	// 	Violet
	// )
	// fmt.Println("----- Basic iota -----")
	// fmt.Printf("The value of Red    is %v\n", Red)
	// fmt.Printf("The value of Orange is %v\n", Orange)
	// fmt.Printf("The value of Yellow is %v\n", Yellow)
	// fmt.Printf("The value of Green  is %v\n", Green)
	// fmt.Printf("The value of Blue   is %v\n", Blue)
	// fmt.Printf("The value of Indigo is %v\n", Indigo)
	// fmt.Printf("The value of Violet is %v\n", Violet)

	/*
		# -- Order matter
			- If we take the same code as above, but change the order of the constants, we'll see the value
			of the constants change as well.
	*/

	const (
		Blue int = iota
		Green
		Indigo
		Orange
		Red
		Violet
		Yellow
	)
	fmt.Println("----- Order Matter -----")
	fmt.Printf("The value of Red    is %v\n", Red)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Yellow is %v\n", Yellow)
	fmt.Printf("The value of Green  is %v\n", Green)
	fmt.Printf("The value of Blue   is %v\n", Blue)
	fmt.Printf("The value of Indigo is %v\n", Indigo)
	fmt.Printf("The value of Violet is %v\n", Violet)

	/*
		# -- Skipping Values
			- It may be necessary to skip a value. If so, you can use the _ (underscore) operator:
	*/
	const (
		_   int = iota // Skip the first value of 0
		Foo            // Foo = 1
		Bar            // Bar = 2
		_
		_
		Bin // Bin = 5
		// Using a comment or a blank line will not increment the iota value
		Baz // Baz = 6
	)
	fmt.Println("----- Skipping values -----")
	fmt.Printf("The value of Foo    is %v\n", Foo)
	fmt.Printf("The value of Bar    is %v\n", Bar)
	fmt.Printf("The value of Bin    is %v\n", Bin)
	fmt.Printf("The value of Baz    is %v\n", Baz)

	/*
		# -- Advanced iota
			- Because of the way iota automatically increments, you can use it to calculate more advanced scenarios. For
			instance, if you have worked with bitmasks , Iota can be used to quickly create the correct values by using the
			bit shift operator.
	*/

	const (
		Execute = 1 << iota
		Read
		Write
		// full permission for admin
		admin = Execute | Read | Write
	)
	fmt.Println("----- Advanced iota - bitwise -----")
	fmt.Printf("The value of Execute    is %v\n", Execute)
	fmt.Printf("The value of Read       is %v\n", Read)
	fmt.Printf("The value of Write      is %v\n", Write)
	fmt.Printf("The permission of admin is %v\n", admin)

}

// the resources: https://www.gopherguides.com/articles/how-to-use-iota-in-golang
