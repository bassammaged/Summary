package main

import "fmt"

/*
	# ----- What is interface
		- Before we dive into the code, we need to first take a look at what an interface is and what it provides for us.
		Well, I like to think of interfaces as contracts that my code must follow. They allow us to effectively define
		what our dependencies should look like.
		- For example, if I’m implementing a comment service that talks to a database, I would want to implement
		 an interface within that comment service that effectively describes the methods I’ll need any database
		  implementation to have.

		# -- Interface in golang
			1st step: Declare interface
				type GetData interface {
					GetName() string
				}

				type Engineer struct {
					// block of attributes
				}
				type Emp struct {
					// block of attributes
				}
			2nd step: Create a method with same specific of interface method
				func (e Emp) GetData string{
					// block of code
				}
				func (e *Engineer) GetData string{
					// block of code
				}
			3rd step: Create a method recieve the interface as parameter
				func getdata(gd GetData) {
					gd.GetName()
				}
			4th step: Using it
				e 	:= Emp{""}
				en 	:= Engineer{""}
				getdata(&e)
				getdata(&en)
*/

// Person attributes
type person struct {
	firstName, lastName string
	age                 uint8
	department          string
	salary              float64
}

type employee struct {
	basicInfo person
}

type manager struct {
	basicInfo person
	carType   string
}

type commonThings interface {
	salary()
}

func Newmanager(fn, ln, depart string, age uint8, cT string, salary float64) *manager {
	return &manager{
		basicInfo: person{fn, ln, age, depart, salary},
		carType:   cT,
	}
}

func Newemployee(fn, ln, depart string, age uint8, salary float64) *employee {
	return &employee{
		basicInfo: person{fn, ln, age, depart, salary},
	}
}

func (m *manager) salary() {
	fmt.Printf("[+] The salary of %v manager is %v USD \n", m.basicInfo.firstName, m.basicInfo.salary)
}
func (e *employee) salary() {
	fmt.Printf("[+] The salary of %v employee is %v EGP \n", e.basicInfo.firstName, e.basicInfo.salary)
}

func salary(comm commonThings) {
	comm.salary()
}

func Interface() {
	var emp employee = *Newemployee("Bassam", "Maged", "CS", 26, 1000)
	var man manager = *Newmanager("Kemet", "TEMEK", "QA", 26, "BMW", 2100)
	salary(&emp)
	salary(&man)
}
