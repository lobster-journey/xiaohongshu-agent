# 📚 API参考文档

> xiaohongshu-agent API完整参考

---

## 🔗 基础信息

**Base URL**: `http://localhost:18060`

**版本**: v1

**认证**: Cookie认证

---

## 📊 健康检查

### GET /health

检查服务状态

**响应**:
```json
{
  "success": true,
  "data": {
    "account": "ai-report",
    "service": "xiaohongshu-mcp",
    "status": "healthy"
  }
}
```

---

## 📝 发布接口

### POST /api/v1/publish/image

发布图文内容

**请求体**:
```json
{
  "title": "标题",
  "content": "内容（支持Markdown）",
  "images": [
    "/path/to/image1.jpg",
    "/path/to/image2.jpg"
  ],
  "tags": ["标签1", "标签2"],
  "draft": false
}
```

**响应**:
```json
{
  "success": true,
  "data": {
    "post_id": "1234567890",
    "url": "https://www.xiaohongshu.com/post/1234567890"
  }
}
```

---

### POST /api/v1/publish/video

发布视频内容

**请求体**:
```json
{
  "title": "标题",
  "content": "内容",
  "video": "/path/to/video.mp4",
  "cover": "/path/to/cover.jpg",
  "tags": ["标签1", "标签2"],
  "draft": false
}
```

**响应**:
```json
{
  "success": true,
  "data": {
    "post_id": "1234567890",
    "url": "https://www.xiaohongshu.com/post/1234567890"
  }
}
```

---

## 🔍 搜索接口

### GET /api/v1/search

搜索内容

**参数**:
- `keyword`: 关键词（必需）
- `limit`: 数量限制（默认20）
- `sort`: 排序方式（general/hot/newest）

**响应**:
```json
{
  "success": true,
  "data": {
    "notes": [
      {
        "id": "1234567890",
        "title": "标题",
        "content": "内容摘要",
        "likes": 100,
        "comments": 20,
        "author": {
          "id": "user123",
          "nickname": "作者昵称"
        }
      }
    ],
    "total": 100
  }
}
```

---

## 💬 互动接口

### POST /api/v1/interact/like

点赞内容

**请求体**:
```json
{
  "post_id": "1234567890"
}
```

---

### POST /api/v1/interact/collect

收藏内容

**请求体**:
```json
{
  "post_id": "1234567890"
}
```

---

### POST /api/v1/interact/comment

评论内容

**请求体**:
```json
{
  "post_id": "1234567890",
  "content": "评论内容"
}
```

---

## 📊 数据接口

### GET /api/v1/stats/account

获取账号统计

**响应**:
```json
{
  "success": true,
  "data": {
    "followers": 1000,
    "following": 100,
    "posts": 50,
    "likes": 5000
  }
}
```

---

### GET /api/v1/stats/post/:post_id

获取单篇内容统计

**响应**:
```json
{
  "success": true,
  "data": {
    "views": 1000,
    "likes": 100,
    "comments": 20,
    "collects": 30,
    "shares": 10
  }
}
```

---

### GET /api/v1/stats/recent

获取近期统计

**参数**:
- `start_date`: 开始日期（YYYY-MM-DD）
- `end_date`: 结束日期（YYYY-MM-DD）

---

### GET /api/v1/stats/overview

获取数据概览

**响应**:
```json
{
  "success": true,
  "data": {
    "total_posts": 50,
    "total_views": 10000,
    "total_likes": 500,
    "avg_engagement": 0.05
  }
}
```

---

## ⚠️ 错误响应

```json
{
  "success": false,
  "error": {
    "code": "AUTH_FAILED",
    "message": "Cookie已过期，请重新登录"
  }
}
```

**常见错误码**:
- `AUTH_FAILED`: 认证失败
- `RATE_LIMIT`: 频率限制
- `CONTENT_REVIEW`: 内容审核失败
- `NETWORK_ERROR`: 网络错误

---

## 📚 相关文档

- [使用指南](./USAGE_GUIDE.md)
- [故障排查](./TROUBLESHOOTING.md)
- [架构设计](../docs/ARCHITECTURE.md)

---

**Created by 🦞 Lobster Journey Studio**
