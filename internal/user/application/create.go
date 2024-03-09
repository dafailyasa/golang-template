package application

import (
	"time"

	"github.com/dafailyasa/golang-template/internal/user/domain/models"
	customErr "github.com/dafailyasa/golang-template/pkg/custom-errors"
	"github.com/dafailyasa/golang-template/tools"
	"github.com/gofiber/fiber/v2"
)

func (app *UserApp) Create(body *models.RegisterRequest) *fiber.Error {
	if body.Password != body.PasswordConfirmation {
		return fiber.NewError(fiber.StatusUnprocessableEntity, customErr.ErrPasswordConfirmation.Error())
	}

	isExist, _ := app.repo.FindByEmail(body.Email)
	if isExist != nil {
		return fiber.NewError(fiber.StatusConflict, customErr.ErrEmailRegistered.Error())
	}

	user := new(models.User)
	user.Email = body.Email
	user.Password = body.Password
	user.CreatedAt = time.Now()
	user.Password = tools.Hash256Password(body.Password)

	err := app.repo.Create(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, customErr.ErrInternalServer.Error())
	}

	return nil
}
