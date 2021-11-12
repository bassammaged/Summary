package main

import (
	"fmt"
	"sync"
)

/*
	# ----- Closing Channels:
		- Closing a channel indicates that no more values will be sent on it. This can be useful to communicate
		completion to the channel’s receivers.
*/

var wgChannelClose sync.WaitGroup

// func randomGoRoutineFunc(ch chan string) {
// 	ch <- "This message from randomGoRoutineFunc"
// 	ch <- "Done! randomGoRoutineFunc finsihed its work"
// 	wgChannelClose.Done()
// }

func closingExample(ch chan int) {
	for {
		j, more := <-ch
		if more {
			fmt.Println("Recieved jobs", j)
		} else {
			fmt.Println("All jobs has been recieved")
		}
		wgChannelClose.Done()
	}
}

func ChannelClose() {
	fmt.Println("===== Start Channel Close =====")
	// ch := make(chan string)
	ch1 := make(chan int)
	// wgChannelClose.Add(1)
	// go randomGoRoutineFunc(ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// close(ch)
	// fmt.Println("The channel has been closed")

	wgChannelClose.Add(5)
	go closingExample(ch1)
	for j := 1; j <= 5; j++ {
		ch1 <- j
		fmt.Println("sent job", j)
	}
	wgChannelClose.Wait()

}
