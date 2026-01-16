package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var Instance *AppConfig

type AppConfig struct {
	StoreType string `env:"STORE_TYPE" envDefault:"JSON"`
}

func Load() {
	_ = godotenv.Load()

	cfg := &AppConfig{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Config parsing failed: %v", err)
	}

	Instance = cfg
}
