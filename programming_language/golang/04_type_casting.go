package main

/*
	# ----- strconv package
		- Quick overview: Package strconv implements conversions to and from string representations of basic data types.
		# -- functions
			strconv.ParseInt(arg1,arg2,arg3)    # arg1 for value, arg2 assign base, arg3 determine variable size, ParseInt return 2 values (value and error)

	# -- start buildIn function
		- float64(arg1)				# covert arg1 to float64
		- int8(arg1)				# convert arg1 to int8
		- and so on ....
	# --------------------------------   end buildIn function -------------------------------- #
*/

import (
	"fmt"
	"strconv" //
)

// splitExampleName attempts to split example name s at index i,
// and reports if that produces a valid split. The suffix may be
// absent. Otherwise, it must start with a lower-case letter and
// be preceded by '_'.
//
// One of i == len(s) or s[i] == '_' must be true.
//
//Kemet,
func TypeCasting() {
	fmt.Println("===== Type Casting =====")
	var num int = 1001
	converted := strconv.FormatInt(int64(num), 10)
	fmt.Printf("[+] The type of num: %T, value of num: %v\n", num, num)
	fmt.Printf("[+] The type of num after casting: %T, value of num: %v\n", converted, converted)
}
