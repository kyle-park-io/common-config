package main

import "zap-logger/logger"

func main() {
	logger.InitLogger()
	logger.Log.Info("hi! i'm kyle!")
}
