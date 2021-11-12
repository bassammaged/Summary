package main

/*
	# ----- What is a package
		- “Don't Repeat Yourself.”, packages is one of the mechanism for code reuse: we've seen import "fmt".
			- fmt is the name of a package that includes a variety of functions related to formatting and output to the screen.

	# ----- Type of packages
		1. pre-defined packages.
		2. third-party package.
		3. custom package.

		# -- pre-defined package:
			import ("fmt" "math")

		# -- custom package
			- 1st: configure the go env or using go mod
			- 1st: go mod init <path_of_package>
				go mod init github.com/bassammaged/learn-go
			- 2nd: import package and function
				import "github.com/bassammaged/learn-go/string_util"
			- 3rd: import the public variables or public functions
				string_util.Reverse()
				string.util.Name

			# --import the module of the same package
				- 1st: call the function of the other module directly.
				- 2nd: while compling provide the module file along with the name of the script you are going to run.
					- ex: go run main.go welcome.go #OR
					- ex: go run .

	# ----- Private vs Public accessibility/visability
		- Private var/func has to start with lowercase letter such as reverseIt() or nameOfCustomer
		- Public var/func that accessable by other package has to start with uppercase letter such as ReverseIt()
*/

import (
	"fmt"
	"math"

	"github.com/bassammaged/learn-go/string_util"
)

func Packages() {
	fmt.Println("======= Pre-defined package =======")
	fmt.Println("[+] The math.Floor: ", math.Floor(1.7), "The math.Ceil: ", math.Ceil(1.7))
	fmt.Println("======= Custom Package =======")
	fmt.Println("[+] Reverse it ", string_util.Reverse("Reverse me!"))
}
