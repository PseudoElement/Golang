package chats

import (
	"fmt"
	"net/http"
)

func (m *ChatsModule) SetRoutes() {
	m.router.HandleFunc("/chats/create-chat", m._createChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/delete-chat", m._deleteChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/listen-chat", m._listenChatsController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/disconnect-chats", m._disconnectAllChats).Methods(http.MethodGet)

	fmt.Println("ChatsModule started!")
}
