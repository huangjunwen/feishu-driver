package conf

// AppConfig 是应用配置
type AppConfig interface {
	// FeishuAppId 返回飞书应用唯一标识
	FeishuAppId() string

	// FeishuAppSecret 返回飞书应用秘钥
	FeishuAppSecret() string
}

type defaultAppConfig struct {
	appId     string
	appSecret string
}

// NewAppConfig 创建一个 AppConfig
func NewAppConfig(appId, appSecret string) AppConfig {
	return &defaultAppConfig{
		appId:     appId,
		appSecret: appSecret,
	}
}

func (cnf *defaultAppConfig) FeishuAppId() string {
	return cnf.appId
}

func (cnf *defaultAppConfig) FeishuAppSecret() string {
	return cnf.appSecret
}
