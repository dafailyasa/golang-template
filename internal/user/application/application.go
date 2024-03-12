package application

import (
	user "github.com/dafailyasa/golang-template/internal/user/domain/ports"
	logger "github.com/dafailyasa/golang-template/pkg/logger/ports"
	"github.com/spf13/viper"
)

type UserApp struct {
	repo   user.UserRepository
	logger logger.LoggerApplication
	viper  viper.Viper
}

func NewUserApp(
	repo user.UserRepository,
	logger logger.LoggerApplication,
	viper viper.Viper,
) *UserApp {
	return &UserApp{
		repo:   repo,
		logger: logger,
		viper:  viper,
	}
}
