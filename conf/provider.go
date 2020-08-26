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

// AppAccessTokenProviderFunc 是函数形式的 AppAccessTokenProvider
type AppAccessTokenProviderFunc func() (string, error)

// TenantAccessTokenProviderFunc 是函数形式的 TenantAccessTokenProvider
type TenantAccessTokenProviderFunc func() (string, error)

// AppTicketProviderFunc 是函数形式的 AppTicketProvider
type AppTicketProviderFunc func() (string, error)

var (
	_ AppAccessTokenProvider    = (AppAccessTokenProviderFunc)(nil)
	_ TenantAccessTokenProvider = (TenantAccessTokenProviderFunc)(nil)
	_ AppTicketProvider         = (AppTicketProviderFunc)(nil)
)

func (f AppAccessTokenProviderFunc) FeishuAppAccessToken() (string, error)       { return f() }
func (f TenantAccessTokenProviderFunc) FeishuTenantAccessToken() (string, error) { return f() }
func (f AppTicketProviderFunc) FeishuAppTicket() (string, error)                 { return f() }
