package main

import "fmt"

func main() {

	const (
		first = iota
		second
	)
	const (
		third = iota
	)
	println(first)
	println(second)
	println(third)

	const (
		a = 1 << (10 * iota)
		b
		c
	)
	println(a)
	println(b)
	println(c)

	var myInt int
	myInt = 42
	println(myInt)

	var myFloat32 float32 = 42
	println(myFloat32)

	myString := "Hello Go"
	println(myString)

	myComplex := complex(2, 3)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))

	// myArray := [...]int{42, 27, 92}
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 92
	fmt.Println(myArray)

	mySlice := make([]float32, 100)
	mySlice[0] = 12.
	mySlice[1] = 15.
	mySlice[2] = 18.
	fmt.Println(mySlice)

	myMap := make(map[int]string)
	myMap[32] = "Foo"
	myMap[14] = "Bar"
	fmt.Println(myMap)
}
