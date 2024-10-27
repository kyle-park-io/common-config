package cpubound

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func BasicFor() {
	for i := 0; i < 1e9; i++ {
	}
	fmt.Println("For 0~1e9")
}

func Mu() {
	var mu sync.Mutex

	mu.Lock()
	mu.Unlock()
}

func Channel() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	val := <-ch
	fmt.Println(val)
}

// cpu + io bound
func Context() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Completed")
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	}
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
