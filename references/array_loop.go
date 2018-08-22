package main

import (
	"fmt"
)

func main() {
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

	var c [58]string
	for i := 65; i <= 122; i++ {
		c[i-65] = string(i)
	}
	fmt.Println(c)
}
