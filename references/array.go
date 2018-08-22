package main

import (
	"fmt"
	"strconv"
)

func main() {
	// The type of elements and length are both part of the array’s type.
	// By default an array is zero-valued, which for ints means 0s,
	// for strings " " (empty string)
	// We see slices much more often than arrays in typical Go.

	// create an array a that will hold exactly 5 ints
	var a [5]int
	a[3] = 777          // assigning value to array
	fmt.Println(a[3])   // to get value from an array
	fmt.Println(a)      // print complete array
	fmt.Println(len(a)) // The builtin len returns the length of an array

	// declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// Array types are one-dimensional, but we can compose types to build multi-dimensional data structures
	var twoD [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = strconv.Itoa(i) + strconv.Itoa(j)
		}
	}
	fmt.Println("2d: ", twoD)

	var c [58]string
	for i := 65; i <= 122; i++ {
		c[i-65] = string(i)
	}
	fmt.Println(c)

	var d [256]byte
	for i := 0; i < 256; i++ {
		d[i] = byte(i)
	}
	// loop through an array
	for i, v := range d {
		fmt.Printf("%v - %T - %b\n", v, v, v) // type will be uint8 (alis for byte)
		if i > 10 {
			break
		}
	}
}
