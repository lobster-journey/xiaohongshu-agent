package handler

import (
	"net/http"
	"strconv"

	"github.com/Cody-Chan/xiaohongshu-agent/internal/xiaohongshu"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetUserNotes 获取用户笔记列表
func GetUserNotes(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "user_id参数必需",
		})
		return
	}

	limit := 20 // 默认20篇
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	// 获取浏览器实例
	browser := GetBrowser()
	if browser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "浏览器实例未初始化",
		})
		return
	}

	// 创建笔记管理器
	manager := xiaohongshu.NewNoteManager(browser)

	// 获取笔记列表
	notes, err := manager.GetUserNotes(userID, limit)
	if err != nil {
		logrus.Errorf("获取笔记列表失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"notes": notes,
			"total": len(notes),
		},
	})
}

// HideNote 隐藏单篇笔记
func HideNote(c *gin.Context) {
	var req struct {
		NoteID string `json:"note_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 获取浏览器实例
	browser := GetBrowser()
	if browser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "浏览器实例未初始化",
		})
		return
	}

	// 创建笔记管理器
	manager := xiaohongshu.NewNoteManager(browser)

	// 隐藏笔记
	if err := manager.HideNote(req.NoteID); err != nil {
		logrus.Errorf("隐藏笔记失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "笔记已隐藏",
		"data": gin.H{
			"note_id": req.NoteID,
		},
	})
}

// BatchHideNotes 批量隐藏笔记
func BatchHideNotes(c *gin.Context) {
	var req struct {
		NoteIDs []string `json:"note_ids" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 获取浏览器实例
	browser := GetBrowser()
	if browser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "浏览器实例未初始化",
		})
		return
	}

	// 创建笔记管理器
	manager := xiaohongshu.NewNoteManager(browser)

	// 批量隐藏
	successCount, errors := manager.BatchHideNotes(req.NoteIDs)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "批量隐藏完成",
		"data": gin.H{
			"total":     len(req.NoteIDs),
			"success":   successCount,
			"failed":    len(errors),
			"errors":    errors,
		},
	})
}

// DeleteNote 删除单篇笔记
func DeleteNote(c *gin.Context) {
	var req struct {
		NoteID string `json:"note_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 获取浏览器实例
	browser := GetBrowser()
	if browser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "浏览器实例未初始化",
		})
		return
	}

	// 创建笔记管理器
	manager := xiaohongshu.NewNoteManager(browser)

	// 删除笔记
	if err := manager.DeleteNote(req.NoteID); err != nil {
		logrus.Errorf("删除笔记失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "笔记已删除",
		"data": gin.H{
			"note_id": req.NoteID,
		},
	})
}
