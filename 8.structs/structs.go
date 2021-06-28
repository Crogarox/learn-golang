package main

import "fmt"

// Create new struct of type Vertex
type Vertex struct {
	X int
	Y int
}

func main() {

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 4}
	// Struct fields are accessed using a dot.
	v.X = 7
	fmt.Println(v)

	// Struct fields can be accessed through a struct pointer.
	pv := &v
	pv.Y = 10

	fmt.Println(v)

	// To access the field X of a struct when we have the struct pointer pv we could write (*pv).X.
	// However, that notation is cumbersome, so the language permits us instead to write just pv.X, without the explicit dereference.
	(*pv).X = 10

	fmt.Println(v)

	// A struct literal denotes a newly allocated struct value by listing the values of its fields.
	var (
		v1 = Vertex{1, 2} // has type Vertex
		v2 = Vertex{X: 1} // Y:0 is implicit
		v3 = Vertex{}     // X:0 and Y:0
		// The special prefix & returns a pointer to the struct value.
		p = &Vertex{1, 2} // has type *Vertex

		// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p)

}
