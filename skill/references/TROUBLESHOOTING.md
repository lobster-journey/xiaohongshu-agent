# 🔧 故障排查

> 常见问题与解决方案

---

## 🚨 常见错误

### 1. Cookie已过期

**错误信息**：
```
AUTH_FAILED: Cookie已过期，请重新登录
```

**解决方案**：
```bash
# 重新登录
xiaohongshu-login

# 或
make login
```

---

### 2. 发布频率限制

**错误信息**：
```
RATE_LIMIT: 发布频率过高，请稍后再试
```

**解决方案**：
- 等待30分钟后再发布
- 降低发布频率
- 使用定时发布分散时间

---

### 3. 内容审核失败

**错误信息**：
```
CONTENT_REVIEW: 内容审核未通过
```

**排查步骤**：
1. 检查标题是否包含敏感词
2. 检查内容是否合规
3. 检查图片是否有水印
4. 检查是否有违规链接

**解决方案**：
- 修改内容后重试
- 使用内容审核工具预检

---

### 4. 网络连接失败

**错误信息**：
```
NETWORK_ERROR: 无法连接到小红书服务器
```

**排查步骤**：
1. 检查网络连接
2. 检查防火墙设置
3. 检查代理配置

**解决方案**：
```bash
# 测试网络连接
ping xiaohongshu.com

# 检查代理设置
echo $HTTP_PROXY
echo $HTTPS_PROXY

# 临时关闭代理
unset HTTP_PROXY HTTPS_PROXY
```

---

### 5. 图片上传失败

**错误信息**：
```
UPLOAD_FAILED: 图片上传失败
```

**排查步骤**：
1. 检查图片格式（JPG/PNG）
2. 检查图片大小（<5MB）
3. 检查图片路径是否正确

**解决方案**：
```bash
# 检查图片
file /path/to/image.jpg
ls -lh /path/to/image.jpg

# 压缩图片
convert original.jpg -quality 85 compressed.jpg
```

---

### 6. 视频上传超时

**错误信息**：
```
UPLOAD_TIMEOUT: 视频上传超时
```

**解决方案**：
- 检查视频大小（建议<500MB）
- 检查网络速度
- 增加超时时间

---

## 🔍 日志查看

### 查看MCP日志

```bash
# 实时查看
make logs

# 或
tail -f ~/.xiaohongshu-agent/logs/mcp.log

# 查看最近100行
tail -100 ~/.xiaohongshu-agent/logs/mcp.log
```

---

### 查看发布日志

```bash
# 查看发布历史
cat ~/.xiaohongshu-agent/logs/publish.log

# 查看错误日志
grep ERROR ~/.xiaohongshu-agent/logs/*.log
```

---

## 🛠️ 调试模式

### 启用详细日志

```bash
# 设置日志级别
export LOG_LEVEL=debug

# 重启服务
make run
```

---

### 测试API连接

```bash
# 健康检查
curl http://localhost:18060/health

# 测试搜索
curl "http://localhost:18060/api/v1/search?keyword=test&limit=1"
```

---

## 🔧 服务管理

### 重启服务

```bash
# 停止服务
pkill xiaohongshu-mcp

# 启动服务
make run
```

---

### 清理缓存

```bash
# 清理Cookie
rm -rf ~/.xiaohongshu-agent/cookies/*

# 清理日志
rm -rf ~/.xiaohongshu-agent/logs/*

# 清理缓存
rm -rf ~/.xiaohongshu-agent/cache/*
```

---

## 📊 性能优化

### 1. 发布优化

- 批量发布使用队列
- 图片预压缩
- 使用CDN加速

### 2. 数据采集优化

- 增加缓存层
- 使用连接池
- 异步处理

---

## 🆘 获取帮助

### 查看文档

- [API参考](./API_REFERENCE.md)
- [使用指南](./USAGE_GUIDE.md)
- [开发指南](../../docs/DEVELOPMENT.md)

### 提交Issue

如果问题仍未解决，请提交Issue：
https://github.com/lobster-journey/xiaohongshu-agent/issues

**Issue模板**：
```
## 问题描述
[描述遇到的问题]

## 复现步骤
1. ...
2. ...
3. ...

## 期望结果
[期望的正常结果]

## 实际结果
[实际的错误结果]

## 环境信息
- 操作系统：
- Python版本：
- Go版本：
- xiaohongshu-agent版本：

## 日志信息
```
[相关日志]
```
```

---

**Created by 🦞 Lobster Journey Studio**
