package main

import "fmt"

// pass by value
func change(a *int) *float64 {
	*a = 100
	b := 6.7

	return &b
}

func changeValues(q int, p float64, n string, s bool) {
	q = 512
	p = 24.99
	n = "Traktor Henio"
	s = true
}

func changeValuesByPointer(q *int, p *float64, n *string, s *bool) {
	*q = 512
	*p = 24.99
	*n = "Traktor Henio"
	*s = true
}

type Product struct {
	id    string
	price float64
}

func changeProduct(p *Product) {
	(*p).id = "bq-23-35"
	p.price = 256.99
}

func changeSlice(s []int) {
	for i, _ := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

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

	a := 5.5
	ptrA := &a
	pPtrA := &ptrA

	fmt.Printf("value of ptrA: %v, address po ptrA: %p\n", ptrA, &ptrA)
	fmt.Printf("value of pPtrA: %v, address po pPtrA: %p\n", pPtrA, &pPtrA)
	fmt.Println(*pPtrA, *ptrA, **pPtrA)

	**pPtrA++

	fmt.Println(a)

	// Comparin pointers
	// pointers are equal if they point to same address or are nil

	var ptrB *int
	var ptrC *int

	fmt.Println(ptrB == ptrC) // true

	c := 5

	ptrB = &c
	ptrC = &c

	fmt.Println(ptrB == ptrC) // true

	d := 5
	ptrC = &d
	fmt.Println(ptrB == ptrC) // false

	q := 8
	w := &q

	fmt.Printf("value of q before calling \"change\": %v\n", q)
	change(w) // copy of this pointer is created but value of the pointer is unchanged. Original, unchanged pointer and new pointer copy (pass by value) point to the same memory address
	fmt.Printf("value of q after calling \"change\": %v\n", q)

	// pass by value, changeValues(...) didnt change original values, only copies that are limited to its scope
	quantity, price, name, sold := 256, 19.99, "opaska", false
	fmt.Println("before changeValues", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("after changeValues", quantity, price, name, sold)

	fmt.Println("before changeValuesByPointer", quantity, price, name, sold)
	// pass a pointer = pass address of variable x notation: "&x"
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("after change values", quantity, price, name, sold)

	originalProduct := Product{
		id:    "po-11-l2",
		price: 99.99,
	}

	changeProduct(&originalProduct)

	fmt.Println(originalProduct)

	prices := []int{1, 2, 3}
	changeSlice(prices) // function modifies the slice!
	fmt.Println("slice after calling changeSlice(): ", prices)

	myMap := map[string]int{"a": 9, "b": 8, "c": 7}
	changeMap(myMap) // function will modify map
	fmt.Println("map after calling changeMap(): ", myMap)

}
