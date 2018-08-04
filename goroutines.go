package main

import (
	"runtime"
	"time"
)

// func main() {
// 	abcGen()
// }
// func abcGen() {
// 	for l := byte('a'); l <= byte('z'); l++ {
// 		println(string(l))
// 	}
// }

// // main funcion is entry and exit point
// func main() {
// 	go abcGen()
// }
// func abcGen() {
// 	for l := byte('a'); l <= byte('z'); l++ {
// 		println(string(l))
// 	}
// }

// import "time"
// func main() {
// 	go abcGen()
//  println("This comes first!")
// 	time.Sleep(100 * time.Millisecond)
// }
// func abcGen() {
// 	for l := byte('a'); l <= byte('z'); l++ {
// 		println(string(l))
// 	}
// }

func main() {
	runtime.GOMAXPROCS(8)
	go abcGen()
	println("This comes first!")
	time.Sleep(100 * time.Millisecond)
}
func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
