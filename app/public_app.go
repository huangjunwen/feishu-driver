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
	_ conf.AppAccessTokenProvider = (*PublicApp)(nil)
)

// PublicApp 是应用商店应用 (对外公开故叫 public), 它会提供其最新获得 app access token,
// 并定期检查，如果接近过期或已过期则更新之.
//
// 满足 AppAccessTokenProvider 接口，
// 使用者也可通过 PAOnUpdateAppAccessToken 回调选项获得 app access token
type PublicApp struct {
	appConfig             conf.AppConfig
	ticketProvider        conf.AppTicketProvider
	ctx                   context.Context
	appAccessTokenUpdator tokenUpdator
}

// NewPublicApp 创建 PublicApp 并开启自动更新
func NewPublicApp(cnf conf.AppConfig, ticketProvider conf.AppTicketProvider, opts ...PublicAppOption) (*PublicApp, error) {
	a := &PublicApp{
		appConfig:      cnf,
		ticketProvider: ticketProvider,
		ctx:            context.Background(),
		appAccessTokenUpdator: tokenUpdator{
			name:           fmt.Sprintf("PA-%s-app", cnf.FeishuAppId()),
			onUpdate:       func(string) error { return nil },
			updateInterval: DefaultUpdateInterval,
			retryInterval:  DefaultRetryInterval,
			logger:         logr.Nop,
		},
	}
	a.appAccessTokenUpdator.tokenGetter = a.appAccessTokenGetter
	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}

	a.appAccessTokenUpdator.Start()
	return a, nil
}

func (a *PublicApp) appAccessTokenGetter() (token string, expire time.Time, err error) {
	ticket, err := a.ticketProvider.FeishuAppTicket()
	if err != nil {
		return
	}
	res, err := authz.GetPublicAppAccessToken(a.ctx, a.appConfig, ticket)
	if err != nil {
		return
	}
	err = res.ResultError()
	if err != nil {
		return
	}
	return res.AppAccessToken, time.Now().Add(time.Duration(res.Expire) * time.Second), nil
}

// FeishuAppAccessToken 返回已知最新的 app access token, 或如果还没有获得到则返回错误
func (a *PublicApp) FeishuAppAccessToken() (string, error) {
	return a.appAccessTokenUpdator.Get()

}

// Stop 停止更新
func (a *PublicApp) Stop() {
	a.appAccessTokenUpdator.Stop()
}
