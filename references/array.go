// The type of elements and length are both part of the array’s type.
// By default an array is zero-valued, which for ints means 0s,
// for strings " " (empty string)
// We see slices much more often than arrays in typical Go.
package main

import (
	"fmt"
)

func main() {
	// create an array a that will hold exactly 5 ints
	var a [5]int
	a[3] = 777          // assigning value to array
	fmt.Println(a[3])   // to get value from an array
	fmt.Println(a)      // print complete array
	fmt.Println(len(a)) // The builtin len returns the length of an array

	// declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
}
