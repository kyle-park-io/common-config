package pattern

import (
	"context"
	"fmt"
	"math"
	"time"
)

func forGorountine(id int, ch1 chan int, ch2 chan int) {
	fmt.Printf("Run node! forGorountine_id: %d\n", id)

	channels := []chan int{ch1, ch2}
	for i, ch := range channels {
		go func(j int, c chan int) {
			for job := range c {
				fmt.Printf("Received job%d: %d\n", j, job)

			}
		}(i, ch)
	}

	// fmt.Printf("End node! s_id: %d\n", id)
}

func forSelect(id int, ch1 chan int, ch2 chan int, quit chan bool) {
	fmt.Printf("Run node! forSelect_id: %d\n", id)

forLoop:
	for {
		select {
		case job1 := <-ch1:
			fmt.Println("Received job1: ", job1)
		case job2 := <-ch2:
			fmt.Println("Received job2: ", job2)
		case <-quit:
			break forLoop
		}
	}

	fmt.Printf("End node! forSelect_id: %d\n", id)
}

func forSelectContext(id int, ch1 chan int, ch2 chan int, quit chan bool) {
	fmt.Printf("Run node! forSelectContext_id: %d\n", id)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

forLoop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received cancel signal")
			break forLoop
		case job1 := <-ch1:
			fmt.Println("Received job1:", job1)
		case job2 := <-ch2:
			fmt.Println("Received job2:", job2)
		case <-quit:
			cancel()
		}
	}

	fmt.Printf("End node! forSelectContext_id: %d\n", id)
}

func A() {

	ex1_ch1 := make(chan int)
	ex1_ch2 := make(chan int)

	ex2_ch1 := make(chan int)
	ex2_ch2 := make(chan int)
	ex2_quit := make(chan bool)

	ex3_ch1 := make(chan int)
	ex3_ch2 := make(chan int)
	ex3_quit := make(chan bool)

	for i := 0; i < 1; i++ {
		go forGorountine(i, ex1_ch1, ex1_ch2)
		go forSelect(i, ex2_ch1, ex2_ch2, ex2_quit)
		go forSelectContext(i, ex3_ch1, ex3_ch2, ex3_quit)
	}

	go func() {

		ex1_ch1 <- 1
		ex1_ch2 <- 2
		// close(ex1_ch1)
		// close(ex1_ch2)

		ex2_ch1 <- 3
		ex2_ch2 <- 4
		ex2_quit <- true
		// close(ex2_ch1)
		// close(ex2_ch2)

		ex3_ch1 <- 5
		ex3_ch2 <- 6
		ex3_quit <- true
		// close(ex3_ch1)
		// close(ex3_ch2)
	}()
	time.Sleep(time.Duration(math.MaxInt64))
}
