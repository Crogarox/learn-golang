package main

import (
	"fmt"
	"math"
	"time"
)

// An interface type is defined as a set of method signatures.
type Abser interface {
	Abs() float64
}

type T struct {
	S string
}
type I interface {
	M()
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// A value of interface type can hold any value that implements those methods.
	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())

	// The interface type that specifies zero methods is known as the empty interface:
	var iss interface{}

	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type
	iss = 645
	iss = "haha"

	TypeAssertions()
	TypeSwitches(iss)
	Stringers() // One of the most ubiquitous interfaces is Stringer defined by the fmt package.
	Errors()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A type implements an interface by implementing its methods
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func TypeAssertions() {
	var i interface{} = "hello"

	// A type assertion provides access to an interface value's underlying concrete value.
	s := i.(string)
	fmt.Println(s)

	// To test whether an interface value holds a specific type.
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func TypeSwitches(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func Stringers() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

type MyError struct {
	When time.Time
	What string
}

// The error type is a built-in interface similar to fmt.Stringer:
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func Errors() {
	// A nil error denotes success; a non-nil error denotes failure.
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
