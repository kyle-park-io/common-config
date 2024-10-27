package patomic

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	x int64
	y int64
}

var count int64
var counter Counter

func worker() {
	atomic.AddInt64(&count, 1)

	atomic.AddInt64(&counter.x, 1)
	atomic.AddInt64(&counter.y, 1)
}

func A() {
	for i := 0; i < 10; i++ {
		go worker()
	}
	time.Sleep(time.Second)

	fmt.Println(atomic.LoadInt64(&count))
	fmt.Println("Possibility that councurrency may be violated: ", counter.x, counter.y)
}
