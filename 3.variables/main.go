package main

import "fmt"

//  The var statement declares a list of variables; as in function argument lists, the type is last.
var a, b bool

//  A var declaration can include initializers, one per variable.
var c, d int = 1, 2

//  If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var e, f, g, h = 2.9, false, 3, "hello"

//  Constants are declared like variables, but with the const keyword.
//  Constants can be character, string, boolean, or numeric values.
//  Constants cannot be declared using the := syntax.
const Pi = 3.14

// Numeric constants are high-precision values.
// An untyped constant takes the type needed by its context.
// (An int can store at maximum a 64-bit integer, and sometimes less.)
const big_int = 1 << 100
const small_int = big_int >> 99

func main() {
	//  A var statement can be at package or function level.
	var i, j int
	fmt.Println(a, b, c, d, e, f, g, h, i, j)

	//  Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// this is called: Short variable declarations
	k, l := "hi!", false

	fmt.Println(k, l)

	const greeting = "Hello"

	fmt.Println(greeting)
	fmt.Println(small_int)

	// go's basic type are:
	var bool_type bool = true       // default: false
	var string_type string = "haha" // default: ""

	fmt.Println(bool_type, string_type)

	var int_type int = 124    // default: 0. The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
	var int8_type int8 = 40   // default: 0.
	var int16_type int16 = 9  // default: 0.
	var int32_type int32 = 17 // default: 0.
	var int64_type int64 = 91 // default: 0.

	fmt.Println(int_type, int8_type, int16_type, int32_type, int64_type)

	var uint_type uint = 123123   // default: 0. The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems
	var uint8_type uint8 = 3      // default: 0.
	var uint16_type uint16 = 19   // default: 0.
	var uint32_type uint32 = 62   // default: 0.
	var uint64_type uint64 = 36   // default: 0.
	var uintptr_type uintptr = 34 // default: 0. The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems

	fmt.Println(uint_type, uint8_type, uint16_type, uint32_type, uint64_type, uintptr_type)

	//  When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

	var byte_type byte = 19 // An alias for int8

	var rune_type rune = 21 // An alias for int32 represents a Unicode code point

	fmt.Println(byte_type, rune_type)

	var float32_type float32 = 3.14               // default: 0.
	var float64_type float64 = 66.666666663453453 // default: 0.

	fmt.Println(float32_type, float64_type)

	var complex64_type complex64 = -5 + 22i   // default: 0 + 0i.
	var complex128_type complex128 = 6 - 231i // default: 0 + 0i.

	fmt.Println(complex64_type, complex128_type)

	// To convert between type The expression T(v) is use to converts the value v to the type T.
	var m float64 = 623.3453453476
	var n int = int(m)
	o := uint(m)
	p := rune(n)

	fmt.Println(m, n, o, p)

	// When declaring a variable without specifying an explicit type, the variable's type is inferred from the value on the right hand side.
	var q = o // uint
	r := p    // rune

	fmt.Println(q, r)

	// But when the right hand side contains an untyped numeric constant,
	// the new variable may be an int, float64, or complex128 depending on the precision of the constant:
	s := 42           // int
	t := 3.142        // float64
	u := 0.867 + 0.5i // complex128

	fmt.Println(s, t, u)
}
