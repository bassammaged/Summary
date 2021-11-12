package main

import "fmt"

/*
	# ------ struct in golang
		- A struct (short for "structure") is a collection of data fields with declared data types. Golang has
		the ability to declare and create own data types by combining one or more types, including both built-in
		and user-defined types. Each data field in a struct is declared with a known type, which could be a built-in
		type or another user-defined type.

		- Structs are the only way to create concrete user-defined types in Golang. Struct types are declared by
		composing a fixed set of unique fields. Structs can improve modularity and allow to create and pass complex
		data structures around the system. You can also consider Structs as a template for creating a data record, like
		an employee record or an e-commerce product.

		# -- How to use struct in golang
			- The declaration starts with the keyword type, then a name for the new struct, and finally the
			keyword struct. Within the curly brackets, a series of data fields are specified with a name and a type.
				type identifier struct{
					field1 data_type
					field2 data_type
					field3 data_type
				}


*/

// Adress struct is a structure for detailed address information.
type Address struct {
	floor    int
	building int
	street   string
	district string
	city     string
	country  string
}

// Person struct is a structure for detailed information about the person.
type Person struct {
	firstName  string
	middleName string
	lastName   string
	age        uint
	address    Address
	title      string
}

// ----- Constructor for the struct
func NewPerson(firstName, middleName, lastName, title string, age uint, address Address) *Person {
	return &Person{
		firstName:  firstName,
		middleName: middleName,
		lastName:   lastName,
		title:      title,
		age:        age,
		address:    address,
	}

}

// ----- Create method for Person struct (GetData)
func (p Person) GetFullName() string {
	return p.firstName + " " + p.middleName + " " + p.lastName
}

// ----- Create method for Person struct (Change/add Data)
func (p *Person) SetFirstName(fName string) {
	p.firstName = fName
}

func Struct() {

	fmt.Println("===== Start struct =====")
	// ----- Create an object of Person struct
	var person1 Person = Person{
		"Bassam",
		"Maged",
		"Mohammed",
		25,
		Address{4, 35, "G73", "New Gersy", "Cairo", "UK"},
		"Cyber Security Engineer",
	}

	fmt.Println("[+] the values for 1st struct are: ", person1)
	fmt.Printf("[+] The firstname of 1st struct %v & the city name: %v\n", person1.firstName, person1.address.city)
	fmt.Printf("[+] The type of the struct: %T\n", person1)

	// ----- Declare struct by using constructor
	person2 := NewPerson("Kemet", "Bassam", "Maged", "CyberSec Engineer", 28, Address{})
	fmt.Println("[+] Using method of struct: ", person2.GetFullName())
	person2.SetFirstName("Qabdai")
	fmt.Println("[+] Using getData method of struct after change firstname by SetData: ", person2.GetFullName())

}
