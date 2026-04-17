# 小红书 Agent 一站式能力包

> 让 OpenClaw Agent 拥有完整的小红书自动化能力

## 🎯 项目愿景

创建一个**开箱即用**的小红书自动化解决方案，无需额外配置skill，一键安装即可使用。

## ✨ 核心特性

- 🚀 **一键安装** - 单命令完成所有依赖安装
- 🔐 **自动登录** - 扫码登录，Cookie自动管理
- 📝 **图文发布** - 支持多图、Markdown内容
- 🎬 **视频发布** - 支持视频上传与自动发布
- 💬 **互动管理** - 评论、点赞、收藏
- 🔍 **内容搜索** - 关键词搜索与数据获取
- 📊 **数据统计** - 发布效果追踪
- ⏰ **定时任务** - 定时发布与自动化

## 🏗️ 架构设计

```
┌─────────────────────────────────────────────────────────┐
│                    OpenClaw Agent                        │
│                  (Claude Sonnet 4.6)                     │
└────────────────────┬────────────────────────────────────┘
                     │
                     │ 调用
                     ▼
┌─────────────────────────────────────────────────────────┐
│              xiaohongshu-agent Skill                     │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │  发布模块     │  │  互动模块     │  │  数据模块     │  │
│  │  - 图文发布   │  │  - 评论管理   │  │  - 搜索内容   │  │
│  │  - 视频发布   │  │  - 点赞收藏   │  │  - 数据统计   │  │
│  │  - 批量发布   │  │  - 关注管理   │  │  - 推荐列表   │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
└────────────────────┬────────────────────────────────────┘
                     │
                     │ HTTP API
                     ▼
┌─────────────────────────────────────────────────────────┐
│              xiaohongshu-mcp Service                     │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │  认证服务     │  │  API网关      │  │  任务队列     │  │
│  │  - 登录管理   │  │  - REST API   │  │  - 异步处理   │  │
│  │  - Cookie池   │  │  - WebSocket  │  │  - 定时任务   │  │
│  │  - 会话保持   │  │  - 认证中间件 │  │  - 重试机制   │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
│                                                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │  浏览器引擎   │  │  数据存储     │  │  监控日志     │  │
│  │  - Playwright│  │  - SQLite     │  │  - 结构化日志 │  │
│  │  - 反检测     │  │  - 缓存层     │  │  - 性能监控   │  │
│  │  - 自动化     │  │  - 配置管理   │  │  - 错误追踪   │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
└────────────────────┬────────────────────────────────────┘
                     │
                     │ 网络请求
                     ▼
┌─────────────────────────────────────────────────────────┐
│                  小红书平台 API                           │
│              (xiaohongshu.com)                          │
└─────────────────────────────────────────────────────────┘
```

## 📁 项目结构

```
xiaohongshu-agent/
├── 📦 mcp/                          # MCP 服务层
│   ├── cmd/                         # 命令行工具
│   │   ├── server/                  # MCP服务器
│   │   └── login/                   # 登录工具
│   ├── internal/                    # 内部模块
│   │   ├── api/                     # API处理
│   │   ├── auth/                    # 认证管理
│   │   ├── browser/                 # 浏览器引擎
│   │   ├── cache/                   # 缓存层
│   │   ├── config/                  # 配置管理
│   │   ├── models/                  # 数据模型
│   │   ├── queue/                   # 任务队列
│   │   ├── service/                 # 业务逻辑
│   │   └── storage/                 # 数据存储
│   ├── pkg/                         # 公共包
│   │   ├── xiaohongshu/             # 小红书API客户端
│   │   └── utils/                   # 工具函数
│   ├── go.mod
│   ├── go.sum
│   └── Makefile
│
├── 🎨 skill/                        # OpenClaw Skill
│   ├── SKILL.md                     # Skill定义
│   ├── scripts/                     # 脚本工具
│   │   ├── install.sh               # 一键安装
│   │   ├── uninstall.sh             # 卸载脚本
│   │   ├── publish_image.py         # 图文发布
│   │   ├── publish_video.py         # 视频发布
│   │   ├── search.py                # 搜索工具
│   │   └── stats.py                 # 数据统计
│   ├── references/                  # 参考文档
│   │   ├── API_REFERENCE.md         # API文档
│   │   ├── USAGE_GUIDE.md           # 使用指南
│   │   └── TROUBLESHOOTING.md       # 故障排查
│   └── tests/                       # 测试用例
│
├── 🔧 configs/                      # 配置文件
│   ├── config.yaml                  # 默认配置
│   ├── config.dev.yaml              # 开发配置
│   └── config.prod.yaml             # 生产配置
│
├── 📚 docs/                         # 项目文档
│   ├── ARCHITECTURE.md              # 架构设计
│   ├── DEVELOPMENT.md               # 开发指南
│   ├── DEPLOYMENT.md                # 部署指南
│   └── CHANGELOG.md                 # 更新日志
│
├── 🐳 docker/                       # Docker支持
│   ├── Dockerfile
│   └── docker-compose.yaml
│
├── ⚡ install.sh                    # 一键安装入口
├── 📄 README.md                     # 项目说明
├── 📄 LICENSE                       # 开源协议
└── 📄 Makefile                      # 构建脚本
```

## 🚀 快速开始

### 安装

```bash
# 一键安装
curl -fsSL https://raw.githubusercontent.com/Cody-Chan/xiaohongshu-agent/main/install.sh | bash

# 或者克隆后安装
git clone https://github.com/Cody-Chan/xiaohongshu-agent.git
cd xiaohongshu-agent
./install.sh
```

### 登录

```bash
# 扫码登录
xiaohongshu-login
```

### 使用

在 OpenClaw 中直接使用：

```
帮我发布一条小红书图文：
标题：AI新技术分享
内容：今天学习了Claude Sonnet 4.6的新特性...
图片：/path/to/image.jpg
```

## 🛠️ API 接口

### 发布接口

**图文发布**
```bash
POST /api/v1/publish/image
{
  "title": "标题",
  "content": "内容",
  "images": ["/path/to/image.jpg"],
  "tags": ["标签1", "标签2"]
}
```

**视频发布**
```bash
POST /api/v1/publish/video
{
  "title": "标题",
  "content": "内容",
  "video": "/path/to/video.mp4",
  "tags": ["标签1", "标签2"]
}
```

### 搜索接口

```bash
POST /api/v1/search
{
  "keyword": "关键词",
  "limit": 20
}
```

### 数据接口

```bash
GET /api/v1/stats
```

## 📊 技术栈

### MCP 服务层
- **语言**: Go 1.24+
- **Web框架**: Gin
- **浏览器**: Playwright (go-rod)
- **存储**: SQLite + BadgerDB
- **队列**: Asynq

### Skill 层
- **语言**: Python 3.10+ / Bash
- **依赖**: requests, playwright-python

## 🗺️ 开发路线

### Phase 1: 基础重构 (Week 1-2)
- [ ] 代码结构优化
- [ ] 错误处理改进
- [ ] 日志系统完善
- [ ] 配置管理重构

### Phase 2: 功能增强 (Week 3-4)
- [ ] Cookie自动刷新
- [ ] 任务队列实现
- [ ] 批量发布功能
- [ ] 定时发布功能

### Phase 3: 稳定性优化 (Week 5-6)
- [ ] 反检测机制
- [ ] 重试策略
- [ ] 性能优化
- [ ] 监控告警

### Phase 4: 文档与发布 (Week 7)
- [ ] 完善文档
- [ ] 测试覆盖
- [ ] 发布 v1.0

## 🤝 贡献指南

欢迎贡献代码、报告问题或提出建议！

## 📄 开源协议

MIT License

## 🙏 致谢

本项目基于 [xpzouying/xiaohongshu-mcp](https://github.com/xpzouying/xiaohongshu-mcp) 开发，感谢原作者的贡献！
