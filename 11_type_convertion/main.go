package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3   // type inference - int
	var y = 3.1 // float

	// x = x * y // error

	x = x * int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)
	fmt.Printf("%v\n", x)

	var a float32 = 1.2
	var b float64 = .1

	_, _ = a, b
	// x = x * y  type mismatch

	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 35122)
	fmt.Println(myStr1)
	fmt.Println(string(35122))

	s1 := "3.123"
	fmt.Printf("%T \n", s1) // string
	var f1, _ = strconv.ParseFloat(s1, 64)
	fmt.Printf("%T \n", f1) // string

	i, _ := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)
	fmt.Printf("%T %v \n", i, i)
	fmt.Printf("%T %q \n", s2, s2)
}
