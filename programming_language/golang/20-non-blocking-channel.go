package main

import "fmt"

/*
	# ----- What is meant by non blocking?
		- Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even
		non-blocking multi-way selects.
*/

func NonBlockingWithDefault() {
	ch1NonBlocking := make(chan string)
	ch2NonBlocking := make(chan string)

	select {
	case msg := <-ch1NonBlocking:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case ch2NonBlocking <- msg: // There is no receiver that why the message will not be sent.
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-ch1NonBlocking:
		fmt.Println("received message", msg)
	case sig := <-ch2NonBlocking:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
