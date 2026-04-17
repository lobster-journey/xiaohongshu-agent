package xiaohongshu

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Cody-Chan/xiaohongshu-agent/internal/browser"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// 发布页面URL
	publishURL = "https://creator.xiaohongshu.com/publish/publish?source=official"
	// 标题最大长度
	titleMaxLength = 20
	// 内容最大长度
	contentMaxLength = 1000
)

// PublishContent 发布内容结构
type PublishContent struct {
	Title        string     // 标题
	Content      string     // 内容
	Images       []string   // 图片路径（本地路径或URL）
	Tags         []string   // 标签
	ScheduleTime *time.Time // 定时发布时间（nil表示立即发布）
	IsOriginal   bool       // 是否原创
	Visibility   string     // 可见范围：公开可见、仅自己可见、仅互关好友可见
	Products     []string   // 商品关键词
}

// Publisher 发布器
type Publisher struct {
	browser *browser.Browser
	page    *rod.Page
}

// NewPublisher 创建发布器实例
func NewPublisher(b *browser.Browser) *Publisher {
	return &Publisher{
		browser: b,
		page:    b.Page(),
	}
}

// PublishImage 发布图文
func (p *Publisher) PublishImage(ctx context.Context, content PublishContent) error {
	logrus.Infof("开始发布图文: %s", content.Title)

	// 1. 验证内容
	if err := p.validateContent(content); err != nil {
		return err
	}

	// 2. 导航到发布页面
	if err := p.navigateToPublishPage(); err != nil {
		return err
	}

	// 3. 上传图片
	if err := p.uploadImages(content.Images); err != nil {
		return errors.Wrap(err, "上传图片失败")
	}

	// 4. 填写标题
	if err := p.fillTitle(content.Title); err != nil {
		return errors.Wrap(err, "填写标题失败")
	}

	// 5. 填写内容
	if err := p.fillContent(content.Content); err != nil {
		return errors.Wrap(err, "填写内容失败")
	}

	// 6. 添加标签
	if len(content.Tags) > 0 {
		if err := p.addTags(content.Tags); err != nil {
			logrus.Warnf("添加标签失败: %v", err)
		}
	}

	// 7. 设置定时发布
	if content.ScheduleTime != nil {
		if err := p.setScheduleTime(*content.ScheduleTime); err != nil {
			logrus.Warnf("设置定时发布失败: %v", err)
		}
	}

	// 8. 设置可见范围
	if content.Visibility != "" && content.Visibility != "公开可见" {
		if err := p.setVisibility(content.Visibility); err != nil {
			logrus.Warnf("设置可见范围失败: %v", err)
		}
	}

	// 9. 设置原创声明
	if content.IsOriginal {
		if err := p.setOriginal(); err != nil {
			logrus.Warnf("设置原创声明失败: %v", err)
		}
	}

	// 10. 点击发布
	if err := p.clickPublish(); err != nil {
		return errors.Wrap(err, "点击发布按钮失败")
	}

	logrus.Info("发布完成")
	return nil
}

// validateContent 验证内容
func (p *Publisher) validateContent(content PublishContent) error {
	if content.Title == "" {
		return errors.New("标题不能为空")
	}

	if len([]rune(content.Title)) > titleMaxLength {
		return errors.Errorf("标题长度超过限制（%d字）", titleMaxLength)
	}

	if content.Content == "" {
		return errors.New("内容不能为空")
	}

	if len([]rune(content.Content)) > contentMaxLength {
		return errors.Errorf("内容长度超过限制（%d字）", contentMaxLength)
	}

	if len(content.Images) == 0 {
		return errors.New("至少需要上传1张图片")
	}

	if len(content.Images) > 9 {
		logrus.Warnf("图片数量超过9张，只使用前9张")
		content.Images = content.Images[:9]
	}

	if len(content.Tags) > 10 {
		logrus.Warnf("标签数量超过10个，只使用前10个")
		content.Tags = content.Tags[:10]
	}

	return nil
}

// navigateToPublishPage 导航到发布页面
func (p *Publisher) navigateToPublishPage() error {
	logrus.Info("导航到发布页面...")

	if err := p.browser.Navigate(publishURL); err != nil {
		return err
	}

	// 等待页面加载
	time.Sleep(3 * time.Second)

	// 点击"上传图文"标签
	if err := p.clickTab("上传图文"); err != nil {
		return errors.Wrap(err, "点击上传图文标签失败")
	}

	time.Sleep(1 * time.Second)
	return nil
}

// clickTab 点击标签
func (p *Publisher) clickTab(tabName string) error {
	// 查找包含指定文本的标签
	elems, err := p.page.Elements("div.creator-tab")
	if err != nil {
		return err
	}

	for _, elem := range elems {
		text, err := elem.Text()
		if err != nil {
			continue
		}

		if strings.TrimSpace(text) == tabName {
			if err := elem.Click(proto.InputMouseButtonLeft, 1); err != nil {
				return err
			}
			logrus.Infof("已点击标签: %s", tabName)
			return nil
		}
	}

	return errors.Errorf("未找到标签: %s", tabName)
}

// uploadImages 上传图片
func (p *Publisher) uploadImages(imagePaths []string) error {
	logrus.Infof("上传%d张图片...", len(imagePaths))

	for i, imgPath := range imagePaths {
		logrus.Infof("上传第%d张图片: %s", i+1, imgPath)

		// 检查文件是否存在
		if _, err := os.Stat(imgPath); os.IsNotExist(err) {
			logrus.Warnf("图片文件不存在: %s", imgPath)
			continue
		}

		// 查找上传输入框
		selector := "input[type='file']"
		if i == 0 {
			selector = ".upload-input"
		}

		uploadInput, err := p.page.Element(selector)
		if err != nil {
			return errors.Wrapf(err, "查找上传输入框失败（第%d张）", i+1)
		}

		// 上传文件
		if err := uploadInput.SetFiles([]string{imgPath}); err != nil {
			return errors.Wrapf(err, "上传第%d张图片失败", i+1)
		}

		// 等待上传完成
		if err := p.waitForImageUpload(i + 1); err != nil {
			return errors.Wrapf(err, "第%d张图片上传超时", i+1)
		}

		logrus.Infof("第%d张图片上传成功", i+1)
		time.Sleep(1 * time.Second)
	}

	return nil
}

// waitForImageUpload 等待图片上传完成
func (p *Publisher) waitForImageUpload(expectedCount int) error {
	maxWait := 60 * time.Second
	start := time.Now()

	for time.Since(start) < maxWait {
		// 查找已上传的图片预览
		images, err := p.page.Elements(".img-preview-area .pr")
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if len(images) >= expectedCount {
			logrus.Infof("图片上传完成，共%d张", len(images))
			return nil
		}

		time.Sleep(500 * time.Millisecond)
	}

	return errors.Errorf("等待图片上传超时（%ds）", int(maxWait.Seconds()))
}

// fillTitle 填写标题
func (p *Publisher) fillTitle(title string) error {
	logrus.Infof("填写标题: %s", title)

	titleInput, err := p.page.Element("div.d-input input")
	if err != nil {
		return err
	}

	if err := titleInput.Input(title); err != nil {
		return err
	}

	// 检查标题长度
	time.Sleep(500 * time.Millisecond)
	hasError, _, err := p.page.Has(`div.title-container div.max_suffix`)
	if err == nil && hasError {
		return errors.New("标题长度超过限制")
	}

	return nil
}

// fillContent 填写内容
func (p *Publisher) fillContent(content string) error {
	logrus.Info("填写内容...")

	// 查找内容编辑器
	contentEditor, err := p.page.Element("div.ql-editor")
	if err != nil {
		// 尝试其他选择器
		contentEditor, err = p.page.Element(`div[role="textbox"]`)
		if err != nil {
			return errors.New("未找到内容编辑器")
		}
	}

	if err := contentEditor.Input(content); err != nil {
		return err
	}

	// 检查内容长度
	time.Sleep(500 * time.Millisecond)
	hasError, _, err := p.page.Has(`div.edit-container div.length-error`)
	if err == nil && hasError {
		return errors.New("内容长度超过限制")
	}

	return nil
}

// addTags 添加标签
func (p *Publisher) addTags(tags []string) error {
	logrus.Infof("添加%d个标签...", len(tags))

	// 找到内容编辑器
	contentEditor, err := p.page.Element("div.ql-editor")
	if err != nil {
		contentEditor, err = p.page.Element(`div[role="textbox"]`)
		if err != nil {
			return errors.New("未找到内容编辑器")
		}
	}

	// 移动到内容末尾
	for i := 0; i < 20; i++ {
		ka, err := contentEditor.KeyActions()
		if err != nil {
			continue
		}
		ka.Type(input.ArrowDown).MustDo()
		time.Sleep(10 * time.Millisecond)
	}

	// 按回车换行
	ka, _ := contentEditor.KeyActions()
	ka.Press(input.Enter).Press(input.Enter).MustDo()

	time.Sleep(1 * time.Second)

	// 输入标签
	for _, tag := range tags {
		tag = strings.TrimLeft(tag, "#")

		// 输入#
		if err := contentEditor.Input("#"); err != nil {
			logrus.Warnf("输入#失败: %v", err)
			continue
		}
		time.Sleep(200 * time.Millisecond)

		// 输入标签文本
		for _, char := range tag {
			contentEditor.Input(string(char))
			time.Sleep(50 * time.Millisecond)
		}

		time.Sleep(1 * time.Second)

		// 尝试点击标签联想
		topicContainer, err := p.page.Element("#creator-editor-topic-container")
		if err == nil {
			firstItem, err := topicContainer.Element(".item")
			if err == nil {
				firstItem.Click(proto.InputMouseButtonLeft, 1)
				logrus.Infof("已添加标签: #%s", tag)
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

// setScheduleTime 设置定时发布
func (p *Publisher) setScheduleTime(t time.Time) error {
	logrus.Infof("设置定时发布: %s", t.Format("2006-01-02 15:04"))

	// 点击定时发布开关
	switchBtn, err := p.page.Element(".post-time-wrapper .d-switch")
	if err != nil {
		return err
	}

	if err := switchBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return err
	}

	time.Sleep(800 * time.Millisecond)

	// 设置日期时间
	dateTimeStr := t.Format("2006-01-02 15:04")
	input, err := p.page.Element(".date-picker-container input")
	if err != nil {
		return err
	}

	if err := input.SelectAllText(); err != nil {
		logrus.Warnf("选择文本失败: %v", err)
	}

	if err := input.Input(dateTimeStr); err != nil {
		return err
	}

	logrus.Infof("已设置定时发布: %s", dateTimeStr)
	return nil
}

// setVisibility 设置可见范围
func (p *Publisher) setVisibility(visibility string) error {
	logrus.Infof("设置可见范围: %s", visibility)

	// 点击可见范围下拉框
	dropdown, err := p.page.Element("div.permission-card-wrapper div.d-select-content")
	if err != nil {
		return err
	}

	if err := dropdown.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	// 查找并点击目标选项
	opts, err := p.page.Elements("div.d-options-wrapper div.d-grid-item div.custom-option")
	if err != nil {
		return err
	}

	for _, opt := range opts {
		text, err := opt.Text()
		if err != nil {
			continue
		}

		if strings.Contains(text, visibility) {
			if err := opt.Click(proto.InputMouseButtonLeft, 1); err != nil {
				return err
			}
			logrus.Infof("已设置可见范围: %s", visibility)
			return nil
		}
	}

	return errors.Errorf("未找到可见范围选项: %s", visibility)
}

// setOriginal 设置原创声明
func (p *Publisher) setOriginal() error {
	logrus.Info("设置原创声明...")

	// 查找原创声明开关
	switchCards, err := p.page.Elements("div.custom-switch-card")
	if err != nil {
		return err
	}

	for _, card := range switchCards {
		text, err := card.Text()
		if err != nil {
			continue
		}

		if !strings.Contains(text, "原创声明") {
			continue
		}

		// 找到原创声明卡片，点击开关
		switchElem, err := card.Element("div.d-switch")
		if err != nil {
			continue
		}

		if err := switchElem.Click(proto.InputMouseButtonLeft, 1); err != nil {
			return err
		}

		logrus.Info("已开启原创声明")
		time.Sleep(500 * time.Millisecond)

		// 处理确认弹窗
		p.handleOriginalConfirm()

		return nil
	}

	return errors.New("未找到原创声明选项")
}

// handleOriginalConfirm 处理原创声明确认弹窗
func (p *Publisher) handleOriginalConfirm() {
	time.Sleep(800 * time.Millisecond)

	// 勾选须知
	checkbox, err := p.page.Element("div.footer div.d-checkbox input[type='checkbox']")
	if err == nil {
		checkbox.Click(proto.InputMouseButtonLeft, 1)
		time.Sleep(300 * time.Millisecond)
	}

	// 点击声明原创按钮
	btn, err := p.page.Element("div.footer button.custom-button")
	if err == nil {
		btn.Click(proto.InputMouseButtonLeft, 1)
		logrus.Info("已确认原创声明")
	}
}

// clickPublish 点击发布按钮
func (p *Publisher) clickPublish() error {
	logrus.Info("点击发布按钮...")

	submitBtn, err := p.page.Element(".publish-page-publish-btn button.bg-red")
	if err != nil {
		return err
	}

	if err := submitBtn.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return err
	}

	// 等待发布完成
	time.Sleep(3 * time.Second)

	logrus.Info("发布按钮点击成功")
	return nil
}
