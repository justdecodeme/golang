package main

import "fmt"

func say_hello(msg string) {
	fmt.Println(msg)
}

func return_anony_func() func(string) {
	// return an anonymous function
	return func(msg string) {
		fmt.Println(msg)
	}
}

// closure
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	say_hello("Hello")

	// anonlymous function declared
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymous")

	print_func := return_anony_func()
	print_func("Hello returned from anonymous")

	next_int := int_seq()
	fmt.Println(next_int())
	fmt.Println(next_int())

	next_int2 := int_seq()
	fmt.Println(next_int2())
}
