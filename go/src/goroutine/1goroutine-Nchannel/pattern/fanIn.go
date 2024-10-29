package pattern

import (
	"fmt"
	"math"
	"time"
)

// Make sure close the channel because the sender(out channel) is located
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	done := make(chan struct{})

	for i, ch := range channels {
		go func(j int, c <-chan int) {
			for val := range c {
				fmt.Printf("Received job%d: %d\n", j, val)
				// pattern
				out <- val
			}
			// pattern
			done <- struct{}{}
		}(i, ch)
	}

	go func() {
		for i := 0; i < len(channels); i++ {
			// Wait until data arrives on the done channel(unbufferd).
			// pattern
			<-done
		}
		// pattern
		close(out)
	}()

	return out
}

func worker(id int, ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- id*10 + i
		time.Sleep(500 * time.Millisecond)
	}
	// pattern
	close(ch)
}

func B() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go worker(1, ch1)
	go worker(2, ch2)
	go worker(3, ch3)

	merged := fanIn(ch1, ch2, ch3)
	for val := range merged {
		fmt.Printf("Received on merged channel: %d\n", val)
	}
	time.Sleep(time.Duration(math.MaxInt64))
}
