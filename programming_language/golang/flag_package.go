package main

/*
	# ----- flag package
		- Command-line flags are a common way to specify options for command-line programs.
		For example, in wc -l the -l is a command-line flag.
		- Go provides a flag package supporting basic command-line flag parsing. We’ll use this package
		to implement our example command-line program.

		# -- How to use it
		 	- Basic flag declarations are available for string, integer, and boolean options
			    wordPtr := flag.String("word", "foo", "a string")	// declare a flag -word with description "a string" and default value is foo
*/

import (
	"flag"
	"fmt"
)

func PackageFlag() {
	wordlist := flag.String("w", "/path/names.txt", "Wordlist path for bruteforcing")
	thread := flag.Int("t", 3, "Number of threads")
	flag.Parse()

	fmt.Println("===== Start package flag =====")
	fmt.Println("[+] The default vaule for -w:", *wordlist)
	fmt.Println("[+] The default vaule for -t:", *thread)

}
