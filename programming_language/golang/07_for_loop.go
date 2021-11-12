package main

import "fmt"

func ForLoop() {
	// ----- we have different way to define for loop
	// 1st way
	var x int8 = 1
	for x <= 5 {
		fmt.Println("1st: Here is paraphrase:", x)
		x++
	}

	// 2nd way
	for z := 1; z <= 5; z++ {
		fmt.Println("2nd: Here is paraphrase:", z)
	}

	// ----- interation over array
	arr1 := [6]string{"Bassam", "Fadel", "Junior", "Maze", "Waleed", "Hamada"}
	for i := 0; i < len(arr1); i++ {
		fmt.Println("Iteration over Array: ", arr1[i])
	}

	/*
		# -------------------------------- start advanced control structure -------------------------------- #
			- break keyword
			- continue keyword
		# --------------------------------   end advanced control structure -------------------------------- #
	*/
}
