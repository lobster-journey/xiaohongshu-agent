# API 参考文档

## 📡 API 基础信息

- **Base URL**: `http://localhost:18060`
- **版本**: v1
- **认证**: Cookie-based

## 🔐 认证

### 登录状态检查

```http
GET /api/v1/auth/status
```

**响应示例**:
```json
{
  "logged_in": true,
  "user_id": "5f8d9a...",
  "username": "用户名",
  "expires_at": "2026-05-17T22:00:00Z"
}
```

---

## 📝 发布接口

### 发布图文笔记

```http
POST /api/v1/publish/image
Content-Type: application/json

{
  "title": "标题（最多20字）",
  "content": "正文内容（最多1000字）",
  "images": [
    "/path/to/image1.jpg",
    "/path/to/image2.jpg"
  ],
  "tags": ["标签1", "标签2"]
}
```

**响应示例**:
```json
{
  "success": true,
  "data": {
    "post_id": "笔记ID",
    "url": "https://www.xiaohongshu.com/explore/...",
    "published_at": "2026-04-17T22:00:00Z"
  }
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| title | string | 是 | 标题，最多20字 |
| content | string | 是 | 正文，最多1000字 |
| images | []string | 是 | 图片路径列表，1-9张 |
| tags | []string | 否 | 标签列表 |

---

### 发布视频笔记

```http
POST /api/v1/publish/video
Content-Type: application/json

{
  "title": "标题（最多20字）",
  "content": "正文内容（最多1000字）",
  "video": "/path/to/video.mp4",
  "tags": ["标签1", "标签2"]
}
```

**响应示例**:
```json
{
  "success": true,
  "data": {
    "post_id": "笔记ID",
    "url": "https://www.xiaohongshu.com/explore/...",
    "published_at": "2026-04-17T22:00:00Z",
    "video_duration": 120
  }
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| title | string | 是 | 标题，最多20字 |
| content | string | 是 | 正文，最多1000字 |
| video | string | 是 | 视频文件路径 |
| tags | []string | 否 | 标签列表 |

---

### 批量发布

```http
POST /api/v1/publish/batch
Content-Type: application/json

{
  "posts": [
    {
      "title": "标题1",
      "content": "内容1",
      "images": ["/path/to/image1.jpg"]
    },
    {
      "title": "标题2",
      "content": "内容2",
      "images": ["/path/to/image2.jpg"]
    }
  ],
  "interval": 60
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| posts | []Post | 是 | 发布内容列表 |
| interval | int | 否 | 发布间隔（秒），默认60 |

---

## 💬 互动接口

### 发布评论

```http
POST /api/v1/interact/comment
Content-Type: application/json

{
  "post_id": "笔记ID",
  "xsec_token": "安全Token",
  "content": "评论内容"
}
```

### 点赞笔记

```http
POST /api/v1/interact/like
Content-Type: application/json

{
  "post_id": "笔记ID",
  "xsec_token": "安全Token"
}
```

### 收藏笔记

```http
POST /api/v1/interact/collect
Content-Type: application/json

{
  "post_id": "笔记ID",
  "xsec_token": "安全Token"
}
```

---

## 🔍 搜索接口

### 搜索内容

```http
POST /api/v1/search
Content-Type: application/json

{
  "keyword": "搜索关键词",
  "limit": 20,
  "sort": "relevance"
}
```

**响应示例**:
```json
{
  "success": true,
  "data": {
    "total": 100,
    "items": [
      {
        "post_id": "笔记ID",
        "title": "标题",
        "content": "内容摘要",
        "author": {
          "user_id": "用户ID",
          "username": "用户名",
          "avatar": "头像URL"
        },
        "images": ["图片URL"],
        "stats": {
          "likes": 100,
          "comments": 20,
          "collects": 50
        },
        "xsec_token": "安全Token"
      }
    ]
  }
}
```

---

## 📊 数据接口

### 获取推荐列表

```http
GET /api/v1/feed/recommend?limit=20
```

### 获取笔记详情

```http
GET /api/v1/post/{post_id}?xsec_token={token}
```

### 获取用户统计

```http
GET /api/v1/stats/user
```

**响应示例**:
```json
{
  "success": true,
  "data": {
    "posts_count": 10,
    "likes_received": 1000,
    "followers": 500,
    "following": 200
  }
}
```

---

## 📋 任务接口

### 查询任务状态

```http
GET /api/v1/task/{task_id}
```

**响应示例**:
```json
{
  "success": true,
  "data": {
    "task_id": "任务ID",
    "type": "publish_image",
    "status": "completed",
    "progress": 100,
    "result": {
      "post_id": "笔记ID",
      "url": "笔记链接"
    },
    "created_at": "2026-04-17T22:00:00Z",
    "completed_at": "2026-04-17T22:00:30Z"
  }
}
```

### 取消任务

```http
DELETE /api/v1/task/{task_id}
```

---

## 🔔 WebSocket 接口

### 连接

```javascript
ws://localhost:18060/ws
```

### 消息格式

```json
{
  "type": "task_update",
  "data": {
    "task_id": "任务ID",
    "status": "processing",
    "progress": 50,
    "message": "正在上传图片..."
  }
}
```

---

## ❌ 错误响应

### 错误格式

```json
{
  "success": false,
  "error": {
    "code": "PUBLISH_FAILED",
    "message": "发布失败：标题过长",
    "details": "标题长度超过20字限制"
  }
}
```

### 错误码

| 错误码 | 说明 |
|--------|------|
| NOT_LOGGED_IN | 未登录 |
| COOKIE_EXPIRED | Cookie已过期 |
| INVALID_PARAMS | 参数错误 |
| PUBLISH_FAILED | 发布失败 |
| TASK_NOT_FOUND | 任务不存在 |
| RATE_LIMITED | 请求过于频繁 |

---

## 🚦 限流

- **默认限制**: 60请求/分钟
- **发布接口**: 10请求/分钟
- **搜索接口**: 30请求/分钟

---

## 📝 请求示例

### cURL

```bash
# 发布图文
curl -X POST http://localhost:18060/api/v1/publish/image \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试标题",
    "content": "测试内容",
    "images": ["/path/to/image.jpg"]
  }'
```

### Python

```python
import requests

response = requests.post(
    "http://localhost:18060/api/v1/publish/image",
    json={
        "title": "测试标题",
        "content": "测试内容",
        "images": ["/path/to/image.jpg"]
    }
)

print(response.json())
```

### JavaScript

```javascript
const response = await fetch('http://localhost:18060/api/v1/publish/image', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    title: '测试标题',
    content: '测试内容',
    images: ['/path/to/image.jpg']
  })
});

const data = await response.json();
console.log(data);
```
