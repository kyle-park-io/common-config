package nestedlisteners

import (
	"fmt"
	"time"
)

func worker(ch1 chan int, ch2 chan int) {
	fmt.Println("Run worker!")

	for job1 := range ch1 {
		fmt.Printf("layer1: %d\n", job1)

		for job2 := range ch2 {
			fmt.Printf("layer2: %d\n", job2)
		}
	}

	fmt.Println("End worker!")
}

func A() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int)
	go worker(ch1, ch2)

	go func() {
		go func() {
			for i := 0; i < 3; i++ {
				ch1 <- i
				ch2 <- i + 3
			}
			close(ch2)
			// close(ch1)
		}()
	}()
	time.Sleep(2 * time.Second)
}
