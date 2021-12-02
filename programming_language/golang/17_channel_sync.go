package main

import "time"

func anotherGoRoutine(ch chan bool) {
	time.Sleep(time.Second * time.Duration(4))
	ch <- true
}

func ChannelSync() {
	var chSync = make(chan bool, 1)
	go anotherGoRoutine(chSync)
	<-chSync
}
