package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is simple Echo CLI. Type anything and I will repeat")
	fmt.Println("Type !exit to stop the program")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "!exit" {
			break
		}

		fmt.Println("Echo: ", input)
	}
	fmt.Println("Bye bye")
}
