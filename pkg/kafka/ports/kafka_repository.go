package ports

type KafkaRepository interface {
	Send(topic string, data any) error
}
