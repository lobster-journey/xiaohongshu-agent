# 开发指南

## 🛠️ 开发环境搭建

### 前置要求

- Go 1.24+
- Python 3.10+
- Node.js 18+ (用于Playwright)
- Make
- Git

### 本地开发

```bash
# 1. 克隆项目
git clone https://github.com/Cody-Chan/xiaohongshu-agent.git
cd xiaohongshu-agent

# 2. 安装依赖
make install-deps

# 3. 安装Playwright浏览器
make install-browser

# 4. 启动开发服务器
make dev
```

### 项目结构说明

```
xiaohongshu-agent/
├── mcp/                    # Go MCP服务
│   ├── cmd/               # 命令入口
│   ├── internal/          # 内部模块（不对外暴露）
│   └── pkg/               # 公共包（可被外部引用）
│
├── skill/                 # Python/Bash Skill
│   ├── scripts/          # 脚本工具
│   └── tests/            # 测试用例
│
└── docs/                  # 文档
```

## 📝 代码规范

### Go 代码规范

```go
// 包注释
package service

// Service 服务接口
type Service interface {
    // Publish 发布内容
    Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error)
}

// PublishRequest 发布请求
type PublishRequest struct {
    Title   string   `json:"title"`
    Content string   `json:"content"`
    Images  []string `json:"images"`
}

// 错误处理
if err != nil {
    return nil, fmt.Errorf("failed to publish: %w", err)
}
```

### Python 代码规范

```python
"""模块文档字符串"""

from typing import List, Optional

class Publisher:
    """发布器类"""
    
    def publish_image(
        self,
        title: str,
        content: str,
        images: List[str]
    ) -> dict:
        """
        发布图文笔记
        
        Args:
            title: 标题
            content: 内容
            images: 图片列表
            
        Returns:
            发布结果
            
        Raises:
            PublishError: 发布失败
        """
        pass
```

## 🧪 测试

### 单元测试

```bash
# Go 测试
make test-go

# Python 测试
make test-python

# 全部测试
make test
```

### 集成测试

```bash
# 需要先登录
xiaohongshu-login

# 运行集成测试
make test-integration
```

### 测试覆盖率

```bash
make coverage
```

## 🔨 构建与发布

### 构建

```bash
# 构建所有平台
make build-all

# 仅构建当前平台
make build

# 构建Docker镜像
make docker
```

### 发布

```bash
# 发布新版本
make release VERSION=v1.0.0
```

## 🐛 调试

### 启用调试日志

```yaml
# config.yaml
log:
  level: "debug"
```

### 查看日志

```bash
# 实时日志
tail -f logs/app.log

# 过滤错误
grep "ERROR" logs/app.log
```

### 浏览器调试

```yaml
# config.yaml
browser:
  headless: false  # 显示浏览器
```

## 📦 依赖管理

### Go 依赖

```bash
# 添加依赖
go get github.com/gin-gonic/gin

# 更新依赖
go mod tidy

# 查看依赖
go mod graph
```

### Python 依赖

```bash
# 安装依赖
pip install -r requirements.txt

# 添加依赖
pip install requests
pip freeze > requirements.txt
```

## 🔧 常见问题

### Q: 浏览器启动失败？

```bash
# 安装浏览器依赖
make install-browser-deps
```

### Q: Cookie 过期？

```bash
# 重新登录
xiaohongshu-login
```

### Q: 端口被占用？

```bash
# 修改配置
# config.yaml
server:
  port: 18061
```

## 📚 相关文档

- [API文档](./API_REFERENCE.md)
- [架构设计](./ARCHITECTURE.md)
- [部署指南](./DEPLOYMENT.md)
