package main

import "fmt"

func main() {
	// i exists only in for scope
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("mock while loop below")
	// no while loop but can make it with for
	// j exists outside for loop
	j := 10

	for j >= 0 {
		fmt.Println(j)
		j--
	}

	fmt.Println("continue statemet")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("break statemet")
	// break terminates inner most for loop or switch

	count := 0

	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 {
			break
		}
	}
}
