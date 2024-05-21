package chats

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
)

type ChatsModule struct {
	chatsQueries     *chats_queries.ChatsQueries
	router           *mux.Router
	chats            map[string]map[string]*ChatClient
	updatesListeners map[string]*websocket.Conn
}

func NewModule(chatsQueries *chats_queries.ChatsQueries, router *mux.Router) *ChatsModule {
	return &ChatsModule{
		chatsQueries:     chatsQueries,
		router:           router,
		chats:            make(map[string]map[string]*ChatClient),
		updatesListeners: make(map[string]*websocket.Conn),
	}
}
