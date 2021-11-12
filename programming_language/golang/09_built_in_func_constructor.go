package main

/*
	# ---- What is the constructor function
		- There are times, when creating applications in Go, that you need to be able to set up some form of state
		on the initial startup of your program. This could involve creating connections to databases, or loading in
		configuration from locally stored configuration files. When it comes to doing this in Go, this is where
		your init() functions come into play.

		# -- The init Function
			- In Go, the init() function is incredibly powerful and compared to some other languages, is a lot
			easier to use within your Go programs. These init() functions can be used within a package block and
			regardless of how many times that package is imported, the init() function will only be called once.
*/

import "fmt"

func Constructor() {
	fmt.Println("===== Constructor ======")
	fmt.Println("Hello From Constructor module")
}

// func init() {
// 	fmt.Println("[+] This will get called on initialization")
// }
