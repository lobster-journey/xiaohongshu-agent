package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchPosts 搜索笔记
func SearchPosts(c *gin.Context) {
	keyword := c.Query("keyword")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "20")

	// TODO: 实现搜索逻辑

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"keyword": keyword,
			"page":    page,
			"page_size": pageSize,
			"total":   0,
			"items":   []interface{}{},
		},
	})
}

// SearchUsers 搜索用户
func SearchUsers(c *gin.Context) {
	keyword := c.Query("keyword")

	// TODO: 实现用户搜索逻辑

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"keyword": keyword,
			"items":   []interface{}{},
		},
	})
}
