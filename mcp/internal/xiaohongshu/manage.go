package xiaohongshu

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Cody-Chan/xiaohongshu-agent/internal/browser"
	"github.com/go-rod/rod"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// 用户主页URL
	userProfileURL = "https://www.xiaohongshu.com/user/profile/"
	// 创作者中心URL
	creatorCenterURL = "https://creator.xiaohongshu.com/"
)

// NoteInfo 笔记信息
type NoteInfo struct {
	ID          string    `json:"id"`           // 笔记ID
	Title       string    `json:"title"`        // 标题
	Type        string    `json:"type"`         // 类型：normal/video
	Status      string    `json:"status"`       // 状态：published/draft/hidden
	PublishTime time.Time `json:"publish_time"` // 发布时间
	Likes       int       `json:"likes"`        // 点赞数
	Comments    int       `json:"comments"`     // 评论数
	Collects    int       `json:"collects"`     // 收藏数
	URL         string    `json:"url"`          // 笔记链接
}

// NoteManager 笔记管理器
type NoteManager struct {
	browser *browser.Browser
	page    *rod.Page
}

// NewNoteManager 创建笔记管理器实例
func NewNoteManager(b *browser.Browser) *NoteManager {
	return &NoteManager{
		browser: b,
		page:    b.Page(),
	}
}

// GetUserNotes 获取用户笔记列表
func (nm *NoteManager) GetUserNotes(userID string, limit int) ([]NoteInfo, error) {
	ctx := context.Background()

	// 访问用户主页
	url := userProfileURL + userID
	if err := nm.page.Navigate(url); err != nil {
		return nil, errors.Wrap(err, "访问用户主页失败")
	}

	// 等待页面加载
	if err := nm.page.WaitLoad(); err != nil {
		return nil, errors.Wrap(err, "等待页面加载失败")
	}

	time.Sleep(2 * time.Second) // 额外等待

	// 滚动加载更多笔记
	var notes []NoteInfo

	// 获取笔记列表元素
	elements, err := nm.page.Elements("section.note-item")
	if err != nil {
		return nil, errors.Wrap(err, "获取笔记元素失败")
	}

	for i, elem := range elements {
		if i >= limit {
			break
		}

		note, err := nm.parseNoteElement(elem)
		if err != nil {
			logrus.Warnf("解析笔记元素失败: %v", err)
			continue
		}

		notes = append(notes, *note)
	}

	logrus.Infof("获取到 %d 篇笔记", len(notes))
	return notes, nil
}

// parseNoteElement 解析笔记元素
func (nm *NoteManager) parseNoteElement(elem *rod.Element) (*NoteInfo, error) {
	note := &NoteInfo{}

	// 获取笔记链接
	link, err := elem.Element("a")
	if err == nil {
		href, err := link.Property("href")
		if err == nil {
			note.URL = href.String()
			// 从URL中提取笔记ID
			if len(note.URL) > 0 {
				// https://www.xiaohongshu.com/explore/xxxxx
				parts := strings.Split(note.URL, "/")
				if len(parts) > 0 {
					note.ID = parts[len(parts)-1]
				}
			}
		}
	}

	// 获取标题
	titleElem, err := elem.Element(".title")
	if err == nil {
		title, err := titleElem.Text()
		if err == nil {
			note.Title = title
		}
	}

	// 获取点赞数等数据
	// 这里需要根据实际页面结构解析

	return note, nil
}

// HideNote 隐藏笔记（设置为仅自己可见）
func (nm *NoteManager) HideNote(noteID string) error {
	ctx := context.Background()

	logrus.Infof("开始隐藏笔记: %s", noteID)

	// 访问笔记详情页
	noteURL := fmt.Sprintf("https://www.xiaohongshu.com/explore/%s", noteID)
	if err := nm.page.Navigate(noteURL); err != nil {
		return errors.Wrap(err, "访问笔记页面失败")
	}

	if err := nm.page.WaitLoad(); err != nil {
		return errors.Wrap(err, "等待页面加载失败")
	}

	time.Sleep(2 * time.Second)

	// 点击右上角"..."按钮
	moreBtn, err := nm.page.Element(".more-btn")
	if err != nil {
		return errors.Wrap(err, "找不到更多按钮")
	}

	if err := moreBtn.Click(); err != nil {
		return errors.Wrap(err, "点击更多按钮失败")
	}

	time.Sleep(1 * time.Second)

	// 点击"设置为仅自己可见"选项
	hideOption, err := nm.page.Element("text=设置为仅自己可见")
	if err != nil {
		return errors.Wrap(err, "找不到隐藏选项")
	}

	if err := hideOption.Click(); err != nil {
		return errors.Wrap(err, "点击隐藏选项失败")
	}

	time.Sleep(1 * time.Second)

	// 确认操作
	confirmBtn, err := nm.page.Element(".confirm-btn")
	if err == nil {
		if err := confirmBtn.Click(); err != nil {
			logrus.Warnf("点击确认按钮失败: %v", err)
		}
	}

	logrus.Infof("笔记 %s 已隐藏", noteID)
	return nil
}

// DeleteNote 删除笔记
func (nm *NoteManager) DeleteNote(noteID string) error {
	ctx := context.Background()

	logrus.Infof("开始删除笔记: %s", noteID)

	// 访问笔记详情页
	noteURL := fmt.Sprintf("https://www.xiaohongshu.com/explore/%s", noteID)
	if err := nm.page.Navigate(noteURL); err != nil {
		return errors.Wrap(err, "访问笔记页面失败")
	}

	if err := nm.page.WaitLoad(); err != nil {
		return errors.Wrap(err, "等待页面加载失败")
	}

	time.Sleep(2 * time.Second)

	// 点击右上角"..."按钮
	moreBtn, err := nm.page.Element(".more-btn")
	if err != nil {
		return errors.Wrap(err, "找不到更多按钮")
	}

	if err := moreBtn.Click(); err != nil {
		return errors.Wrap(err, "点击更多按钮失败")
	}

	time.Sleep(1 * time.Second)

	// 点击"删除"选项
	deleteOption, err := nm.page.Element("text=删除")
	if err != nil {
		return errors.Wrap(err, "找不到删除选项")
	}

	if err := deleteOption.Click(); err != nil {
		return errors.Wrap(err, "点击删除选项失败")
	}

	time.Sleep(1 * time.Second)

	// 确认删除
	confirmBtn, err := nm.page.Element(".confirm-btn")
	if err == nil {
		if err := confirmBtn.Click(); err != nil {
			logrus.Warnf("点击确认按钮失败: %v", err)
		}
	}

	logrus.Infof("笔记 %s 已删除", noteID)
	return nil
}

// BatchHideNotes 批量隐藏笔记
func (nm *NoteManager) BatchHideNotes(noteIDs []string) (int, []error) {
	var errors []error
	successCount := 0

	for i, noteID := range noteIDs {
		logrus.Infof("处理第 %d/%d 篇笔记: %s", i+1, len(noteIDs), noteID)

		if err := nm.HideNote(noteID); err != nil {
			errors = append(errors, fmt.Errorf("隐藏笔记 %s 失败: %v", noteID, err))
		} else {
			successCount++
		}

		// 避免操作过快
		if i < len(noteIDs)-1 {
			time.Sleep(3 * time.Second)
		}
	}

	logrus.Infof("批量隐藏完成: 成功 %d 篇，失败 %d 篇", successCount, len(errors))
	return successCount, errors
}
