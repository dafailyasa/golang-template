package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/viper"
)

type Server struct {
	viper *viper.Viper
}

func NewServer(viper *viper.Viper) *Server {
	return &Server{
		viper: viper,
	}
}

func (s *Server) Run() error {
	s.viper.SetDefault("APP_NAME", "Golang Template👋")
	s.viper.SetDefault("APP_PORT", 3030)

	appName := s.viper.GetString("APP_NAME")
	appPort := s.viper.GetInt("APP_PORT")

	app := fiber.New(fiber.Config{
		AppName:           appName,
		EnablePrintRoutes: true,
	})

	// middlewares
	app.Use(cors.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: appName + " Metrics"}))

	err := app.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		return err
	}

	return nil
}
