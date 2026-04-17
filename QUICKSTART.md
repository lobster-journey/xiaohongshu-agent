# 快速开始

## 🚀 5分钟启动服务

### 方式1：直接运行（推荐开发调试）

```bash
# 1. 进入项目目录
cd /home/gem/.openclaw/workspace/xiaohongshu-agent

# 2. 安装依赖
make deps

# 3. 安装Playwright（首次运行）
make install-playwright

# 4. 运行服务
make run
```

### 方式2：Docker运行（推荐生产环境）

```bash
# 1. 构建镜像
make docker-build

# 2. 启动容器
make docker-run

# 3. 查看日志
docker logs -f xiaohongshu-agent
```

---

## 📝 使用服务

### 1. 检查服务状态

```bash
curl http://localhost:18060/health
```

**返回：**
```json
{
  "status": "ok",
  "service": "xiaohongshu-agent",
  "version": "1.0.0"
}
```

### 2. 登录小红书

```bash
curl -X POST http://localhost:18060/api/v1/auth/login
```

**返回：**
```json
{
  "success": true,
  "message": "请在浏览器中完成登录",
  "login_url": "http://localhost:18060/login"
}
```

### 3. 发布图文

```bash
curl -X POST http://localhost:18060/api/v1/publish/image \
  -H "Content-Type: application/json" \
  -d '{
    "title": "我的第一篇笔记",
    "content": "这是内容",
    "images": ["https://example.com/image1.jpg"],
    "tags": ["测试", "小红书"]
  }'
```

---

## 🧪 测试API

```bash
# 运行测试脚本
bash test_api.sh
```

---

## 📚 API文档

访问：http://localhost:18060/api/v1

### 认证
- `GET /api/v1/auth/status` - 获取登录状态
- `POST /api/v1/auth/login` - 登录
- `POST /api/v1/auth/logout` - 登出

### 发布
- `POST /api/v1/publish/image` - 发布图文
- `POST /api/v1/publish/video` - 发布视频
- `POST /api/v1/publish/batch` - 批量发布
- `GET /api/v1/publish/task/:id` - 查询任务状态

### 搜索
- `GET /api/v1/search/posts?keyword=xxx` - 搜索笔记
- `GET /api/v1/search/users?keyword=xxx` - 搜索用户

### 互动
- `POST /api/v1/interaction/comment` - 评论
- `POST /api/v1/interaction/like/:id` - 点赞
- `POST /api/v1/interaction/follow/:id` - 关注

### 统计
- `GET /api/v1/stats/overview` - 数据概览
- `GET /api/v1/stats/post/:id` - 笔记统计

---

## 🔧 开发指南

### 项目结构

```
xiaohongshu-agent/
├── mcp/                    # MCP服务
│   ├── main.go            # 入口
│   ├── internal/          # 内部模块
│   │   ├── browser/       # 浏览器自动化
│   │   ├── cookie/        # Cookie管理
│   │   ├── queue/         # 任务队列
│   │   └── service/       # 业务逻辑
│   └── go.mod             # 依赖
├── skill/                  # OpenClaw Skill
├── docs/                   # 文档
├── docker/                 # Docker配置
├── configs/               # 配置文件
└── Makefile               # 构建脚本
```

### 开发流程

1. **修改代码**
2. **本地测试**：`make run`
3. **运行测试**：`make test`
4. **提交代码**：`git add . && git commit -m "xxx"`

---

## ❓ 常见问题

### Q: Playwright安装失败？

A: 手动安装：
```bash
cd mcp
go get github.com/playwright-community/playwright-go
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium
```

### Q: 端口被占用？

A: 修改端口：
```bash
PORT=8080 make run
```

### Q: Cookie过期怎么办？

A: 重新登录：
```bash
curl -X POST http://localhost:18060/api/v1/auth/login
```

---

## 📖 下一步

1. **完善浏览器自动化** - 实现完整的发布流程
2. **添加错误处理** - 提高稳定性
3. **实现Cookie刷新** - 自动续期
4. **添加日志系统** - 方便调试
5. **编写单元测试** - 提高代码质量

---

**祝你使用愉快！** 🦞
