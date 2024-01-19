package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("sum: %.2f\n", a+b)
	wg.Done()
}

func main() {
	data := [][]float64{{1.0, 2.0}, {3.0, 4.0}, {5.0, 6.0}}
	var wg sync.WaitGroup

	wg.Add(len(data))
	for i := 0; i < len(data); i++ {
		go sum(data[i][0], data[i][1], &wg)
	}

	wg.Wait()
}
