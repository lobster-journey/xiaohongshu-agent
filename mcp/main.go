package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Cody-Chan/xiaohongshu-agent/cmd"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	if err := initConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听系统信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动服务
	go func() {
		if err := cmd.Execute(); err != nil {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 等待退出信号
	<-sigChan
	fmt.Println("\n正在关闭服务...")
	cancel()
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("$HOME/.openclaw/mcp")

	// 设置默认值
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.port", 18060)
	viper.SetDefault("server.mode", "release")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		// 配置文件不存在，使用默认值
	}

	return nil
}
