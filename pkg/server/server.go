package server

import (
	"fmt"

	user "github.com/dafailyasa/golang-template/internal/user/domain/ports"
	"github.com/dafailyasa/golang-template/pkg/server/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	routeConfig := route.RouteConfig{
		App:     app,
		UserHdl: s.userHdl,
		Viper:   s.viper,
	}

	routeConfig.InitRoute()

	err := app.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		return err
	}

	return nil
}
