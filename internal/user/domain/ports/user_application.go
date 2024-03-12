package ports

import (
	"github.com/dafailyasa/golang-template/internal/user/domain/models"
	"github.com/gofiber/fiber/v2"
)

type UserApplication interface {
	Create(registerRequest *models.RegisterRequest) *fiber.Error
	//Login(credentials *models.AuthRequest) (authToken *models.AuthResponse, err error)
}
