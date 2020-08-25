package authz

import (
	"context"

	"github.com/huangjunwen/feishu-driver/conf"
	"github.com/huangjunwen/feishu-driver/utils"
)

// AppAccessTokenResult 包含应用维度授权凭证，开放平台可据此识别调用方的应用身份
type AppAccessTokenResult struct {
	utils.APIResultBase

	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

// TenantAccessTokenResult 包含应用的企业授权凭证，开放平台据此识别调用方的应用身份和企业身份
type TenantAccessTokenResult struct {
	utils.APIResultBase

	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// GetInternalAppAccessToken 调接口获得企业自建应用 (internal) 的应用授权凭证,
// 见：https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
func GetInternalAppAccessToken(ctx context.Context, cnf conf.AppConfig) (*AppAccessTokenResult, error) {
	body := &struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppId:     cnf.FeishuAppId(),
		AppSecret: cnf.FeishuAppSecret(),
	}
	result := &AppAccessTokenResult{}
	err := utils.PostJSON(ctx, "/auth/v3/app_access_token/internal", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicAppAccessToken 调接口获得应用商店应用 (public) 的应用授权凭证,
// 见：https://open.feishu.cn/document/ukTMukTMukTM/uEjNz4SM2MjLxYzM
func GetPublicAppAccessToken(ctx context.Context, cnf conf.AppConfig, appTicket string) (*AppAccessTokenResult, error) {
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
	err := utils.PostJSON(ctx, "/auth/v3/app_access_token", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInternalTenantAccessToken 调接口获得企业自建应用 (internal) 的企业授权凭证,
// 见：https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
func GetInternalTenantAccessToken(ctx context.Context, cnf conf.AppConfig) (*TenantAccessTokenResult, error) {
	body := &struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppId:     cnf.FeishuAppId(),
		AppSecret: cnf.FeishuAppSecret(),
	}
	result := &TenantAccessTokenResult{}
	err := utils.PostJSON(ctx, "/auth/v3/tenant_access_token/internal", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicTenantAccessToken 调接口获得应用商店应用 (public) 的企业授权凭证,
// 见：https://open.feishu.cn/document/ukTMukTMukTM/uMjNz4yM2MjLzYzM
func GetPublicTenantAccessToken(ctx context.Context, appAccessToken, tenantKey string) (*TenantAccessTokenResult, error) {
	body := &struct {
		AppAccessToken string `json:"app_access_token"`
		TenantKey      string `json:"tenant_key"`
	}{
		AppAccessToken: appAccessToken,
		TenantKey:      tenantKey,
	}
	result := &TenantAccessTokenResult{}
	err := utils.PostJSON(ctx, "/auth/v3/tenant_access_token", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResendAppTicket 触发重新推送 app ticket: https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM
func ResendAppTicket(ctx context.Context, cnf conf.AppConfig) (*utils.APIResultBase, error) {
	body := &struct {
		AppId     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{
		AppId:     cnf.FeishuAppId(),
		AppSecret: cnf.FeishuAppSecret(),
	}
	result := &utils.APIResultBase{}
	err := utils.PostJSON(ctx, "/auth/v3/app_ticket/resend", body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
