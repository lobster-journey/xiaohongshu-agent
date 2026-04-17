package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Cody-Chan/xiaohongshu-agent/internal/service"
	"github.com/Cody-Chan/xiaohongshu-agent/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// App 应用结构体
type App struct {
	xhsService *service.XiaohongshuService
	router     *gin.Engine
	httpServer *http.Server
}

// NewApp 创建新应用
func NewApp() *App {
	app := &App{
		xhsService: service.NewXiaohongshuService(),
	}

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)
	app.router = gin.New()

	return app
}

// Start 启动服务
func (a *App) Start(port int) error {
	// 设置路由
	a.setupRoutes()

	// 创建HTTP服务器
	a.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: a.router,
	}

	// 启动HTTP服务（goroutine）
	go func() {
		logrus.Infof("🚀 小红书Agent服务启动: http://0.0.0.0:%d", port)
		logrus.Infof("📚 API文档: http://0.0.0.0:%d/api/v1", port)
		logrus.Infof("💚 健康检查: http://0.0.0.0:%d/health", port)

		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("正在关闭服务...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		logrus.Warnf("服务关闭错误: %v", err)
		return err
	}

	logrus.Info("服务已优雅关闭")
	return nil
}

// setupRoutes 设置路由
func (a *App) setupRoutes() {
	// 中间件
	a.router.Use(gin.Logger())
	a.router.Use(gin.Recovery())

	// 健康检查
	a.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "xiaohongshu-agent",
			"version": "1.0.0",
		})
	})

	// WebSocket
	a.router.GET("/ws", server.HandleWebSocket)

	// API v1
	v1 := a.router.Group("/api/v1")
	{
		// 认证
		auth := v1.Group("/auth")
		{
			auth.GET("/status", a.xhsService.GetAuthStatus)
			auth.POST("/login", a.xhsService.Login)
			auth.POST("/logout", a.xhsService.Logout)
		}

		// 发布
		publish := v1.Group("/publish")
		{
			publish.POST("/image", a.xhsService.PublishImage)
			publish.POST("/video", a.xhsService.PublishVideo)
			publish.POST("/batch", a.xhsService.PublishBatch)
			publish.GET("/task/:id", a.xhsService.GetTaskStatus)
		}

		// 搜索
		search := v1.Group("/search")
		{
			search.GET("/posts", a.xhsService.SearchPosts)
			search.GET("/users", a.xhsService.SearchUsers)
		}

		// 互动
		interaction := v1.Group("/interaction")
		{
			interaction.POST("/comment", a.xhsService.AddComment)
			interaction.POST("/like/:id", a.xhsService.LikePost)
			interaction.POST("/follow/:id", a.xhsService.FollowUser)
		}

		// 统计
		stats := v1.Group("/stats")
		{
			stats.GET("/overview", a.xhsService.GetOverview)
			stats.GET("/post/:id", a.xhsService.GetPostStats)
		}
	}
}

func main() {
	app := NewApp()

	port := 18060
	if p := os.Getenv("PORT"); p != "" {
		fmt.Sscanf(p, "%d", &port)
	}

	if err := app.Start(port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
