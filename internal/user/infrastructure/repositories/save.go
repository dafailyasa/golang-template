package repositories

import (
	"context"

	models "github.com/dafailyasa/golang-template/internal/user/domain/models"
)

func (repo *UserMongoDB) Create(user *models.User) error {
	_, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		repo.logger.Error("error when save user", err)
		return err
	}

	return nil
}
