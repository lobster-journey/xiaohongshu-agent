package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PublishImage 发布图文笔记
func PublishImage(c *gin.Context) {
	var req struct {
		Title   string   `json:"title" binding:"required,max=20"`
		Content string   `json:"content" binding:"required,max=1000"`
		Images  []string `json:"images" binding:"required,min=1,max=9"`
		Tags    []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 实现发布逻辑
	// 1. 验证图片URL
	// 2. 调用小红书API发布
	// 3. 返回结果

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "图文笔记发布成功",
		"data": gin.H{
			"post_id": "xxx",
			"url":     "https://www.xiaohongshu.com/explore/xxx",
		},
	})
}

// PublishVideo 发布视频笔记
func PublishVideo(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required,max=20"`
		Description string `json:"description" binding:"max=1000"`
		VideoURL    string `json:"video_url" binding:"required"`
		CoverURL    string `json:"cover_url"`
		Tags        []string `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 实现视频发布逻辑

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "视频笔记发布成功",
		"data": gin.H{
			"post_id": "xxx",
			"url":     "https://www.xiaohongshu.com/explore/xxx",
		},
	})
}

// PublishBatch 批量发布
func PublishBatch(c *gin.Context) {
	var req struct {
		Posts []interface{} `json:"posts" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 实现批量发布逻辑
	// 1. 添加到任务队列
	// 2. 返回任务ID

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "批量发布任务已创建",
		"data": gin.H{
			"task_id": "xxx",
			"count":   len(req.Posts),
		},
	})
}
