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
}
