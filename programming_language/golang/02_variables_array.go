package main

import "fmt"

/*
	# ----- Arrays in golang
		# -- 1st way
			- Declare the array
				var <array_variable_name> [<length_array>]<array_type>
				var names [6]string
			- add values to the array
				names[0] = "Bassam"
				names[1] = "Maged"
				...

		# -- 2nd way
			- Declare and assign value for the array
				var <arr_var_name> [<arr_length>]<arr_var_type>{<value>,<value>,....}
				var names [5]string{"Bassam","Maged",....}

		# -- 3rd way
			- Declare array with no specific number of index.
				var <arr_var_name> []<arr_var_type>{<value>,<value>,...}

		# -- 4th way (array multidimensional)
			var <arr_var_name> [<arr_length>][<arr_length>]<arr_var_type>
*/

func VariablesArray() {
	// ----- Declare the array with specific index
	var names = [2]string{"First", "Second"}
	fmt.Printf("[+] The names array:%v\n", names)

	// ----- Declare the array with no index identification
	var ages = []int{1, 2, 3, 4}
	fmt.Printf("[+] The ages array:%v\n", ages)

	// ----- Declare array multidimensional
	var multiDimensionalArray [5][2]int
	for i := 0; i < len(multiDimensionalArray); i++ {
		for j := 0; j < len(multiDimensionalArray[i]); j++ {
			multiDimensionalArray[i][j] = i + j
		}
	}
	fmt.Println("[+] The value of multiDimensionalArray:", multiDimensionalArray)

}
