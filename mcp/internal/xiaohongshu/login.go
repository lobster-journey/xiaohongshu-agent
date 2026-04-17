package xiaohongshu

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Cody-Chan/xiaohongshu-agent/internal/browser"
	"github.com/go-rod/rod"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// 登录检查URL
	loginCheckURL = "https://creator.xiaohongshu.com"
	// 登录URL
	loginURL = "https://www.xiaohongshu.com"
	// 二维码选择器
	qrcodeSelector = "img.login-qrcode"
)

// Login 登录操作
type Login struct {
	browser *browser.Browser
	page    *rod.Page
}

// NewLogin 创建登录操作实例
func NewLogin(b *browser.Browser) *Login {
	return &Login{
		browser: b,
		page:    b.Page(),
	}
}

// CheckLoginStatus 检查登录状态
func (l *Login) CheckLoginStatus(ctx context.Context) (bool, error) {
	logrus.Info("检查小红书登录状态...")

	// 访问创作者中心
	if err := l.browser.Navigate(loginCheckURL); err != nil {
		return false, err
	}

	// 等待页面加载
	time.Sleep(2 * time.Second)

	// 检查是否存在登录按钮
	hasLoginBtn, _, err := l.page.Has("text=登录")
	if err != nil {
		return false, errors.Wrap(err, "检查登录按钮失败")
	}

	// 如果存在登录按钮，说明未登录
	if hasLoginBtn {
		logrus.Info("未登录")
		return false, nil
	}

	// 检查是否存在用户头像或用户名
	hasUserAvatar, _, err := l.page.Has(".user-avatar, .username, [class*='user']")
	if err != nil {
		return false, errors.Wrap(err, "检查用户信息失败")
	}

	if hasUserAvatar {
		logrus.Info("已登录")
		return true, nil
	}

	// 默认认为未登录
	logrus.Info("登录状态未知，默认未登录")
	return false, nil
}

// GetLoginQrcode 获取登录二维码
func (l *Login) GetLoginQrcode(ctx context.Context) (qrcodeBase64 string, err error) {
	logrus.Info("获取小红书登录二维码...")

	// 访问小红书首页
	if err := l.browser.Navigate(loginURL); err != nil {
		return "", err
	}

	// 等待页面加载
	time.Sleep(2 * time.Second)

	// 点击登录按钮
	if err := l.browser.Click("text=登录"); err != nil {
		logrus.Warnf("点击登录按钮失败: %v", err)
	}

	// 等待二维码出现
	time.Sleep(2 * time.Second)

	// 查找二维码图片
	qrcodeImg, err := l.page.Element(qrcodeSelector)
	if err != nil {
		return "", errors.Wrap(err, "查找二维码失败")
	}

	// 获取二维码图片的src属性
	src, err := qrcodeImg.Attribute("src")
	if err != nil {
		return "", errors.Wrap(err, "获取二维码src失败")
	}

	if src == nil || *src == "" {
		// 如果没有src，尝试截图
		screenshot, err := qrcodeImg.Screenshot()
		if err != nil {
			return "", errors.Wrap(err, "二维码截图失败")
		}
		qrcodeBase64 = base64.StdEncoding.EncodeToString(screenshot)
	} else {
		// 如果src是base64格式
		if len(*src) > 22 && (*src)[:22] == "data:image/png;base64," {
			qrcodeBase64 = (*src)[22:]
		} else {
			// 如果是URL，下载图片并转base64（简化处理，直接返回URL）
			qrcodeBase64 = *src
		}
	}

	logrus.Info("二维码获取成功")
	return qrcodeBase64, nil
}

// WaitForLogin 等待用户扫码登录
func (l *Login) WaitForLogin(ctx context.Context, timeout time.Duration) error {
	logrus.Info("等待用户扫码登录...")

	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// 检查是否登录成功
		loggedIn, err := l.CheckLoginStatus(ctx)
		if err != nil {
			logrus.Warnf("检查登录状态失败: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if loggedIn {
			logrus.Info("登录成功！")
			return nil
		}

		time.Sleep(3 * time.Second)
	}

	return errors.New("等待登录超时")
}

// LoginWithQrcode 使用二维码登录（完整流程）
func (l *Login) LoginWithQrcode(ctx context.Context) error {
	// 1. 获取二维码
	qrcode, err := l.GetLoginQrcode(ctx)
	if err != nil {
		return err
	}

	logrus.Info("请扫描二维码登录")
	// 在实际使用中，应该将二维码显示给用户或发送到前端
	logrus.Infof("二维码: %s...", qrcode[:50])

	// 2. 等待登录
	if err := l.WaitForLogin(ctx, 4*time.Minute); err != nil {
		return err
	}

	return nil
}

// Logout 登出
func (l *Login) Logout(ctx context.Context) error {
	logrus.Info("执行登出操作...")

	// 访问小红书首页
	if err := l.browser.Navigate(loginURL); err != nil {
		return err
	}

	// 点击用户头像
	if err := l.browser.Click(".user-avatar, [class*='user']"); err != nil {
		return errors.Wrap(err, "点击用户头像失败")
	}

	time.Sleep(1 * time.Second)

	// 点击登出
	if err := l.browser.Click("text=退出登录, text=退出"); err != nil {
		return errors.Wrap(err, "点击退出登录失败")
	}

	// 确认登出
	if err := l.browser.Click("text=确定, text=确认"); err != nil {
		logrus.Warnf("确认登出失败: %v", err)
	}

	logrus.Info("登出成功")
	return nil
}

// GetUserInfo 获取用户信息
func (l *Login) GetUserInfo(ctx context.Context) (username string, err error) {
	logrus.Info("获取用户信息...")

	// 访问创作者中心
	if err := l.browser.Navigate(loginCheckURL); err != nil {
		return "", err
	}

	time.Sleep(2 * time.Second)

	// 查找用户名元素
	usernameElem, err := l.page.Element(".username, [class*='username'], [class*='name']")
	if err != nil {
		return "", errors.Wrap(err, "查找用户名失败")
	}

	username, err = usernameElem.Text()
	if err != nil {
		return "", errors.Wrap(err, "获取用户名文本失败")
	}

	logrus.Infof("用户名: %s", username)
	return username, nil
}
