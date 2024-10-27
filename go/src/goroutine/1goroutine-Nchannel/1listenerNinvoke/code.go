package listenerninvoke

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	fmt.Println("Run worker!")

	for job := range ch {
		fmt.Println(job)
	}

	fmt.Println("End worker!")
}

func A() {
	ch := make(chan int)
	go worker(ch)
	// for i := 0; i < 2; i++ {
	// 	go worker(ch)
	// }

	// for i := 0; i < 3; i++ {
	// 	ch <- i
	// }
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(2 * time.Second)

	ch = make(chan int)
	// important
	go worker(ch)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	time.Sleep(2 * time.Second)
}
