package app

import (
	"context"
	"fmt"
	"time"

	"github.com/huangjunwen/golibs/logr"
)

// PublicAppTenantOption 是创建 PublicAppTenant 的选项
type PublicAppTenantOption func(*PublicAppTenant) error

// PATContext 设置基础 context.Context, 会在调用接口时用到.
func PATContext(ctx context.Context) PublicAppTenantOption {
	return func(t *PublicAppTenant) error {
		if ctx == nil {
			ctx = context.Background()
		}
		t.ctx = ctx
		return nil
	}
}

// PATOnUpdateTenantAccessToken 在 tenant access token 更新时回调 （值不一定有变化）, 用户可以使用该回调实现自己的 TenantAccessTokenProvider.
// 如果回调返回错误，则会在 PATRetryInterval 后重试.
func PATOnUpdateTenantAccessToken(fn func(string) error) PublicAppTenantOption {
	return func(t *PublicAppTenant) error {
		if fn == nil {
			fn = func(_ string) error { return nil }
		}
		t.tenantAccessTokenUpdator.onUpdate = fn
		return nil
	}
}

// PATUpdateInterval 设置正常的更新间隔，取值应该大于等于 1 分钟并小于等于 10 分钟 (默认 DefaultUpdateInterval).
// 正常的更新间隔是指：在没有出错的情况下，下一次执行检查/更新的间隔时间.
func PATUpdateInterval(interval time.Duration) PublicAppTenantOption {
	return func(t *PublicAppTenant) error {
		if interval < time.Minute {
			return fmt.Errorf("PATUpdateInterval should be at least 1 minute")
		}
		if interval > 10*time.Minute {
			return fmt.Errorf("PATUpdateInterval should be at most 10 miniutes")
		}
		t.tenantAccessTokenUpdator.updateInterval = interval
		return nil
	}
}

// PATRetryInterval 设置出错时的重试间隔，取值应该大于等于 1 秒并小于等于 1 分钟 (默认 DefaultRetryInterval).
// 出错时的重试间隔是指：调用接口错误/回调错误后，下一次执行检查/更新的间隔时间
func PATRetryInterval(interval time.Duration) PublicAppTenantOption {
	return func(t *PublicAppTenant) error {
		if interval < time.Second {
			return fmt.Errorf("PATRetryInterval should be at least 1 second")
		}
		if interval > time.Minute {
			return fmt.Errorf("PATRetryInterval should be at moust 1 minute")
		}
		t.tenantAccessTokenUpdator.retryInterval = interval
		return nil
	}
}

// PATLogger 设置日志
func PATLogger(logger logr.Logger) PublicAppTenantOption {
	return func(t *PublicAppTenant) error {
		if logger == nil {
			logger = logr.Nop
		}
		t.tenantAccessTokenUpdator.logger = logger
		return nil
	}
}
