package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)

	// channel only for recieving
	c1 := make(<-chan string)

	// channel only for sending
	c2 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)

	x := <-c

	fmt.Println("recieved value: ", x)
	fmt.Println("exiting main")
}
