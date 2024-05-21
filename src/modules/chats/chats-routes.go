package chats

import (
	"fmt"
	"net/http"
)

func (m *ChatsModule) SetRoutes() {
	m.router.HandleFunc("/chats/create-chat", m._createChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/delete-chat", m._deleteChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/connect-to-chat", m._conectToChatController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/listen-to-updates", m._listenToChatCreationOrDeletionController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/get-messages", m._getMessagesInChatByIdController).Methods(http.MethodGet)
	m.router.HandleFunc("/chats/template", m._htmlTemplateController).Methods(http.MethodGet)

	fmt.Println("ChatsModule started!")
}
