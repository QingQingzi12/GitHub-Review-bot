package main

import (
	"log"

	"github.com/yourname/GitHub-Review-bot/internal/config"
	"github.com/yourname/GitHub-Review-bot/internal/httpserver"
)

// main 是程序入口：读配置 → 创建路由 → 启动 HTTP 服务。
// Day1 先把“能跑起来”做稳，不接 GitHub Webhook。
func main() {
	cfg := config.Load()
	r := httpserver.NewRouter()

	log.Printf("server listening on %s", cfg.HTTPAddr)
	if err := r.Run(cfg.HTTPAddr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
