package bootstrap

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Debug    bool `envconfig:"DEBUG" default:"false"`
	Services Services
}

type Services struct {
	PoeticImgDescription PoeticImgDescription
}

type PoeticImgDescription struct {
	MaxWorkers int `envconfig:"POETIC_IMAGE_DESCRIPTION_MAX_WORKERS" required:"true"`
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

type Kafkas struct {
	PoeticImgDescription *Kafka
}

type Kafka struct {
	Brokers  []string `envconfig:"KAFKA_BROKERS"  required:"true"`
	TopicIn  string   `envconfig:"KAFKA_TOPIC_IN" required:"true"`
	TopicOut string   `envconfig:"KAFKA_TOPIC_OUT" required:"true"`
	GroupID  string   `envconfig:"KAFKA_GROUP_ID" required:"true"`
}

type Config struct {
	App      App
	Kafkas   Kafkas
	S3       S3
	Database Database
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	var cfg Config

	if err := envconfig.Process("", &cfg.App); err != nil {
		return nil, err
	}
	if err := envconfig.Process("", &cfg.S3); err != nil {
		return nil, err
	}
	if err := envconfig.Process("", &cfg.Database); err != nil {
		return nil, err
	}

	cfg.Kafkas.PoeticImgDescription = new(Kafka)
	if err := envconfig.Process("POETIC_IMAGE_DESCRIPTION", cfg.Kafkas.PoeticImgDescription); err != nil {
		return nil, err
	}

	return &cfg, nil
}
