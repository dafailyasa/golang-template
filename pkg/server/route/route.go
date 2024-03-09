package route

import (
	user "github.com/dafailyasa/golang-template/internal/user/domain/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/spf13/viper"
)

type RouteConfig struct {
	App     *fiber.App
	UserHdl user.UserHandlers
	Viper   *viper.Viper
}

func (r *RouteConfig) InitRoute() {
	appName := r.Viper.GetString("APP.NAME")

	// public routes
	r.App.Get("/metrics", monitor.New(monitor.Config{Title: appName + " Metrics"}))
	v1App := r.App.Group("/api/v1")

	r.authRoutes(v1App)
}

func (r *RouteConfig) authRoutes(prefix fiber.Router) {
	auth := prefix.Group("/auth")
	auth.Post("/register", r.UserHdl.Register)
}
