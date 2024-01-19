package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	const n = 50

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 100; i < 150; i++ {
		go func(a float64) {
			fmt.Println(math.Sqrt(a))
			wg.Done()
		}(float64(i))
	}

	wg.Wait()
}
