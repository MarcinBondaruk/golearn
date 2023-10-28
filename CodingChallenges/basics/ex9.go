package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	i1 := float64(i)
	f1 := int(f)
	fmt.Printf("%T, %f\n", i1, i1)
	fmt.Printf("%T, %d\n", f1, f1)
}
