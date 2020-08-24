package conf

// WebhookConfig 是应用事件订阅配置
type WebhookConfig interface {
	// FeishuWebhookVerifToken 返回飞书应用事件订阅的 Verification Token
	FeishuWebhookVerifToken() string

	// FeishuWebhookEncryptKey 返回飞书应用事件订阅的加密密钥
	FeishuWebhookEncryptKey() string
}

type defaultWebhookConfig struct {
	verifToken string
	encryptKey string
}

// NewWebhookConfig 创建一个 WebhookConfig
func NewWebhookConfig(verifToken, encryptKey string) WebhookConfig {
	return &defaultWebhookConfig{
		verifToken: verifToken,
		encryptKey: encryptKey,
	}
}

func (cnf *defaultWebhookConfig) FeishuWebhookVerifToken() string {
	return cnf.verifToken
}

func (cnf *defaultWebhookConfig) FeishuWebhookEncryptKey() string {
	return cnf.encryptKey
}
