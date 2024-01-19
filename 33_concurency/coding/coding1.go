package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("Hello")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
}
