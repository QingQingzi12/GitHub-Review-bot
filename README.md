# GitHub-Review-bot

用 Go 学习搭建的 GitHub Webhook AI PR 审查机器人（Day1：项目骨架 + 健康检查）。

## 环境要求

- Go 1.22+
- Git
- （可选）GoLand 2026

## Day1：本地启动

在项目根目录执行：

```powershell
go mod tidy
go run .\cmd\server
```

浏览器或终端访问：

```powershell
curl http://127.0.0.1:8080/healthz
```

期望返回类似：

```json
{"status":"ok","service":"GitHub-Review-bot"}
```

## 环境变量

| 变量 | 默认值 | 含义 |
|------|--------|------|
| `HTTP_ADDR` | `:8080` | HTTP 监听地址 |

示例：

```powershell
$env:HTTP_ADDR=":9090"
go run .\cmd\server
```

## 目录说明

```text
cmd/server          程序入口
internal/config     配置读取
internal/httpserver HTTP 路由
```
