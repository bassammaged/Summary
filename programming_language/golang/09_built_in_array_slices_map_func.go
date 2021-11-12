package main

import "fmt"

func BuiltInArraySlicesMapFunc() {
	fmt.Println("===== Built in Functions for array & slices =====")
	var hosts = []string{"secranix.com", "beta.secranix.com"}
	// ----- Length of the array/slice/map
	arrayLength := len(hosts)
	fmt.Printf("[+] The length of the slices/array: %v", arrayLength)

	// ----- (slices) Append a new vaule
	fmt.Printf("[+] The hosts: %v\n", hosts)
	hosts = append(hosts, "test.secranix.com")
	fmt.Printf("[+] The hosts (slices) after append: %v\n", hosts)

	// ----- (map) Delete the index
	var idsNames map[int]string = make(map[int]string)
	idsNames[19951] = "Kemet B"
	idsNames[12211] = "Junior A"
	idsNames[33221] = "Weloo A"
	fmt.Printf("[+] The values of map: %v\n", idsNames)
	delete(idsNames, 33221)
	fmt.Printf("[+] The values of map after deletion: %v\n", idsNames)

}
