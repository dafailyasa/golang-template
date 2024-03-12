package ports

import "github.com/dafailyasa/golang-template/internal/user/domain/models"

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}
