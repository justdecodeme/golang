package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Array types are one-dimensional, but we can compose types to build multi-dimensional data structures
	var twoD [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = strconv.Itoa(i) + strconv.Itoa(j)
		}
	}
	fmt.Println("2d: ", twoD)
}
