package main

import "fmt"

func Variables() {
	/*
		# -------------------------------- start variables -------------------------------- #
			// ------- Unsigned integer types (Not negative)
				- uint8
				- uint16
				- uint32
				- uint64

			// ------- integer types
				- int8
				- int16
				- int32
				- int64

			// ------- machine dependent Types
				- unit (32 or 64 bits based on arch.)
				- int (same size of unit)
				uintptr (as unsigned integer to stire the uniterpreted bits of a pointer value)

			// ------- float types
				- float32
				- float64

			// ------- complex types
				- complex64 	(complex numbers with float32 real and imaginary parts)
				- complex128	(complex numbers with float64 real and imaginary parts)

			// ------- strings types
				- "Hello Word"

			// ------- Booleans types
				- true
				- false

			// ------- Array type
				- identify array
				- array slice
					- slices
					- capacity of slice
				- Mapping array (equal to Dict in python)

		# --------------------------------   end variables -------------------------------- #

	*/

	// ----------- Declaring variable with var
	// declare variable
	var Number int
	var Name string

	// assign variable
	Number = 10
	Name = "Kemet"

	// print the variable
	fmt.Printf("Number: %v, Name: %v\n", Number, Name)

	// ----------- Declaring with initializing
	var Number2 float64 = 20
	var Status bool = false
	// print the variable
	fmt.Printf("Number: %v, Status: %v\n", Number2, Status)

	// ----------- Declaring with no type (Omit type)
	var Number3 = 40
	var Status2 = true
	fmt.Printf("Number: %v, Status: %v\n", Number3, Status2)

	// ----------- Declaring with no type (Omit type) Walrus operator
	Number6 := 50
	Status3, Status4 := true, false
	fmt.Printf("Number6: %v, Status3: %v, Status4: %v\n", Number6, Status3, Status4)

	// ----------- Declare multiple variables
	var Number4, Number5 int = 50, 60
	Name1, Name2 := "Moscu", "Moscu2"
	fmt.Printf("Number4: %v, Number5: %v\n", Number4, Number5)
	fmt.Printf("Name1: %v, Name2: %v\n", Name1, Name2)
}
