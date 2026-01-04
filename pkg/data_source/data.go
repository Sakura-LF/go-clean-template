package data_source

import (
	"fmt"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// DataSource -.
type DataSource struct {
	db *gorm.DB

	rdb *redis.Client
}

// NewDataSource -.
// init Postgres And Redis
func NewDataSource(data config.Data, log *logger.Logger) (*DataSource, error) {
	db, err := NewDB(data.DataBase, log)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("app - Run - postgres.NewUseCase: %w", err))
	}
	rdb, err := NewRedis(data.Redis, log)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("app - Run - redis.NewUseCase: %w", err))
	}
	return &DataSource{
		db:  db,
		rdb: rdb,
	}, nil
}
