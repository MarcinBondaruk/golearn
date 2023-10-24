package main

import "fmt"

func main() {
	// INT TYPE
	var i1 int8 = -128
	fmt.Printf("%T \n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T \n", i2)

	// FLOAT
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T \n", f1, f2, f3)

	// COMPLEX complex64, complex128

	// RUNE Type
	var r rune = 'f'
	fmt.Printf("%T \n", r)
	fmt.Println(r) // f = 102 in ASCII

	// BOOLEAN
	var b bool = true
	fmt.Printf("%T \n", b)

	// STRING
	var s string = "Hello GO"
	fmt.Printf("%T \n", s)

}
