package ports

type ProducerApplication interface {
	Publish(topic string, data any) error
}
