package config

import (
	"fmt"
	"runtime"
)

func CPUInfo() {
	// num of cpu
	numCPU := runtime.NumCPU()
	fmt.Printf("noc: %d\n", numCPU)
	fmt.Printf("nocg: %d\n", runtime.NumCgoCall())
	fmt.Printf("nog: %d\n", runtime.NumGoroutine())
}

func SetGOMAXPROCS() {
	// set GOMAXPROCS
	num := 1 // single core
	prev := runtime.GOMAXPROCS(num)
	fmt.Printf("Previous GOMAXPROCS: %d\n", prev)
	fmt.Printf("Current GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
