package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Debug                          bool `envconfig:"DEBUG" default:"false"`
	PoeticImgDescriptionMaxWorkers int  `envconfig:"POETIC_IMAGE_DESCRIPTION_MAX_WORKERS" required:"true"`
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

type Kafka struct {
	Brokers  []string `envconfig:"KAFKA_BROKERS" required:"true"`
	TopicIn  string   `envconfig:"KAFKA_TOPIC_POETIC_IMAGE_DESCRIPTION_IN" required:"true"`
	TopicOut string   `envconfig:"KAFKA_TOPIC_POETIC_IMAGE_DESCRIPTION_OUT" required:"true"`
	GroupId  string   `envconfig:"KAFKA_GROUP_POETIC_IMAGE_DESCRIPTION_ID" required:"true"`
}

type Config struct {
	App      App
	Kafka    Kafka
	S3       S3
	Database Database
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
