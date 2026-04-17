.PHONY: all build install test clean dev run

# 变量
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.0.0")
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION := $(shell go version | awk '{print $$3}')
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# 默认目标
all: build

# 安装依赖
install-deps:
	@echo "📥 安装 Go 依赖..."
	cd mcp && go mod download
	@echo "📥 安装 Python 依赖..."
	pip install -r skill/requirements.txt
	@echo "✅ 依赖安装完成"

# 安装浏览器
install-browser:
	@echo "🌐 安装 Playwright 浏览器..."
	playwright install chromium
	playwright install-deps chromium
	@echo "✅ 浏览器安装完成"

# 构建
build:
	@echo "🔨 构建 MCP 服务..."
	cd mcp && go build $(LDFLAGS) -o bin/xiaohongshu-mcp ./cmd/server
	cd mcp && go build $(LDFLAGS) -o bin/xiaohongshu-login ./cmd/login
	@echo "✅ 构建完成"

# 构建所有平台
build-all:
	@echo "🔨 构建所有平台..."
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/xiaohongshu-mcp-linux-amd64 ./mcp/cmd/server
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/xiaohongshu-login-linux-amd64 ./mcp/cmd/login
	
	# Linux ARM64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/xiaohongshu-mcp-linux-arm64 ./mcp/cmd/server
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/xiaohongshu-login-linux-arm64 ./mcp/cmd/login
	
	# Darwin AMD64
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/xiaohongshu-mcp-darwin-amd64 ./mcp/cmd/server
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/xiaohongshu-login-darwin-amd64 ./mcp/cmd/login
	
	# Darwin ARM64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/xiaohongshu-mcp-darwin-arm64 ./mcp/cmd/server
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/xiaohongshu-login-darwin-arm64 ./mcp/cmd/login
	
	@echo "✅ 全平台构建完成"

# 开发模式
dev:
	@echo "🚀 启动开发服务器..."
	cd mcp && go run ./cmd/server --config ../configs/config.dev.yaml

# 运行服务
run:
	@echo "🚀 启动服务..."
	cd mcp && ./bin/xiaohongshu-mcp --config ../configs/config.yaml

# 测试
test:
	@echo "🧪 运行测试..."
	make test-go
	make test-python

test-go:
	@echo "🧪 运行 Go 测试..."
	cd mcp && go test -v -race -coverprofile=coverage.out ./...

test-python:
	@echo "🧪 运行 Python 测试..."
	cd skill && python -m pytest tests/ -v

# 测试覆盖率
coverage:
	@echo "📊 生成测试覆盖率报告..."
	cd mcp && go tool cover -html=coverage.out -o coverage.html
	@echo "✅ 覆盖率报告: mcp/coverage.html"

# 代码检查
lint:
	@echo "🔍 代码检查..."
	cd mcp && golangci-lint run
	cd skill && flake8 scripts/ tests/

# 格式化代码
fmt:
	@echo "📝 格式化代码..."
	cd mcp && go fmt ./...
	cd skill && black scripts/ tests/

# 清理
clean:
	@echo "🧹 清理构建文件..."
	rm -rf mcp/bin
	rm -rf dist
	rm -f mcp/coverage.out
	@echo "✅ 清理完成"

# Docker 构建
docker:
	@echo "🐳 构建 Docker 镜像..."
	docker build -t xiaohongshu-agent:$(VERSION) -f docker/Dockerfile .
	@echo "✅ Docker 镜像构建完成"

# 发布
release:
	@echo "📦 发布新版本 $(VERSION)..."
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)
	make build-all
	@echo "✅ 发布完成"

# 安装
install:
	@echo "📦 安装到系统..."
	./install.sh

# 帮助
help:
	@echo "可用的 make 目标:"
	@echo "  make install-deps    - 安装依赖"
	@echo "  make install-browser - 安装浏览器"
	@echo "  make build           - 构建项目"
	@echo "  make build-all       - 构建所有平台"
	@echo "  make dev             - 开发模式运行"
	@echo "  make run             - 运行服务"
	@echo "  make test            - 运行测试"
	@echo "  make test-go         - 运行 Go 测试"
	@echo "  make test-python     - 运行 Python 测试"
	@echo "  make coverage        - 测试覆盖率"
	@echo "  make lint            - 代码检查"
	@echo "  make fmt             - 格式化代码"
	@echo "  make clean           - 清理构建"
	@echo "  make docker          - Docker 构建"
	@echo "  make release         - 发布新版本"
	@echo "  make install         - 一键安装"
