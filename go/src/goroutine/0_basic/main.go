package main

import (
	"goroutine-example/config"
	goroutinetasks "goroutine-example/goroutine_tasks"
	"goroutine-example/utils"
)

func main() {
	// config
	config.CPUInfo()
	config.SetGOMAXPROCS()

	// start
	utils.SetStart()

	// body
	goroutinetasks.CountPrimes()

	// end
	utils.SetEnd()
}
