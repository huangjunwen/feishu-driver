package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// PostJSON 使用 POST 方法调用位于 URLBase+urlPath 的接口，body 是请求的 body，result 是响应的 body，
// 两者均用 json 编码/解码; 调用者可使用 APIOptions 附着到 ctx 来调整调用配置
func PostJSON(ctx context.Context, urlPath string, body interface{}, result interface{}) error {
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

	resp, err := opts.Client.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
