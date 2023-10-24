package main

import "fmt"

func main() {
	const a float64 = 5.1 // typed const
	const b = 6.7         // untyped const
	const c float64 = a * b
	const str = "Hello " + "Go!"
	const d = 5 > 10

	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x

	const x = 5
	const y = 2.2 * 5

	fmt.Printf("%T \n", y)

	var i int = x
	var j float64 = x
	var p byte = x

	fmt.Println(i, j, p)

	const r = 5
	var rr = r
	fmt.Printf("%T\n", rr)

	// iota - number generator 4 constants

	const (
		c0 = iota // default 0
		c1
		c2
	)

	fmt.Println(c0, c1, c2) // incremeted

	const (
		f = (iota * 2) + 1
		f1
		f2
	)

	fmt.Println(f, f1, f2)

	const (
		x1 = -(iota + 2) // -2
		_                // -3
		x2               // -4
		x3               // -5
	)
	fmt.Println(x1, x2, x3)
}
