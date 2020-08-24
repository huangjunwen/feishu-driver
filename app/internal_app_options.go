package app

import (
	"context"
	"fmt"
	"time"

	"github.com/huangjunwen/golibs/logr"
)

// InternalAppOption 是创建 InternalApp 的选项
type InternalAppOption func(*InternalApp) error

// IAContext 设置基础 context.Context, 会在调用接口时用到.
func IAContext(ctx context.Context) InternalAppOption {
	return func(a *InternalApp) error {
		if ctx == nil {
			ctx = context.Background()
		}
		a.ctx = ctx
		return nil
	}
}

// IAOnUpdateAppAccessToken 在 app access token 更新时回调 （值不一定有变化）, 用户可以使用该回调实现自己的 AppAccessTokenProvider.
// 如果回调返回错误，则会在 IARetryInterval 后重试.
func IAOnUpdateAppAccessToken(fn func(string) error) InternalAppOption {
	return func(a *InternalApp) error {
		if fn == nil {
			fn = func(_ string) error { return nil }
		}
		a.appAccessTokenUpdator.onUpdate = fn
		return nil
	}
}

// IAOnUpdateTenantAccessToken 在 tenant access token 更新时回调 （值不一定有变化）, 用户可以使用该回调实现自己的 TenantAccessTokenProvider.
// 如果回调返回错误，则会在 IARetryInterval 后重试.
func IAOnUpdateTenantAccessToken(fn func(string) error) InternalAppOption {
	return func(a *InternalApp) error {
		if fn == nil {
			fn = func(_ string) error { return nil }
		}
		a.tenantAccessTokenUpdator.onUpdate = fn
		return nil
	}
}

// IAUpdateInterval 设置正常的更新间隔，取值应该大于等于 1 分钟并小于等于 10 分钟 (默认 DefaultUpdateInterval).
// 正常的更新间隔是指：在没有出错的情况下，下一次执行检查/更新的间隔时间.
func IAUpdateInterval(interval time.Duration) InternalAppOption {
	return func(a *InternalApp) error {
		if interval < time.Minute {
			return fmt.Errorf("IAUpdateInterval should be at least 1 minute")
		}
		if interval > 10*time.Minute {
			return fmt.Errorf("IAUpdateInterval should be at most 10 miniutes")
		}
		a.appAccessTokenUpdator.updateInterval = interval
		a.tenantAccessTokenUpdator.updateInterval = interval
		return nil
	}
}

// IARetryInterval 设置出错时的重试间隔，取值应该大于等于 1 秒并小于等于 1 分钟 (默认 DefaultRetryInterval).
// 出错时的重试间隔是指：调用接口错误/回调错误后，下一次执行检查/更新的间隔时间
func IARetryInterval(interval time.Duration) InternalAppOption {
	return func(a *InternalApp) error {
		if interval < time.Second {
			return fmt.Errorf("IARetryInterval should be at least 1 second")
		}
		if interval > time.Minute {
			return fmt.Errorf("IARetryInterval should be at moust 1 minute")
		}
		a.appAccessTokenUpdator.retryInterval = interval
		a.tenantAccessTokenUpdator.retryInterval = interval
		return nil
	}
}

// IALogger 设置日志
func IALogger(logger logr.Logger) InternalAppOption {
	return func(a *InternalApp) error {
		if logger == nil {
			logger = logr.Nop
		}
		a.appAccessTokenUpdator.logger = logger
		a.tenantAccessTokenUpdator.logger = logger
		return nil
	}
}
