/*
# -------------------------------- start introduction -------------------------------- #
	- Workspace:
		In Go, programs are kept in a directory hierarchy that is called a workspace. A workspace is simply a root directory of your Go applications. A workspace
		contains three subdirectories at its root:
		- src – This directory contains source files organized as packages. You will write your Go applications inside this directory.
		- pkg – This directory contains Go package objects.
		- bin – This directory contains executable programs.

	You have to specify the location of the workspace before developing Go programs. The GOPATH environment variable is used for specifying
	the location of Go workspaces.
# --------------------------------   end introduction -------------------------------- #
*/

/*
# -------------------------------- start Packages -------------------------------- #
	In Go, source files are organized into system directories called packages, which enable code reusability across the Go applications.
	The naming convention for Go package is to use the name of the system directory where we are putting our Go source files. Within a single folder,
	the package name will be same for the all source files which belong to that directory.

	Go’s standard library comes with lot of useful packages which can be used for building real-world applications. For example, the standard
	library provides a “net/http” package which can be used for building web applications and web services.
# --------------------------------   end Packages -------------------------------- #
*/

/*
# -------------------------------- start Package Main -------------------------------- #
	When you build reusable pieces of code, you will develop a package as a shared library. But when you develop executable programs,
	you will use the package “main” for making the package as an executable program. The package “main” tells the Go compiler that the package
	should compile as an executable program instead of a shared library. The main function in the package “main” will be the entry point of
	our executable program. When you build shared libraries, you will not have any main package and main function in the package.
# --------------------------------   end Package Main -------------------------------- #
*/

// sepcify entry function to the complier
package main

/*
# -------------------------------- start Import Packages -------------------------------- #
	- The keyword “import” is used for importing a package into other packages. In the Code Listing -1, we have imported the package “fmt” into
	the sample program for using the function Println. The package “fmt” comes from the Go standard library. When you import packages,
	the Go compiler will look on the locations specified by the environment variable GOROOT and GOPATH. Packages from the standard library are
	available in the GOROOT location. The packages that are created by yourself, and third-party packages which you have imported, are
	available in the GOPATH location.

	- Install Third-Party Packages:
		We can download and install third-party Go packages by using “Go get” command. The Go get command will fetch the packages from the
		source repository and put the packages on the GOPATH location.

# --------------------------------   end Import Packages   -------------------------------- #
*/

// import modules/libraries
import "fmt"

// main function
func MainFunction() {
	fmt.Println("test")
}
