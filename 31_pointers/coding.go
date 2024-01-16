package main

import "fmt"

func switchValues(a *float64, b *float64) {
	*a, *b = *b, *a
}

func main() {
	x := 10.10

	fmt.Println(&x)
	ptr := &x
	fmt.Printf("%T, %v\n", ptr, ptr)
	fmt.Println(&ptr, *ptr)

	m, n := 10, 2
	ptrM, ptrN := &m, &n
	z := (*ptrM) / (*ptrN)

	fmt.Println(z)

	q, w := 5.5, 8.8
	switchValues(&q, &w)
	fmt.Println(q, w)
}
