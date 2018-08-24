package main

import "fmt"

func main() {
	// ----------
	// LOOPS
	// ----------

	// Like a C for
	// for init; condition; post { }

	// Like a C while
	// for condition { }

	//  Like a C for(;;)
	// for { }

	// The most basic type, with a single condition.
	m := 1
	for m <= 3 {
		fmt.Println(m)
		m = m + 1
	}

	// using continue and break within loop
	for n := 0; n <= 5; n++ {
		if n == 4 {
			break
		}
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// nested loops
	var i, j int
	for i = 2; i < 20; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // if factor found, not prime
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
