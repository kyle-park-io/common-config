package utils

import (
	"fmt"
	"runtime"
	"time"
)

type Duration struct {
	StartTime   time.Time
	EndTime     time.Time
	ElapsedTime time.Duration
}

var D Duration

func SetStart() {
	D.StartTime = time.Now()
}

func SetEnd() {
	D.EndTime = time.Now()
	D.ElapsedTime = time.Since(D.StartTime)

	fmt.Printf("%+v\n", D)
	fmt.Printf("Total time taken with GOMAXPROCS(%d): %v\n", runtime.NumCPU(), D.ElapsedTime)
}
