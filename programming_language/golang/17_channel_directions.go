package main

import (
	"errors"
	"fmt"
)

/*
	# ----- What is channel direction?
		- Channels (as function parameters) can have a direction. By default a channel can both send and receive data, like func f(c chan string).
		But you can define a channel to be receive-only or send-only. If you then use it in the other direction, it would show a compile-time error. This
		improves the type-safety of the program.
*/

// var wg_ch_dir sync.WaitGroup

// Function "ping" is desngined to send msg over the channel
func ping(ch chan<- string, msg string) error {
	if msg == "" {
		return errors.New("ERROR: The message is empty")
	} else {
		ch <- msg
		fmt.Println("Data has been added to the channel.")
	}
	return nil

}

// Function "pong" recieve message over the channel
func pong(ch <-chan string) error {
	fmt.Println(len(ch))
	if len(ch) == 0 {
		return errors.New("ERROR: No message in the queue")
	}
	fmt.Println("Your message :", <-ch)
	return nil
}
func ChannelDirection() {
	fmt.Println("===== Channel Direction =====")
	chDirect := make(chan string, 1)

	ping(chDirect, "Test")
	pong(chDirect)
	fmt.Println("Done...")
}
