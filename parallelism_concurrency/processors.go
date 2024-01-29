package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

const (
	numberRange     = 1 * 1000 * 1000
	goroutineNumber = 4
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func findPrimesInRange(start, end int) []int {
	var primes []int

	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func findInRange(wg *sync.WaitGroup) {
	start := 0
	end := numberRange

	for i := 0; i < goroutineNumber; i++ {
		go func(i int) {
			fmt.Println(findPrimesInRange(start, end))
			wg.Done()
		}(i)
		start += numberRange
		end += numberRange
	}
}

func main() {
	file, _ := os.Create("trace-8core.out")
	trace.Start(file)
	defer trace.Stop()

	runtime.GOMAXPROCS(8)

	wg := sync.WaitGroup{}
	wg.Add(goroutineNumber)
	findInRange(&wg)
	wg.Wait()
}
