package app

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/huangjunwen/golibs/logr"
)

type tokenUpdator struct {
	// 以下必须填写
	name           string                            // 用于标识该 updator
	tokenGetter    func() (string, time.Time, error) // 用于获得 token 以及其过期时间
	onUpdate       func(string) error                // 回调
	updateInterval time.Duration
	retryInterval  time.Duration
	logger         logr.Logger

	token   atomic.Value // string, nil 表示未有 token
	stopped atomic.Value // bool, nil 表示未 stopped
}

func (updator *tokenUpdator) Start() error {
	return updator.update(true, "", time.Now())
}

func (updator *tokenUpdator) Get() (string, error) {
	v := updator.token.Load()
	if v == nil {
		return "", fmt.Errorf("Updator(%s) has no token yet", updator.name)
	}
	return v.(string), nil
}

// NOTE: Stop 会等到下一次调用 update 时才起作用
func (updator *tokenUpdator) Stop() {
	updator.stopped.Store(true)
}

func (updator *tokenUpdator) update(atStart bool, currToken string, currExpire time.Time) (err error) {

	if v := updator.stopped.Load(); v != nil && v.(bool) {
		return errors.New("Updator stopped")
	}

	// 有错误时（包括调接口错误，更新回调错误等），均在 retryInterval 后重试，
	// 否则在 updateInterval 后重试
	defer func() {
		var interval time.Duration
		if err != nil {
			if atStart {
				// 如果是在 Start 中调用（第一次），不触发计时器
				return
			}
			interval = updator.retryInterval
			updator.logger.Error(err, "Token update error", "updator", updator.name)
		} else {
			interval = updator.updateInterval
		}
		time.AfterFunc(interval, func() {
			updator.update(false, currToken, currExpire)
		})
	}()

	// 若尚未有 token 或者当前时间接近过期或已经过期, 则重新调用接口
	//
	// NOTE: https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
	// Token 有效期为 2 小时，在此期间调用该接口 token 不会改变。当 token 有效期小于 10 分的时候，
	// 再次请求获取 token 的时候，会生成一个新的 token，与此同时老的 token 依然有效。
	if currToken == "" || time.Now().After(currExpire.Add(-10*time.Minute)) {
		token, expire, err := updator.tokenGetter()
		if err != nil {
			return err
		}

		currToken = token
		currExpire = expire
		updator.token.Store(currToken)
		updator.logger.Info("Token update ok", "updator", updator.name, "expiredAt", currExpire.String())
	}

	// 回调
	return updator.onUpdate(currToken)
}
