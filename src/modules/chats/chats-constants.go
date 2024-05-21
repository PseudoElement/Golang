package chats

var (
	ALLOWED_ORIGINS []string = []string{"http://localhost:8080", "http://localhost:4200", "https://websocketking.com", "https://piehost.com"}
)

const (
	MSG_DISCONNECT_TYPE = "disconnect"
	MSG_CONNECT_TYPE    = "connect"
	MSG_MESSAGE_TYPE    = "message"
	MSG_BAN_TYPE        = "ban"
)

const (
	UPD_CHAT_CREATED = "create"
	UPD_CHAT_DELETED = "delete"
)
