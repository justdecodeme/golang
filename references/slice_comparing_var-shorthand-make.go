package main

import (
	"fmt"
)

func main() {
	var v1 []int
	var v12 [][]int
	fmt.Println(v1)        // []
	fmt.Println(v12)       // []
	fmt.Println(v1 == nil) // true
	// v1[0] = 1              // invalid
	v1 = append(v1, 7)                // valid
	fmt.Println(v1, len(v1), cap(v1)) // [7] 1 1
	fmt.Println("-----------------")

	s1 := []int{}
	s12 := [][]int{}
	fmt.Println(s1)        // []
	fmt.Println(s12)       // []
	fmt.Println(s1 == nil) // false
	// s1[0] = 1              // invalid
	s1 = append(s1, 8)                // valid
	fmt.Println(s1, len(s1), cap(s1)) // [8] 1 1
	fmt.Println("-----------------")

	m1 := make([]int, 3)
	m12 := make([][]int, 3)
	fmt.Println(m1)        // [0 0 0]
	fmt.Println(m12)       // [[] [] []]
	fmt.Println(m1 == nil) // false
	// m1[0] = 1              // valid
	// m1[3] = 9 // invalid (runtime error: index out of range)
	m1 = append(m1, 9)                // valid
	fmt.Println(m1, len(m1), cap(m1)) // [0 0 0 9] 4 6
	fmt.Println("-----------------")
}
