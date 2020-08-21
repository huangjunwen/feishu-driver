package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"

	fsdconf "github.com/huangjunwen/feishu-driver/conf"
)

// Payload 是订阅事件的 payload
type Payload struct {
	// Type 是类型: event_callback-事件推送，url_verification-url地址验证
	Type string `json:"type"`

	// Token 即 Verification Token
	Token string `json:"token"`

	// Challenge 需要原样返回. type 为 url_verification 时有
	Challenge string `json:"challenge"`

	// Timestamp 是事件发送时间，一般近似于事件发生的时间. type 为 event_callback 时有
	Timestamp string `json:"ts"`

	// UUID 是事件的唯一标识, 主要用于保证幂等性. type 为 event_callback 时有
	UUID string `json:"uuid"`

	// RawEvent 是未解析的事件内容. type 为 event_callback 时有
	RawEvent json.RawMessage `json:"event"`

	event interface{}
}

// PayloadHandler 用于处理 webhook payload
type PayloadHandler func(w http.ResponseWriter, r *http.Request, payload *Payload)

// NewHandler 创建一个订阅事件的处理器, 它会自动处理 url_verification 回调，
// payloadHandler 主要用于处理各类事件
func NewHandler(cnf fsdconf.WebhookConfig, payloadHandler PayloadHandler) http.Handler {

	token := cnf.FeishuWebhookVerifToken()
	if token == "" {
		panic(fmt.Errorf("Empty feishu verify token"))
	}

	var d *decrypter
	if key := cnf.FeishuWebhookEncryptKey(); key != "" {
		d = newDecrypter(key)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		invalidPayload := func(code int) {
			http.Error(w, "Invalid payload", code)
		}

		// 加密模式
		if d != nil {
			encryptPayload := &struct {
				Encrypt string `json:"encrypt"`
			}{}
			err := json.Unmarshal(body, encryptPayload)
			if err != nil {
				invalidPayload(461)
				return
			}

			body, err = d.Decrypt(encryptPayload.Encrypt)
			if err != nil {
				invalidPayload(461)
				return
			}
		}

		payload := new(Payload)
		err = json.Unmarshal(body, payload)
		if err != nil {
			invalidPayload(462)
			return
		}

		// 验证 token
		if payload.Token != token {
			invalidPayload(463)
			return
		}

		switch payload.Type {
		default:
			payloadHandler(w, r, payload)
			return

		case PayloadTypeURLVerification:
			// NOTE: 这里忽略错误
			json.NewEncoder(w).Encode(map[string]interface{}{
				"challenge": payload.Challenge,
			})
			return

		case PayloadTypeEventCallback:
			newEv := newEvMaps[gjson.GetBytes(payload.RawEvent, "type").Str]
			if newEv == nil {
				newEv = newEvMaps[""]
			}
			ev := newEv()

			if err := json.Unmarshal(payload.RawEvent, ev); err != nil {
				panic(err)
			}
			payload.event = ev
			payloadHandler(w, r, payload)
			return
		}

	})
}

// GetEvent 返回事件，仅当 type 是 event_callback 时返回非 nil
func (payload *Payload) GetEvent() interface{} {
	if payload.Type != PayloadTypeEventCallback {
		return nil
	}
	return payload.event
}
