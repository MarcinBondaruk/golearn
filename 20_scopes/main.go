package main

// imports are file scoped
import "fmt"

// consts are package scoped
const done = false

// func is package scoped
func main() {
	fmt.Println("elo")
	// local scoped
	const done = true

	//local/block scoped
	n, m := 1, 2

	_, _ = n, m
}
