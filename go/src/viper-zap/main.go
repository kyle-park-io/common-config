package main

import (
	"viper-zap/configs"
	"viper-zap/logger"
)

func main() {
	logger.InitLogger()
	logger.Log.Info("hi! i'm kyle!")

	if err := configs.InitConfig(); err != nil {
		logger.Log.Fatalf("Check Errors, %v", err)
	}
}
