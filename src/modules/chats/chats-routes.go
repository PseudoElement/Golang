package chats

import (
	"fmt"
	"net/http"
)

func (m *ChatsModule) SetRoutes() {
	m.router.HandleFunc("/chats/create-chat", m._createChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/listen-chat", m._listenChatController).Methods(http.MethodGet)

	fmt.Println("ChatsModule started!")
}
