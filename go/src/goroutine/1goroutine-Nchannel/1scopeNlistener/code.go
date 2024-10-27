package scopenlistener

import (
	"fmt"
	"time"
)

func worker(ch1 chan int, ch2 chan int) {
	fmt.Println("Run worker!")

	for job := range ch1 {
		fmt.Println(job)
	}
	for job := range ch2 {
		fmt.Println(job)
	}

	fmt.Println("End worker!")
}

func A() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go worker(ch1, ch2)

	go func() {
		go func() {
			for i := 0; i < 3; i++ {
				ch1 <- i
				// ch2 <- i + 3
			}
			// close(ch1)
		}()
		go func() {
			for i := 0; i < 3; i++ {
				// ch1 <- i
				ch2 <- i + 3
			}
		}()
		// will panic
		close(ch1)
	}()
	time.Sleep(2 * time.Second)
}
