#!/bin/bash

# 测试小红书Agent API

BASE_URL="http://localhost:18060"

echo "🧪 小红书Agent API测试"
echo "======================"
echo ""

# 测试健康检查
echo "1. 测试健康检查..."
curl -s "$BASE_URL/health" | jq .
echo ""

# 测试登录状态
echo "2. 测试登录状态..."
curl -s "$BASE_URL/api/v1/auth/status" | jq .
echo ""

# 测试发布图文
echo "3. 测试发布图文..."
curl -s -X POST "$BASE_URL/api/v1/publish/image" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "测试标题",
    "content": "这是测试内容",
    "images": ["https://example.com/image1.jpg"],
    "tags": ["测试", "小红书"]
  }' | jq .
echo ""

# 测试搜索
echo "4. 测试搜索笔记..."
curl -s "$BASE_URL/api/v1/search/posts?keyword=AI" | jq .
echo ""

# 测试统计
echo "5. 测试数据统计..."
curl -s "$BASE_URL/api/v1/stats/overview" | jq .
echo ""

echo "✅ 测试完成"
