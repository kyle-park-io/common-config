package goroutinetasks

import (
	"fmt"
	"goroutine-example/cpubound"
	"sync"
)

func CountPrimes() {
	var wg sync.WaitGroup

	primeCount := 0
	numTasks := 100000

	var mu sync.Mutex
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if cpubound.IsPrime(n) {
				mu.Lock()
				primeCount++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Num of prime: ", primeCount)
}
