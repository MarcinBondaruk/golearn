package main

import "fmt"

func f() {}

func main() {
	// ARRAY
	var numbers = [4]int{5, 6, -2, 100}
	fmt.Printf("%T \n", numbers) // [4]int - 4 element int array

	// SLICE
	var cities = []string{"London", "Tokio", "NY"}
	fmt.Printf("%T \n", cities) // []string - slice of strings

	// MAP
	balances := map[string]float64{
		"USD": 2332.1,
		"EUR": 71223.2,
	}
	fmt.Printf("%T \n", balances) // map[string]float64

	// STRUCT
	type Person struct {
		name string
		age  int
	}

	var me Person
	fmt.Printf("%T \n", me)

	// POINTERS
	var x int = 2
	ptr := &x // address of x

	fmt.Printf("%T %v \n", ptr, ptr) // type *int (pointer to memory holding integer)

	// FUNCTION TYPE
	fmt.Printf("%T \n", f) // func()

}
