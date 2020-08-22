package utils

import (
	"context"
	"net/http"
	"strings"
)

var (
	// DefaultURLBase 是默认的 api 地址前缀
	DefaultURLBase = "https://open.feishu.cn/open-apis"
)

var (
	defaultAPIOptions = &APIOptions{
		URLBase: DefaultURLBase,
		Client:  http.DefaultClient,
	}
)

// APIOptions 是调用 api 时的配置
type APIOptions struct {
	// URLBase 是 api 地址前缀，若空使用 DefaultURLBase.
	// NOTE: URLBase 结尾的 '/' 会被截取掉
	URLBase string

	// Client 是使用的 http client, 若空使用 http.DefaultClient
	Client HTTPClient
}

// HTTPClient 是泛化的 http.Client
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type apiOptionsCtxKey struct{}

func (opts *APIOptions) fillDefault() {
	if opts.URLBase == "" {
		opts.URLBase = DefaultURLBase
	}
	opts.URLBase = strings.TrimRight(opts.URLBase, "/")

	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
}

// WithCtx 将 APIOptions 附着到 context.Context 中并返回一个新的 context.Context
func (opts APIOptions) WithCtx(ctx context.Context) context.Context {
	opts.fillDefault()
	return context.WithValue(ctx, apiOptionsCtxKey{}, &opts)
}

// AsDefault 将 APIOptions 设置成全局默认
func (opts APIOptions) AsDefault() {
	opts.fillDefault()
	defaultAPIOptions = &opts
}

// CtxAPIOptions 从 context.Context 中获得 APIOptions, 若 context.Context 没有，则返回全局默认
func CtxAPIOptions(ctx context.Context) APIOptions {
	v := ctx.Value(apiOptionsCtxKey{})
	if v == nil {
		return DefaultAPIOptions()
	}
	return *(v.(*APIOptions))
}

// DefaultAPIOptions 获得全局默认的 APIOptions
func DefaultAPIOptions() APIOptions {
	return *defaultAPIOptions
}
