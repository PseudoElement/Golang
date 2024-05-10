package chats

import (
	"fmt"
	"net/http"
)

func (m *ChatsModule) SetRoutes() {
	m.router.HandleFunc("/chats/create-chat", m._createChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/delete-chat", m._deleteChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/conect-to-chat", m._conectChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/listen-updates", m._listenToUpdatesController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/get-messages", m._getMessagesInChatByIdController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/template", m._htmlTemplateController).Methods(http.MethodGet)

	m.router.HandleFunc("/chats/connect-to-chat-2", m._connectToChat2Controller).Methods(http.MethodGet)

	fmt.Println("ChatsModule started!")
}
