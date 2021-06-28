package main

import (
	"fmt"
	"strings"
)

func main() {

	//An array has a fixed size.
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// The type []T is a slice with elements of type T.
	var s []int
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	s = primes[1:4] // This selects a half-open range which includes the first element, but excludes the last one.
	fmt.Println(s)

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	s2 := primes[0:6]
	s2[0] = 10
	s2[1] = 9
	s2[2] = 7
	s2[3] = 4
	s2[4] = 1
	s2[5] = 1
	fmt.Println(s2)

	// Other slices that share the same underlying array will see those changes.
	fmt.Println(s)

	// A slice literal is like an array literal without the length.
	// This creates an array, then builds a slice that references it:
	sa := []bool{true, true, false}
	fmt.Println(sa)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.

	// Given an array:
	var ai [10]int
	// These are equivalent.
	ae := ai[0:10]
	al := ai[:10]
	af := ai[0:]
	afe := ai[:]

	fmt.Println(ae, al, af, afe)

	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	// We can obtain these properties by using len() and cap()
	aaaa := ai[6:8]
	aa_len := len(aaaa)
	aa_cap := cap(aaaa) // 6 -> 10
	fmt.Println(aa_len, aa_cap)

	//  The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var sy []int
	fmt.Println(sy, len(sy), cap(sy))
	if sy == nil {
		fmt.Println("nil!")
	}

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	ad := make([]int, 5) // len(a)=5
	fmt.Println(len(ad), cap(ad))
	// To specify a capacity, pass a third argument to make:
	bd := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(len(bd), cap(bd))

	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// It is common to append new elements to a slice, and so Go provides a built-in append function.
	var ap []int
	fmt.Printf("len=%d cap=%d %v\n", len(ap), cap(ap), ap)
	// append works on nil slices and grows as needed.
	ap = append(ap, 0)
	// We can add more than one element at a time.
	ap = append(ap, 2, 3, 4)
	fmt.Printf("len=%d cap=%d %v\n", len(ap), cap(ap), ap)

}
