package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	err = os.Rename("info.txt", "data.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentsInBytes, err := os.ReadFile("songs.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q\n", string(contentsInBytes))
	for _, v := range contentsInBytes {
		fmt.Printf("%q\n", string(v))
	}
}
