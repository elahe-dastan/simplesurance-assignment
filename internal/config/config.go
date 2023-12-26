package config

import (
	"os"
	"time"
)

type Config struct {
	WriteInterval time.Duration
}

func New() Config {
	cfg := Config{
		WriteInterval: 10 * time.Second,
	}

	wi := os.Getenv("SAA_WRITE_INTERVAL")
	if wi != "" {
		wi, err := time.ParseDuration(wi)
		if err == nil {
			cfg.WriteInterval = wi
		}
	}

	return cfg
}
