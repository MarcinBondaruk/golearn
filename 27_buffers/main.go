package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	succ := scanner.Scan()
	if succ == false {
		err = scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Reached EOF")
	}

	fmt.Println("First Line: ", scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
