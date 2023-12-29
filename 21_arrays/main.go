package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{1, 2, 3}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"dan", "diana", "paul", "john"}
	fmt.Printf("%#v\n", a3)

	a4 := [3]string{"asda", "dsa"}
	fmt.Printf("%#v\n", a4)

	// elipsis operator
	a5 := [...]int{1, 2, 3, 5, 6, -20, 123}
	fmt.Printf("%v", len(a5))

}
