package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func task() {
	fmt.Println(rand.Int() * rand.Int())
}

func periodicTask(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			task()
		case <-ctx.Done():
			fmt.Println("Stopping periodic task")
			ticker.Stop()
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go periodicTask(ctx)
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM)
	<-sigCh
}
