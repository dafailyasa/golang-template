package ports

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/dafailyasa/golang-template/pkg/kafka/models"
)

type ConsumerHandler func(msg *kafka.Message) error

type KafkaRepository interface {
	Send(payload *models.Producer) error
	//Consume(ctx context.Context, topic string, handler ConsumerHandler)
	Consume(ctx context.Context, topic string)
}
