package main

/*

	# ----- Concurrency vs Parallelism
		- Concurrency is the composition of independently executing processes, while parallelism is the
		simultaneous execution of (possibly related) computations. Concurrency is dealing with lots of things
		at once. Parallelism is about doing things at once
	# ----- What is goroutine?
		- Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can
		be thought of as lightweight threads. The cost of creating a Goroutine is tiny when compared to a thread.
		Hence it's common for Go applications to have thousands of Goroutines running concurrently.

		# -- Advantages of Goroutines over threads
			1. Goroutines are extremely cheap when compared to threads. They are only a few kb in stack size and
			the stack can grow and shrink according to the needs of the application whereas in the case of threads
			the stack size has to be specified and is fixed.
			2. The Goroutines are multiplexed to a fewer number of OS threads. There might be only one thread in a program
			with thousands of Goroutines. If any Goroutine in that thread blocks say waiting for user input,
			then another OS thread is created and the remaining Goroutines are moved to the new OS thread. All
			these are taken care of by the runtime and we as programmers are abstracted from these intricate details
			and are given a clean API to work with concurrency.
			3.Goroutines communicate using channels. Channels by design prevent race conditions from happening
			when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which
			Goroutines communicate. We will discuss channels in detail in the next tutorial.

		# -- Disadvantage of Goroutine
			- Race Condition:
				- A race condition occurs when two or more routines try to access the resource like variable or
				data structure and attempt to read or write the resources without regard to other routines. So it
				will create tremendous problems.

				- So Golang tooling introduced race detector. Race detector is code that is built in your
				program during the build process. Once your program starts, it will start to detect race condition. It
				is a really superb tool and does a great job



	# ----- What is synchronization
		- Golang synchronizes multiple goroutines via channels and sync packages to prevent multiple
		goroutines from competing for data.

		# -- Most important functions that provided by sync
			- WaitGroup
				- WaitGroup is the good concept in goroutines that will wait for other goroutines to finish their
				execution. Sometimes, for executing one activity, other activities need to complete.
				- As WaitGroup waits for no. of Goroutines to complete their execution, the main goroutine calls
				Add to set no. of goroutines to wait for. Then each goroutine runs and calls Done when finished
				their execution. At same time it will call Wait to block and wait until all goroutines finish
				their execution.
				# -- To use sync.WaitGroup :
					1. Create instance of sync.WaitGroup → var wg sync.WaitGroup
					2. Call Add(n) where n is no of goroutines to wait for → wg.Add(1)
					3. Execute defer Wg.Done( ) in each goroutine to indicate that goroutine is finished
					4. executing to the WaitGroup.
					5. Call wg.Wait( ) where we want to block

	# ----- Resources:
		- Read more about channel, Concurrency, Parallelism & synchronization: https://www.mindbowser.com/golang-concurrency/
		- in advance: https://lebum.medium.com/use-of-synchronization-techniques-in-golang-53d75bc0a646
*/

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Function to repeat the value based on the user willing
// Function support sync done method
// Usage: repeatAfterMe("Your Message",No_of_repeat), No return value
// Function has it own sleeping time (Interval time 1-3 seconds) to simulate waiting period.
func repeatAfterMe(msg string, times uint8) {

	if times < 1 {
		fmt.Println("[+] ERROR: Something wrong!")
		return
	}
	var i uint8
	for i = 1; i <= uint8(times); i++ {
		fmt.Printf("[+] %v, %v \n", i, msg)
		var t int = rand.Intn(3)
		time.Sleep(time.Second * time.Duration(t))
	}
	// Done function for sync module.
	wg.Done()
}

func GoRoutine() {
	fmt.Println("===== Start GoRoutine =====")
	fmt.Printf("[+] INPUT: Please insert your message: ")
	msg := bufio.NewScanner(os.Stdin)
	msg.Scan()

	fmt.Printf("[+] INPUT: Please insert the repeat times: ")
	input2 := bufio.NewScanner(os.Stdin)
	input2.Scan()
	times, _ := strconv.ParseInt(input2.Text(), 10, 8)

	// Adding WaitGroup number
	wg.Add(2)
	go repeatAfterMe(msg.Text(), uint8(times))
	go repeatAfterMe(msg.Text(), uint8(times))
	// Instruct the main process/thread to wait work goroutine
	wg.Wait()
}
