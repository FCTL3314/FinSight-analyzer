package config

import (
	"github.com/FCTL3314/imagination-go-sdk/pkg/brokers/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"strings"
)

type App struct {
	Debug bool `envconfig:"DEBUG" default:"false"`
}

type S3 struct {
	EndpointURL     string `envconfig:"S3_ENDPOINT_URL" required:"true"`
	AccessKeyID     string `envconfig:"S3_ACCESS_KEY_ID" required:"true"`
	SecretAccessKey string `envconfig:"S3_SECRET_ACCESS_KEY" required:"true"`
}

type Database struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     uint32 `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

type Analyzer struct {
	MaxImageSizeMB uint32 `envconfig:"MAX_IMAGE_SIZE_MB" default:"100"`
}

type Config struct {
	App         App
	KafkaPoetic config.Kafka
	S3          S3
	Database    Database
	Analyzer    Analyzer
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	cfg.KafkaPoetic = config.Kafka{
		Brokers:      strings.Split(os.Getenv("KAFKA_BROKERS"), ","), // TODO: Add split with trim func
		InputTopics:  []string{os.Getenv("KAFKA_TOPIC_IN_POETIC")},
		OutputTopics: []string{os.Getenv("KAFKA_TOPIC_OUT_POETIC")},
		GroupID:      os.Getenv("KAFKA_GROUP_ID_POETIC"),
	}

	return &cfg, nil
}
