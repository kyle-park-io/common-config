package pbasic

import (
	"fmt"
	"time"
)

var count int64

func worker() {
	count += 1
}

func A() {
	for i := 0; i < 10; i++ {
		go worker()
	}
	time.Sleep(time.Second)

	fmt.Println(count)
}
