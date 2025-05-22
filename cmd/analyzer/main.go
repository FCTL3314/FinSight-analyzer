package main

import (
	"context"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/bootstrap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("Received signal: %v, shutting down gracefully...", sig)
		cancel()
	}()

	app := bootstrap.NewApplication()

	app.Logger.Info("Running consumers...")
	app.Consumers.ImgPoeticDescriptionConsumer.Start()

	<-ctx.Done()
	app.Logger.Info("Consumers stopped.")
}
