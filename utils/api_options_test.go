package utils

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIOptions(t *testing.T) {
	assert := assert.New(t)

	newURLBase := "https://close.feishu.cn/close-apis"
	newClient := &http.Client{}

	// 将 opts 设置为全局默认后修改应该不影响设置过去的值
	opts := &APIOptions{
		URLBase: newURLBase + "/", // 测试如果结尾有 "/"
	}
	opts.AsDefault()
	opts.URLBase = DefaultURLBase

	// 全局
	{
		opts2 := DefaultAPIOptions()

		assert.Equal(newURLBase, opts2.URLBase)
		assert.Equal(http.DefaultClient, opts2.Client) // 没有设置，应该是默认的
	}

	// ctx 没有附着的情况，应该使用全局的
	{
		opts2 := CtxAPIOptions(context.Background())

		assert.Equal(newURLBase, opts2.URLBase)
		assert.Equal(http.DefaultClient, opts2.Client)
	}

	// ctx 有附着 APIOptions 的情况
	{
		ctx := APIOptions{
			Client: newClient,
		}.WithCtx(context.Background())
		opts2 := CtxAPIOptions(ctx)

		assert.Equal(DefaultURLBase, opts2.URLBase)
		assert.Equal(newClient, opts2.Client)
	}
}
