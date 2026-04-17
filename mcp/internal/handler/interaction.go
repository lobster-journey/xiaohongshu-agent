package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddComment 添加评论
func AddComment(c *gin.Context) {
	var req struct {
		PostID  string `json:"post_id" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 实现评论逻辑

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "评论成功",
	})
}

// LikePost 点赞笔记
func LikePost(c *gin.Context) {
	postID := c.Param("id")

	// TODO: 实现点赞逻辑

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "点赞成功",
		"post_id": postID,
	})
}

// FollowUser 关注用户
func FollowUser(c *gin.Context) {
	userID := c.Param("id")

	// TODO: 实现关注逻辑

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "关注成功",
		"user_id": userID,
	})
}
