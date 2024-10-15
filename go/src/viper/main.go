package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	appName := viper.GetString("app.name")
	appVersion := viper.GetString("app.version")
	serverPort := viper.GetInt("server.port")
	serverDebug := viper.GetBool("server.debug")

	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("App Version: %s\n", appVersion)
	fmt.Printf("Server Port: %d\n", serverPort)
	fmt.Printf("Debug Mode: %v\n", serverDebug)
}

func bingConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("can't read config: %v", err)
	}

	viper.BindEnv("db.DB_USER", "db.DB_USER")
	viper.BindEnv("db.DB_PASS", "db.DB_PASS")
	viper.BindEnv("DB_NAME")

	fmt.Println("DB User:", viper.GetString("db.DB_USER"))
	fmt.Println("DB Pass:", viper.GetString("db.DB_PASS"))
	fmt.Println("DB Name:", viper.GetString("DB_NAME"))
}

func setDefault() {
	viper.SetDefault("DB_HOST", "localhost")

	dbHost := viper.GetString("DB_HOST")
	fmt.Printf("Database Host: %s\n", dbHost)
}

func main() {

	initConfig()
	bingConfig()
	setDefault()

}
