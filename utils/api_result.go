package utils

// APIResultBase 是 api 返回结果基础字段
type APIResultBase struct {
	// Code 是错码，非 0 表示错误
	Code int `json:"code"`

	// Msg 是错误描述
	Msg string `json:"msg"`
}
