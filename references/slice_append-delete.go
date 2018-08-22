package main

import "fmt"

func main() {
	// slice to slice
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9}

	s1 = append(s1, s2...)
	fmt.Println(s1)

	// delete
	s1 = append(s1[:2], s2[3:]...)
	fmt.Println(s1)
}
