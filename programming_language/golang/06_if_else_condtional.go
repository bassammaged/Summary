package main

/*
	# ----- If, else, else if
		- how to write decision-making conditional statements used to perform different actions in Golang.
		- The if statement - executes some code if one condition is true
		- The if...else statement - executes some code if a condition is true and another code if that condition is false
		- The if...else if....else statement - executes different codes for more than two conditions
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IfCondtional() {
	fmt.Println("===== Condtional statekejts =====")
	fmt.Printf("[+] Enter your age: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	ageInput, _ := strconv.ParseInt(scanner.Text(), 10, 8)
	// ----- if, else if, else statements.
	if ageInput < 12 {
		fmt.Println("[+] INFO: You are so young.")
	} else if ageInput >= 12 && ageInput <= 18 {
		fmt.Println("[+] INFO: You are allowed to watch some movies.")
	} else {
		fmt.Println("[+] INFO: You own it.")
	}

	// ------ statement supports a composite (initialization) syntax, if  var declaration;  condition {}
	if x := 100; x == 100 {
		println("[+] INFO: x has been declared and compared with 100 in the same syntax.")
	}

}
