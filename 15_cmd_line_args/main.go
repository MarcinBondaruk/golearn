package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// args are alwasy strings!
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("Path:", os.Args[1])
	fmt.Println("numb of args:", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(result)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	_ = err
}
