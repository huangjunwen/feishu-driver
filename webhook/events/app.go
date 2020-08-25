package events

// AppOpen 首次开通应用 (只有应用商店应用才能订阅此事件) https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN#%E9%A6%96%E6%AC%A1%E5%BC%80%E9%80%9A%E5%BA%94%E7%94%A8
type AppOpen struct {
	AppId      string `json:"app_id"`
	TenantKey  string `json:"tenant_key"`
	Applicants []struct {
		OpenId string `json:"open_id"`
		UserId string `json:"user_id"`
	} `json:"applicants"`
	Installer struct {
		OpenId string `json:"open_id"`
		UserId string `json:"user_id"`
	} `json:"installer"`
}

// AppStatusChange 应用停启用 (只有应用商店应用才能订阅此事件) https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN#%E5%BA%94%E7%94%A8%E5%81%9C%E5%90%AF%E7%94%A8
type AppStatusChange struct {
	AppId     string `json:"app_id"`
	TenantKey string `json:"tenant_key"`
	Status    string `json:"status"`
}

// OrderPaid 应用商店应用购买 (只有应用商店应用才能订阅此事件) https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN#%E5%BA%94%E7%94%A8%E5%95%86%E5%BA%97%E5%BA%94%E7%94%A8%E8%B4%AD%E4%B9%B0
type OrderPaid struct {
	AppId         string `json:"app_id"`
	TenantKey     string `json:"tenant_key"`
	OrderId       string `json:"order_id"`
	PricePlanId   string `json:"price_plan_id"`
	PricePlanType string `json:"price_plan_type"`
	Seats         int    `json:"seats"`
	BuyCount      int    `json:"buy_count"`
	CreateTime    string `json:"create_time"`
	PayTime       string `json:"pay_time"`
	BuyType       string `json:"buy_type"`
	SrcOrderId    string `json:"src_order_id"`
	OrderPayPrice uint64 `json:"order_pay_price"`
}

// AppTicket app_ticket 事件 ( 应用商店应用自动订阅此事件；企业自建应用不需要此事件) https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN#app_ticket%20%E4%BA%8B%E4%BB%B6
type AppTicket struct {
	AppId     string `json:"app_id"`
	AppTicket string `json:"app_ticket"`
}

// AppUninstalled 应用卸载 (应用商店应用开发者应订阅此事件，并在应用卸载后进行相应的账户注销、数据清理等处理) https://open.feishu.cn/document/ukTMukTMukTM/uQTNxYjL0UTM24CN1EjN#%E5%BA%94%E7%94%A8%E5%8D%B8%E8%BD%BD
type AppUninstalled struct {
	AppId     string `json:"app_id"`
	TenantKey string `json:"tenant_key"`
}
