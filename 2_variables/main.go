package main

import "fmt"

func main() {
	var age int = 33
	fmt.Printf("Age: %d \n", age)

	var name = "Martin"
	_ = name // blackhole

	s := "Learning Golang"
	fmt.Println(s)

	car, cost := "Audi", 5000
	fmt.Println(car, cost)
	car, year := "BWM", 2023
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstname string
		gender    bool
	)

	fmt.Println(salary, firstname, gender)

	var a, b, c int

	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	fmt.Println(i, j)

	j, i = i, j // swapping
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

}
