package v1

import f "fmt"

func Welcome(n string) {
	f.Printf("Welcome %v\n", n)
}

func sayHi() {
	f.Println("====== Custom Package ======")
	f.Println("Hi, There!")
}

func init() {
	sayHi()
}
