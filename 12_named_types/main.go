package main

import "fmt"

type km float64
type mile float64

func main() {
	type age int // type age, underlying type is int
	type oldAge age
	type veryOldAge age

	// [!] can attach methods to newly defined types

	type speed uint
	var s1 speed = 10
	var s2 speed = 20
	fmt.Println(s2 - s1)

	var x int
	_ = x
	// x = s1 error
	x = int(s1) // no error

	// var s3 speed = x error
	var s3 speed = speed(x)
	_ = s3

	var parisToLondon km = 465
	var distanceInMiles mile

	// distanceInMiles = parisToLondon / 0.621  error, different types
	distanceInMiles = mile(parisToLondon) * 0.621
	fmt.Println(distanceInMiles)
}
