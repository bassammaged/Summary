package main

import (
	"fmt"
)

var foo int = 314

func recap() {
	foo := 3.14
	// ----- Variables
	var name string = "Marta"
	age := 26
	var address = "Madinety"

	var oneChar byte = 'b'
	var oneChar2 byte = '~'

	fmt.Println(name, age)
	fmt.Printf("The type of address %T\n", address)
	fmt.Printf("The value of foo: %v\n", foo)
	fmt.Printf("The %c:%U & the %c:%U\n", oneChar, oneChar, oneChar2, oneChar2)

	// var b1 byte = 'a'
	var b2 []byte
	b2 = append(b2, 'a')
	b2 = append(b2, 'b')
	PrintSlince(b2)

}

func PrintSlince(s []byte) {
	fmt.Printf("len:%d, cap:%d, val:%c\n", len(s), cap(s), s[0])
}
