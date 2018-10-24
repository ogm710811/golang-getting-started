package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	// create the channels
	ch := make(chan string)
	doneCh := make(chan bool)

	// pass the channel to the abcGen()
	go abcGen(ch)
	go printer(ch, doneCh)

	println("This comes first")

	<-doneCh
}

func abcGen(ch chan string) {
	// fill the channel with the for loop
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}
	doneCh <- true
}
