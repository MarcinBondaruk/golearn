package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1 i=", i)
	}
	fmt.Println("f1 stoped")

	wg.Done()
}

func f2() {
	fmt.Println("f2 started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2 i=", i)
	}
	fmt.Println("f2 stoped")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	fmt.Println("Main Execution Started")
	fmt.Println("num of cpus: ", runtime.NumCPU())
	fmt.Println("num of goroutine", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)

	fmt.Println("Gomaxprocs:", runtime.GOMAXPROCS(0))
	go f1(&wg)
	fmt.Println("num of goroutine", runtime.NumGoroutine())

	f2()
	wg.Wait()
	fmt.Println("main execution stopped")
}
