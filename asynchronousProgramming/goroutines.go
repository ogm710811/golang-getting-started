package main

import (
	"runtime"
	"time"
)

func main() {
	// for parallel threats using multiple CPU
	// 8 is the number of processors running
	// so the abcGen() will run parallel in multiple
	// processors
	runtime.GOMAXPROCS(8)
	go abcGen()
	println("This comes first")

	// Since go routines are asynchronuos we need
	// to pause the main function in order to wait
	// for the abcGen() to finish so we use the
	// time package
	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
