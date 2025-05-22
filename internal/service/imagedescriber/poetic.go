package imagedescriber

import (
	"context"
	"encoding/json"
	"github.com/FCTL3314/imagination-analyzer/pkg/models"
	kafkasdk "github.com/FCTL3314/imagination-go-sdk/pkg/brokers/kafka"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func DescribeImagePoetically(
	ctx context.Context,
	logger *zap.Logger,
	metadata *kafkasdk.MessageMetadata,
	msg kafka.Message,
) error {
	var payload models.ImageToDescribe
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		logger.Error("failed to unmarshal payload", zap.Error(err))
		return err
	}
	logger.Info("parsed message payload", zap.Any("image_to_describe", payload))

	// … здесь остальная логика DescribeImagePoetically …

	return nil
}
