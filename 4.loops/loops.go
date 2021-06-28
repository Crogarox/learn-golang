package main

import "fmt"

func main() {

	// Print 0 to 10.
	for i := 0; i != 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// The init and post statements are optional.
	// and you can drop the semicolons, C's while is spelled for in Go.
	i := 0
	for i != 10 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	// Remove the condition give a loop, loops forever.
	for {

	}
}
