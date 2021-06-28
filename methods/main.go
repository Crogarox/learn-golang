package main

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.
type Vertex struct {
	X, Y float64
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// the Abs method has a receiver of type Vertex named v.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods are functions
// Remember: a method is just a function with a receiver argument.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare methods with pointer receivers to change the value it pointed to.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	// methods with pointer receivers take either a value or a pointer as the receiver when they are called:
	p := &v
	p.Scale(10)
	fmt.Println(p.Abs())

}
