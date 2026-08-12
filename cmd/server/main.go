package main

import (
	"context"
	"log"
	"time"

	"github.com/QingQingzi12/GitHub-Review-bot/internal/config"
	"github.com/QingQingzi12/GitHub-Review-bot/internal/httpserver"
	"github.com/QingQingzi12/GitHub-Review-bot/internal/store"
)

// main：读配置 → 连接 Postgres/Redis → 探活 → 启动 HTTP。
func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pg, err := store.NewPostgres(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("postgres init failed: %v", err)
	}
	defer pg.Close()

	if err := pg.Ping(ctx); err != nil {
		log.Fatalf("postgres ping failed: %v (请先执行: docker compose up -d)", err)
	}
	log.Println("postgres: ok")

	rdb, err := store.NewRedis(cfg.RedisURL)
	if err != nil {
		log.Fatalf("redis init failed: %v", err)
	}
	defer rdb.Close()

	if err := rdb.Ping(ctx); err != nil {
		log.Fatalf("redis ping failed: %v (请先执行: docker compose up -d)", err)
	}
	log.Println("redis: ok")

	r := httpserver.NewRouter(httpserver.Deps{
		Postgres: pg,
		Redis:    rdb,
	})

	log.Printf("server listening on %s", cfg.HTTPAddr)
	if err := r.Run(cfg.HTTPAddr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
