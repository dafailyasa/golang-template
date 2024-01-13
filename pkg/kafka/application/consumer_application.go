package application

import (
	"context"

	"github.com/dafailyasa/golang-template/pkg/kafka/ports"
)

type KafkaConsumer struct {
	ConsumerRepo ports.KafkaRepository
}

var _ ports.ProducerApplication = (*KafkaProducer)(nil)

func NewKafkaConsumer(repo ports.KafkaRepository) *KafkaConsumer {
	return &KafkaConsumer{
		ConsumerRepo: repo,
	}
}

func (c *KafkaConsumer) Consumer(topic string, ctx context.Context) {
	c.ConsumerRepo.Consume(ctx, topic)
}
