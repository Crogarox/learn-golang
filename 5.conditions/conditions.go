package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// The if condition.
	i := 5
	if i == 5 {
		fmt.Println(i)
	}

	//  Like for, the if statement can start with a short statement to execute before the condition.
	//  Variables declared by the statement are only in scope until the end of the if.
	if v := 5; v < 10 {
		fmt.Println(v)
	} // v ends here.

	//  Variables declared inside an if short statement are also available inside any of the else blocks.
	if v := 10; v != 10 {
		fmt.Print("v not equal to 10: ")
		fmt.Println(v)
	} else {
		fmt.Println("v equal to 10")
	}

	// A switch statement is a shorter way to write a sequence of if - else statements.
	// It runs the first case whose value is equal to the condition expression.
	// Go only runs the selected case, not all the cases that follow.
	// Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
	//  Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd, plan9, windows...
		fmt.Printf("%s\n", os)
	}

	//  Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Printf("Good everning")
	}
}
