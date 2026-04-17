# 📖 使用指南

> xiaohongshu-agent 完整使用指南

---

## 🚀 快速开始

### 1. 安装

```bash
# 方式一：一键安装
curl -fsSL https://github.com/lobster-journey/xiaohongshu-agent/main/install.sh | bash

# 方式二：克隆后安装
git clone https://github.com/lobster-journey/xiaohongshu-agent.git
cd xiaohongshu-agent
make install
```

---

### 2. 登录账号

```bash
# 启动登录工具
xiaohongshu-login

# 或使用make命令
make login
```

**登录流程**：
1. 打开小红书登录页面
2. 使用手机APP扫码登录
3. 等待登录成功提示
4. Cookie自动保存

---

### 3. 启动服务

```bash
# 启动MCP服务
make run

# 或直接运行
cd mcp && ./xiaohongshu-mcp --port=18060
```

---

### 4. 检查状态

```bash
# 检查服务状态
curl http://localhost:18060/health
```

---

## 📝 发布内容

### 图文发布

**命令行方式**：
```bash
python3 skill/scripts/publish_image.py \
  --title "AI新技术分享" \
  --content "今天学习了Claude的新特性..." \
  --images /path/to/image1.jpg /path/to/image2.jpg \
  --tags AI Claude 效率
```

**API方式**：
```bash
curl -X POST http://localhost:18060/api/v1/publish/image \
  -H "Content-Type: application/json" \
  -d '{
    "title": "标题",
    "content": "内容",
    "images": ["/path/to/image.jpg"],
    "tags": ["标签"]
  }'
```

**OpenClaw中使用**：
```
帮我发布一条小红书图文：
标题：AI新技术分享
内容：今天学习了Claude的新特性...
图片：/path/to/image.jpg
```

---

### 视频发布

**命令行方式**：
```bash
python3 skill/scripts/publish_video.py \
  --title "AI教程视频" \
  --content "一个完整的AI使用教程..." \
  --video /path/to/video.mp4 \
  --cover /path/to/cover.jpg \
  --tags AI 教程
```

**API方式**：
```bash
curl -X POST http://localhost:18060/api/v1/publish/video \
  -H "Content-Type: application/json" \
  -d '{
    "title": "标题",
    "content": "内容",
    "video": "/path/to/video.mp4",
    "cover": "/path/to/cover.jpg"
  }'
```

---

## 🔍 搜索内容

**命令行方式**：
```bash
# 普通搜索
python3 skill/scripts/search.py "AI技巧"

# 热门排序
python3 skill/scripts/search.py "AI技巧" --sort hot

# JSON输出
python3 skill/scripts/search.py "AI技巧" --output json
```

**API方式**：
```bash
curl "http://localhost:18060/api/v1/search?keyword=AI技巧&limit=20"
```

---

## 📊 数据统计

**命令行方式**：
```bash
# 数据概览
python3 skill/scripts/stats.py

# 账号统计
python3 skill/scripts/stats.py --type account

# 近7天统计
python3 skill/scripts/stats.py --type recent --days 7

# 单篇内容统计
python3 skill/scripts/stats.py --type post --post-id 1234567890
```

---

## 💬 互动管理

### 点赞

```bash
curl -X POST http://localhost:18060/api/v1/interact/like \
  -H "Content-Type: application/json" \
  -d '{"post_id": "1234567890"}'
```

### 收藏

```bash
curl -X POST http://localhost:18060/api/v1/interact/collect \
  -H "Content-Type: application/json" \
  -d '{"post_id": "1234567890"}'
```

### 评论

```bash
curl -X POST http://localhost:18060/api/v1/interact/comment \
  -H "Content-Type: application/json" \
  -d '{
    "post_id": "1234567890",
    "content": "很棒的内容！"
  }'
```

---

## ⏰ 定时任务

### 使用OpenClaw定时功能

```python
# 设置定时发布
from openclaw import cron

# 每天上午10点发布
cron.schedule("0 10 * * *", publish_morning_post)

# 每天下午2点发布
cron.schedule("0 14 * * *", publish_afternoon_post)
```

---

## 🔧 配置管理

### 配置文件

**位置**: `configs/config.yaml`

**内容**:
```yaml
server:
  port: 18060
  host: "0.0.0.0"

xiaohongshu:
  cookie_path: "~/.xiaohongshu-agent/cookies"
  user_agent: "Mozilla/5.0..."

publish:
  daily_limit: 5
  min_interval: 30  # 分钟

log:
  level: "info"
  path: "~/.xiaohongshu-agent/logs"
```

---

## 📚 相关文档

- [API参考](./API_REFERENCE.md)
- [故障排查](./TROUBLESHOOTING.md)
- [开发指南](../../docs/DEVELOPMENT.md)

---

**Created by 🦞 Lobster Journey Studio**
