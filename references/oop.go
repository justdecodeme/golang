package main

import "fmt"

func main() {
	// foo := myStruct{}
	// foo.myField = "bar"

	// foo := myStruct{"bar"}

	foo := new(myStruct)
	foo.myField = "bar"

	println(foo.myField)

	bar := newMyStruct()
	bar.myMap["bar"] = "baz"
	fmt.Println(bar)

	mp := messagePrinter{"foo"}
	mp.printMessage()
}

// Structs and Fields
type myStruct struct {
	myField string
}

// constructor functions
type myStruct2 struct {
	myMap map[string]string
}

func newMyStruct() *myStruct2 {
	result := myStruct2{}
	result.myMap = map[string]string{}

	return &result
}

// methods
type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}
