package bootstrap

import (
	"github.com/FCTL3314/imagination-analyzer/internal/service/imagedescriber"
	"github.com/FCTL3314/imagination-go-sdk/pkg/brokers/kafka"
	"github.com/FCTL3314/imagination-go-sdk/pkg/brokers/kafka/workerpool"
	"go.uber.org/zap"
	"log"
	"time"
)

type Application struct {
	Cfg       *Config
	Logger    *zap.Logger
	Consumers Consumers
	Producers Producers
}

type Consumers struct {
	ImgPoeticDescriptionConsumer *kafka.Consumer
}

type Producers struct {
}

func NewApplication() *Application {
	var app Application
	app.initConfig()
	app.initLogger()
	app.initImgPoeticDescriptionConsumer()
	return &app
}

func (app *Application) initConfig() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	app.Cfg = cfg
}

func (app *Application) initLogger() {
	var (
		logger *zap.Logger
		err    error
	)
	if app.Cfg.App.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	app.Logger = logger
}

func (app *Application) initImgPoeticDescriptionConsumer() {
	cfg := kafka.ConsumerConfig{
		Brokers:        app.Cfg.Kafkas.PoeticImgDescription.Brokers,
		Topic:          app.Cfg.Kafkas.PoeticImgDescription.TopicIn,
		GroupID:        app.Cfg.Kafkas.PoeticImgDescription.GroupID,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		Logger:         app.Logger,
	}

	poolHandler := workerpool.NewPoolHandler(
		app.Cfg.App.Services.PoeticImgDescription.MaxWorkers,
		kafka.HandlerFunc(imagedescriber.DescribeImagePoetically),
		app.Logger,
	)
	consumer := kafka.NewConsumer(cfg, poolHandler)
	app.Consumers.ImgPoeticDescriptionConsumer = consumer
}
