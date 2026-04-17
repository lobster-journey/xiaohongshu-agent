# 安装指南

## 环境要求

- **Go**: 1.24 或更高版本
- **Chrome/Chromium**: 任何现代浏览器
- **操作系统**: Linux / macOS / Windows

---

## 安装步骤

### 1. 安装 Go 环境

#### Linux (Ubuntu/Debian)

```bash
# 下载 Go 1.24
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz

# 解压到 /usr/local
sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz

# 添加到 PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# 验证安装
go version
```

#### macOS

```bash
# 使用 Homebrew
brew install go

# 或下载官方安装包
# https://go.dev/dl/go1.24.0.darwin-amd64.pkg
```

#### Windows

```powershell
# 使用 Chocolatey
choco install golang

# 或下载官方安装包
# https://go.dev/dl/go1.24.0.windows-amd64.msi
```

---

### 2. 克隆项目

```bash
# 克隆仓库
git clone GitHub 搜索：xiaohongshu-agent.git

# 进入项目目录
cd xiaohongshu-agent
```

---

### 3. 安装依赖

```bash
# 进入 MCP 目录
cd mcp

# 下载依赖
go mod download

# 验证依赖
go mod verify
```

---

### 4. 构建项目

```bash
# 返回项目根目录
cd ..

# 使用 Makefile 构建
make deps
make build

# 或手动构建
cd mcp && go build -o ../bin/xiaohongshu-agent .
```

---

### 5. 运行服务

```bash
# 使用 Makefile
make run

# 或直接运行
./bin/xiaohongshu-agent

# 或使用 Docker
make docker-build
make docker-run
```

---

## 快速测试

### 测试健康检查

```bash
curl http://localhost:18060/health
```

**期望输出：**
```json
{
  "status": "ok",
  "service": "xiaohongshu-agent",
  "version": "1.0.0"
}
```

### 测试登录状态

```bash
curl http://localhost:18060/api/v1/auth/status
```

### 测试发布接口

```bash
curl -X POST http://localhost:18060/api/v1/publish/image \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试标题",
    "content": "测试内容",
    "images": ["/path/to/image.jpg"],
    "tags": ["测试", "小红书"]
  }'
```

---

## 环境变量

创建 `.env` 文件或设置环境变量：

```bash
# 浏览器配置
BROWSER_HEADLESS=true          # 是否无头模式
BROWSER_BIN_PATH=/usr/bin/chromium-browser  # 浏览器路径（可选）
BROWSER_PROXY=http://127.0.0.1:8080  # 代理地址（可选）

# Cookie 配置
XHS_COOKIES_PATH=/tmp/xhs_cookies.json  # Cookie 存储路径

# 服务配置
PORT=18060  # 服务端口
```

---

## 常见问题

### Q1: Go 安装失败？

**A:**
- 检查网络连接
- 使用国内镜像：
  ```bash
  export GOPROXY=https://goproxy.cn,direct
  go mod download
  ```

### Q2: 依赖下载慢？

**A:**
```bash
# 设置代理
go env -w GOPROXY=https://goproxy.cn,direct

# 重新下载
go mod download
```

### Q3: 浏览器启动失败？

**A:**
- 确保安装了 Chrome 或 Chromium
- 设置浏览器路径：
  ```bash
  export BROWSER_BIN_PATH=/usr/bin/google-chrome
  ```

### Q4: Cookie 无法保存？

**A:**
- 检查文件权限
- 设置可写路径：
  ```bash
  export XHS_COOKIES_PATH=/tmp/xhs_cookies.json
  ```

### Q5: 发布失败？

**A:**
- 先登录：访问 http://localhost:18060/api/v1/auth/login
- 检查登录状态：访问 http://localhost:18060/api/v1/auth/status
- 查看日志输出

---

## Docker 部署

### 构建镜像

```bash
make docker-build
```

### 运行容器

```bash
make docker-run
```

### 查看日志

```bash
docker logs -f xiaohongshu-agent
```

### 停止容器

```bash
make docker-stop
```

---

## 下一步

1. ✅ 安装完成
2. ✅ 服务启动
3. ⏳ 完善浏览器自动化
4. ⏳ 实现完整发布流程
5. ⏳ 添加更多功能（搜索、互动、统计）

---

## 项目结构

```
xiaohongshu-agent/
├── mcp/                           # MCP 服务
│   ├── main.go                   # 入口文件
│   ├── internal/
│   │   ├── browser/              # 浏览器自动化
│   │   ├── xiaohongshu/          # 小红书操作
│   │   │   ├── login.go         # 登录
│   │   │   └── publish.go       # 发布
│   │   ├── cookie/               # Cookie 管理
│   │   └── service/              # 业务逻辑
│   └── go.mod                    # 依赖配置
├── skill/                         # OpenClaw Skill
├── docker/                        # Docker 配置
├── docs/                          # 文档
└── Makefile                       # 构建脚本
```

---

## 开发指南

### 本地开发

```bash
# 开发模式（需要安装 air）
make dev

# 运行测试
make test

# 清理构建
make clean
```

### 代码结构

- **browser/**: 浏览器自动化核心
- **xiaohongshu/**: 小红书平台操作
- **cookie/**: Cookie 持久化
- **service/**: REST API 服务

---

**祝你使用愉快！** 🦞

如有问题，请提交 Issue：GitHub 搜索：xiaohongshu-agent/issues
