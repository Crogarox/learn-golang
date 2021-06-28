package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// A map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
var m map[string]Vertex

func main() {
	// The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals are like struct literals, but the keys are required.
	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var v = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(v)

	// Insert or update an element in map m:
	m["Answer"] = Vertex{42, 25}
	// Retrieve an element:
	answer := m["Answer"]
	fmt.Println(answer)

	// Test that a key is present with a two-value assignment:
	elem, ok := m["Answer"]
	if ok {
		fmt.Println(elem)
	}

}
