package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	# ------ What is buffered channel:
		- Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
			ch := make(chan int, 100)	# Can handle till 100 variable
		- Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
*/

var wg_chB sync.WaitGroup

// Create function to write data to the channel.
// writeToCh(channel, buffer_number)
func writeToCh(c chan uint, ch_count uint) {
	var i uint
	for i = 1; i <= ch_count; i++ {
		c <- i
		fmt.Println("Successfully wrote", i, "to ch.")
	}
	wg_chB.Done()
}

// Create Function to loop over the channel to read values
func ReadFromCh(c chan uint, ch_count uint) {
	var i uint = 1
	for val := range c {

		fmt.Println("Read value from ch: ", val)
		time.Sleep(time.Second * time.Duration(2))
		i++

		// Avoid wg error!
		if i == ch_count+1 {
			close(c)
			wg_chB.Done()
			break
		}
	}
}

func ChannelBuffer() {
	var ch_count uint = 10
	var channelBuffered = make(chan uint, ch_count)

	wg_chB.Add(2)
	go writeToCh(channelBuffered, ch_count)
	go ReadFromCh(channelBuffered, ch_count)
	wg_chB.Wait()
}
