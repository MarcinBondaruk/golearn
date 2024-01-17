package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup
	var n int = 0
	wg.Add(gr * 2)

	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			// m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final value of n:", n)
}
