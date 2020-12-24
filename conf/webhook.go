package conf

// WebhookConfig 是应用事件订阅配置
type WebhookConfig interface {
	// FeishuWebhookVerifToken 返回飞书应用事件订阅的 Verification Token
	FeishuWebhookVerifToken() string

	// FeishuWebhookEncryptKey 返回飞书应用事件订阅的加密密钥
	FeishuWebhookEncryptKey() string
}

// DefaultAppConfig 是默认的事件订阅配置
type DefaultWebhookConfig struct {
	VerifToken string `json:"verifToken"`
	EncryptKey string `json:"encryptKey"`
}

// NewWebhookConfig 创建一个 WebhookConfig
func NewWebhookConfig(verifToken, encryptKey string) WebhookConfig {
	return &DefaultWebhookConfig{
		VerifToken: verifToken,
		EncryptKey: encryptKey,
	}
}

// FeishuWebhookVerifToken 满足 WebhookConfig 接口
func (cnf *DefaultWebhookConfig) FeishuWebhookVerifToken() string {
	return cnf.VerifToken
}

// FeishuWebhookEncryptKey 满足 WebhookConfig 接口
func (cnf *DefaultWebhookConfig) FeishuWebhookEncryptKey() string {
	return cnf.EncryptKey
}
