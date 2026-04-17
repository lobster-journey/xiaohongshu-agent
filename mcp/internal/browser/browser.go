package browser

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/playwright-community/playwright-go"
)

// Browser 浏览器实例
type Browser struct {
	pw      playwright.Playwright
	browser playwright.Browser
	context playwright.BrowserContext
	page    playwright.Page
	mu      sync.Mutex
}

// NewBrowser 创建浏览器实例
func NewBrowser() (*Browser, error) {
	// 安装Playwright（如果未安装）
	if err := playwright.Install(); err != nil {
		return nil, fmt.Errorf("安装Playwright失败: %v", err)
	}

	// 启动Playwright
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("启动Playwright失败: %v", err)
	}

	return &Browser{
		pw: pw,
	}, nil
}

// Launch 启动浏览器
func (b *Browser) Launch(headless bool) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// 启动浏览器
	browser, err := b.pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(headless),
	})
	if err != nil {
		return fmt.Errorf("启动浏览器失败: %v", err)
	}
	b.browser = browser

	// 创建上下文
	context, err := b.browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
		Viewport: &playwright.Size{
			Width:  1920,
			Height: 1080,
		},
	})
	if err != nil {
		return fmt.Errorf("创建上下文失败: %v", err)
	}
	b.context = context

	// 创建页面
	page, err := context.NewPage()
	if err != nil {
		return fmt.Errorf("创建页面失败: %v", err)
	}
	b.page = page

	return nil
}

// Login 登录小红书
func (b *Browser) Login(ctx context.Context) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.page == nil {
		return fmt.Errorf("浏览器未启动")
	}

	// 访问小红书
	_, err := b.page.Goto("https://www.xiaohongshu.com")
	if err != nil {
		return fmt.Errorf("访问小红书失败: %v", err)
	}

	// 点击登录按钮
	err = b.page.Click("text=登录")
	if err != nil {
		return fmt.Errorf("点击登录按钮失败: %v", err)
	}

	// 等待用户扫码登录
	// TODO: 检测登录成功状态

	return nil
}

// GetCookies 获取Cookie
func (b *Browser) GetCookies() ([]playwright.Cookie, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.context == nil {
		return nil, fmt.Errorf("浏览器上下文未创建")
	}

	cookies, err := b.context.Cookies()
	if err != nil {
		return nil, fmt.Errorf("获取Cookie失败: %v", err)
	}

	return cookies, nil
}

// SetCookies 设置Cookie
func (b *Browser) SetCookies(cookies []playwright.Cookie) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.context == nil {
		return fmt.Errorf("浏览器上下文未创建")
	}

	err := b.context.AddCookies(cookies)
	if err != nil {
		return fmt.Errorf("设置Cookie失败: %v", err)
	}

	return nil
}

// PublishImage 发布图文
func (b *Browser) PublishImage(title, content string, images []string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.page == nil {
		return fmt.Errorf("浏览器未启动")
	}

	// 访问发布页面
	_, err := b.page.Goto("https://creator.xiaohongshu.com/publish/publish")
	if err != nil {
		return fmt.Errorf("访问发布页面失败: %v", err)
	}

	// TODO: 实现发布逻辑
	// 1. 上传图片
	// 2. 填写标题和内容
	// 3. 添加标签
	// 4. 点击发布

	time.Sleep(2 * time.Second)

	return nil
}

// Close 关闭浏览器
func (b *Browser) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.page != nil {
		b.page.Close()
	}
	if b.context != nil {
		b.context.Close()
	}
	if b.browser != nil {
		b.browser.Close()
	}
	if b.pw != nil {
		b.pw.Stop()
	}

	return nil
}
