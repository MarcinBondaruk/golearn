package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func cube(a float64) float64 {
	return math.Pow(a, 3)
}

func f1(a uint) (uint, uint) {
	var factorial uint = 1
	var sum uint = 0

	for i := uint(1); i <= a; i++ {
		factorial *= i
		sum += i
	}

	return factorial, sum
}

func myFunc(n string) int {
	i, _ := strconv.Atoi(n)
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)
	return i + ii + iii
}

// with naked return
func sum(a ...int) (sum int) {
	for _, n := range a {
		sum += n
	}

	return
}

func searchItem(a []string, b string) bool {
	for _, v := range a {
		if strings.EqualFold(b, v) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(cube(2))
	fmt.Println(f1(5))
	fmt.Println(myFunc("8"))
	fmt.Println(myFunc("2"))
	fmt.Println(sum(1, 2, 3))

	fmt.Println(searchItem([]string{"red", "blue", "green"}, "green"))
	fmt.Println(searchItem([]string{"red", "blue", "green"}, "Green"))
}
