package main

import "fmt"

func main() {
	// The type [n]T is an array of n values of type T.
	var a1 [5]int // Declare an array of 5 int and init with default value of int (which is 0)
	// Note: An array's length is part of its type, so arrays cannot be resized.
	fmt.Println(a1)

	// Access array's element by using subscript operator.
	a1[0] = 10
	a1[1] = 9
	a1[2] = 7
	a1[3] = 4
	a1[4] = 1
	fmt.Println(a1)

	// Init value of array when array is created is supports. And it called an array literal:
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
