package kafka

import (
	"context"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/config"
	"github.com/FCTL3314/ExerciseManager-Backend/internal/service/imagedescriber"
	"github.com/segmentio/kafka-go"
	"log"
)

type HandlerFunc func(ctx context.Context, msg kafka.Message) error

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) RegisterHandler(topic string, handler HandlerFunc) {
	r.handlers[topic] = handler
}

func (r *Router) GetHandler(topic string) HandlerFunc {
	return r.handlers[topic]
}

func (r *Router) Consume(ctx context.Context, reader *kafka.Reader) {
	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Println("[kafka] consumer shutdown")
				return
			}
			log.Printf("[kafka] read error: %v", err)
			continue
		}

		handler, ok := r.handlers[m.Topic]
		if !ok {
			log.Printf("[kafka] no handler for topic %s, skipping offset=%d", m.Topic, m.Offset)
			continue
		}

		go func(msg kafka.Message) {
			if err := handler(ctx, msg); err != nil {
				log.Printf("[kafka] handler error for topic %s: %v", msg.Topic, err)
			}
		}(m)
	}
}

func NewReader(cfg config.Kafka) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   cfg.TopicInput,
		GroupID: cfg.GroupID,
	})
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(reader)
	return reader
}

func RegisterSomeConsumer(cfg config.Kafka, ctx context.Context, reader *kafka.Reader) {
	router := NewRouter()
	router.RegisterHandler(cfg.TopicOutput, imagedescriber.HandlerFunc)
	router.Consume(ctx, reader)
}
