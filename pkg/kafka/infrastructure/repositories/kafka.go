package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dafailyasa/golang-template/pkg/constants"
	"github.com/dafailyasa/golang-template/pkg/kafka/models"
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

func NewKafkaClient(logger *loggerApp.Logger, viper *viper.Viper, action string) (*Kafka, error) {
	var c *kafka.Consumer
	var p *kafka.Producer
	var err error

	config := kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA.SERVERS"),
		"security.protocol": viper.GetString("KAFKA.SECURITY_PROTOCOL"),
		"sasl.mechanisms":   viper.GetString("KAFKA.SASL_MECHANISMS"),
		"sasl.username":     viper.GetString("KAFKA.USERNAME"),
		"sasl.password":     viper.GetString("KAFKA.PASSWORD"),
		"auto.offset.reset": viper.GetString("KAFKA.AUTO_OFFSET_RESET"),
	}

	if action == constants.ConsumerAction {
		c, err = kafka.NewConsumer(&config)
		if err != nil {
			logger.Error("Failed to create consumer", err)
			return nil, err
		}
	} else {
		config["group.id"] = viper.GetString("KAFKA.GROUP_ID")

		p, err = kafka.NewProducer(&config)
		if err != nil {
			logger.Error("Failed to create consumer", err)
			return nil, err
		}
	}

	return &Kafka{
		Producer: p,
		Consumer: c,
		Logger:   logger,
	}, nil
}

func (k *Kafka) Send(payload *models.Producer) error {
	val, err := json.Marshal(payload.Data)
	if err != nil {
		k.Logger.Error("Failed marshal data", err)
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &payload.Topic,
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
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
				if err != nil {
					k.Logger.Error("Failed to process message", err)
				} else {
					_, err = k.Consumer.CommitMessage(msg)
					if err != nil {
						k.Logger.Error("Failed to commit message", err)
					}
				}
			} else if !err.(kafka.Error).IsTimeout() {
				k.Logger.Warn("Consumer error", err)
			}
		}
	}

	msg := fmt.Sprintf("Closing consumer for topic %s", topic)
	k.Logger.Info(msg, nil)

	err = k.Consumer.Close()
	if err != nil {
		panic(err)
	}
}
