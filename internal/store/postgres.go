package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres 包装连接池，后面 Day3 写表也会用它。
type Postgres struct {
	Pool *pgxpool.Pool
}

// NewPostgres 根据 DATABASE_URL 创建连接池。
func NewPostgres(ctx context.Context, databaseURL string) (*Postgres, error) {
	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database url: %w", err)
	}

	// 学习阶段连接数开小一点即可
	cfg.MaxConns = 5

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}

	return &Postgres{Pool: pool}, nil
}

// Ping 检查数据库是否可连通。
func (p *Postgres) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	return p.Pool.Ping(ctx)
}

// Close 关闭连接池（进程退出时调用）。
func (p *Postgres) Close() {
	p.Pool.Close()
}
