package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.POST("/enter", Reply)

	r.Run(":8080")

}

type RequestMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        int64
	Content      string
	Event        string
	EventKey     string
}

// Reply 微信的接收消息
func Reply(c *gin.Context) {
	// 解析一下消息 目前只接受文本 其他全部丢掉
	base := RequestMessage{}
	if err := c.BindXML(&base); err != nil {

		base.FromUserName = "test"
		base.ToUserName = "test"

		return
	}
	log.Println(base)

}
