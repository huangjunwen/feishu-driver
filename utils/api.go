package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/huangjunwen/feishu-driver/conf"
)

// PostJSON 使用 POST 方法调用位于 URLBase+urlPath 的接口，body 是请求的 body，result 是响应的 body，
// 两者均用 json 编码/解码; 调用者可使用 APIOptions 附着到 ctx 来调整调用配置
func PostJSON(ctx context.Context, urlPath string, body interface{}, result interface{}) error {
	return postJSON(ctx, urlPath, nil, body, result)
}

// PostJSONWithAppAccessToken 类似于 PostJSON，不过 Authorization 头部会添加 app access token
func PostJSONWithAppAccessToken(ctx context.Context, urlPath string, provider conf.AppAccessTokenProvider, body interface{}, result interface{}) error {
	token, err := provider.FeishuAppAccessToken()
	if err != nil {
		return err
	}
	return postJSON(ctx, urlPath, func(req *http.Request) *http.Request {
		req.Header.Add("Authorization", "Bearer "+token)
		return req
	}, body, result)
}

// PostJSONWithTenantAccessToken 类似于 PostJSON，不过 Authorization 头部会添加 tenant access token
func PostJSONWithTenantAccessToken(ctx context.Context, urlPath string, provider conf.TenantAccessTokenProvider, body interface{}, result interface{}) error {
	token, err := provider.FeishuTenantAccessToken()
	if err != nil {
		return err
	}
	return postJSON(ctx, urlPath, func(req *http.Request) *http.Request {
		req.Header.Add("Authorization", "Bearer "+token)
		return req
	}, body, result)
}

func postJSON(ctx context.Context, urlPath string, reqModify func(*http.Request) *http.Request, body interface{}, result interface{}) error {
	opts := CtxAPIOptions(ctx)

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		opts.URLBase+urlPath,
		buf,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if reqModify != nil {
		req = reqModify(req)
	}

	resp, err := opts.Client.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
