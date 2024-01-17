package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, ch chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down\n", url)
		s += fmt.Sprintf("Error: %v", err)
		ch <- s
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s : Code: %d\n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			if err != nil {
				log.Fatal(err)
			}

			s += fmt.Sprintf("Writing response body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0644)

			if err != nil {
				s += fmt.Sprintf("Error: %v", err)
				ch <- s
			}

			s += fmt.Sprintf("%s is UP!", url)
			ch <- s
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://google.com"}
	ch := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, ch)
		fmt.Println(strings.Repeat("##", 20))
		fmt.Println(<-ch)
	}

	fmt.Println("num of goroutines:", runtime.NumGoroutine())

}
