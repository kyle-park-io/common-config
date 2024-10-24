package pmutex

import (
	"fmt"
	"sync"
	"time"
)

var data int
var mu sync.Mutex

func worker() {
	mu.Lock()
	defer mu.Unlock()

	data++
}

func A() {
	for i := 0; i < 10; i++ {
		go worker()
	}
	time.Sleep(time.Second)
	fmt.Println(data)
}
