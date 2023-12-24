package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	loggerApp "github.com/dafailyasa/golang-template/pkg/logger/application"
	"github.com/spf13/viper"
)

func NewKafkaConsumer(viper *viper.Viper, logger *loggerApp.Logger) *kafka.Consumer {
	kafkaCgf := &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          viper.GetString("KAFKA_GROUP_ID"),
		"auto.offset.reset": viper.GetString("KAFKA_AUTO_OFFSET_RESET"),
	}

	c, err := kafka.NewConsumer(kafkaCgf)
	if err != nil {
		logger.Error("Failed to create consumer:", err)
	}

	return c
}

func NewKafkaProducer(viper *viper.Viper, logger *loggerApp.Logger) *kafka.Producer {
	kafkaCgf := &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA_BOOTSTRAP_SERVERS"),
	}

	p, err := kafka.NewProducer(kafkaCgf)
	if err != nil {
		logger.Error("Failed to create producer", err)
	}

	return p
}
