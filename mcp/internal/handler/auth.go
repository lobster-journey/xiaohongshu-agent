package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAuthStatus 获取登录状态
func GetAuthStatus(c *gin.Context) {
	// TODO: 检查Cookie是否有效

	c.JSON(http.StatusOK, gin.H{
		"logged_in": true,
		"username":  "ai-report",
		"expires_at": "2026-05-17T00:00:00Z",
	})
}

// Login 登录小红书
func Login(c *gin.Context) {
	// TODO: 实现登录逻辑
	// 1. 启动浏览器
	// 2. 用户扫码登录
	// 3. 获取Cookie
	// 4. 保存Cookie

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "请在浏览器中完成登录",
		"login_url": "http://localhost:18060/login",
	})
}

// Logout 登出
func Logout(c *gin.Context) {
	// TODO: 清除Cookie

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "已登出",
	})
}
