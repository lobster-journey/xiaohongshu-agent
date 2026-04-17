#!/bin/bash

# 小红书 Agent 快速启动脚本

set -e

echo "🦞 小红书 Agent - 快速启动"
echo "============================"
echo ""

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 检查 Go 环境
echo -e "${YELLOW}检查 Go 环境...${NC}"
if ! command -v go &> /dev/null; then
    echo -e "${RED}✗ Go 未安装${NC}"
    echo ""
    echo "请先安装 Go 1.24 或更高版本："
    echo ""
    echo "Linux:"
    echo "  wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz"
    echo "  sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz"
    echo "  export PATH=\$PATH:/usr/local/go/bin"
    echo ""
    echo "macOS:"
    echo "  brew install go"
    echo ""
    echo "Windows:"
    echo "  choco install golang"
    echo ""
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}')
echo -e "${GREEN}✓ Go 已安装: $GO_VERSION${NC}"
echo ""

# 设置 Go 代理（加速下载）
if [[ -z "$GOPROXY" ]]; then
    echo -e "${YELLOW}设置 Go 代理...${NC}"
    go env -w GOPROXY=https://goproxy.cn,direct
    echo -e "${GREEN}✓ 已设置 GOPROXY=https://goproxy.cn,direct${NC}"
    echo ""
fi

# 检查项目目录
if [[ ! -f "mcp/go.mod" ]]; then
    echo -e "${RED}✗ 请在项目根目录运行此脚本${NC}"
    exit 1
fi

# 下载依赖
echo -e "${YELLOW}下载依赖...${NC}"
cd mcp
go mod download
echo -e "${GREEN}✓ 依赖下载完成${NC}"
echo ""

# 编译项目
echo -e "${YELLOW}编译项目...${NC}"
mkdir -p ../bin
go build -o ../bin/xiaohongshu-agent .
echo -e "${GREEN}✓ 编译完成: bin/xiaohongshu-agent${NC}"
echo ""

# 返回根目录
cd ..

# 运行服务
echo -e "${YELLOW}启动服务...${NC}"
echo ""
echo "服务地址:"
echo "  - HTTP: http://localhost:18060"
echo "  - 健康检查: http://localhost:18060/health"
echo "  - API文档: http://localhost:18060/api/v1"
echo ""
echo "按 Ctrl+C 停止服务"
echo ""

./bin/xiaohongshu-agent
