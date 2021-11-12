package main

/*
	# ----- Error handling in go
		- Golang has an unconventional approach and that’s the reason why people like to talk about Golang. Golang
		has its way of Go Error handling, while other languages handle the error by try... catch block. As Golang has
		its unique way of handling errors, it is challenging for developers to get all information about Go’s Error
		handling method.

		# -- What are Error and Error handling?
			- Error: Those unusual and unwanted conditions that occur in the program are known as errors. The error
			can exist at the time of compile or run time. Some examples of errors are - failed db connection,
			invalid user inputs, or file does not exist.
			- Now, if your program behaves abnormally, you try to implement the solution and predict where things
			 go wrong - it’s known as Error Handling. Many languages like Java, Python, or PHP use try...catch block
			 for Error Handling.

		# -- Golang Error Handling Patterns
			- Dealing with new technology or language has always been difficult; it doesn’t matter how easy and
			straightforward it is. And, when you get frustrated while learning about new things, you start
			criticizing. That’s exactly what Go has to face. It’s challenging for developers to understand Go’s way
			of error handling. Go error handling has so much to learn about, but first, let’s look at the built-in
			error type of Golang.

			- In Go’s, the type of interface is known as error type. It declares itself as a string.
			Its looks like this:
				type error interface {
					Error() string
				}
			- Golang doesn’t overlook the errors, and they take them very important. If you just start learning about Go,
			The syntax of func f() (value, error) is very easy to understand and smooth to implement.
			- Golang uses errors as first-class values of the functions. If you missed returning the error string
			from your function like this way:
				func getUsers() (*Users, error) { .... }
					func main() {
					users, _ := getUsers()
				}

			Let’s have one example of Error handling practice in Go.

				if error := criticalOperation(); error != nil {
				// Not returning anything is a bad practice.
				log.Printf("Oops! Something went wrong in the program", error)
				// return your error message hereafter this line!
				}
				if error := saveData(data); error != nil {
					return fmt.Errorf("Data has been lost", error)
				}

			# -- Why didn’t Golang utilize exceptions, a conventional way to handle errors?
				Keep these two hey points in mind while Golang error handling is:
					- Keep it simple.
					- Plan where it can go wrong.
				# -- the advantages of Golang new error handling.
						No interruption of sudden uncaught exceptions.
						Simple syntax.
						You have complete control over the errors
						Transparent control-flow.
						Easy implementation of error chains to take action on the error.


	# ----- Resources:
				-https://dev.to/amelias26018837/prepare-yourself-to-deal-with-golang-error-handling-3jj5
				-
*/

import (
	"errors"
	"fmt"
	"os"
)

// Create a function to return one specific error
func oneException() error {
	return errors.New("CustomError: There is an error")
}

// Create a function to return a dynamic error
func oneExceptionCustom(err string) error {
	return errors.New(err)
}

// an example 1 of error handling
func openFile(fullPath string) error {
	// Check the filename
	if fullPath == "" {
		return errors.New("ERROR: The filename is empty")
	}

	// Open a file
	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("ERROR: Unable to open the file %v", err)
	}
	// Check the filename expectation
	if file.Name() != "./02_variables.go" {
		return fmt.Errorf("ERROR: The filename is not whitelisted %v", file.Name())
	}

	return nil
}

func ErrorHandling() {
	var v1 bool = false
	if v1 {
		fmt.Println("[+] Okay")
	} else {
		err := oneException()
		fmt.Println("[+] Raise an error:", err)
	}

	if v1 {
		fmt.Println("[+] Okay")
	} else {
		err := oneExceptionCustom("CustomERROR: written by me")
		fmt.Println("[+] Raise an error:", err)
	}

	// Use an example 1
	err := openFile("./02_variables.go")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		print("Okay")
	}

}
