package handlers

import (
	"github.com/dafailyasa/golang-template/internal/user/domain/models"
	"github.com/dafailyasa/golang-template/tools"
	"github.com/gofiber/fiber/v2"
)

func (hdl *UserHdl) Register(ctx *fiber.Ctx) error {
	body := new(models.RegisterRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			tools.ApiErrorResponse(fiber.StatusBadRequest, "failed", nil),
		)
	}

	if errs := tools.Validate(body); len(errs) > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			tools.ApiErrorResponse(fiber.StatusUnprocessableEntity, "failed", errs),
		)
	}

	err := hdl.app.Create(body)
	if err != nil {
		return ctx.Status(err.Code).JSON(
			tools.ApiErrorResponse(err.Code, "failed", err.Message),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		tools.ApiResponse(fiber.StatusCreated, "success", nil),
	)
}
