package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dafailyasa/golang-template/pkg/kafka/ports"
	loggerApp "github.com/dafailyasa/golang-template/pkg/logger/application"
	"github.com/spf13/viper"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Kafka struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
	Logger   *loggerApp.Logger
}

var _ ports.KafkaRepository = (*Kafka)(nil)

func NewKafkaProducer(logger *loggerApp.Logger, viper *viper.Viper) *Kafka {
	kafkaCgf := &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA.SERVERS"),
		"security.protocol": viper.GetString("KAFKA.SECURITY_PROTOCOL"),
		"sasl.mechanisms":   viper.GetString("KAFKA.SASL_MECHANISMS"),
		"sasl.username":     viper.GetString("KAFKA.USERNAME"),
		"sasl.password":     viper.GetString("KAFKA.PASSWORD"),
		"auto.offset.reset": viper.GetString("KAFKA.AUTO_OFFSET_RESET"),
	}

	p, err := kafka.NewProducer(kafkaCgf)
	if err != nil {
		logger.Error("Failed to create producer", err)
	}

	return &Kafka{
		Producer: p,
		Logger:   logger,
	}
}

func NewKafkaConsumer(logger *loggerApp.Logger, viper *viper.Viper) *Kafka {
	kafkaCgf := &kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA.SERVERS"),
		"security.protocol": viper.GetString("KAFKA.SECURITY_PROTOCOL"),
		"sasl.mechanisms":   viper.GetString("KAFKA.SASL_MECHANISMS"),
		"sasl.username":     viper.GetString("KAFKA.USERNAME"),
		"sasl.password":     viper.GetString("KAFKA.PASSWORD"),
		"auto.offset.reset": viper.GetString("KAFKA.AUTO_OFFSET_RESET"),
	}

	c, err := kafka.NewConsumer(kafkaCgf)
	if err != nil {
		logger.Error("Failed to create consumer", err)
	}

	return &Kafka{
		Consumer: c,
		Logger:   logger,
	}
}

func (k *Kafka) Send(topic string, data any) error {
	val, err := json.Marshal(data)
	if err != nil {
		k.Logger.Error("Failed marshal data", err)
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: val,
	}

	err = k.Producer.Produce(msg, nil)
	if err != nil {
		k.Logger.Error("Failed to produce message", err)
		return err
	}

	return nil
}

func (k *Kafka) Consume(ctx context.Context, topic string) {
	err := k.Consumer.Subscribe(topic, nil)
	if err != nil {
		msg := fmt.Sprintf("Failed to subscribe topic %s", topic)
		k.Logger.Error(msg, err)
	}

	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			msg, err := k.Consumer.ReadMessage(time.Second)
			if err == nil {
				_, err = k.Consumer.CommitMessage(msg)

				if err != nil {
					k.Logger.Error("Failed to commit message", err)
				}
			} else if !err.(kafka.Error).IsTimeout() {
				k.Logger.Warn("Consumer error", err)
			}
		}
	}

	k.Logger.Info("Closing consumer for topic", topic)

	err = k.Consumer.Close()
	if err != nil {
		panic(err)
	}
}
