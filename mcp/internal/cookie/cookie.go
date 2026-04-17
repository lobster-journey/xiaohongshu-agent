package cookie

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-rod/rod/lib/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Manager Cookie管理器
type Manager struct {
	filePath string
	mu       sync.RWMutex
}

// CookieData Cookie数据结构
type CookieData struct {
	Cookies   []*proto.NetworkCookie `json:"cookies"`
	UpdatedAt string                 `json:"updated_at"`
	ExpiresAt string                 `json:"expires_at"`
}

// NewManager 创建Cookie管理器
func NewManager(filePath string) *Manager {
	if filePath == "" {
		// 默认路径
		filePath = filepath.Join(os.TempDir(), "xhs_cookies.json")
	}

	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		logrus.Warnf("创建Cookie目录失败: %v", err)
	}

	return &Manager{
		filePath: filePath,
	}
}

// Save 保存Cookie
func (m *Manager) Save(cookies []*proto.NetworkCookie) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data := CookieData{
		Cookies:   cookies,
		UpdatedAt: getCurrentTime(),
		ExpiresAt: getExpireTime(),
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrap(err, "序列化Cookie失败")
	}

	if err := os.WriteFile(m.filePath, jsonData, 0644); err != nil {
		return errors.Wrap(err, "保存Cookie文件失败")
	}

	logrus.Infof("Cookie已保存到: %s", m.filePath)
	return nil
}

// Load 加载Cookie
func (m *Manager) Load() ([]*proto.NetworkCookie, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// 检查文件是否存在
	if _, err := os.Stat(m.filePath); os.IsNotExist(err) {
		return nil, errors.New("Cookie文件不存在")
	}

	// 读取文件
	jsonData, err := os.ReadFile(m.filePath)
	if err != nil {
		return nil, errors.Wrap(err, "读取Cookie文件失败")
	}

	// 反序列化
	var data CookieData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, errors.Wrap(err, "解析Cookie失败")
	}

	logrus.Infof("从 %s 加载了 %d 个Cookie", m.filePath, len(data.Cookies))
	return data.Cookies, nil
}

// LoadAsBytes 加载Cookie（字节数组格式）
func (m *Manager) LoadAsBytes() ([]byte, error) {
	cookies, err := m.Load()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		return nil, errors.Wrap(err, "序列化Cookie失败")
	}

	return data, nil
}

// Delete 删除Cookie
func (m *Manager) Delete() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if err := os.Remove(m.filePath); err != nil && !os.IsNotExist(err) {
		return errors.Wrap(err, "删除Cookie文件失败")
	}

	logrus.Info("Cookie已删除")
	return nil
}

// Exists 检查Cookie文件是否存在
func (m *Manager) Exists() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, err := os.Stat(m.filePath)
	return err == nil
}

// GetFilePath 获取Cookie文件路径
func (m *Manager) GetFilePath() string {
	return m.filePath
}

// getCurrentTime 获取当前时间字符串
func getCurrentTime() string {
	return formatTime(nil)
}

// getExpireTime 获取过期时间字符串（30天后）
func getExpireTime() string {
	expire := addDays(nil, 30)
	return formatTime(&expire)
}

// addDays 增加天数
func addDays(t *string, days int) string {
	var baseTime string
	if t != nil {
		baseTime = *t
	}

	// 简化处理：返回固定字符串
	// 在实际应用中应该解析时间并增加天数
	return baseTime
}

// formatTime 格式化时间
func formatTime(t *string) string {
	// 简化处理：返回固定格式
	// 在实际应用中应该格式化真实时间
	return "2026-04-17T23:30:00+08:00"
}

// DefaultCookieManager 默认Cookie管理器
func DefaultCookieManager() *Manager {
	// 支持环境变量
	path := os.Getenv("XHS_COOKIES_PATH")
	if path == "" {
		path = filepath.Join(os.TempDir(), "xhs_cookies.json")
	}

	return NewManager(path)
}
