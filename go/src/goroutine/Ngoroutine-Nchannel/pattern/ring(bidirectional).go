package pattern

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
)

// 2-node circular structure
// warning: depending on the processing order, each sends values to its respective channel,
// which can lead to a deadlock state where the sent values cannot be received.
func Ring1() {
	// 1 goroutine: ch2 <- val
	go func() {
		for {
			fmt.Println("Check the start of 1goroutine scope")
			val := <-ch1
			fmt.Println("Received from ch1:", val)
			// warning
			ch2 <- val
		}
	}()

	// 2 goroutine: ch1 <- val
	go func() {
		for {
			fmt.Println("Check the start of 2goroutine scope")
			val := <-ch2
			fmt.Println("Received from ch2:", val)
			// warning
			ch1 <- val
		}
	}()

	// init
	ch1 <- 1
	ch1 <- 2

	// wait
	select {}
}

// Augment1 - buffered channel
func Ring2() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	// 1 goroutine: ch2 <- val
	go func() {
		for {
			fmt.Println("Check the start of 1goroutine scope")
			val := <-ch1
			fmt.Println("Received from ch1:", val)
			ch2 <- val
		}
	}()

	// 2 goroutine: ch1 <- val
	go func() {
		for {
			fmt.Println("Check the start of 2goroutine scope")
			val := <-ch2
			fmt.Println("Received from ch2:", val)
			ch1 <- val
		}
	}()

	// init
	ch1 <- 1
	ch1 <- 2

	// wait
	select {}
}

// Augment2 - timeout
// release blocking using a timeout and a queue
func Ring3() {
	type job struct {
		ch  int
		val int
	}
	queue := make([]job, 0)

	// 1 goroutine: ch2 <- val
	go func() {
		for {
			val := <-ch1
			fmt.Println("Received from ch1:", val)

			select {
			case ch2 <- val:
			case <-time.After(3 * time.Second):
				fmt.Println("Timeout: could not send to ch2\ntry to release blocking!")
				queue = append(queue, job{ch: 2, val: val})
			}
		}
	}()

	// 2 goroutine: ch1 <- val
	go func() {
		for {
			val := <-ch2
			fmt.Println("Received from ch2:", val)

			select {
			case ch1 <- val:
			case <-time.After(3 * time.Second):
				fmt.Println("Timeout: could not send to ch1\ntry to release blocking!")
				queue = append(queue, job{ch: 1, val: val})
			}
		}
	}()

	// 3 goroutine: The queue retransmits values to the channel
	go func() {
		i := 0
		for {
			if len(queue) != 0 {
				fmt.Printf("Queue: %+v\n", queue)
				fmt.Println("Try to send value to channel!")

				switch queue[i].ch {
				case 1, 2:
					select {
					case getChannel(queue[i].ch) <- queue[i].val:
						queue = append(queue[:i], queue[i+1:]...)
						i = 0
					case <-time.After(3 * time.Second):
						i++
					}
				}
				if i >= len(queue) {
					i = 0
				}
			} else {
				time.Sleep(3 * time.Second)
			}
		}
	}()

	// init
	ch1 <- 1
	ch1 <- 2

	// wait
	select {}
}
func getChannel(ch int) chan int {
	switch ch {
	case 1:
		return ch1
	case 2:
		return ch2
	default:
		return nil
	}
}
