// Package postgres implements postgres connection.
package data_source

import (
	"context"
	"fmt"
	log2 "log"
	"os"
	"time"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

const (
	// DB
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second * 5

	// Redis
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

//func NewDataSource1(data config.Data, opts ...Option) (*DataSource, error) {
//	pg := &DataSource{
//		maxPoolSize:  _defaultMaxPoolSize,
//		connAttempts: _defaultConnAttempts,
//		connTimeout:  _defaultConnTimeout,
//	}
//
//	// Custom options
//	for _, opt := range opts {
//		opt(pg)
//	}
//
//	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
//
//	poolConfig, err := pgxpool.ParseConfig(url)
//	if err != nil {
//		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
//	}
//
//	poolConfig.MaxConns = int32(pg.maxPoolSize) //nolint:gosec // skip integer overflow conversion int -> int32
//
//	for pg.connAttempts > 0 {
//		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
//		if err == nil {
//			break
//		}
//
//		log.Printf("DataSource is trying to connect, attempts left: %d", pg.connAttempts)
//
//		time.Sleep(pg.connTimeout)
//
//		pg.connAttempts--
//	}
//
//	if err != nil {
//		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
//	}
//
//	return pg, nil
//}

// Close -.
func (p *DataSource) Close() {
	if p.rdb != nil {
		p.rdb.Close()
	}
}

func NewDB(c config.DBConfig, log *logger.Logger) (*gorm.DB, error) {

	// 终端打印输入 sql 执行记录
	newLogger := gorm_logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
		gorm_logger.Config{
			SlowThreshold:             300 * time.Millisecond, // 慢查询 SQL 阈值
			Colorful:                  true,                   // 是否启动彩色打印
			IgnoreRecordNotFoundError: false,
			LogLevel:                  gorm_logger.Error, // Log lever
		},
	)
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal().Msg("failed to connect database")
	} else {
		log.Info().Msg("mysql connect success")
	}
	return db
}

// NewRedis 创建 Redis 连接
func NewRedis(c config.RedisConfig, log *logger.Logger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Username:     c.Username, // default user
		Password:     c.Password,
		WriteTimeout: c.WriteTimeout,
		ReadTimeout:  c.ReadTimeout,
		DB:           0,               // use default DB
		DialTimeout:  1 * time.Second, // 1 second
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Error().Err(err)
		panic("Redis 连接失败")
	} else {
		log.Info().Msg("Redis connect success")
	}
	return rdb
}
