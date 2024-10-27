package pchannel

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func basicBufferd() {
	ch := make(chan int, 10)
	for i := 0; i < 5; i++ {
		ch <- 1
	}

	// recommend
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// close(ch)
	}()
}

var count int64
var wg sync.WaitGroup

func b_worker(id int, ch chan int, done chan bool) {
	fmt.Println("Run worker! id: ", id)

	for job := range ch {
		fmt.Printf("Worker %d received job %d\n", id, job)

		// time.Sleep(1 * time.Second)
		// time.Sleep(time.Duration(id) * time.Second)
		// break

		atomic.AddInt64(&count, 1)
		if count == 10 {
			fmt.Println("Received all the data!")
			done <- true
		}
	}

	fmt.Println("End worker! id: ", id)
}
func b_worker2(id int, ch chan int) {
	fmt.Println("Run standby-status-worker! id: ", id)
	job := <-ch
	fmt.Printf("Worker %d received job %d\n", id, job)
	for {
		fmt.Println("Repeat")
		time.Sleep(1 * time.Second)
	}
}
func b_worker3(id int) {
	fmt.Println("Run simply-repeating-worker! id: ", id)
	for {
		fmt.Println("Repeat")
		time.Sleep(1 * time.Second)
	}
}

func B() {
	ch := make(chan int, 30)
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go b_worker(i, ch, done)
	}
	// go b_worker2(6, ch)
	// go b_worker3(7)
	time.Sleep(1 * time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Scope for sending values into a channel id: ", i)
			ch <- i
		}
		// close or done channel pattern
		close(ch)
		_ = <-done
		close(done)
	}()
	fmt.Println("main scope!")
	time.Sleep(100 * time.Second)
}
