package main

import "fmt"

func main() {
	const secPerDay = 24 * 60 * 60
	const daysInYear = 365

	fmt.Printf("%d\n", secPerDay*daysInYear)
}
