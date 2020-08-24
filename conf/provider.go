package conf

// AppAccessTokenProvider 提供 app access token
type AppAccessTokenProvider interface {
	// FeishuAppAccessToken 返回 app access token 或错误
	FeishuAppAccessToken() (string, error)
}

// TenantAccessTokenProvider 提供 tenant access token
type TenantAccessTokenProvider interface {
	// FeishuTenantAccessToken 返回 tenant access token 或错误
	FeishuTenantAccessToken() (string, error)
}

// AppTicketProvider 提供 app ticket
type AppTicketProvider interface {
	// FeishuAppTicket 返回 app ticket 或错误
	FeishuAppTicket() (string, error)
}
