package chats

type NewChatData struct {
	FromEmail string `json:"from_email"`
	ToEmail   string `json:"to_email"`
}

type CreateAction struct {
	FromEmail string
	ToEmail   string
}

type ChatCreatedMessage struct {
	ChatId string `json:"chatId"`
}

type MessageFromClient struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	Type    string `json:"type"`
}

type MessageActionToClient struct {
	Message string `json:"message"`
	Email   string `json:"email"`
	Type    string `json:"type"`
}

type DisconnectActionToClient struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type ConnectActionToClient struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type NewChatUpdatesAction struct {
	Members []string `json:"members"`
	ChatId  string   `json:"chatId"`
}
