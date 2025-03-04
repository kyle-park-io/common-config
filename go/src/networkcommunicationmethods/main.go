package main

import (
	"math"
	"time"

	"network/methods"
)

func main() {

	// // run
	methods.RunTCP()
	methods.RunUDP()
	methods.RunHTTP()
	methods.RunHTTP2()
	methods.RunJSONRPC()
	methods.RunWebsocket()

	// wait
	time.Sleep(time.Duration(math.MaxInt64))
}
