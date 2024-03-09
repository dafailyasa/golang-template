package main

import (
	"os"

	"github.com/dafailyasa/golang-template/pkg/constants"
	"github.com/dafailyasa/golang-template/pkg/factories"
	"github.com/dafailyasa/golang-template/pkg/server"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	factories := factories.NewFactory(
		constants.ConfigPath,
		constants.LogPath,
	)

	viper := factories.InitializeViper()
	factories.InitializeZapLogger()
	factories.InitializeMongoDB()

	userHdl := factories.BuildUserHandler()

	server := server.NewServer(viper, userHdl)

	err := server.Run()
	if err != nil {
		log.Errorf("Failed to start server. ", err.Error())
		os.Exit(1)
	}
}
