package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dafailyasa/golang-template/pkg/factories"
)

func main() {
	factories := factories.NewFactory(
		"./config/config.yaml",
		"logs/log.csv",
	)

	_ = factories.InitializeViper()
	logger := factories.InitializeZapLogger()
	//factories.InitializeMongoDB()

	logger.Info("Starting Initialize kafka consumer", nil)

	ctx, cancel := context.WithCancel(context.Background())
	consumer := factories.InitializeKafkaConsumer(ctx)

	logger.Info("Worker is running", nil)

	// consumer
	go consumer.ConsumerRepo.Consume(ctx, "datagen-topic")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	stop := false
	for !stop {
		s := <-terminateSignals
		fmt.Println("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME:", s)
		cancel()
		stop = true
	}

	time.Sleep(5 * time.Second)
}
