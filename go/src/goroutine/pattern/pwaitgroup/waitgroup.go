package pwaitgroup

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker(id int) {
	defer wg.Done()
	fmt.Println("id: ", id)
}

func A() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
	fmt.Println("Worker finished")
}
