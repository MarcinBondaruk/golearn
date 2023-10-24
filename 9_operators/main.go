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

	// COMPARASION OPERATORS
	c, d := 5, 10

	fmt.Println(c == d) // false
	fmt.Println(c != d) // true
	fmt.Println(c > d)  // false
	fmt.Println(c < d)  // true

	// LOGICAL OPERATORS
	fmt.Println(c > 1 && d == 10)   // true short circuit - second expression not evaluated when first one is false
	fmt.Println(c == 5 || b == 100) // true
	fmt.Println(!(c > 0))           // false

}
