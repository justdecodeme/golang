package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(8)
	ch := make(chan string)
	doneCh := make(chan bool)

	go abcGen(ch)
	go printer(ch)
	println("This comes first!")
	time.Sleep(100 * time.Millisecond)
}
func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
	close(ch)
}
func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}
	doneCh <- true
}
