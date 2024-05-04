package chats

import (
	"net/http"

	"github.com/gorilla/mux"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
	errors_module "github.com/pseudoelement/go-server/src/errors"
	"github.com/pseudoelement/go-server/src/utils"
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

func (m *ChatsModule) CreateChat(w http.ResponseWriter, req *http.Request, chatId string) {
	newChat := NewChatSocket(chatSocketInitParams{
		chatsQueries: m.chatsQueries,
		w:            w,
		req:          req,
		chatId:       chatId,
	})
	m.chats = append(m.chats, newChat)
	go newChat.Broadcast()
}

func (m *ChatsModule) DeleteChat(chatId string) errors_module.ErrorWithStatus {
	found, _ := utils.Find(m.chats, func(_chat *ChatSocket) bool {
		return _chat.chatId == chatId
	})
	chat, ok := found.(*ChatSocket)
	if !ok {
		return errors_module.ChatNotFound()
	}

	err := chat.Disconnect()
	if err != nil {
		return err
	}

	m.chats = utils.Filter(m.chats, func(_chat *ChatSocket, i int) bool {
		return _chat.chatId != chatId
	})

	return nil
}
