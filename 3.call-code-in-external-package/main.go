// Declare a main package
package main

// Import fmt package give us Println function and more.
// rsc.io/quote give us Go function. Since we need rsc.io package, we need to enable dependency tracking!
import (
	"fmt"

	"rsc.io/quote"
)

// A main function executes by default when you run the main package.
func main() {
	// quote.Go() function, printing a clever message about communication.
	fmt.Println(quote.Go())
}
