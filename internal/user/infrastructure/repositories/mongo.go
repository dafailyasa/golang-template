package repositories

import (
	"fmt"

	"github.com/dafailyasa/golang-template/internal/user/domain/ports"
	logger "github.com/dafailyasa/golang-template/pkg/logger/ports"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	logger     logger.LoggerApplication
	viper      *viper.Viper
}

var _ ports.UserRepository = (*UserMongoDB)(nil)

func NewUserMongoDB(logger logger.LoggerApplication, client *mongo.Client, viper *viper.Viper) *UserMongoDB {
	dbName := viper.GetString("MONGO.NAME")
	if dbName == "" {
		logger.Error("Failed to get database name", nil)
		return &UserMongoDB{}
	}

	fmt.Println(client.Database(dbName))

	return &UserMongoDB{
		logger:     logger,
		viper:      viper,
		client:     client,
		database:   client.Database(dbName),
		collection: client.Database(dbName).Collection("users"),
	}
}
