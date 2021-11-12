package main

/*
	# ----- switch statement
	 - Golang also supports a switch statement similar to that found in other languages such as, Php or Java. Switch
	 statements are an alternative way to express lengthy if else comparisons into more readable code based on the
	 state of a variable.

		# -- Basic example:
			 i := 2
			fmt.Print("Write ", i, " as ")
			switch i {
			case 1:
				fmt.Println("one")
			case 2:
				fmt.Println("two")
			case 3:
				fmt.Println("three")
			}
		# -- Comma separate multiple experssions example:
			  switch time.Now().Weekday() {
				case time.Saturday, time.Sunday:
					fmt.Println("It's the weekend")
				default:
					fmt.Println("It's a weekday")
				}

*/
import (
	"bufio"
	"fmt"
	"os"
)

func SwitchCondtional() {
	fmt.Println("===== Start Switch Conditional")
	fmt.Printf("[+] Please Insert your Secret word: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	secretWord := scanner.Text()

	switch secretWord {
	case "Eh ya rays":
		fmt.Println("[+] Result: Welcome, Kemet.")
	case "He5a boom":
		fmt.Println("[+] Result: Welcome, Bassam.")
	default:
		fmt.Println("Access Denied.")
	}

}
