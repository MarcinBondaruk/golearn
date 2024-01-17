package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s : Code: %d\n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Writing response body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0644)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://google.com", "https://google.com/a.html"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("##", 20))
	}

	fmt.Println("num of goroutines:", runtime.NumGoroutine())
	wg.Wait()
}
