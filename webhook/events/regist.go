package events

func Regist(regist func(string, func() interface{})) {
	// app
	regist("app_open", func() interface{} { return new(AppOpen) })
	regist("app_status_change", func() interface{} { return new(AppStatusChange) })
	regist("order_paid", func() interface{} { return new(OrderPaid) })
	regist("app_ticket", func() interface{} { return new(AppTicket) })
	regist("app_uninstalled", func() interface{} { return new(AppUninstalled) })
	// bot and msg
	regist("add_bot", func() interface{} { return new(EvAddBot) })
	regist("remove_bot", func() interface{} { return new(EvRemoveBot) })
	regist("p2p_chat_create", func() interface{} { return new(EvP2pChatCreate) })
	regist("message", func() interface{} { return new(EvMessage) })
	regist("message_read", func() interface{} { return new(EvMessageRead) })
	// unsupported
	regist("", func() interface{} { return new(EvUnsupported) })
}
