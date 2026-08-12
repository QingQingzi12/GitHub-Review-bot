package config

import "os"

// Config 存放程序启动时需要的配置。
type Config struct {
	// HTTPAddr 例如 ":8080"
	HTTPAddr string
	// DatabaseURL Postgres 连接串
	DatabaseURL string
	// RedisURL Redis 连接串
	RedisURL string
}

// Load 从环境变量读取配置；没设置则用本地 Docker Compose 的默认值。
func Load() Config {
	return Config{
		HTTPAddr:    getEnv("HTTP_ADDR", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://reviewbot:reviewbot@127.0.0.1:5432/reviewbot?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "redis://127.0.0.1:6379/0"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
