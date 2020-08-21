package events

func Regist(regist func(string, func() interface{})) {
	// bot and msg
	regist("add_bot", func() interface{} { return new(EvAddBot) })
	regist("remove_bot", func() interface{} { return new(EvRemoveBot) })
	regist("p2p_chat_create", func() interface{} { return new(EvP2pChatCreate) })
	regist("message", func() interface{} { return new(EvMessage) })
	regist("message_read", func() interface{} { return new(EvMessageRead) })
	// unsupported
	regist("", func() interface{} { return new(EvUnsupported) })
}
