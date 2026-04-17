# 贡献指南

感谢你考虑为小红书 Agent 项目做出贡献！🎉

## 🤝 如何贡献

### 报告问题

如果你发现了bug或有功能建议：

1. 检查 [Issues](https://github.com/Cody-Chan/xiaohongshu-agent/issues) 中是否已有相同问题
2. 如果没有，创建新的 Issue，包含：
   - 问题描述
   - 复现步骤
   - 期望行为
   - 实际行为
   - 环境信息（OS、Go版本等）

### 提交代码

1. **Fork 项目**
   ```bash
   git clone https://github.com/your-username/xiaohongshu-agent.git
   cd xiaohongshu-agent
   ```

2. **创建分支**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **编写代码**
   - 遵循代码规范
   - 添加必要的测试
   - 更新相关文档

4. **提交变更**
   ```bash
   git add .
   git commit -m "feat: add your feature"
   ```

   提交信息格式：
   - `feat:` 新功能
   - `fix:` 修复bug
   - `docs:` 文档更新
   - `style:` 代码格式
   - `refactor:` 重构
   - `test:` 测试
   - `chore:` 构建/工具

5. **推送分支**
   ```bash
   git push origin feature/your-feature-name
   ```

6. **创建 Pull Request**
   - 描述你的变更
   - 关联相关 Issue
   - 等待代码审查

## 📝 代码规范

### Go 代码

- 使用 `gofmt` 格式化代码
- 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 添加必要的注释
- 编写单元测试

### Python 代码

- 遵循 [PEP 8](https://peps.python.org/pep-0008/)
- 使用 `black` 格式化
- 添加类型注解
- 编写测试用例

### 文档

- 使用 Markdown 格式
- 保持简洁清晰
- 添加必要的示例

## 🧪 测试

```bash
# 运行所有测试
make test

# 运行特定测试
make test-go
make test-python
```

## 📋 检查清单

提交 PR 前，请确保：

- [ ] 代码通过所有测试
- [ ] 代码符合规范
- [ ] 添加了必要的测试
- [ ] 更新了相关文档
- [ ] 提交信息格式正确

## 🙏 感谢

感谢所有贡献者的付出！

## 📧 联系方式

如有问题，可以通过以下方式联系：
- 创建 Issue
- 发送邮件至：your-email@example.com
