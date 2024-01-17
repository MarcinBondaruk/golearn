package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, ch chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down\n", url)
		s += fmt.Sprintf("Error: %v", err)
		fmt.Println(s)
		ch <- url
	} else {
		s := fmt.Sprintf("%s : Code: %d\n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP!", url)
		fmt.Println(s)
		ch <- url
	}
}

func main() {
	urls := []string{"https://golang.org", "https://google.com"}
	ch := make(chan string)

	for _, url := range urls {
		go checkUrl(url, ch)
	}

	fmt.Println("num of goroutines:", runtime.NumGoroutine())

	// for {
	// 	go checkUrl(<-ch, ch)
	// 	fmt.Println(strings.Repeat("#", 20))
	// 	time.Sleep(time.Second * 2)
	// }

	for url := range ch {
		time.Sleep(time.Second * 2)
		go checkUrl(url, ch)
	}

}
