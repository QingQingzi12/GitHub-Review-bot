package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis 包装客户端，后面做幂等、队列、限流都会用它。
type Redis struct {
	Client *redis.Client
}

// NewRedis 根据 REDIS_URL 创建客户端。
func NewRedis(redisURL string) (*Redis, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}

	return &Redis{Client: redis.NewClient(opt)}, nil
}

// Ping 检查 Redis 是否可连通。
func (r *Redis) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return r.Client.Ping(ctx).Err()
}

// Close 关闭客户端。
func (r *Redis) Close() error {
	return r.Client.Close()
}
