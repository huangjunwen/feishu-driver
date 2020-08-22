package app

import (
	"context"

	fsdconf "github.com/huangjunwen/feishu-driver/conf"
	fsdutils "github.com/huangjunwen/feishu-driver/utils"
)

// AppAccessTokenResult 包含应用维度授权凭证，开放平台可据此识别调用方的应用身份
type AppAccessTokenResult struct {
	fsdutils.APIResultBase

	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

// TenantAccessTokenResult 包含应用的企业授权凭证，开放平台据此识别调用方的应用身份和企业身份
type TenantAccessTokenResult struct {
	fsdutils.APIResultBase

	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// PublicAppAccessToken 调接口获得应用商店应用 (public) 的应用授权凭证
func PublicAppAccessToken(ctx context.Context, cnf fsdconf.AppConfig, appTicket string) (*AppAccessTokenResult, error) {
	body := &struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
		AppTicket string `json:"app_ticket"`
	}{
		AppId:     cnf.FeishuAppId(),
		AppSecret: cnf.FeishuAppSecret(),
		AppTicket: appTicket,
	}
	result := &AppAccessTokenResult{}
	err := fsdutils.PostJSON(ctx, "/auth/v3/app_access_token", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InternalTenantAccessToken 调接口获得企业自建应用 (internal) 的企业授权凭证
func InternalTenantAccessToken(ctx context.Context, cnf fsdconf.AppConfig) (*TenantAccessTokenResult, error) {
	body := &struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppId:     cnf.FeishuAppId(),
		AppSecret: cnf.FeishuAppSecret(),
	}
	result := &TenantAccessTokenResult{}
	err := fsdutils.PostJSON(ctx, "/auth/v3/tenant_access_token/internal", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PublicTenantAccessToken 调接口获得应用商店应用 (public) 的企业授权凭证
func PublicTenantAccessToken(ctx context.Context, appAccessToken, tenantKey string) (*TenantAccessTokenResult, error) {
	body := &struct {
		AppAccessToken string `json:"app_access_token"`
		TenantKey      string `json:"tenant_key"`
	}{
		AppAccessToken: appAccessToken,
		TenantKey:      tenantKey,
	}
	result := &TenantAccessTokenResult{}
	err := fsdutils.PostJSON(ctx, "/auth/v3/tenant_access_token", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
