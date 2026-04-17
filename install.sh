#!/bin/bash

# xiaohongshu-agent 一键安装脚本
# Created by 🦞 Lobster Journey Studio

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logo
echo -e "${BLUE}"
echo "🦞 xiaohongshu-agent 安装程序"
echo "   Lobster Journey Studio"
echo -e "${NC}"

# 检测操作系统
OS="$(uname -s)"
case "$OS" in
    Linux*)     OS_TYPE="Linux";;
    Darwin*)    OS_TYPE="Mac";;
    CYGWIN*)    OS_TYPE="Cygwin";;
    MINGW*)     OS_TYPE="MinGW";;
    *)          OS_TYPE="Unknown";;
esac

echo -e "${GREEN}[✓]${NC} 检测到操作系统: $OS_TYPE"

# 检查依赖
check_dependencies() {
    echo -e "\n${YELLOW}[→]${NC} 检查依赖..."

    # Python
    if ! command -v python3 &> /dev/null; then
        echo -e "${RED}[✗]${NC} 未找到 Python3，请先安装"
        exit 1
    fi
    echo -e "${GREEN}[✓]${NC} Python3: $(python3 --version)"

    # pip
    if ! command -v pip3 &> /dev/null; then
        echo -e "${RED}[✗]${NC} 未找到 pip3，请先安装"
        exit 1
    fi
    echo -e "${GREEN}[✓]${NC} pip3 已安装"

    # Go (可选)
    if command -v go &> /dev/null; then
        echo -e "${GREEN}[✓]${NC} Go: $(go version)"
    else
        echo -e "${YELLOW}[!]${NC} Go 未安装（可选，用于MCP服务）"
    fi
}

# 安装Python依赖
install_python_deps() {
    echo -e "\n${YELLOW}[→]${NC} 安装Python依赖..."

    if [ -f "skill/requirements.txt" ]; then
        pip3 install -r skill/requirements.txt
        echo -e "${GREEN}[✓]${NC} Python依赖安装完成"
    else
        echo -e "${YELLOW}[!]${NC} 未找到 requirements.txt"
    fi
}

# 安装MCP服务
install_mcp_service() {
    echo -e "\n${YELLOW}[→]${NC} 安装MCP服务..."

    if [ -f "mcp/go.mod" ]; then
        cd mcp
        if command -v go &> /dev/null; then
            go mod download
            go build -o xiaohongshu-mcp ./cmd/server
            echo -e "${GREEN}[✓]${NC} MCP服务编译完成"
        else
            echo -e "${YELLOW}[!]${NC} 跳过MCP服务编译（Go未安装）"
        fi
        cd ..
    else
        echo -e "${YELLOW}[!]${NC} 未找到MCP源码"
    fi
}

# 配置环境
setup_config() {
    echo -e "\n${YELLOW}[→]${NC} 配置环境..."

    if [ ! -f "configs/config.yaml" ]; then
        cp configs/config.yaml configs/config.yaml.bak
        echo -e "${GREEN}[✓]${NC} 配置文件已备份"
    fi

    # 创建必要目录
    mkdir -p ~/.xiaohongshu-agent
    mkdir -p ~/.xiaohongshu-agent/cookies
    mkdir -p ~/.xiaohongshu-agent/logs

    echo -e "${GREEN}[✓]${NC} 环境配置完成"
}

# 安装Skill
install_skill() {
    echo -e "\n${YELLOW}[→]${NC} 安装OpenClaw Skill..."

    SKILL_DIR="$HOME/.openclaw/skills/xiaohongshu-agent"

    if [ -d "$SKILL_DIR" ]; then
        echo -e "${YELLOW}[!]${NC} Skill已存在，跳过安装"
    else
        mkdir -p "$SKILL_DIR"
        cp -r skill/* "$SKILL_DIR/"
        echo -e "${GREEN}[✓]${NC} Skill安装完成: $SKILL_DIR"
    fi
}

# 完成
show_success() {
    echo -e "\n${GREEN}"
    echo "╔══════════════════════════════════════════════╗"
    echo "║                                              ║"
    echo "║       🦞 安装完成！                         ║"
    echo "║                                              ║"
    echo "║  下一步：                                    ║"
    echo "║  1. 运行 xiaohongshu-login 登录账号         ║"
    echo "║  2. 在OpenClaw中使用小红书功能              ║"
    echo "║                                              ║"
    echo "║  文档：https://github.com/lobster-journey/   ║"
    echo "║       xiaohongshu-agent                      ║"
    echo "║                                              ║"
    echo "╚══════════════════════════════════════════════╝"
    echo -e "${NC}"
}

# 主流程
main() {
    check_dependencies
    install_python_deps
    install_mcp_service
    setup_config
    install_skill
    show_success
}

main
