package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter 创建并返回 Gin 路由引擎。
// 以后 Webhook、查询 API 都往这里注册路径。
func NewRouter() *gin.Engine {
	// gin.Default() 会带上日志和 panic 恢复中间件，适合学习阶段
	r := gin.Default()

	// GET /healthz：探活接口，用来确认服务是否活着
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "GitHub-Review-bot",
		})
	})

	return r
}
