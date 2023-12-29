package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%d\n", i)
			counter++
		}

		if counter == 3 {
			break
		}
	}
	fmt.Println("next exercise")
	for i := 1; i <= 500; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("print years")

	j := 1990
	for j <= time.Now().Year() {
		fmt.Println(j)
		j++
	}

	fmt.Println("fixing code")
	age := 20

	switch {
	case age < 0, age > 100:
		fmt.Println("invalid age")
	case age < 18:
		fmt.Println("You are minor")
	case age == 18:
		fmt.Println("Congratulation you've just become an adult")
	default:
		fmt.Println("You are major")
	}

}
