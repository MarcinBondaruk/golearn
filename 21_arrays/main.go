package main

import (
	"fmt"
	"strings"
)

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
	fmt.Printf("%v\n", a5)
	fmt.Printf("%d\n", len(a5))

	newNumbers := [3]int{10, 20, 30}
	newNumbers[0] = 7

	fmt.Printf("%v\n", newNumbers)

	for i, v := range newNumbers {
		fmt.Printf("%d => %d\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 20))
}
