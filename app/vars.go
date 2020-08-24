package app

import (
	"time"
)

var (
	// DefaultRetryInterval 时默认的错误重试间隔
	DefaultRetryInterval = 10 * time.Second

	// DefaultUpdateInterval 时默认的更新间隔
	DefaultUpdateInterval = 5 * time.Minute
)
