package utils

import (
	"fmt"
)

// APIResultBase 是 api 返回结果基础字段
type APIResultBase struct {
	// Code 是错码，非 0 表示错误
	Code int `json:"code"`

	// Msg 是错误描述
	Msg string `json:"msg"`
}

// APIError 是 API 错误
type APIError APIResultBase

// ResultError 在 Code 非 0 的时候返回非 nil error
func (result *APIResultBase) ResultError() error {
	if result.Code == 0 {
		return nil
	}
	return (*APIError)(result)
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
