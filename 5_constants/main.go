package main

import "fmt"

func main() {
	const days int = 7
	var i int

	fmt.Println(i)

	// const pi float64 - const must be initialized!
	const pi float64 = 3.1415

	x, y := 5, 0

	_, _ = x, y

	//  fmt.Println(x / y) <- runtime error

	const a int = 5
	const b int = 0

	// fmt.Println(a / b) compile time error since b is constant

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	// multi declaration
	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	// multi declaration with same value
	const (
		min4 = -500
		min5
		min6
	)

	fmt.Println(min4, min5, min6)

	/*
		const rules
		1. cant change constant
		cont x int = 10
		x = 20 <- error

		2. cant initialize const at runtime
		const power = math.Pow(2,3) <- error

		3. cant use variable to init constant
		t := 5
		const tc = t <- error

		4.
		const ln = len("hello") no error, len is available at compile time

		variables belong to runtime!

	*/
}
