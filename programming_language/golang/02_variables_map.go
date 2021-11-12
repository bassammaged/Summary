package main

import "fmt"

/*
	# ----- What is map in golang>
		- There's no specific data type in Golang called map; instead, we use the map keyword to create a map
		with keys of a certain type, and values of another type (or the same type).
		- Ex:
			var menu map[string]float64
			- In this example, we declare a map that has strings for its keys, and float64s for its values. The
			result is a variable of a very specific type: map[string]float64.

		# -- how to use map
			- 1st way:
				var <map_var_name> map[<key_var_type>]<value_var_type> 	= make(map[<key_var_type>]<value_var_type>)
			- 2nd way:
				<map_var_name> := map[<key_var_type>]<value_var_type>{<key>:<value>,<key>:<value>,....}
*/

func VariablesMaps() {

	fmt.Println("=====  Start map =====")
	// ----- Delcare a map
	var map1 map[string]string = make(map[string]string)
	// assign a value
	map1["name"] = "Kemet"
	fmt.Printf("[+] The value of key name is %v\n", map1["name"])

}
