package application

import (
	"github.com/dafailyasa/golang-template/pkg/kafka/ports"
)

type KafkaConsumer struct {
	ProducerRepo ports.KafkaRepository
}

var _ ports.ProducerApplication = (*KafkaProducer)(nil)

func NewKafkaConsumer(repo ports.KafkaRepository) *KafkaProducer {
	return &KafkaProducer{
		ProducerRepo: repo,
	}
}

// func (p *KafkaProducer) Consume(topic string) error {
// 	err := p.ProducerRepo
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
