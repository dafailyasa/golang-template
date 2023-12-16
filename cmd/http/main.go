package main

import (
	"os"

	"github.com/dafailyasa/golang-template/pkg/factories"
	"github.com/dafailyasa/golang-template/pkg/server"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	factories := factories.NewFactory(
		"./config/config.yaml",
		"logs/log.csv",
	)

	viper := factories.InitializeViper()
	factories.InitializeZapLogger()
	factories.InitializeMongoDB()

	server := server.NewServer(viper)

	err := server.Run()
	if err != nil {
		log.Errorf("Failed to start server. ", err.Error())
		os.Exit(1)
	}
}
