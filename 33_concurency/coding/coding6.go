package main

func main() {
	var ch chan float64
	ch = make(chan float64)

	// recieve only
	c2 := make(<-chan float64)

	// send only
	c3 := make(chan<- float64)

	c4 := make(chan float64, 10)
}
