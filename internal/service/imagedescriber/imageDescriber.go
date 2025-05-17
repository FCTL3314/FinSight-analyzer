package imagedescriber

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func HandlerFunc(ctx context.Context, msg kafka.Message) error {
	log.Printf("Broker message: %+v", msg)
	return nil
}
