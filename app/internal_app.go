package app

import (
	"context"
	"fmt"
	"time"

	"github.com/huangjunwen/golibs/logr"

	"github.com/huangjunwen/feishu-driver/authz"
	"github.com/huangjunwen/feishu-driver/conf"
)

var (
	_ conf.TenantAccessTokenProvider = (*InternalApp)(nil)
	_ conf.AppAccessTokenProvider    = (*InternalApp)(nil)
)

// InternalApp 是企业自建应用 (只对企业内部开放所以叫 internal), 它会提供其最新获得的 app/tenant access token,
// 并定期检查，如果接近过期或已过期则更新之.
//
// 满足 AppAccessTokenProvider/TenantAccessTokenProvider 接口，
// 使用者也可通过 IAOnUpdateAppAccessToken/IAOnUpdateTenantAccessToken 这些回调选项获得 access token
type InternalApp struct {
	appConfig                conf.AppConfig
	ctx                      context.Context
	appAccessTokenUpdator    tokenUpdator
	tenantAccessTokenUpdator tokenUpdator
}

// NewInternalApp 创建 InternalApp 并开启自动更新
func NewInternalApp(cnf conf.AppConfig, opts ...InternalAppOption) (*InternalApp, error) {
	a := &InternalApp{
		appConfig: cnf,
		ctx:       context.Background(),
		appAccessTokenUpdator: tokenUpdator{
			name:           fmt.Sprintf("IA-%s-app", cnf.FeishuAppId()),
			onUpdate:       func(string) error { return nil },
			updateInterval: DefaultUpdateInterval,
			retryInterval:  DefaultRetryInterval,
			logger:         logr.Nop,
		},
		tenantAccessTokenUpdator: tokenUpdator{
			name:           fmt.Sprintf("IA-%s-tenant", cnf.FeishuAppId()),
			onUpdate:       func(string) error { return nil },
			updateInterval: DefaultUpdateInterval,
			retryInterval:  DefaultRetryInterval,
			logger:         logr.Nop,
		},
	}
	a.appAccessTokenUpdator.tokenGetter = a.appAccessTokenGetter
	a.tenantAccessTokenUpdator.tokenGetter = a.tenantAccessTokenGetter
	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}

	a.appAccessTokenUpdator.Start()
	a.tenantAccessTokenUpdator.Start()
	return a, nil
}

func (a *InternalApp) appAccessTokenGetter() (token string, expire time.Time, err error) {
	res, err := authz.GetInternalAppAccessToken(a.ctx, a.appConfig)
	if err != nil {
		return
	}
	err = res.ResultError()
	if err != nil {
		return
	}
	return res.AppAccessToken, time.Now().Add(time.Duration(res.Expire) * time.Second), nil
}

func (a *InternalApp) tenantAccessTokenGetter() (token string, expire time.Time, err error) {
	res, err := authz.GetInternalTenantAccessToken(a.ctx, a.appConfig)
	if err != nil {
		return
	}
	err = res.ResultError()
	if err != nil {
		return
	}
	return res.TenantAccessToken, time.Now().Add(time.Duration(res.Expire) * time.Second), nil
}

// FeishuAppAccessToken 返回已知最新的 app access token, 或如果还没有获得到则返回错误
func (a *InternalApp) FeishuAppAccessToken() (string, error) {
	return a.appAccessTokenUpdator.Get()
}

// FeishuTenantAccessToken 返回已知最新的 tenant access token, 或如果还没有获得到则返回错误
func (a *InternalApp) FeishuTenantAccessToken() (string, error) {
	return a.tenantAccessTokenUpdator.Get()
}

// Stop 停止更新
func (a *InternalApp) Stop() {
	a.appAccessTokenUpdator.Stop()
	a.tenantAccessTokenUpdator.Stop()
}
