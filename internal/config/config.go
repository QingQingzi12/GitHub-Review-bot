package config

import "os"

// Config 存放程序启动时需要的配置。
// Day1 只需要 HTTP 监听地址，后面再加数据库、Redis、Webhook 密钥等。
type Config struct {
	// HTTPAddr 例如 ":8080"，表示在本机 8080 端口提供服务
	HTTPAddr string
}

// Load 从环境变量读取配置。
// 初学者理解：环境变量像“系统级开关”，改端口不用改代码。
func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		// 没设置时给一个默认值，方便第一次直接运行
		addr = ":8080"
	}
	return Config{HTTPAddr: addr}
}
