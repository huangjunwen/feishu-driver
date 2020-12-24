package conf

// AppConfig 是应用配置
type AppConfig interface {
	// FeishuAppId 返回飞书应用唯一标识
	FeishuAppId() string

	// FeishuAppSecret 返回飞书应用秘钥
	FeishuAppSecret() string
}

// DefaultAppConfig 是默认应用配置
type DefaultAppConfig struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

// NewAppConfig 创建一个 AppConfig
func NewAppConfig(appId, appSecret string) AppConfig {
	return &DefaultAppConfig{
		AppId:     appId,
		AppSecret: appSecret,
	}
}

// FeishuAppId 满足 AppConfig 接口
func (cnf *DefaultAppConfig) FeishuAppId() string {
	return cnf.AppId
}

// FeishuAppSecret 满足 AppConfig 接口
func (cnf *DefaultAppConfig) FeishuAppSecret() string {
	return cnf.AppSecret
}
