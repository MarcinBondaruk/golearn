package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 3
	var f float64 = 3.2
	var s1, s2 = "3.14", "5"

	i1 := strconv.Itoa(i)
	s21, _ := strconv.Atoi(s2)
	f1 := fmt.Sprintf("%f", f)
	s11, _ := strconv.ParseFloat(s1, 64)

	fmt.Printf("%T, %T, %T,%T\n", i1, s21, f1, s11)
}
