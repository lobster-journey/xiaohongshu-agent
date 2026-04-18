package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/Cody-Chan/xiaohongshu-agent/internal/handler"
	"github.com/Cody-Chan/xiaohongshu-agent/internal/middleware"
)

func Start(port int) error {
	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)

	// 创建路由
	r := gin.New()

	// 中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "xiaohongshu-mcp",
		})
	})

	// WebSocket
	r.GET("/ws", handler.HandleWebSocket)

	// API v1
	v1 := r.Group("/api/v1")
	{
		// 登录相关
		auth := v1.Group("/auth")
		{
			auth.GET("/status", handler.GetAuthStatus)
			auth.POST("/login", handler.Login)
			auth.POST("/logout", handler.Logout)
		}

		// 发布相关
		publish := v1.Group("/publish")
		{
			publish.POST("/image", handler.PublishImage)
			publish.POST("/video", handler.PublishVideo)
			publish.POST("/batch", handler.PublishBatch)
		}

		// 搜索相关
		search := v1.Group("/search")
		{
			search.GET("/posts", handler.SearchPosts)
			search.GET("/users", handler.SearchUsers)
		}

		// 互动相关
		interaction := v1.Group("/interaction")
		{
			interaction.POST("/comment", handler.AddComment)
			interaction.POST("/like", handler.LikePost)
			interaction.POST("/follow", handler.FollowUser)
		}

		// 数据统计
		stats := v1.Group("/stats")
		{
			stats.GET("/overview", handler.GetOverview)
			stats.GET("/posts/:id", handler.GetPostStats)
		}

		// 笔记管理
		notes := v1.Group("/notes")
		{
			notes.GET("", handler.GetUserNotes)
			notes.POST("/hide", handler.HideNote)
			notes.POST("/batch-hide", handler.BatchHideNotes)
			notes.POST("/delete", handler.DeleteNote)
		}
	}

	// 启动服务
	addr := fmt.Sprintf(":%d", port)
	log.Printf("🚀 服务启动在 http://0.0.0.0%s", addr)
	log.Printf("📚 API文档: http://0.0.0.0%s/api/v1", addr)
	log.Printf("💚 健康检查: http://0.0.0.0%s/health", addr)

	return r.Run(addr)
}
