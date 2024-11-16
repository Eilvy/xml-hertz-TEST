package main

import (
	"context"
	"encoding/xml"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
)

func main() {
	r := *server.Default()

	r.POST("/enter", HReply)

	r.Run()

}

type HRequestMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        int64
	Content      string
}

func HReply(ctx context.Context, c *app.RequestContext) {
	// 解析一下消息 目前只接受文本 其他全部丢掉
	base := HRequestMessage{}

	data := c.Request.Body()

	err := xml.Unmarshal(data, &base)
	if err != nil {
		log.Printf("xml unmarshal err: %v", err)
		return
	}

	log.Println(string(data))

	log.Println(base)

}
