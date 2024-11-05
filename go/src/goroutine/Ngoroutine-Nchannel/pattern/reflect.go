package pattern

import (
	"Ngoroutine-Nchannel/utils"
	"fmt"
	"reflect"
	"time"
)

func worker2(id int, shared_channel chan int, individual_channel chan int) {
	fmt.Println("Run worker! id: ", id)

	for job := range shared_channel {
		fmt.Printf("Worker %d received job %d\n", id, job)
		individual_channel <- job
	}

	fmt.Println("End worker! id: ", id)
}

func B() {
	numGoroutines := 10
	shared_channel := make(chan int)
	individual_channel := make([]chan int, numGoroutines)
	cases := make([]reflect.SelectCase, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		individual_channel[i] = make(chan int)
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(individual_channel[i])}
		go worker2(i, shared_channel, individual_channel[i])
	}

	go func() {
		for {
			r := utils.GetRandom(1000)
			shared_channel <- r
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		chosen, value, ok := reflect.Select(cases)
		if ok {
			fmt.Printf("Received: %d from channel %d\n", value.Int(), chosen)
		} else {
			fmt.Printf("Channel %d closed\n", chosen)
		}
	}
}
