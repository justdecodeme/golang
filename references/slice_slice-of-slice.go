package main

import (
	"fmt"
)

func main() {
	// EXAMPLE 01

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD) // [[0] [1 2] [2 3 4]]

	// EXAMPLE 02

	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions) // [[0 1 2 3] [0 1 2 3] [0 1 2 3]]

	// EXAMPLE 03

	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records) // [[Foster Nathan 100.00 74.00] [Gomez Lisa 92.00 96.00]]
}
