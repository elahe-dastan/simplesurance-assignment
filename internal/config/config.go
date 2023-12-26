package config

import (
	"os"
	"time"
)

type Config struct {
	WriteInterval time.Duration
	Filename      string
}

func New() Config {
	cfg := Config{
		WriteInterval: 10 * time.Second,
		Filename:      "state.json",
	}

	wi := os.Getenv("SAA_WRITE_INTERVAL")
	if wi != "" {
		wi, err := time.ParseDuration(wi)
		if err == nil {
			cfg.WriteInterval = wi
		}
	}

	fn := os.Getenv("SAA_FILENAME")
	if fn != "" {
		cfg.Filename = fn
	}

	return cfg
}
