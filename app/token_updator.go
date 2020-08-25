package app

import (
	"fmt"
	"sync"
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

	token atomic.Value // string, nil 表示未有 token

	mu    sync.Mutex
	timer *time.Timer // nil 表示未启动（停止），非 nil （即使 timer 停止）表示已经启动
}

// Get 获得当前已知最新的 token，无论处于停止还是启动状态
func (updator *tokenUpdator) Get() (string, error) {
	v := updator.token.Load()
	if v == nil {
		return "", fmt.Errorf("Updator(%s) has no token yet", updator.name)
	}
	return v.(string), nil
}

// Start 启动已经停止了的 updator，如果已经启动了则 nop
func (updator *tokenUpdator) Start() error {
	updator.mu.Lock()
	defer updator.mu.Unlock()

	if updator.timer != nil {
		return nil
	}

	updator.timer = time.NewTimer(time.Hour)
	updator.stopTimer()

	return updator.update("", time.Now())
}

// Stop 停止已经启动了的 updator，如果已经停止了则 nop
func (updator *tokenUpdator) Stop() {
	updator.mu.Lock()
	defer updator.mu.Unlock()

	if updator.timer == nil {
		return
	}

	updator.stopTimer()
	updator.timer = nil
}

// NOTE: 该函数必须由 mutex 包裹
func (updator *tokenUpdator) update(currToken string, currExpire time.Time) (err error) {

	if updator.timer == nil {
		// 已经停止了，则直接返回, 这是有可能的（虽然可能性比较微小）:
		// 在计时器触发 update 时刚好 Stop 了
		updator.logger.Info("Token update but stopped", "updator", updator.name)
		return nil
	}
	updator.stopTimer()

	// 有错误时（包括调接口错误，更新回调错误等），均在 retryInterval 后重试，
	// 否则在 updateInterval 后重试
	defer func() {
		var interval time.Duration
		if err != nil {
			interval = updator.retryInterval
			updator.logger.Error(err, "Token update error", "updator", updator.name)
		} else {
			interval = updator.updateInterval
		}
		updator.timer = time.AfterFunc(interval, func() {
			updator.mu.Lock()
			defer updator.mu.Unlock()
			updator.update(currToken, currExpire)
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

func (updator *tokenUpdator) stopTimer() {
	if !updator.timer.Stop() {
		<-updator.timer.C
	}
}
