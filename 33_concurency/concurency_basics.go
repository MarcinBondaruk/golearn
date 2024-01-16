package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1 i=", i)
	}
	fmt.Println("f1 stoped")
}

func f2() {
	fmt.Println("f2 started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2 i=", i)
	}
	fmt.Println("f2 stoped")
}

func main() {
	fmt.Println("Main Execution Started")
	fmt.Println("num of cpus: ", runtime.NumCPU())
	fmt.Println("num of goroutine", runtime.NumGoroutine())

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)

	fmt.Println("Gomaxprocs:", runtime.GOMAXPROCS(0))
	go f1()
	fmt.Println("num of goroutine", runtime.NumGoroutine())

	f2()

	time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped")
}
