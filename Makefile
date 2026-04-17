# xiaohongshu-agent Makefile
# Created by 🦞 Lobster Journey Studio

.PHONY: all build install clean test help

# 默认目标
all: build

# 构建MCP服务
build:
	@echo "🦞 构建 xiaohongshu-mcp..."
	cd mcp && go build -o xiaohongshu-mcp ./cmd/server
	@echo "✅ 构建完成"

# 安装
install:
	@echo "🦞 安装 xiaohongshu-agent..."
	./install.sh

# 清理
clean:
	@echo "🧹 清理构建产物..."
	rm -f mcp/xiaohongshu-mcp
	rm -rf __pycache__
	rm -rf .pytest_cache
	find . -type d -name "__pycache__" -exec rm -rf {} +
	@echo "✅ 清理完成"

# 运行测试
test:
	@echo "🧪 运行测试..."
	cd skill && python3 -m pytest tests/ -v

# 运行MCP服务
run:
	@echo "🚀 启动 xiaohongshu-mcp..."
	cd mcp && ./xiaohongshu-mcp --port=18060

# 登录小红书
login:
	@echo "🔐 启动登录工具..."
	cd mcp && go run ./cmd/login

# 查看日志
logs:
	@echo "📋 查看日志..."
	tail -f ~/.xiaohongshu-agent/logs/mcp.log

# 帮助
help:
	@echo "🦞 xiaohongshu-agent 命令："
	@echo ""
	@echo "  make build      - 构建MCP服务"
	@echo "  make install    - 安装所有依赖"
	@echo "  make clean      - 清理构建产物"
	@echo "  make test       - 运行测试"
	@echo "  make run        - 启动MCP服务"
	@echo "  make login      - 登录小红书"
	@echo "  make logs       - 查看日志"
	@echo ""
	@echo "🦞 Lobster Journey Studio"

# Docker构建
docker-build:
	@echo "🐳 构建Docker镜像..."
	docker build -t xiaohongshu-agent:latest -f docker/Dockerfile .

# Docker运行
docker-run:
	@echo "🐳 启动Docker容器..."
	docker-compose -f docker/docker-compose.yaml up -d
