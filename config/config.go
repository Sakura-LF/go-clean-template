package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		App     App
		HTTP    HTTP
		Data    Data
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

	Data struct {
		DataBase DBConfig
		Redis    RedisConfig
	}

	DBConfig struct {
		Driver           string        `mapstructure:"driver" yaml:"driver"`
		Source           string        `mapstructure:"source" yaml:"source"`
		ReconnectionNum  int           `mapstructure:"reconnection_num" yaml:"reconnection_num"`
		ReconnectionTime time.Duration `mapstructure:"reconnection_time" yaml:"reconnection_time"`
	}

	RedisConfig struct {
		Addr         string        `json:"addr"`
		Username     string        `json:"username"`
		Password     string        `json:"password"`
		ReadTimeout  time.Duration `mapstructure:"readTimeout" yaml:"readTimeout" json:"readTimeout"`
		WriteTimeout time.Duration `mapstructure:"writeTimeout" yaml:"writeTimeout" json:"writeTimeout"`
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
