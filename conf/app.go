package conf

// AppConfig 是应用配置
type AppConfig interface {
	// FeishuAppId 返回飞书应用唯一标识
	FeishuAppId() string

	// FeishuAppSecret 返回飞书应用秘钥
	FeishuAppSecret() string
}
