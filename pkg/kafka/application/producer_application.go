package application

import (
	"github.com/dafailyasa/golang-template/pkg/kafka/models"
	"github.com/dafailyasa/golang-template/pkg/kafka/ports"
)

type KafkaProducer struct {
	ProducerRepo ports.KafkaRepository
}

var _ ports.ProducerApplication = (*KafkaProducer)(nil)

func NewKafkaProducer(repo ports.KafkaRepository) *KafkaProducer {
	return &KafkaProducer{
		ProducerRepo: repo,
	}
}

func (p *KafkaProducer) Publish(topic string, data any) error {
	payload := models.Producer{
		Topic: topic,
		Data:  data,
	}

	err := p.ProducerRepo.Send(&payload)
	if err != nil {
		return err
	}

	return nil
}
