package main

import "fmt"

func main() {
	name := "Marcin"

	// get address
	fmt.Println(&name)

	var x int = 2
	// pointer to int
	ptr := &x

	fmt.Printf("%T, %v and address of a pointer %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	// uninitialized pointer with 0-value = nil
	var ptr1 *float64
	_ = ptr1

	p := new(int)

	x = 100
	p = &x
	fmt.Printf("%T, %v\n", p, p)

	// override value of a memory cell that p point to. X value should change 100 -> 90
	*p = 90

	fmt.Println(x, *p)

	*p = 10
	*p = *p / 2
	fmt.Println(x) // should be 5
}
