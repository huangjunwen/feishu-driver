package webhook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"

	"github.com/tidwall/gjson"

	"github.com/huangjunwen/feishu-driver/conf"
	"github.com/huangjunwen/feishu-driver/webhook/events"
)

var (
	_ http.Handler           = (*Handler)(nil)
	_ conf.AppTicketProvider = (*Handler)(nil)
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

// Handler 是一个 http.Handler，用于处理订阅事件回调/自动处理 url_verification，
// 同时它也满足 AppTicketProvider, 可为应用商店应用提供 app ticket
type Handler struct {
	verifToken string
	decrypter  *decrypter
	handler    PayloadHandler

	appTicket atomic.Value // string
}

// GetEvent 返回事件，仅当 type 是 event_callback 时返回非 nil
func (payload *Payload) GetEvent() interface{} {
	if payload.Type != PayloadTypeEventCallback {
		return nil
	}
	return payload.event
}

// New 创建一个 Handler，cnf 必须提供，handler 可用于处理感兴趣的事件，
// 也可以传入 nil
func New(cnf conf.WebhookConfig, handler PayloadHandler) *Handler {

	verifToken := cnf.FeishuWebhookVerifToken()
	if verifToken == "" {
		panic(fmt.Errorf("Empty feishu verify token"))
	}

	var decrypter *decrypter
	if key := cnf.FeishuWebhookEncryptKey(); key != "" {
		decrypter = newDecrypter(key)
	}

	if handler == nil {
		handler = func(w http.ResponseWriter, r *http.Request, payload *Payload) {}
	}

	return &Handler{
		verifToken: verifToken,
		decrypter:  decrypter,
		handler:    handler,
	}

}

// ServeHTTP 满足 http.Handler 接口
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	invalidPayload := func(code int) {
		http.Error(w, "Invalid payload", code)
	}

	// 加密模式
	if h.decrypter != nil {
		encryptPayload := &struct {
			Encrypt string `json:"encrypt"`
		}{}
		err := json.Unmarshal(body, encryptPayload)
		if err != nil {
			invalidPayload(461)
			return
		}

		body, err = h.decrypter.Decrypt(encryptPayload.Encrypt)
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
	if payload.Token != h.verifToken {
		invalidPayload(463)
		return
	}

	switch payload.Type {
	default:
		h.handler(w, r, payload)
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

		switch e := ev.(type) {
		case *events.AppTicket:
			h.appTicket.Store(e.AppTicket)
		}

		h.handler(w, r, payload)
		return
	}

}

// FeishuAppTicket 满足 AppTicketProvider 接口
func (h *Handler) FeishuAppTicket() (string, error) {
	v := h.appTicket.Load()
	if v == nil {
		return "", fmt.Errorf("No app ticket yet")
	}
	return v.(string), nil
}
