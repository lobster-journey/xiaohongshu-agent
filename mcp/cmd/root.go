package cmd

import (
	"fmt"
	"os"

	"github.com/Cody-Chan/xiaohongshu-agent/internal/server"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	port    int
)

var rootCmd = &cobra.Command{
	Use:   "xiaohongshu-mcp",
	Short: "小红书 MCP 服务",
	Long: `小红书 MCP 服务 - 提供小红书自动化能力的 MCP 服务

支持功能：
- 图文发布
- 视频发布
- 评论互动
- 内容搜索
- 数据统计`,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动服务
		if err := server.Start(port); err != nil {
			fmt.Fprintf(os.Stderr, "服务启动失败: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件路径")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 18060, "服务端口")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录小红书账号",
	Long:  `通过浏览器登录小红书账号，获取Cookie`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("正在启动浏览器登录...")
		// TODO: 实现登录逻辑
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
