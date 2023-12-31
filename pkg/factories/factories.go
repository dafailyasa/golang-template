package factories

import (
	"context"
	"log"
	"time"

	customErr "github.com/dafailyasa/golang-template/pkg/custom-errors"
	loggerApp "github.com/dafailyasa/golang-template/pkg/logger/application"
	loggerRepo "github.com/dafailyasa/golang-template/pkg/logger/infrastructure/repositories"
	"github.com/spf13/viper"
	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MongoClientTimeout = 10

type Factory struct {
	logFilePath     string
	configFilePath  string
	viper           *viper.Viper
	dbClient        *mongo.Client
	logger          *loggerApp.Logger
	httpRateLimiter throttled.HTTPRateLimiterCtx
}

func NewFactory(configFilePath string, logFilePath string) *Factory {
	return &Factory{
		configFilePath: configFilePath,
		logFilePath:    logFilePath,
	}
}

func (f *Factory) InitializeViper() *viper.Viper {
	cgf := viper.New()

	cgf.SetConfigFile(f.configFilePath)
	cgf.SetConfigType("yaml")

	if err := cgf.ReadInConfig(); err != nil {
		log.Fatal("Fatal error reading config file. ", err)
	}

	f.viper = cgf
	return cgf
}

func (f *Factory) InitializeZapLogger() *loggerApp.Logger {
	if f.logger != nil {
		return f.logger
	}

	path := f.logFilePath
	repo := loggerRepo.NewCSVFile(path)
	logger := loggerApp.NewLogger(repo)
	f.logger = logger

	return logger
}

func (f *Factory) InitializeMongoDB() *mongo.Client {
	if f.dbClient != nil {
		return f.dbClient
	}

	uri := f.viper.GetString("MONGO_URI")
	if uri == "" {
		log.Fatal(customErr.ErrMongoUrlRequired)
	}

	ctx, cancel := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	f.dbClient = client
	return client
}

func (f *Factory) InitializeThrottled() throttled.HTTPRateLimiterCtx {
	store, err := memstore.NewCtx(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(20),
		MaxBurst: 5,
	}
	rateLimiter, err := throttled.NewGCRARateLimiterCtx(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiterCtx{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	f.httpRateLimiter = httpRateLimiter

	return httpRateLimiter
}
