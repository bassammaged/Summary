package main

import (
	"fmt"
)

/*
	# ----- Empty Interface
		- The interface type that specifies zero methods is known as the empty interface:
			interface{}
		- An empty interface may hold values of any type. (Every type implements at least zero methods.)
		- Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any
		number of arguments of type interface{}.
*/

func anyWayPrint(i interface{}) {
	fmt.Printf("The value: %v, the type: %T\n", i, i)
}

func EmptyInterface() {
	v1 := 5
	v2 := "5"
	v3 := true

	fmt.Println("===== Start empty interface =====")
	anyWayPrint(v1)
	anyWayPrint(v2)
	anyWayPrint(v3)
}
