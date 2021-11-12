package main

import "fmt"

/*
	# ----- Slices in golang:
		- Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
		- Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
		To create an empty slice with non-zero length, use the builtin make.

		# -- How to use slices
			- 1st step: declare slice by using make func, Here we make a slice of strings of length 3 (initially zero-valued).
				var <slices_name> []<slices_type> = make(<slices_type>,<slices_length>)
			- 2nd step: We can set and get just like with arrays.
				<slices_name>[<index>] = <value>

			# -- slices is much more rich than array
				- In addition to these basic operations, slices support several more that make them richer than
				arrays. One is the builtin append, which returns a slice containing one or more new values. Note
				that we need to accept a return value from append as we may get a new slice value.
					var s []string = make([]string,3)
					s = append(s, "d")
    				s = append(s, "e", "f")

			# -- Declare and assign value for slices
				t := []string{"g", "h", "i"}

*/

// Function for chage the default null values of slices
func changeValues(s []string) []string {
	s[0] = "Bassam"
	s[1] = "Maged"
	s[2] = "Mohammed"
	return s
}

func SlicesFunc() {
	fmt.Println("===== Start Slices Func =====")
	names := make([]string, 3)
	names = changeValues(names)
	// Append new value for slices
	names = append(names, "Kemet")
	fmt.Println("[+] The values of slices: ", names)

}
