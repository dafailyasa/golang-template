package handlers

import (
	user "github.com/dafailyasa/golang-template/internal/user/domain/ports"
	logger "github.com/dafailyasa/golang-template/pkg/logger/ports"
)

type UserHdl struct {
	app    user.UserApplication
	logger logger.LoggerApplication
}

var _ user.UserHandlers = (*UserHdl)(nil)

func NewUserHandler(app user.UserApplication, logger logger.LoggerApplication) *UserHdl {
	return &UserHdl{
		app:    app,
		logger: logger,
	}
}
