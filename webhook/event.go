package webhook

import (
	"github.com/huangjunwen/feishu-driver/webhook/events"
)

var (
	newEvMaps = map[string]func() interface{}{}
)

// RegistEv 注册一种订阅事件类型，typ 是 payload["event"]["type"] 字段, newEv 是工厂函数，
// 用于创建一个新的该类型事件
func RegistEv(typ string, newEv func() interface{}) {
	newEvMaps[typ] = newEv
}

func init() {
	events.Regist(RegistEv)
}
