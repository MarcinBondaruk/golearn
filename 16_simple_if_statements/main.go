package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("45x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if args := os.Args; len(args) < 2 {
		fmt.Println("Require at least 1 argument")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("argument must be an intiger. Error:", err)
	} else {
		fmt.Printf("%d km in miles is about %.3f miles\n", km, float64(km)/1.609)
	}
}
