package pattern

import (
	"fmt"
	"time"
)

func broadcast(value int, chans []chan int) {
	for _, ch := range chans {
		ch <- value
	}
}

func worker(id int, ch chan int) {
	fmt.Println("Run worker! id: ", id)

	for job := range ch {
		fmt.Printf("Worker %d received job %d\n", id, job)
	}

	fmt.Println("End worker! id: ", id)
}

func A() {
	numWorkers := 10
	chans := make([]chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		chans[i] = make(chan int)
		go worker(i, chans[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Broadcasting value: %d\n", i)
		broadcast(i, chans)
	}
	time.Sleep(2 * time.Second)
}
