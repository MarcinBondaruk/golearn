package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int)

	fmt.Printf("%T %v\n", ch, ch)

	c := make(chan int)

	// send operation
	c <- 10

	// recieve
	num := <-c

	fmt.Println(<-c)
	_ = num

	close(c)
}
