package main

import "fmt"

/*
	# ---- How can iterate over the array?
		- The simplest answer is using range
			- range iterates over elements in a variety of data structures. Let’s see how to use range with some
			of the data structures we’ve already learned.
*/

func Range() {
	fmt.Println("==== Start range =====")
	arr_map := map[string]string{"id": "1415155", "name": "Kemet"}
	for kargs, v := range arr_map {
		fmt.Printf("key: %v => value: %v\n", kargs, v)
	}
}
