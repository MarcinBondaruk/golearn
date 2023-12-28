package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python")
		// no need to add break statement in go
	case "Go", "golang":
		// and statement ^
		fmt.Println("Learning Go")
	default:
		fmt.Println("Whatever.")
	}

	n := 5

	switch true {
	case n%2 == 0:
		fmt.Println("even number")
	case n%2 != 0:
		fmt.Println("odd number")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
