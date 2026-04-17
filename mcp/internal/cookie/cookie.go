package cookie

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// CookieManager Cookie管理器
type CookieManager struct {
	filePath string
	mu       sync.RWMutex
}

// CookieData Cookie数据
type CookieData struct {
	Cookies   []Cookie `json:"cookies"`
	UpdatedAt string   `json:"updated_at"`
	ExpiresAt string   `json:"expires_at"`
}

// Cookie Cookie结构
type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Expires  int    `json:"expires"`
	HttpOnly bool   `json:"httpOnly"`
	Secure   bool   `json:"secure"`
}

// NewCookieManager 创建Cookie管理器
func NewCookieManager(filePath string) *CookieManager {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("创建目录失败: %v", err))
	}

	return &CookieManager{
		filePath: filePath,
	}
}

// Save 保存Cookie
func (cm *CookieManager) Save(cookies []Cookie) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	data := CookieData{
		Cookies:   cookies,
		UpdatedAt: time.Now().Format(time.RFC3339),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour).Format(time.RFC3339),
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化Cookie失败: %v", err)
	}

	if err := os.WriteFile(cm.filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("保存Cookie文件失败: %v", err)
	}

	return nil
}

// Load 加载Cookie
func (cm *CookieManager) Load() ([]Cookie, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	// 检查文件是否存在
	if _, err := os.Stat(cm.filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Cookie文件不存在")
	}

	// 读取文件
	jsonData, err := os.ReadFile(cm.filePath)
	if err != nil {
		return nil, fmt.Errorf("读取Cookie文件失败: %v", err)
	}

	// 反序列化
	var data CookieData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("解析Cookie失败: %v", err)
	}

	// 检查是否过期
	expiresAt, err := time.Parse(time.RFC3339, data.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("解析过期时间失败: %v", err)
	}

	if time.Now().After(expiresAt) {
		return nil, fmt.Errorf("Cookie已过期")
	}

	return data.Cookies, nil
}

// Delete 删除Cookie
func (cm *CookieManager) Delete() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if err := os.Remove(cm.filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除Cookie文件失败: %v", err)
	}

	return nil
}

// IsValid 检查Cookie是否有效
func (cm *CookieManager) IsValid() bool {
	cookies, err := cm.Load()
	if err != nil {
		return false
	}

	return len(cookies) > 0
}
