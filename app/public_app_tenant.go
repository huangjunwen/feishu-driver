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
	_ conf.TenantAccessTokenProvider = (*PublicAppTenant)(nil)
)

// PublicAppTenant 是应用商店应用租赁 (给特定企业租用), 它会提供其最新获得的 tenant access token,
// 并定期检查，如果接近过期或已过期则更新之.
//
// 满足 TenantAccessTokenProvider 接口，
// 使用者也可通过 PATOnUpdateTenantAccessToken 回调选项获得 tenant access token
type PublicAppTenant struct {
	appAccessTokenProvider   conf.AppAccessTokenProvider
	tenantKey                string
	ctx                      context.Context
	tenantAccessTokenUpdator tokenUpdator
}

// NewPublicAppTenant 创建 PublicAppTenant 并开启自动更新，
// 其中 appAccessTokenProvider 必须是应用商店应用的 app access token provider (不能使用企业自建应用)
func NewPublicAppTenant(appAccessTokenProvider conf.AppAccessTokenProvider, tenantKey string, opts ...PublicAppTenantOption) (*PublicAppTenant, error) {
	t := &PublicAppTenant{
		appAccessTokenProvider: appAccessTokenProvider,
		tenantKey:              tenantKey,
		ctx:                    context.Background(),
		tenantAccessTokenUpdator: tokenUpdator{
			name:           fmt.Sprintf("PA-tenant-%s", tenantKey),
			onUpdate:       func(string) error { return nil },
			updateInterval: DefaultUpdateInterval,
			retryInterval:  DefaultRetryInterval,
			logger:         logr.Nop,
		},
	}
	t.tenantAccessTokenUpdator.tokenGetter = t.tenantAccessTokenGetter
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return nil, err
		}
	}

	if err := t.tenantAccessTokenUpdator.Start(); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *PublicAppTenant) tenantAccessTokenGetter() (token string, expire time.Time, err error) {
	appAccessToken, err := t.appAccessTokenProvider.FeishuAppAccessToken()
	if err != nil {
		return
	}
	res, err := authz.GetPublicTenantAccessToken(t.ctx, appAccessToken, t.tenantKey)
	if err != nil {
		return
	}
	return res.TenantAccessToken, time.Now().Add(time.Duration(res.Expire) * time.Second), nil
}

// FeishuTenantAccessToken 返回已知最新的 tenant access token, 或如果还没有获得到则返回错误
func (t *PublicAppTenant) FeishuTenantAccessToken() (string, error) {
	return t.tenantAccessTokenUpdator.Get()
}

// Stop 停止更新
func (t *PublicAppTenant) Stop() {
	t.tenantAccessTokenUpdator.Stop()
}
