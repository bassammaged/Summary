package main

import "fmt"

/*
	# -------------------------------- start constant -------------------------------- #
		==> Quickview: A constant is a name or an identifier for a fixed value. The value of
					   a variable can vary, but the value of a constant must remain constant.

						The keyword const is used for declaring constants followed by the desired name and the type of value the
						constant will hold. You must assign a value at the time of the constant declaration, you can't assign a value later
						as with variables.
	# --------------------------------   end constant -------------------------------- #
*/

func VariablesConst() {

	// ---------- Declare const variable
	const product string = "TV" // Explicit
	const brand = "Samsung"     // Implicit

	fmt.Printf("product: %v, brand: %v\n", product, brand)

	// ---------- Declare Multiple Cosntants
	const (
		Product  = "TV"
		Brand    = "Samsung"
		Quantity = 200
		Price    = 1500.5
		Stock    = true
	)
	fmt.Println("===== Const variables =====")
	fmt.Printf("Product: %v, Brand: %v, Quantity: %v, Price: %v, Stock: %v \n", Product, Brand, Quantity, Price, Stock)

}
