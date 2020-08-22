package app

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	fsdconf "github.com/huangjunwen/feishu-driver/conf"
	fsdutils "github.com/huangjunwen/feishu-driver/utils"
)

type AppAccessTokenResult struct {
	fsdutils.APIResultBase

	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

type TenantAccessTokenResult struct {
	fsdutils.APIResultBase

	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

func PublicAppAccessToken(ctx context.Context, cnf fsdconf.AppConfig, appTicket string) (*AppAccessTokenResult, error) {
	opts := fsdutils.CtxAPIOptions(ctx)

	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(map[string]interface{}{
		"app_id":     cnf.FeishuAppId(),
		"app_secret": cnf.FeishuAppSecret(),
		"app_ticket": appTicket,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		opts.URLBase+"/auth/v3/app_access_token",
		body,
	)
	if err != nil {
		return nil, err
	}

	resp, err := opts.Client.Do(req)
	if err != nil {
		return nil, err
	}

	res := &AppAccessTokenResult{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}

func InternalTenantAccessToken(ctx context.Context, cnf fsdconf.AppConfig) (*TenantAccessTokenResult, error) {
	opts := fsdutils.CtxAPIOptions(ctx)

	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(map[string]interface{}{
		"app_id":     cnf.FeishuAppId(),
		"app_secret": cnf.FeishuAppSecret(),
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		opts.URLBase+"/auth/v3/tenant_access_token/internal",
		body,
	)
	if err != nil {
		return nil, err
	}

	resp, err := opts.Client.Do(req)
	if err != nil {
		return nil, err
	}

	res := &TenantAccessTokenResult{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}

func PublicTenantAccessToken(ctx context.Context, appAccessToken, tenantKey string) (*TenantAccessTokenResult, error) {
	opts := fsdutils.CtxAPIOptions(ctx)

	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(map[string]interface{}{
		"app_access_token": appAccessToken,
		"tenant_key":       tenantKey,
	}); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		opts.URLBase+"/auth/v3/tenant_access_token",
		body,
	)
	if err != nil {
		return nil, err
	}

	resp, err := opts.Client.Do(req)
	if err != nil {
		return nil, err
	}

	res := &TenantAccessTokenResult{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
