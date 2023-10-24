package main

import "fmt"

func main() {
	// ARYTHMETIC
	a, b := 4, 2
	r := (a + b) / (a - b) * 2

	fmt.Println(r) // 6

	r = 9 % a
	fmt.Println(r) // 1

	// ASSIGNMENT

	a += b
	fmt.Println(a) // 6

	b -= 2
	fmt.Println(b) // 0

	b += 1
	b *= 10
	fmt.Println(b) // 10

	b /= 5
	fmt.Println(b) // 2

	// INCREMENT, DECREMENT
	x := 1
	x += 1 // increment assignments
	x++    // increment statement
	fmt.Println(x)

	x-- // decrement statement
	fmt.Println(x)
	// fmt.Println(5 + x--) illegal expression
}
