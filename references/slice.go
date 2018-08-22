// The value of uninitialized slice is nil

// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

/*
You can do "slicing" on a string because a string is a slice of bytes.
Remember:
	- A string is made up of runes
	- A rune is a unicode code point
	- A unicode code point is 1 to 4 bytes
Strings are made up of runes, runes are made up of bytes, so strings are
made up of bytes. A stringis a bunch of bytes; a slice of bytes.
*/
package main

import "fmt"

func main() {
	var s1 []string
	fmt.Println(s1)        // []
	fmt.Println(s1 == nil) // true (IMPORTANT)
	fmt.Printf("%T\n", s1) // []string

	// We can declare and initialize a variable for slice in a single line as well.
	s2 := []string{}
	fmt.Println(s2)        // []
	fmt.Println(s2 == nil) // false (IMPORTANT)
	fmt.Printf("%T\n", s2) // []string

	t := []string{"g", "h", "i"}
	fmt.Println(t)        // [g h i]
	fmt.Printf("%T\n", t) // []string

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length,
	// use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]int, 3)
	fmt.Println(s) // [0 0 0]

	// append returns a slice containing one or more new values
	s = append(s, 6)
	s = append(s, 9, 84)
	fmt.Println(s) // [0 0 0 6 9 84]

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]int, len(s))
	copy(c, s)
	fmt.Println(c) // [0 0 0 6 9 84]

	// Slices support a “slice” operator with the syntax slice[low:high].
	l := s[2:5]
	fmt.Println(l) // [0 6 9]

	// loop through slice
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}
	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}
	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	// returns value at index i of a string
	fmt.Println("aString"[0]) // 97
}
