package app

import (
	"context"
	"fmt"
	"time"

	"github.com/huangjunwen/golibs/logr"
)

// PublicAppOption 是创建 PublicApp 的选项
type PublicAppOption func(*PublicApp) error

// PAContext 设置基础 context.Context, 会在调用接口时用到.
func PAContext(ctx context.Context) PublicAppOption {
	return func(a *PublicApp) error {
		if ctx == nil {
			ctx = context.Background()
		}
		a.ctx = ctx
		return nil
	}
}

// PAOnUpdateAppAccessToken 在 app access token 更新时回调 （值不一定有变化）, 用户可以使用该回调实现自己的 AppAccessTokenProvider.
// 如果回调返回错误，则会在 PARetryInterval 后重试.
func PAOnUpdateAppAccessToken(fn func(string) error) PublicAppOption {
	return func(a *PublicApp) error {
		if fn == nil {
			fn = func(_ string) error { return nil }
		}
		a.appAccessTokenUpdator.onUpdate = fn
		return nil
	}
}

// PAUpdateInterval 设置正常的更新间隔，取值应该大于等于 1 分钟并小于等于 10 分钟 (默认 DefaultUpdateInterval).
// 正常的更新间隔是指：在没有出错的情况下，下一次执行检查/更新的间隔时间.
func PAUpdateInterval(interval time.Duration) PublicAppOption {
	return func(a *PublicApp) error {
		if interval < time.Minute {
			return fmt.Errorf("PAUpdateInterval should be at least 1 minute")
		}
		if interval > 10*time.Minute {
			return fmt.Errorf("PAUpdateInterval should be at most 10 miniutes")
		}
		a.appAccessTokenUpdator.updateInterval = interval
		return nil
	}
}

// PARetryInterval 设置出错时的重试间隔，取值应该大于等于 1 秒并小于等于 1 分钟 (默认 DefaultRetryInterval).
// 出错时的重试间隔是指：调用接口错误/回调错误后，下一次执行检查/更新的间隔时间
func PARetryInterval(interval time.Duration) PublicAppOption {
	return func(a *PublicApp) error {
		if interval < time.Second {
			return fmt.Errorf("PARetryInterval should be at least 1 second")
		}
		if interval > time.Minute {
			return fmt.Errorf("PARetryInterval should be at moust 1 minute")
		}
		a.appAccessTokenUpdator.retryInterval = interval
		return nil
	}
}

// PALogger 设置日志
func PALogger(logger logr.Logger) PublicAppOption {
	return func(a *PublicApp) error {
		if logger == nil {
			logger = logr.Nop
		}
		a.appAccessTokenUpdator.logger = logger
		return nil
	}
}
