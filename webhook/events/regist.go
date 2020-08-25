package events

func Regist(regist func(string, func() interface{})) {
	// app 应用事件
	regist("app_open", func() interface{} { return new(AppOpen) })
	regist("app_status_change", func() interface{} { return new(AppStatusChange) })
	regist("order_paid", func() interface{} { return new(OrderPaid) })
	regist("app_ticket", func() interface{} { return new(AppTicket) })
	regist("app_uninstalled", func() interface{} { return new(AppUninstalled) })
	// contact 通讯录事件
	regist("user_add", func() interface{} { return new(UserAdd) })
	regist("user_update", func() interface{} { return new(UserUpdate) })
	regist("user_leave", func() interface{} { return new(UserLeave) })
	regist("dept_add", func() interface{} { return new(DeptAdd) })
	regist("dept_update", func() interface{} { return new(DeptUpdate) })
	regist("dept_delete", func() interface{} { return new(DeptDelete) })
	regist("user_status_change", func() interface{} { return new(UserStatusChange) })
	regist("contact_scope_change", func() interface{} { return new(ContactScopeChange) })
	// bot and msg 机器人和消息会话事件
	regist("add_bot", func() interface{} { return new(AddBot) })
	regist("remove_bot", func() interface{} { return new(RemoveBot) })
	regist("p2p_chat_create", func() interface{} { return new(P2pChatCreate) })
	regist("message", func() interface{} { return new(Message) })
	regist("message_read", func() interface{} { return new(MessageRead) })
	// unsupported
	regist("", func() interface{} { return new(Unsupported) })
}
