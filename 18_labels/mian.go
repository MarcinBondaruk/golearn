package main

import "fmt"

func main() {
	outer := 19
	_ = outer
	people := [5]string{"Donek", "Jarek", "Zero", "Czadownia", "Dudi"}
	friends := [2]string{"Dudi", "Sasin"}

outer:
	for idx, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("foudn a friend %q at index %d\n", friend, idx)
				break outer
			}
		}
	}
	fmt.Println("next instruction after break")
}
