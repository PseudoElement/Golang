package chats

import (
	"github.com/gorilla/mux"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
)

type ChatsModule struct {
	chatsQueries *chats_queries.ChatsQueries
	router       *mux.Router
	chats        []*ChatSocket
	actionChan   chan ChatAction
}

func NewModule(chatsQueries *chats_queries.ChatsQueries, router *mux.Router) *ChatsModule {
	return &ChatsModule{
		chatsQueries: chatsQueries,
		router:       router,
		chats:        []*ChatSocket{},
		actionChan:   make(chan ChatAction),
	}
}

func (m *ChatsModule) AddChat(chat *ChatSocket) {
	m.chats = append(m.chats, chat)

}
