package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/QingQingzi12/GitHub-Review-bot/internal/store"
	"github.com/gin-gonic/gin"
)

// Deps 路由依赖。Day2 先放 Postgres/Redis，方便 /healthz 检查连通性。
type Deps struct {
	Postgres *store.Postgres
	Redis    *store.Redis
}

// NewRouter 创建并返回 Gin 路由引擎。
func NewRouter(deps Deps) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		pgOK := deps.Postgres != nil && deps.Postgres.Ping(ctx) == nil
		redisOK := deps.Redis != nil && deps.Redis.Ping(ctx) == nil

		status := "ok"
		code := http.StatusOK
		if !pgOK || !redisOK {
			status = "degraded"
			code = http.StatusServiceUnavailable
		}

		c.JSON(code, gin.H{
			"status":   status,
			"service":  "GitHub-Review-bot",
			"postgres": pgOK,
			"redis":    redisOK,
		})
	})

	return r
}
