package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var vnt int
var mutex sync.Mutex //create mutex type variable
func incrementv(s string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(3 * time.Millisecond)
		mutex.Lock() //Locking while incrementing
		x := vnt
		x++
		vnt = x
		fmt.Println(s, i, "Counter:", vnt)
		mutex.Unlock() //Unlocking
	}
	defer wg1.Done()
}
func GoRoutineRaceMutex() {
	fmt.Println("Starting")
	wg1.Add(2)
	go incrementv("show:")
	go incrementv("display:")
	wg1.Wait()
	fmt.Println("Final Counter:", vnt)
}

/*
	- Now observe the output with a race detector as there is no any race condition and counter
	values for display and show are not mixing.
*/
