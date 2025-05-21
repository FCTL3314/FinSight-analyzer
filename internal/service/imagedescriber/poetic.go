package imagedescriber

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func DescribeImagePoetically(ctx context.Context, msg kafka.Message) error {
	log.Printf("Broker message: %+v", msg)
	log.Printf("Message value: %+v", string(msg.Value))
	return nil
}
