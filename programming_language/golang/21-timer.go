package main

import (
	"fmt"
	"time"
)

/*
	# ----- What is the timer?
		- These are used for one-off tasks. It represents a single event in the future. You tell the timer how long you want to wait, and it provides a channel
		that will be notified at that time.
		- There are 3 type of Timers:
			1. Timeout (Timer)
				- time.After waits for a specified duration and then sends the current time on the returned channel:
					select {
						case news := <-AFP:
							fmt.Println(news)
						case <-time.After(time.Hour):
							fmt.Println("No news in an hour.")
					}
			2. The underlying time.Timer will not be recovered by the garbage collector until the timer fires. If this is a concern, use time.NewTimer instead and
			call its Stop method when the timer is no longer needed:
				for alive := true; alive; {
					timer := time.NewTimer(time.Hour)
					select {
					case news := <-AFP:
						timer.Stop()
						fmt.Println(news)
					case <-timer.C:
						alive = false
						fmt.Println("No news in an hour. Service aborting.")
					}
				}
			3. Repeat (Ticker)
				- time.Tick returns a channel that delivers clock ticks at even intervals:


*/
func backgroundTaskTimer() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Println("Tick, Tock")
	}
}

func Timer() {

	go backgroundTaskTimer()
	select {}

}
