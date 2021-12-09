package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	# ----- What is timeout function?
		- Go is normally used to write backend services. Generally, a request is completed by multiple serial or parallel subtasks. Each subtask may issue
		another internal request. When the request times out, it’s better to return quickly and release the occupied resources, such as goroutines, file
		descriptors, etc.
		- Common timeout control on the server side
			1. In-process processing
			2. Serving client requests, such as HTTP or RPC requests
			3. Calling other services, including calling RPC or accessing DB, etc.

		# -- Example on unhandled timeout golang code:
			- In order to simplify, we take a request function hardWork as an example. It does not matter what it is used for. As the name suggests,
			it may be slow to process.
				func hardWork(job interface{}) error {
					time.Sleep(time.Minute)
					return nil
				}
			​
				func requestWork(ctx context.Context, job interface{}) error {
					return hardWork(job)
				}
			- When we use this kind of code to serve, the familiar image shows up for one minute. I guess most people can’t wait that long, but the server is
			still working on that even the page is closed. And the processing resources are not released to serve other requests.
*/

var wgTimeout sync.WaitGroup

func sendData_server1(ch chan string) error {
	time.Sleep(time.Second * time.Duration(randTimeGenerator()))
	ch <- "Data has been sent by server 1"
	wgTimeout.Done()
	return nil
}

func randTimeGenerator() int {
	return rand.Intn(10)
}

func sendData_server2(ch chan string) error {
	time.Sleep(time.Second * time.Duration(randTimeGenerator()))
	ch <- "Data has been sent by server 2"
	wgTimeout.Done()
	return nil
}

func TimeOut() {
	fmt.Println("==== Timeout practice =====")

	// Make a channel
	channelTimeOut := make(chan string, 2)

	// Apply the goroutine
	wgTimeout.Add(2)
	go sendData_server1(channelTimeOut)
	go sendData_server2(channelTimeOut)

	select {
	case result := <-channelTimeOut:
		fmt.Println(result)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout!")
	}
	// wgTimeout.Wait()

}
