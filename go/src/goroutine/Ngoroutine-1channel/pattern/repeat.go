package pattern

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type done struct {
	id   int
	send bool
}

var count int64
var mu sync.Mutex
var arr [10]bool

func worker2(id int, ch chan int, ch_main chan done) {
	for job := range ch {
		mu.Lock()
		if arr[id] {
			mu.Unlock()
			continue
		}

		if job == id {
			fmt.Printf("Worker %d received %d\n", id, job)
			arr[id] = true
			ch_main <- done{id: id, send: true}
		} else {
			ch_main <- done{id: id, send: false}
		}
		mu.Unlock()
	}
}

func B() {
	ch := make(chan int)
	ch_main := make(chan done)
	for i := 0; i < 10; i++ {
		go worker2(i, ch, ch_main)
	}

	go func() {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
		go func() {
			for job := range ch_main {
				if !job.send {
					ch <- job.id
				} else {
					atomic.AddInt64(&count, 1)
					if count == 10 {
						fmt.Println("All data has arrived!")
						break
					}
				}
			}
		}()
	}()
	time.Sleep(time.Duration(math.MaxInt64))
}
