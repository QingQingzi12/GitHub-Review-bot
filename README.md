# GitHub-Review-bot

用 Go 学习搭建的 GitHub Webhook AI PR 审查机器人。

## 环境要求

- Go 1.22+
- Git
- Docker Desktop（Day2 起需要）
- GoLand 2026（可选）

## Day1：健康检查

```powershell
go run .\cmd\server
curl http://127.0.0.1:8080/healthz
```

## Day2：启动依赖并验证连通

```powershell
# 1. 启动 Postgres + Redis
docker compose up -d

# 2. 看容器是否 healthy
docker compose ps

# 3. 安装/整理依赖后启动服务
go mod tidy
go run .\cmd\server

# 4. 探活（应看到 postgres/redis 均为 true）
curl http://127.0.0.1:8080/healthz
```

期望返回类似：

```json
{"status":"ok","service":"GitHub-Review-bot","postgres":true,"redis":true}
```

### 进容器练手

```powershell
# Postgres
docker exec -it reviewbot-postgres psql -U reviewbot -d reviewbot
# 进入后可输入 \conninfo 然后 \q 退出

# Redis
docker exec -it reviewbot-redis redis-cli ping
# 应返回 PONG
```

## 环境变量

| 变量 | 默认值 | 含义 |
|------|--------|------|
| `HTTP_ADDR` | `:8080` | HTTP 监听地址 |
| `DATABASE_URL` | `postgres://reviewbot:reviewbot@127.0.0.1:5432/reviewbot?sslmode=disable` | Postgres 连接串 |
| `REDIS_URL` | `redis://127.0.0.1:6379/0` | Redis 连接串 |

可复制 `.env.example` 为 `.env`（勿提交 `.env`）。

## 目录说明

```text
cmd/server              程序入口
internal/config         配置读取
internal/httpserver     HTTP 路由
internal/store          Postgres / Redis 连接
docker-compose.yml      本地依赖
```
