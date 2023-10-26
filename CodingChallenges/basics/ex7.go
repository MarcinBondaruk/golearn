package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gopher!"
	score := []int{10, 20, 30}

	fmt.Printf("%d %f %q %#v\n", x, y, z, score)
	fmt.Printf("%T %T", y, score)

}
