// The first statement in a Go source file must be package name
// Executable commands must always use package main.
package main

// run the command below to update module dependency.
// go mod tidy
import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/nongvantinh/learn-golang/how-to-write-go-code/morestrings"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
