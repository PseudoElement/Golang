package chats

import "github.com/gorilla/websocket"

type ChatClient struct {
	email string
	conn  *websocket.Conn
}

type NewChatData struct {
	FromEmail string `json:"from_email"`
	ToEmail   string `json:"to_email"`
}

type DisconnectAction struct {
	ChatId string
	Email  string
}

type ConnectAction struct {
	ChatId string
	Email  string
}

type CreateAction struct {
	FromEmail string
	ToEmail   string
}
