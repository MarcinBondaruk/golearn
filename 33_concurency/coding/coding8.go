package main

import (
	"fmt"
	"math"
)

func power(p int, ch chan int) {
	ch <- int(math.Pow(float64(p), 2))
}

func main() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}
}
