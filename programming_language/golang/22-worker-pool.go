package main

/*
	# ----- What is worker pool?
		- However, there’s one blind spot in Go’s concurrency APIs that I find myself implementing in new Go programs time and time again: the worker
		pool (or otherwise known as a thread pool). Worker pools are a model in which a fixed number of m workers (implemented in Go with goroutines) work their
		way through n tasks in a work queue (implemented in Go with a channel). Work stays in a queue until a worker finishes up its current task and pulls a new
		one off.

		- Traditionally, thread pools have been useful to amortizing the costs of spinning up new threads. Goroutines are lightweight enough that that’s not a
		problem in Go, but a worker pool can still be useful in controlling the number of concurrently running tasks, especially when those tasks are accessing
		resources that can easily be saturated like I/O or remote APIs.

		- Explain by graph
			1. "./images/worker-pool.svg"
			1. "./images/worker-pool.png"

		- Resources:
			1- https://brandur.org/go-worker-pool
			2- https://www.bogotobogo.com/GoLang/GoLang_Channels_Worker_Pools.php
*/
import (
	"fmt"
	"math/rand"
	"time"
)

// Func worker to recieve to jobs through jobs channel and send back the results via results channel
func worker(jobs <-chan int, results chan<- int, id int) {
	t := 0
	for j := range jobs {
		fmt.Printf("[+] PROGRESS: Worker %v, job %v is starting...\n", id, j)
		t = rand.Intn(8)
		time.Sleep(time.Second * time.Duration(t))
		fmt.Printf("[+] DONE: worker: %v, job: %v finished.\n", id, j)
		results <- j * 2
	}
}

func WorkerPool() {
	fmt.Println("===== Start Worker pool =====")
	// Send and recieve channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// number of jobs
	var noOfJobs int = 5

	// start workers - initially blocked because there is no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(jobs, results, w)
	}

	// start jobs, and send them then close the channel
	for j := 1; j <= noOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect and print out the result of jobs
	for r := 1; r <= noOfJobs; r++ {
		fmt.Println("Result:", <-results)
	}
	close(results)

}
