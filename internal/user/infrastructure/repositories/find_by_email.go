package repositories

import (
	"context"

	"github.com/dafailyasa/golang-template/internal/user/domain/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *UserMongoDB) FindByEmail(email string) (*models.User, error) {
	user := new(models.User)

	res := repo.collection.FindOne(context.Background(), bson.D{
		{Key: "email", Value: email},
	})

	if res.Err() != nil {
		return user, res.Err()
	}

	err := res.Decode(user)
	if err != nil {
		return user, nil
	}

	return user, nil
}
