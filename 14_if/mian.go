package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("too expensive")
	}

	if price <= 100 && inStock {
		fmt.Println("buy")
	}

	// if price {
	// 	fmt.Println("sth")
	// }

	if price < 100 {
		fmt.Println("cheap")
	} else if price == 100 {
		fmt.Println("on the edge")
	} else {
		fmt.Println("too expensive")
	}

	age := 7

	if age >= 0 && age < 18 {
		fmt.Printf("you cannot vote. please return in %d years\n", 18-age)
	} else if age == 18 {
		fmt.Printf("its your first vote\n")
	} else if age > 18 && age < 100 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("invalid age")
	}
}
