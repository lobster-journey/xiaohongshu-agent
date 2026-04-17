# 更新日志

本文档记录项目的所有重要变更。

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
版本号遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [Unreleased]

### 新增
- 项目初始化
- 基础架构设计
- 核心文档编写

### 变更
- 无

### 修复
- 无

## [1.0.0] - 2026-04-XX

### 新增
- 🚀 一键安装脚本
- 📝 图文发布功能
- 🎬 视频发布功能
- 💬 评论互动功能
- 🔍 内容搜索功能
- 📊 数据统计功能
- ⏰ 定时发布功能
- 🔄 Cookie自动刷新
- 📡 REST API接口
- 🌐 WebSocket实时推送
- 📚 完整文档

### 技术栈
- Go 1.24+
- Python 3.10+
- Playwright
- SQLite
- Asynq

---

## 版本说明

- **[Unreleased]**: 开发中的功能
- **[1.0.0]**: 首个正式版本

## 升级指南

### 从原项目迁移

如果你之前使用 `xpzouying/xiaohongshu-mcp`：

1. 备份Cookie文件：
   ```bash
   cp ~/.openclaw/mcp/cookies.json ~/cookies.json.backup
   ```

2. 安装新版本：
   ```bash
   curl -fsSL GitHub 搜索：xiaohongshu-agent/main/install.sh | bash
   ```

3. 恢复Cookie：
   ```bash
   cp ~/cookies.json.backup ~/.openclaw/mcp/data/cookies.json
   ```

4. 重启服务：
   ```bash
   pkill -f xiaohongshu-mcp
   cd ~/.openclaw/mcp && ./xiaohongshu-mcp
   ```
