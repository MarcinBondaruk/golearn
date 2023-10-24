package main

import "fmt"

func main() {
	figure := "Circle"
	radius := 5
	pi := 3.14

	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("PI is %.3f \n", pi)
	fmt.Printf("Diameter of a %s  with radius %d is %.3f \n", figure, radius, float64(radius)*2*pi) // convert radius to float64 beacuse cant int * float

	// quoted strings %q
	fmt.Printf("this is a %q \n", figure)

	// replaced by any type %v
	fmt.Printf("Diameter of a %v with radius %v is %.3v \n", figure, radius, float64(radius)*2*pi) // convert radius to float64 beacuse cant int * float

	// print type of variable %T
	fmt.Printf("radius is of type %T, and pi if of type %T \n", radius, pi)

	// format boolean %t
	closed := true
	fmt.Printf("door is closed: %t \n", closed)

	// base 2 -> %b
	number := 55
	fmt.Printf("this is 55 in base2: %b \n", number)

	// print number in base2, display 8bits
	fmt.Printf("this is 55 in base2: %08b \n", number)

}
