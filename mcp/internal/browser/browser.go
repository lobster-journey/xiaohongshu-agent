package browser

import (
	"encoding/json"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Browser 浏览器实例
type Browser struct {
	browser *rod.Browser
	page    *rod.Page
	headless bool
	binPath string
	proxy   string
}

// Config 浏览器配置
type Config struct {
	Headless bool
	BinPath  string
	Proxy    string
	Cookies  []byte
}

// New 创建新的浏览器实例
func New(cfg Config) (*Browser, error) {
	// 启动浏览器
	l := launcher.New().
		Headless(cfg.Headless).
		NoSandbox(true)

	if cfg.BinPath != "" {
		l.Bin(cfg.BinPath)
	}

	if cfg.Proxy != "" {
		l.Proxy(cfg.Proxy)
		logrus.Infof("使用代理: %s", cfg.Proxy)
	}

	url, err := l.Launch()
	if err != nil {
		return nil, errors.Wrap(err, "启动浏览器失败")
	}

	browser := rod.New().ControlURL(url).MustConnect()

	// 创建页面
	page, err := browser.Page(proto.TargetCreateTarget{URL: "about:blank"})
	if err != nil {
		browser.MustClose()
		return nil, errors.Wrap(err, "创建页面失败")
	}

	// 设置超时
	page = page.Timeout(5 * time.Minute)

	// 设置视窗大小
	if err := page.SetViewport(&proto.EmulationSetDeviceMetricsOverride{
		Width:  1920,
		Height: 1080,
		Scale:  proto.Float64(1.0),
	}); err != nil {
		logrus.Warnf("设置视窗大小失败: %v", err)
	}

	// 加载Cookie
	if len(cfg.Cookies) > 0 {
		var cookies []*proto.NetworkCookie
		if err := json.Unmarshal(cfg.Cookies, &cookies); err != nil {
			logrus.Warnf("解析Cookie失败: %v", err)
		} else {
			for _, cookie := range cookies {
				if err := browser.SetCookies([]*proto.NetworkCookieParam{{
					Name:     cookie.Name,
					Value:    cookie.Value,
					Domain:   cookie.Domain,
					Path:     cookie.Path,
					HTTPOnly: cookie.HTTPOnly,
					Secure:   cookie.Secure,
				}}); err != nil {
					logrus.Warnf("设置Cookie失败: %v", err)
				}
			}
			logrus.Info("已加载Cookie")
		}
	}

	return &Browser{
		browser:  browser,
		page:     page,
		headless: cfg.Headless,
		binPath:  cfg.BinPath,
		proxy:    cfg.Proxy,
	}, nil
}

// Page 获取页面实例
func (b *Browser) Page() *rod.Page {
	return b.page
}

// Close 关闭浏览器
func (b *Browser) Close() error {
	if b.browser != nil {
		b.browser.MustClose()
	}
	return nil
}

// GetCookies 获取当前Cookie
func (b *Browser) GetCookies() ([]byte, error) {
	cookies, err := b.browser.GetCookies()
	if err != nil {
		return nil, errors.Wrap(err, "获取Cookie失败")
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		return nil, errors.Wrap(err, "序列化Cookie失败")
	}

	return data, nil
}

// Navigate 导航到指定URL
func (b *Browser) Navigate(url string) error {
	if err := b.page.Navigate(url); err != nil {
		return errors.Wrapf(err, "导航到 %s 失败", url)
	}

	// 等待页面加载
	if err := b.page.WaitLoad(); err != nil {
		logrus.Warnf("等待页面加载失败: %v", err)
	}

	// 额外等待DOM稳定
	time.Sleep(2 * time.Second)

	return nil
}

// Screenshot 截图
func (b *Browser) Screenshot() ([]byte, error) {
	screenshot, err := b.page.Screenshot(false, nil)
	if err != nil {
		return nil, errors.Wrap(err, "截图失败")
	}
	return screenshot, nil
}

// WaitElement 等待元素出现
func (b *Browser) WaitElement(selector string, timeout time.Duration) (*rod.Element, error) {
	page := b.page.Timeout(timeout)
	defer page.CancelTimeout()

	elem, err := page.Element(selector)
	if err != nil {
		return nil, errors.Wrapf(err, "等待元素 %s 失败", selector)
	}

	return elem, nil
}

// Click 点击元素
func (b *Browser) Click(selector string) error {
	elem, err := b.page.Element(selector)
	if err != nil {
		return errors.Wrapf(err, "查找元素 %s 失败", selector)
	}

	if err := elem.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return errors.Wrapf(err, "点击元素 %s 失败", selector)
	}

	return nil
}

// Input 输入文本
func (b *Browser) Input(selector, text string) error {
	elem, err := b.page.Element(selector)
	if err != nil {
		return errors.Wrapf(err, "查找元素 %s 失败", selector)
	}

	if err := elem.Input(text); err != nil {
		return errors.Wrapf(err, "输入文本到 %s 失败", selector)
	}

	return nil
}

// GetText 获取元素文本
func (b *Browser) GetText(selector string) (string, error) {
	elem, err := b.page.Element(selector)
	if err != nil {
		return "", errors.Wrapf(err, "查找元素 %s 失败", selector)
	}

	text, err := elem.Text()
	if err != nil {
		return "", errors.Wrapf(err, "获取元素 %s 文本失败", selector)
	}

	return text, nil
}

// IsLoggedIn 检查是否已登录（通用方法）
func (b *Browser) IsLoggedIn(checkURL, selector string) (bool, error) {
	if err := b.Navigate(checkURL); err != nil {
		return false, err
	}

	// 检查是否存在登录元素
	has, _, err := b.page.Has(selector)
	if err != nil {
		return false, err
	}

	// 如果存在登录元素，说明未登录
	return !has, nil
}

// DefaultBrowserConfig 获取默认浏览器配置
func DefaultBrowserConfig() Config {
	headless := os.Getenv("BROWSER_HEADLESS") != "false"
	binPath := os.Getenv("BROWSER_BIN_PATH")
	proxy := os.Getenv("BROWSER_PROXY")

	return Config{
		Headless: headless,
		BinPath:  binPath,
		Proxy:    proxy,
	}
}
