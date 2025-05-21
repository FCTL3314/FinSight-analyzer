package main

import (
	"context"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/config"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/service/imagedescriber"
	"github.com/FCTL3314/imagination-go-sdk/pkg/brokers/kafka"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("Received signal: %v, shutting down gracefully...", sig)
		cancel()
	}()

	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: cfg.Kafka.Brokers,
		Topic:   cfg.Kafka.TopicIn,
		GroupID: cfg.Kafka.GroupId,
	})
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("failed to close kafka reader: %v", err)
		}
	}()

	log.Printf("Running consumer...")

	router := kafka.NewRouter(reader)
	router.RegisterHandler(cfg.Kafka.TopicIn, imagedescriber.DescribeImagePoetically)
	go router.Consume(ctx)

	<-ctx.Done()

	log.Printf("Consumer stopped.")
}
