package pselect

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Goroutine task running")
	time.Sleep(2 * time.Second)
	// time.Sleep(5 * time.Second)
	done <- true
}

func A() {
	done := make(chan bool)
	go worker(done)

	select {
	case <-done:
		fmt.Println("Goroutine task completed")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout waiting for goroutine task")
	}
}
