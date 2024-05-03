package chats

import (
	"github.com/gorilla/mux"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
)

type ChatsModule struct {
	chatQueries *chats_queries.ChatsQueries
	router      *mux.Router
}

func NewModule(chatsQueries *chats_queries.ChatsQueries, router *mux.Router) *ChatsModule {
	return &ChatsModule{chatQueries: chatsQueries, router: router}
}
