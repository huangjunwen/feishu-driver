package message

import (
	"context"
	"fmt"

	"github.com/huangjunwen/feishu-driver/conf"
	"github.com/huangjunwen/feishu-driver/utils"
)

// BatchSend 批量发送消息
type BatchSend struct {
	DepartmentIds []string `json:"department_ids"`
	OpenIds       []string `json:"open_ids"`
	UserIds       []string `json:"user_ids"`

	// Content 是实际内容，目前仅支持文本消息
	Content SendContent `json:"content"`

	// MsgType 是内容类型，不需要填写，由 Content.SendContentMsgType 获得
	MsgType string `json:"msg_type"`
}

// BatchSendResult 时批量发送消息的结果
type BatchSendResult struct {
	utils.APIResultBase

	Data struct {
		MessageId            string   `json:"message_id"`
		InvalidDepartmentIds []string `json:"invalid_department_ids"`
		InvalidOpenIds       []string `json:"invalid_open_ids"`
		InvalidUserIds       []string `json:"invalid_user_ids"`
	} `json:"data"`
}

// Do 调用 api
func (send BatchSend) Do(ctx context.Context, provider conf.TenantAccessTokenProvider) (*BatchSendResult, error) {
	if send.Content == nil {
		return nil, fmt.Errorf("Missing content")
	}
	_, ok := send.Content.(*SendTextContent)
	if !ok {
		return nil, fmt.Errorf("BatchSend only support text content")
	}
	send.MsgType = send.Content.SendContentMsgType()
	result := &BatchSendResult{}
	err := utils.PostJSONWithTenantAccessToken(ctx, "/message/v4/batch_send", provider, send, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
