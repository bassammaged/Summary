package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	# ----- What are channels?
		- Normally, when we talk about channels, we think of the ones in applications like RabbitMQ, Redis, AWS SQS,
		and so on. Anyone with no or only a small amount of Golang knowledge would think like this. But Channels in
		Golang are different from a work queue system. In the work queue system like above, there are TCP connections
		to the channels, but in Go, the channel is a data structure or  even a design pattern, which we’ll explain
		later. So, what are the channels in Golang exactly?
		Channels are the medium through which goroutines can communicate with each other. In simple terms, a
		channel is a pipe that allows a goroutine to either put or read the data.

		# -- How to define a channel?
			- Let’s see a syntax that will declare a channel. We can do so by using the chan keyword provided by Go.
			- You must specify the data type as the channel can handle data of the same data type.
				//create channel
 				var c chan int
			- Very simple! But this is not useful since it would create a Nil channel. Let’s print it and see.
				fmt.Println(c)
				fmt.Printf("Type of channel: %T\n", c)
				output: Type of channel: chan int
			- As you can see, we have just declared the channel, but we can’t transport data through it. So,
			 to create a useful channel, we must use the make function
			 	//create channel
				c := make(chan int)
				fmt.Printf("Type of `c`: %T\n", c)
				fmt.Printf("Value of `c` is %v\n", c)

				output: Type of `c`: chan int
				output: Value of `c` is 0xc000022120
				- As you may notice here, the value of c is a memory address. Keep in mind that channels are
				nothing but pointers. That’s why we can pass them to goroutines, and we can easily put the data
				or read the data.

			- Write and read operation on a channel:
				- Go provides an easy way to read and write data to a channel by using the left arrow.
					c <- 10
					v1 := <-c

			# -- Blocking:
				Blocking happens when:
					- A goroutine try to receive message, but channel is empty and other some groutine is running

				- In this case, message receiving will be blocked and "Receiver" groutine waits it. Because while
				other groutine is runnning, he hope someone send message to channel. ( Not surrender )

			# -- Deadlock
				- When deadlock happens? A deadlock happens when a group of goroutines are waiting for each other
				and none of them is able to proceed.
				- For example: When one goroutine try to receive message from channel, the channel is empty and
				no other goroutine running.

					- Deadlock example. No sender
						package main

						func main() {
							messages := make(chan string)

							// Do nothing spawned goroutine
							go func() {}()

							// A groutine ( main groutine ) trying to send message to channel
							// But no other groutine runnning
							// And channel has no buffers
							// So it raises deadlock error
							messages <- "I wanna tell you." // fatal error: all goroutines are asleep - deadlock!
						}
					- Deadlock example. No receiver
						package main

						func main() {
							messages := make(chan string)

							// Do nothing spawned goroutine
							go func() {}()

							// A groutine ( main groutine ) trying to receive message from channel
							// But channel has no messages, it is empty.
							// And no other groutine running. ( means no "Sender" exists )
							// So channel will be deadlocking
							<-messages // fatal error: all goroutines are asleep - deadlock!
						}
					- Blocking example : No groutines procceeds
						package main

						import "fmt"

						// Two groutines are running but deadlock happens

						// Output example
						//
						// Try to receive message
						// Try to receive message
						// fatal error: all goroutines are asleep - deadlock!

						func main() {
							messages := make(chan string)

							go func() {
								fmt.Println("Try to receive message") // Printing
								<-messages                            // Blocking
								fmt.Println("Receive message")        // Never reached
							}()

							fmt.Println("Try to receive message") // Printing
							<-messages                            // Blocking
							fmt.Println("Receive message")        // Never reached

						}
				# -- Question. Why deadlock does not happen in spawned groutine ?
						- Because "main goroutine" always running. So until main groutine finishes, it means, one
						or more groutine running anytime. When main groutine finishes, it is end of process, means
						end of all groutines. So even if spawned groutine meets blocking, does not meet deadlock error.
*/

var wg2 sync.WaitGroup

// This is also the way to define the “receive” only type of channels.
func readChannelData(c chan string) {
	v1 := <-c
	// print the data from channel
	fmt.Printf("[+] The value %s\n", v1)
	wg2.Done()

}

// channel blocking Example
func goRoutineChannelBlocking(c chan string) {
	time.Sleep(time.Second * 10)
	c <- "Message from Blocking function! Hey"
	wg2.Done()
}

func GoRoutineChannel() {
	fmt.Println("===== Start Channel =====")
	// Create a channel
	channelGoRoutine := make(chan string)

	var GoRoutineChannelVar string = "A string from GoRoutineChannel"

	wg2.Add(2)
	// pass the channel to the function
	go readChannelData(channelGoRoutine)
	// Pass data to the channel
	channelGoRoutine <- GoRoutineChannelVar

	// Using blocking goroutine function
	go goRoutineChannelBlocking(channelGoRoutine)
	// Get the data from the blocking goroutine function
	/*
		the main goroutine receives message from channel and appears after spawned goroutine awaked from sleeping Channel
		is empty until other goroutine send message. So this receiving will be blocked for a while.
	*/
	var v2 string = <-channelGoRoutine
	fmt.Printf("%v\n", v2)
	// Waiting until wg2.Add() finished.
	wg2.Wait()

}
