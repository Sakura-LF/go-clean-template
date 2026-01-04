package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	redis *redis.Client
}

type Option func(*Client) error

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}
	// 给一个默认的空 Client，或者在 Option 里初始化
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	if c.Client == nil { // 注意：这里判断的是嵌入的字段
		return nil, fmt.Errorf("redis client uninitialized")
	}
	return c, nil
}

// NewRedis 创建 Redis 连接
func NewRedis(c config.RedisConfig, log *logger.Logger) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Username:     c.Username,
		Password:     c.Password,
		WriteTimeout: c.WriteTimeout,
		ReadTimeout:  c.ReadTimeout,
		DB:           0,
		DialTimeout:  1 * time.Second,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Error().Err(err)
		panic("Redis 连接失败")
	} else {
		log.Info().Msg("Redis connect success")
	}
	return rdb, nil
}

func (c *Client) Close() error {
	if c.redis != nil {
		if err := c.redis.Close(); err != nil {
			return err
		}
	}
}
