package pattern

import "fmt"

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
