# 构建阶段
FROM golang:1.24-alpine AS builder

# 安装依赖
RUN apk add --no-cache git make

WORKDIR /build

# 复制依赖文件
COPY mcp/go.mod mcp/go.sum ./
RUN go mod download

# 复制源代码
COPY mcp/ ./

# 构建
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o xiaohongshu-mcp ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o xiaohongshu-login ./cmd/login

# 运行阶段
FROM python:3.11-slim

# 安装系统依赖
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    && rm -rf /var/lib/apt/lists/*

# 安装Playwright
RUN pip install playwright
RUN playwright install chromium
RUN playwright install-deps chromium

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /build/xiaohongshu-mcp /usr/local/bin/
COPY --from=builder /build/xiaohongshu-login /usr/local/bin/

# 复制配置文件
COPY configs/config.yaml /app/config.yaml

# 创建数据目录
RUN mkdir -p /app/data /app/logs

# 暴露端口
EXPOSE 18060 9090

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:18060/health || exit 1

# 启动命令
CMD ["xiaohongshu-mcp", "--config", "/app/config.yaml"]
