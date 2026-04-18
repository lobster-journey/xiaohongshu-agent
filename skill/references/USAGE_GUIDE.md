# 🦞 小红书Agent使用指南

## 📋 快速开始

### 前置要求

1. **Cookie文件** - 小红书创作者平台Cookie
2. **Python 3.8+** - 运行Python脚本
3. **OpenClaw环境**（可选）- 定时任务支持

### 安装步骤

```bash
# 1. 克隆项目
git clone https://github.com/lobster-journey/xiaohongshu-agent.git
cd xiaohongshu-agent

# 2. 安装Python依赖
pip install requests playwright

# 3. 配置Cookie
mkdir -p ~/.openclaw/workspace/config/cookies
# 将Cookie保存到 xiaohongshu.json

# 4. 测试运行
python skill/scripts/publish_simple.py --help
```

---

## 🚀 核心功能

### 1️⃣ 内容发布

#### 图文发布

```bash
# 基础发布
python skill/scripts/publish_simple.py \
  --title "标题" \
  --content "内容" \
  --images "图片1.jpg,图片2.jpg"

# 完整参数
python skill/scripts/publish_simple.py \
  --title "标题" \
  --content "内容\n\n#话题标签" \
  --images "图片1.jpg,图片2.jpg,图片3.jpg" \
  --tags "AI,科技,人工智能" \
  --publish-time "2026-04-19 12:00" \
  --private
```

#### 视频发布

```bash
python skill/scripts/publish_video.py \
  --title "视频标题" \
  --video "video.mp4" \
  --cover "cover.jpg"
```

---

### 2️⃣ 互动管理

#### 单个互动操作

```bash
# 点赞笔记
python skill/scripts/interact.py \
  --note-id "abc123" \
  --action like

# 收藏笔记
python skill/scripts/interact.py \
  --note-id "abc123" \
  --action collect

# 评论笔记
python skill/scripts/interact.py \
  --note-id "abc123" \
  --action comment \
  --content "很棒的内容！"

# 关注用户
python skill/scripts/interact.py \
  --user-id "user456" \
  --action follow
```

#### 批量互动操作

```bash
# 创建笔记ID列表文件
cat > notes.txt << EOF
note_id_1
note_id_2
note_id_3
EOF

# 批量点赞和收藏
python skill/scripts/interact.py \
  --batch-file notes.txt \
  --actions like,collect

# 批量评论
python skill/scripts/interact.py \
  --batch-file notes.txt \
  --actions comment
```

---

### 3️⃣ 数据采集

#### 采集笔记数据

```bash
# 单篇笔记
python skill/scripts/collect_data.py \
  --note-id "abc123"

# 批量采集
cat > notes.txt << EOF
note_id_1
note_id_2
note_id_3
EOF

python skill/scripts/collect_data.py \
  --batch-file notes.txt \
  --type daily-stats
```

#### 采集用户数据

```bash
python skill/scripts/collect_data.py \
  --user-id "user456"
```

#### 采集搜索结果

```bash
python skill/scripts/collect_data.py \
  --keyword "AI人工智能" \
  --limit 50
```

#### 导出数据

```bash
# 导出所有数据为CSV
python skill/scripts/collect_data.py --export-csv

# 导出特定类型数据
python skill/scripts/collect_data.py \
  --export-csv \
  --export-type note
```

---

### 4️⃣ 搜索功能

```bash
# 基础搜索
python skill/scripts/search.py "AI人工智能"

# 限制数量
python skill/scripts/search.py "AI" --limit 50

# 排序方式
python skill/scripts/search.py "AI" --sort hot

# 输出JSON格式
python skill/scripts/search.py "AI" --output json
```

---

### 5️⃣ 数据统计

```bash
# 账号统计
python skill/scripts/stats.py --type account

# 近期统计
python skill/scripts/stats.py --type recent --days 7

# 单篇内容统计
python skill/scripts/stats.py --type post --post-id "abc123"

# 数据概览
python skill/scripts/stats.py --type overview
```

---

## ⚙️ 高级功能

### 定时发布

结合OpenClaw的Cron功能实现定时发布：

```bash
# 创建定时任务
openclaw cron create

# 配置任务
名称：每日内容发布
Cron表达式：0 9 * * *
时区：Asia/Shanghai
命令：python /path/to/publish_simple.py --title "标题" --content "内容"
```

### 自动化运营

创建自动化运营脚本：

```bash
# daily_operation.sh
#!/bin/bash

# 1. 采集热点
python skill/scripts/search.py "AI" --limit 10 > hotspots.json

# 2. 生成内容（需要LLM API）
python generate_content.py --input hotspots.json --output content.json

# 3. 发布内容
python skill/scripts/publish_simple.py \
  --title "$(cat content.json | jq -r '.title')" \
  --content "$(cat content.json | jq -r '.content')" \
  --images "$(cat content.json | jq -r '.images')"

# 4. 采集数据
python skill/scripts/collect_data.py --batch-file published_notes.txt --type daily-stats
```

---

## 📊 数据格式

### 笔记数据格式

```json
{
  "note_id": "abc123",
  "title": "笔记标题",
  "content": "笔记内容...",
  "author": {
    "user_id": "user123",
    "nickname": "作者昵称",
    "avatar": "https://..."
  },
  "stats": {
    "views": 1000,
    "likes": 100,
    "comments": 20,
    "collects": 50,
    "shares": 10
  },
  "publish_time": "2026-04-19 10:00:00",
  "collect_time": "2026-04-19 12:00:00"
}
```

### 每日统计格式

```json
{
  "date": "2026-04-19",
  "total_notes": 10,
  "total_views": 10000,
  "total_likes": 1000,
  "total_comments": 200,
  "total_collects": 500,
  "total_shares": 100,
  "details": [...]
}
```

---

## 🔧 故障排查

### Cookie失效

**症状**：API返回401或403错误

**解决方案**：
1. 重新登录小红书创作者平台
2. 导出新Cookie
3. 更新配置文件

### 发布失败

**症状**：发布返回错误

**可能原因**：
1. Cookie失效
2. 网络问题
3. 内容违规
4. 图片格式错误

**解决方案**：
1. 检查Cookie有效性
2. 检查网络连接
3. 检查内容是否合规
4. 检查图片格式和大小

### 数据采集失败

**症状**：采集不到数据

**可能原因**：
1. MCP服务未启动
2. 笔记ID错误
3. 网络问题

**解决方案**：
1. 启动MCP服务
2. 检查笔记ID是否正确
3. 检查网络连接

---

## ⚠️ 注意事项

### 频率限制

- 点赞：每分钟不超过5次
- 评论：每分钟不超过3次
- 关注：每小时不超过10人
- 发布：每天不超过10篇

### 内容合规

- ✅ 原创内容
- ✅ 真实信息
- ✅ 积极向上
- ❌ 违规内容
- ❌ 虚假信息
- ❌ 侵权内容

### 账号安全

- 定期更换Cookie
- 不要过度自动化
- 遵守平台规则
- 关注账号状态

---

## 📚 进阶使用

### MCP服务（需要Go环境）

```bash
# 安装Go
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# 编译MCP服务
cd mcp
go mod tidy
go build -o xhs-agent

# 运行服务
./xhs-agent

# 测试API
curl http://localhost:18060/health
```

### 集成到OpenClaw

```bash
# 安装Skill
openclaw skill install ./skill

# 使用Skill
openclaw skill run xiaohongshu-publish --title "标题" --content "内容"
```

---

## 🆘 获取帮助

### 命令行帮助

```bash
# 查看帮助
python skill/scripts/publish_simple.py --help
python skill/scripts/interact.py --help
python skill/scripts/collect_data.py --help
```

### 文档资源

- [API文档](./references/API_REFERENCE.md)
- [故障排查](./references/TROUBLESHOOTING.md)
- [开发指南](../docs/DEVELOPMENT.md)

### 问题反馈

- GitHub Issues: 提交问题和建议
- 项目文档: 查看详细文档

---

**祝你使用愉快！** 🦞
