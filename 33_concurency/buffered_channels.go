package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("interation #%d started\n", i)
			c <- i * 5
			fmt.Printf("interation #%d stopped\n", i)
		}
		close(c)
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main goroutine recieved data: ", v)
	}

	fmt.Println(<-c1) // 0 value for closed chanel int
}
