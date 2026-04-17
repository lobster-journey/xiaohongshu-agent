---
name: xiaohongshu-agent
description: 小红书一站式自动化能力包，支持图文发布、视频发布、评论互动、内容搜索等功能。让OpenClaw Agent拥有完整的小红书自动化能力。
metadata:
  openclaw:
    emoji: 📱
    requires:
      bins: [xiaohongshu-mcp, xiaohongshu-login]
    install:
      - id: curl
        kind: curl
        url: GitHub 搜索：xiaohongshu-agent/main/install.sh
        label: 一键安装小红书Agent
---

# 小红书 Agent Skill

让 OpenClaw Agent 拥有完整的小红书自动化能力。

## ✨ 核心功能

### 📝 内容发布
- 图文笔记发布（支持多图）
- 视频笔记发布
- 批量发布
- 定时发布

### 💬 互动管理
- 自动评论
- 点赞收藏
- 关注管理

### 🔍 数据获取
- 内容搜索
- 推荐列表
- 笔记详情
- 用户统计

## 🚀 快速开始

### 1. 安装

```bash
# 一键安装
curl -fsSL GitHub 搜索：xiaohongshu-agent/main/install.sh | bash
```

### 2. 登录

```bash
# 扫码登录
xiaohongshu-login
```

### 3. 使用

在 OpenClaw 中直接使用自然语言：

```
帮我发布一条小红书图文：
标题：AI新技术分享
内容：今天学习了Claude Sonnet 4.6的新特性...
图片：/path/to/image.jpg
```

## 📖 使用示例

### 发布图文

```
发布小红书图文笔记：
- 标题：技术分享
- 内容：Markdown格式的内容
- 图片：使用本地路径
- 标签：AI、技术
```

### 发布视频

```
帮我发一个小红书视频：
- 视频文件：/home/user/video.mp4
- 标题：我的第一个视频
- 描述：视频内容介绍
```

### 搜索内容

```
搜索小红书上关于"AI技术"的内容
```

### 数据统计

```
查看我的小红书数据统计
```

## 🔧 配置

配置文件位置：`~/.openclaw/mcp/config.yaml`

```yaml
server:
  host: "0.0.0.0"
  port: 18060

browser:
  headless: true  # 无头模式
  timeout: "5m"   # 超时时间

log:
  level: "info"   # 日志级别
```

## 📡 API 接口

MCP 服务运行在 `http://localhost:18060`

### 发布接口

```bash
# 图文发布
curl -X POST http://localhost:18060/api/v1/publish/image \
  -H "Content-Type: application/json" \
  -d '{
    "title": "标题",
    "content": "内容",
    "images": ["/path/to/image.jpg"]
  }'

# 视频发布
curl -X POST http://localhost:18060/api/v1/publish/video \
  -H "Content-Type: application/json" \
  -d '{
    "title": "标题",
    "content": "内容",
    "video": "/path/to/video.mp4"
  }'
```

### 搜索接口

```bash
curl -X POST http://localhost:18060/api/v1/search \
  -H "Content-Type: application/json" \
  -d '{"keyword": "关键词"}'
```

## 🛠️ 故障排查

### 服务未启动

```bash
# 检查服务状态
curl http://localhost:18060/health

# 启动服务
cd ~/.openclaw/mcp
./xiaohongshu-mcp
```

### Cookie 过期

```bash
# 重新登录
xiaohongshu-login
```

### 查看日志

```bash
tail -f ~/.openclaw/mcp/logs/app.log
```

## ⚠️ 注意事项

- 标题最多20字
- 正文最多1000字
- 图片支持1-9张
- 视频建议不超过1GB
- 发布间隔建议60秒以上

## 📚 更多文档

- [完整API文档](./references/API_REFERENCE.md)
- [架构设计](./docs/ARCHITECTURE.md)
- [开发指南](./docs/DEVELOPMENT.md)
