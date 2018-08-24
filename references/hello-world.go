package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")     // Hello, 世界
	fmt.Println([]byte("Hello")) // [72 101 108 108 111] | here each letter is taking 1 bytes
	fmt.Println([]byte("世界"))    // [228 184 150 231 149 140] | here each letter is taking 3 bytes
}
