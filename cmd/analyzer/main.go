package main

import (
	"github.com/FCTL3314/ExerciseManager-Backend/internal/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		return
	}
	log.Printf("Config: %+v", cfg)
}
