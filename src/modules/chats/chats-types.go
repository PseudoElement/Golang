package chats

type NewChatData struct {
	FromEmail string `json:"from_email"`
	ToEmail   string `json:"to_email"`
}

type ChatAction struct {
	ActionType string
	ChatId     string
}
