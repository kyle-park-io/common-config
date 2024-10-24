package pcontext

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	fmt.Println("Goroutine task running")
	select {
	case <-time.After(2 * time.Second):
		// case <-time.After(5 * time.Second):
		fmt.Println("Goroutine task completed")
	case <-ctx.Done():
		fmt.Println("Goroutine task canceled")
	}
}

func A() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(4 * time.Second)
	fmt.Println("Main function ends")
}
