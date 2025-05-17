package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type App struct {
	Debug bool `envconfig:"DEBUG" default:"false"`
}
type Kafka struct {
	Brokers     []string `envconfig:"KAFKA_BROKERS" required:"true"`
	TopicInput  string   `envconfig:"KAFKA_TOPIC_INPUT" required:"true"`
	TopicOutput string   `envconfig:"KAFKA_TOPIC_OUTPUT" required:"true"`
	GroupID     string   `envconfig:"KAFKA_GROUP_ID" required:"true"`
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
	App      App
	Kafka    Kafka
	S3       S3
	Database Database
	Analyzer Analyzer
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("No .env file found or error loading it: %v", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
