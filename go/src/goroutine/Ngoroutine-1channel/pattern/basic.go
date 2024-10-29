package pattern

import (
	"Ngoroutine-1channel/utils"
	"fmt"
	"time"
)

type job struct {
	id       int
	job_type string
	job      string
}

func worker(id int, ch chan job) {
	fmt.Println("Run worker! id: ", id)

	for job := range ch {
		fmt.Printf("Worker %d received job %+v\n", id, job)
	}

	fmt.Println("End worker! id: ", id)
}

func A() {
	ch := make(chan job)
	for i := 0; i < 3; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 3; i++ {
		go func() {
			r := utils.GetRandom(3)
			switch r {
			case 0:
				j := job{id: 0, job_type: "A", job: "A"}
				ch <- j
			case 1:
				j := job{id: 1, job_type: "B", job: "B"}
				ch <- j
			case 2:
				j := job{id: 2, job_type: "C", job: "C"}
				ch <- j
			}
		}()
	}
	time.Sleep(2 * time.Second)
}
