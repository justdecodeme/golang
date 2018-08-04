package main

func main() {
	msg1 := "msg1"
	msg2 := "msg2"

	anonymousFun := func() {
		println("Hi from Anonymous function.")
	}
	callByValue(msg1)
	callByRef(&msg2)
	variadicFunc("Hello", "Go", "from", "Rakko")

	println(msg1)
	println(msg2)

	numTerms, sum := add(1, 3, 6)
	println("Added:", numTerms, " terms for a total of", sum)
	anonymousFun()
}

func callByValue(msg string) {
	println(msg)
	msg = "msg12"
}

func callByRef(msg *string) {
	println(*msg)
	*msg = "msg22"
}

func variadicFunc(msg ...string) {
	for _, msg := range msg {
		println(msg)
	}
}

func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
