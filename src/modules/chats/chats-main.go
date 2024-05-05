package chats

import (
	"github.com/gorilla/mux"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
)

type ChatsModule struct {
	chatsQueries          *chats_queries.ChatsQueries
	router                *mux.Router
	chats                 map[string]*ChatSocket
	fullDisconnectionChan chan bool
	connectChan           chan ConnectAction
	disconnectChan        chan DisconnectAction
	createChan            chan CreateAction
}

func NewModule(chatsQueries *chats_queries.ChatsQueries, router *mux.Router) *ChatsModule {
	return &ChatsModule{
		chatsQueries:   chatsQueries,
		router:         router,
		chats:          map[string]*ChatSocket{},
		connectChan:    make(chan ConnectAction),
		disconnectChan: make(chan DisconnectAction),
		createChan:     make(chan CreateAction),
	}
}
