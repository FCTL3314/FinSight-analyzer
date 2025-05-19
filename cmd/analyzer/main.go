package main

import (
	"context"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/brokers/kafka"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/config"
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

	reader := kafka.NewReader(cfg.Kafka)
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("failed to close kafka reader: %v", err)
		}
	}()

	log.Printf("Running consumer...")
	kafka.RegisterSomeConsumer(ctx, cfg.Kafka, reader)

	log.Printf("Consumer stopped.")
}
