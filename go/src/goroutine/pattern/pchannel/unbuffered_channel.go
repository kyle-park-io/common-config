package pchannel

import (
	"fmt"
	"time"
)

func basicUnbufferd() {
	ch := make(chan int)
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

func u_worker(id int, ch chan int) {
	fmt.Println("Run worker! id: ", id)

	for job := range ch {
		fmt.Printf("Worker %d received job %d\n", id, job)

		// time.Sleep(1 * time.Second)
		// time.Sleep(time.Duration(id) * time.Second)
		// break
	}
}
func u_worker2(id int, ch chan int) {
	fmt.Println("Run standby-status-worker! id: ", id)
	job := <-ch
	fmt.Printf("Worker %d received job %d\n", id, job)
	for {
		fmt.Println("Repeat")
		time.Sleep(1 * time.Second)
	}
}
func u_worker3(id int) {
	fmt.Println("Run simply-repeating-worker! id: ", id)
	for {
		fmt.Println("Repeat")
		time.Sleep(1 * time.Second)
	}
}

func A() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go u_worker(i, ch)
	}
	// go u_worker2(6, ch)
	// go u_worker3(7)
	time.Sleep(1 * time.Second)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Scope for sending values into a channel id: ", i)
			ch <- i
		}
	}()
	fmt.Println("main scope!")
	time.Sleep(100 * time.Second)
}
