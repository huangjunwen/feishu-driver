package events

// AddBot 机器人进群 https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN#%E6%9C%BA%E5%99%A8%E4%BA%BA%E8%BF%9B%E7%BE%A4
type AddBot struct {
	AppId               string            `json:"app_id"`
	TenantKey           string            `json:"tenant_key"`
	ChatI18nNames       map[string]string `json:"chat_i18n_names"`
	ChatName            string            `json:"chat_name"`
	ChatOwnerEmployeeId string            `json:"chat_owner_employee_id`
	ChatOwnerName       string            `json:"chat_owner_name"`
	ChatOwnerOpenId     string            `json:"chat_owner_open_id"`
	OpenChatId          string            `json:"open_chat_id"`
	OperatorEmployeeId  string            `json:"operator_employee_id"`
	OperatorName        string            `json:"operator_name"`
	OperatorOpenId      string            `json:"operator_open_id"`
	OwnerIsBot          bool              `json:"owner_is_bot"`
}

// RemoveBot 机器人被移出群 https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN#%E6%9C%BA%E5%99%A8%E4%BA%BA%E8%A2%AB%E7%A7%BB%E5%87%BA%E7%BE%A4
type RemoveBot struct {
	AppId               string            `json:"app_id"`
	TenantKey           string            `json:"tenant_key"`
	ChatI18nNames       map[string]string `json:"chat_i18n_names"`
	ChatName            string            `json:"chat_name"`
	ChatOwnerEmployeeId string            `json:"chat_owner_employee_id`
	ChatOwnerName       string            `json:"chat_owner_name"`
	ChatOwnerOpenId     string            `json:"chat_owner_open_id"`
	OpenChatId          string            `json:"open_chat_id"`
	OperatorEmployeeId  string            `json:"operator_employee_id"`
	OperatorName        string            `json:"operator_name"`
	OperatorOpenId      string            `json:"operator_open_id"`
	OwnerIsBot          bool              `json:"owner_is_bot"`
}

// P2pChatCreate 用户和机器人的会话首次被创建 https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN#%E7%94%A8%E6%88%B7%E5%92%8C%E6%9C%BA%E5%99%A8%E4%BA%BA%E7%9A%84%E4%BC%9A%E8%AF%9D%E9%A6%96%E6%AC%A1%E8%A2%AB%E5%88%9B%E5%BB%BA
type P2pChatCreate struct {
	AppId     string `json:"app_id"`
	TenantKey string `json:"tenant_key"`
	ChatId    string `json:"chat_id"`
	Operator  struct {
		OpenId string `json:"open_id"`
		UserId string `json:"user_id"`
	} `json:"operator"`
	User struct {
		Name   string `json:"name"`
		OpenId string `json:"open_id"`
		UserId string `json:"user_id"`
	} `json:"user"`
}

// Message 接收消息 https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN#%E6%8E%A5%E6%94%B6%E6%B6%88%E6%81%AF
type Message struct {
	AppId            string   `json:"app_id"`
	TenantKey        string   `json:"tenant_key"`
	RootId           string   `json:root_id`
	ParentId         string   `json:"parent_id"`
	OpenChatId       string   `json:"open_chat_id"`
	ChatType         string   `json:"chat_type"`
	MsgType          string   `json:"msg_type"`
	OpenId           string   `json:"open_id"`
	OpenMessageId    string   `json:"open_message_id"`
	IsMention        bool     `json:"is_mention"`
	Text             string   `json:"text"`
	TextWithoutAtBot string   `json:"text_without_at_bot"`
	Title            string   `json:"title"`
	ImageKeys        []string `json:"image_keys"`
	ImageHeight      string   `json:""image_height""`
	ImageWidth       string   `json:"image_width"`
	ImageKey         string   `json:"image_key"`
	FileKey          string   `json:"file_key"`
}

// MessageRead 消息已读 https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN#%E6%B6%88%E6%81%AF%E5%B7%B2%E8%AF%BB
type MessageRead struct {
	AppId          string   `json:"app_id"`
	TenantKey      string   `json:"tenant_key"`
	OpenChatId     string   `json:"open_chat_id"`
	OpenId         string   `json:"open_id"`
	OpenMessageIds []string `json:"open_message_ids"`
}
