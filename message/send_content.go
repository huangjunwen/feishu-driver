package message

type SendContent interface {
	// SendContentMsgType 返回
	SendContentMsgType() string
}

// SendTextContent 代表文本消息
type SendTextContent struct {
	Text string `json:"text"`
}

// SendImageContent 代表图片消息
type SendImageContent struct {
	ImageKey string `json:"image_key"`
}

// SendShareChatContent 代表分享的群名片
type SendShareChatContent struct {
	ShareOpenChatId string `json:"share_open_chat_id"`
}

// SendContentMsgType 返回 text
func (c *SendTextContent) SendContentMsgType() string {
	return "text"
}

// SendContentMsgType 返回 image
func (c *SendImageContent) SendContentMsgType() string {
	return "image"
}

// SendContentMsgType 返回 share_chat
func (c *SendShareChatContent) SendContentMsgType() string {
	return "share_chat"
}
