.PHONY: build run test clean install deps

# 变量
BINARY_NAME=xiaohongshu-agent
MAIN_PATH=./mcp
BUILD_DIR=./bin

# 默认目标
all: deps build

# 安装依赖
deps:
	@echo "📦 安装依赖..."
	cd mcp && go mod download
	@echo "✅ 依赖安装完成"

# 构建
build:
	@echo "🔨 构建项目..."
	mkdir -p $(BUILD_DIR)
	cd mcp && go build -o ../$(BUILD_DIR)/$(BINARY_NAME) .
	@echo "✅ 构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 运行
run: build
	@echo "🚀 启动服务..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# 测试
test:
	@echo "🧪 运行测试..."
	cd mcp && go test -v ./...

# 清理
clean:
	@echo "🧹 清理构建文件..."
	rm -rf $(BUILD_DIR)
	@echo "✅ 清理完成"

# 安装Playwright
install-playwright:
	@echo "🎭 安装Playwright..."
	cd mcp && go run -v -x ./scripts/install_playwright.go
	@echo "✅ Playwright安装完成"

# 开发模式（热重载）
dev:
	@echo "🔄 开发模式..."
	@which air > /dev/null || go install github.com/cosmtrek/air@latest
	cd mcp && air

# Docker构建
docker-build:
	@echo "🐳 构建Docker镜像..."
	docker build -f docker/Dockerfile -t xiaohongshu-agent:latest .

# Docker运行
docker-run:
	@echo "🐳 运行Docker容器..."
	docker-compose -f docker/docker-compose.yaml up -d

# Docker停止
docker-stop:
	@echo "🐳 停止Docker容器..."
	docker-compose -f docker/docker-compose.yaml down

# 帮助
help:
	@echo "小红书Agent - Make命令"
	@echo ""
	@echo "使用方法: make [目标]"
	@echo ""
	@echo "目标:"
	@echo "  deps              安装依赖"
	@echo "  build             构建项目"
	@echo "  run               运行服务"
	@echo "  test              运行测试"
	@echo "  clean             清理构建文件"
	@echo "  install-playwright 安装Playwright"
	@echo "  dev               开发模式（热重载）"
	@echo "  docker-build      构建Docker镜像"
	@echo "  docker-run        运行Docker容器"
	@echo "  docker-stop       停止Docker容器"
	@echo "  help              显示帮助信息"
