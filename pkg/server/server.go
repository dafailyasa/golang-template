package server

import (
	"fmt"

	user "github.com/dafailyasa/golang-template/internal/user/domain/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/viper"
)

type Server struct {
	viper   *viper.Viper
	userHdl user.UserHandlers
}

func NewServer(viper *viper.Viper, user user.UserHandlers) *Server {
	return &Server{
		viper:   viper,
		userHdl: user,
	}
}

func (s *Server) Run() error {
	s.viper.SetDefault("APP.NAME", "Golang Template👋")
	s.viper.SetDefault("APP.PORT", 3030)

	appName := s.viper.GetString("APP.NAME")
	appPort := s.viper.GetInt("APP.PORT")

	app := fiber.New(fiber.Config{
		AppName:           appName,
		EnablePrintRoutes: true,
	})

	// middlewares
	app.Use(cors.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: appName + " Metrics"}))

	v1Api := app.Group("/api/v1")

	v1Api.Post("/register", s.userHdl.Register)

	err := app.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		return err
	}

	return nil
}
