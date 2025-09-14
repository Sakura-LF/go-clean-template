package config

import (
	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		App     App
		HTTP    HTTP
		Log     Log
		PG      PG
		GRPC    GRPC
		RMQ     RMQ
		Metrics Metrics
		Swagger Swagger
	}

	// App -.
	App struct {
		Name    string
		Version string
	}

	// HTTP -.
	HTTP struct {
		Port           string
		UsePreforkMode bool
	}

	// Log -.
	Log struct {
		Level string
	}

	// PG -.
	PG struct {
		PoolMax int
		URL     string
	}

	// GRPC -.
	GRPC struct {
		Port string
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string
		ClientExchange string
		URL            string
	}

	// Metrics -.
	Metrics struct {
		Enabled bool
	}

	// Swagger -.
	Swagger struct {
		Enabled bool
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("settings")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
