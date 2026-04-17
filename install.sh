#!/bin/bash
# 小红书 Agent 一键安装脚本
# 作者: Cody-Chan
# 版本: v1.0.0

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检测操作系统
detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux";;
        Darwin*)    echo "darwin";;
        CYGWIN*)    echo "windows";;
        MINGW*)     echo "windows";;
        *)          echo "unknown";;
    esac
}

# 检测架构
detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64)   echo "amd64";;
        arm64|aarch64)  echo "arm64";;
        *)              echo "unknown";;
    esac
}

# 检查命令是否存在
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# 安装依赖
install_dependencies() {
    log_info "检查依赖..."
    
    # 检查 Go
    if ! command_exists go; then
        log_warn "Go 未安装，正在安装..."
        if command_exists apt-get; then
            sudo apt-get update
            sudo apt-get install -y golang-go
        elif command_exists brew; then
            brew install go
        else
            log_error "请手动安装 Go: https://golang.org/doc/install"
            exit 1
        fi
    fi
    
    # 检查 Python
    if ! command_exists python3; then
        log_warn "Python3 未安装，正在安装..."
        if command_exists apt-get; then
            sudo apt-get install -y python3 python3-pip
        elif command_exists brew; then
            brew install python3
        else
            log_error "请手动安装 Python3"
            exit 1
        fi
    fi
    
    # 检查 Node.js
    if ! command_exists node; then
        log_warn "Node.js 未安装，正在安装..."
        if command_exists apt-get; then
            curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
            sudo apt-get install -y nodejs
        elif command_exists brew; then
            brew install node
        else
            log_error "请手动安装 Node.js: https://nodejs.org/"
            exit 1
        fi
    fi
    
    log_success "依赖检查完成"
}

# 下载二进制文件
download_binaries() {
    local os=$(detect_os)
    local arch=$(detect_arch)
    
    log_info "下载二进制文件 (${os}-${arch})..."
    
    # 创建目录
    mkdir -p ~/.openclaw/mcp
    cd ~/.openclaw/mcp
    
    # 下载 MCP 服务
    local download_url="https://github.com/Cody-Chan/xiaohongshu-agent/releases/latest/download/xiaohongshu-mcp-${os}-${arch}.tar.gz"
    
    if curl -fsSL "$download_url" -o xiaohongshu-mcp.tar.gz; then
        tar -xzf xiaohongshu-mcp.tar.gz
        chmod +x xiaohongshu-mcp xiaohongshu-login
        rm xiaohongshu-mcp.tar.gz
        log_success "二进制文件下载完成"
    else
        log_error "下载失败，请检查网络连接"
        exit 1
    fi
}

# 安装 Playwright
install_playwright() {
    log_info "安装 Playwright..."
    
    # 安装 Playwright Python 包
    pip3 install playwright
    
    # 安装浏览器
    playwright install chromium
    playwright install-deps chromium
    
    log_success "Playwright 安装完成"
}

# 安装 Skill
install_skill() {
    log_info "安装 OpenClaw Skill..."
    
    local skill_dir=~/.openclaw/skills/xiaohongshu-agent
    
    # 创建目录
    mkdir -p "$skill_dir"
    
    # 下载 Skill 文件
    # 这里假设已经克隆了项目，从项目目录复制
    if [ -d "./skill" ]; then
        cp -r ./skill/* "$skill_dir/"
        log_success "Skill 安装完成"
    else
        log_warn "Skill 文件不存在，跳过"
    fi
}

# 配置服务
configure_service() {
    log_info "配置服务..."
    
    # 创建配置目录
    mkdir -p ~/.openclaw/mcp/data
    mkdir -p ~/.openclaw/mcp/logs
    
    # 创建默认配置
    if [ ! -f ~/.openclaw/mcp/config.yaml ]; then
        cat > ~/.openclaw/mcp/config.yaml <<EOF
server:
  host: "0.0.0.0"
  port: 18060

auth:
  cookie_file: "./data/cookies.json"
  refresh_interval: "24h"

browser:
  headless: true
  timeout: "5m"

log:
  level: "info"
  format: "json"
  output: "./logs/app.log"
EOF
        log_success "配置文件创建完成"
    fi
}

# 启动服务
start_service() {
    log_info "启动服务..."
    
    cd ~/.openclaw/mcp
    
    # 检查是否已在运行
    if pgrep -f xiaohongshu-mcp > /dev/null; then
        log_warn "服务已在运行"
        return
    fi
    
    # 启动服务
    nohup ./xiaohongshu-mcp > mcp.log 2>&1 &
    
    # 等待服务启动
    sleep 3
    
    # 检查服务状态
    if curl -s http://localhost:18060/health | grep -q "ok"; then
        log_success "服务启动成功"
        log_info "服务地址: http://localhost:18060"
    else
        log_error "服务启动失败，请检查日志: ~/.openclaw/mcp/mcp.log"
        exit 1
    fi
}

# 登录提示
login_prompt() {
    echo ""
    echo -e "${GREEN}================================${NC}"
    echo -e "${GREEN}  安装完成！${NC}"
    echo -e "${GREEN}================================${NC}"
    echo ""
    echo "下一步："
    echo ""
    echo "1. 登录小红书："
    echo "   ${YELLOW}xiaohongshu-login${NC}"
    echo ""
    echo "2. 在 OpenClaw 中使用："
    echo "   \"帮我发布一条小红书图文...\""
    echo ""
    echo "3. 查看服务状态："
    echo "   curl http://localhost:18060/health"
    echo ""
}

# 主函数
main() {
    echo -e "${BLUE}"
    echo "╔══════════════════════════════════════╗"
    echo "║   小红书 Agent 一键安装脚本 v1.0.0    ║"
    echo "╚══════════════════════════════════════╝"
    echo -e "${NC}"
    
    install_dependencies
    download_binaries
    install_playwright
    install_skill
    configure_service
    start_service
    login_prompt
}

# 执行主函数
main "$@"
