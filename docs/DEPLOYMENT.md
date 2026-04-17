# 部署指南

本文档介绍如何部署小红书 Agent 到生产环境。

## 📋 前置要求

- Docker & Docker Compose
- 或直接安装：Go 1.24+, Python 3.10+, Node.js 18+

## 🐳 Docker 部署（推荐）

### 1. 克隆项目

```bash
git clone GitHub 搜索：xiaohongshu-agent.git
cd xiaohongshu-agent
```

### 2. 配置

```bash
# 复制配置文件
cp configs/config.yaml docker/config.yaml

# 编辑配置
vim docker/config.yaml
```

### 3. 启动服务

```bash
cd docker
docker-compose up -d
```

### 4. 查看日志

```bash
docker-compose logs -f xiaohongshu-agent
```

### 5. 登录

```bash
docker-compose exec xiaohongshu-agent xiaohongshu-login
```

## 🔧 直接部署

### 1. 安装

```bash
curl -fsSL GitHub 搜索：xiaohongshu-agent/main/install.sh | bash
```

### 2. 配置

编辑 `~/.openclaw/mcp/config.yaml`

### 3. 启动服务

```bash
cd ~/.openclaw/mcp
./xiaohongshu-mcp
```

### 4. 登录

```bash
xiaohongshu-login
```

## 🌐 反向代理

### Nginx

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://127.0.0.1:18060;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

## ⚙️ Systemd 服务

创建 `/etc/systemd/system/xiaohongshu-agent.service`:

```ini
[Unit]
Description=Xiaohongshu Agent Service
After=network.target

[Service]
Type=simple
User=your-user
WorkingDirectory=/home/your-user/.openclaw/mcp
ExecStart=/home/your-user/.openclaw/mcp/xiaohongshu-mcp
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

启用服务：

```bash
sudo systemctl daemon-reload
sudo systemctl enable xiaohongshu-agent
sudo systemctl start xiaohongshu-agent
```

## 📊 监控

### Prometheus

服务默认在 `9090` 端口暴露指标。

Prometheus 配置：

```yaml
scrape_configs:
  - job_name: 'xiaohongshu-agent'
    static_configs:
      - targets: ['localhost:9090']
```

### Grafana

导入仪表盘：`docs/grafana-dashboard.json`

## 🔒 安全建议

1. **修改默认端口**
2. **启用API密钥认证**
3. **配置HTTPS**
4. **限制访问IP**
5. **定期更新依赖**

## 🔄 更新

```bash
# Docker
cd xiaohongshu-agent
git pull
docker-compose -f docker/docker-compose.yaml build
docker-compose -f docker/docker-compose.yaml up -d

# 直接安装
curl -fsSL GitHub 搜索：xiaohongshu-agent/main/install.sh | bash
```

## 🐛 故障排查

### 服务无法启动

```bash
# 检查端口
netstat -tlnp | grep 18060

# 查看日志
tail -f ~/.openclaw/mcp/logs/app.log
```

### Cookie过期

```bash
xiaohongshu-login
```

### 性能问题

```bash
# 调整配置
vim ~/.openclaw/mcp/config.yaml

# 增加并发
publish:
  concurrent_limit: 10
```
