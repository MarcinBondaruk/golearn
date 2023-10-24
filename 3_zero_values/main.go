package main

import "fmt"

func main() {
	var a, b = 4, 5.7

	// a = b <- compile time error
	a = int(b) // <- rounds down

	fmt.Println(a, b)

	var x int
	// x = "5" <- x is an int

	_ = x

	// no unitialized variables because of zero values
	// int - 0
	// float64 - 0.0
	// string - ""
	// bool - false

	var values int
	var price float64
	var name string
	var done bool

	fmt.Printf("%d, %f, %s, %t", values, price, name, done)
}
