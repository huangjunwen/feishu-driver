package message

import (
	"context"
	"fmt"

	"github.com/huangjunwen/feishu-driver/conf"
	"github.com/huangjunwen/feishu-driver/utils"
)

// Send 发送消息
type Send struct {
	// 给用户发私聊消息，只需要填 open_id、email、user_id 中的一个即可，向群里发消息使用群的 chat_id
	OpenId string `json"open_id,omitempty"`
	UserId string `json"user_id,omitempty"`
	Email  string `json"email,omitempty"`
	ChatId string `json"chat_id,omitempty"`

	// RootId 是回复某条消息对应消息的消息 ID
	RootId string `json:"root_id,omitempty"`

	// Content 是实际内容
	Content SendContent `json:"content"`

	// MsgType 是内容类型，不需要填写，由 Content.SendContentMsgType 获得
	MsgType string `json:"msg_type"`
}

// SendResult 是发消息接口的结果
type SendResult struct {
	utils.APIResultBase

	Data struct {
		MessageId string `json:"message_id"`
	} `json:"data"`
}

// Do 调用 api
func (send Send) Do(ctx context.Context, provider conf.TenantAccessTokenProvider) (*SendResult, error) {
	if send.Content == nil {
		return nil, fmt.Errorf("Missing content")
	}
	send.MsgType = send.Content.SendContentMsgType()
	result := &SendResult{}
	err := utils.PostJSONWithTenantAccessToken(ctx, "/message/v4/send", provider, send, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
