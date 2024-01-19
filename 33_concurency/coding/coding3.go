package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	a := 3.14
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a float64) {
		fmt.Println(math.Sqrt(a))
		wg.Done()
	}(a)

	wg.Wait()
}
