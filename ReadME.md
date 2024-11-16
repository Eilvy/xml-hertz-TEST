### 2024.11.16记录

使用`cloudwego/hertz`作为服务端框架

写公众号自动回复接收微信服务器发来的消息（为`xml`信息）时没有直接绑定`xml`的方法

可以直接通过`app.RequestContext`中的`Body()`拿到body之后进行`xml`解析

```go
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
```