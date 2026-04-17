package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HandleWebSocket WebSocket处理器
func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket升级失败: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("WebSocket连接建立: %s", c.ClientIP())

	// 读取消息
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		log.Printf("收到消息: %s", string(message))

		// 回显消息
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("发送消息失败: %v", err)
			break
		}
	}
}
