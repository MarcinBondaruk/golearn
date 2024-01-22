package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan string)

	if len(os.Args) >= 2 {
		go func(val string) {
			ch <- val
		}(os.Args[1])
	} else {
		return
	}

	v := <-ch
	fmt.Println(v)
}
