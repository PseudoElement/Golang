package chats

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	api_main "github.com/pseudoelement/go-server/src/api"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
)

type ChatSocket struct {
	conn         *websocket.Conn
	chatsQueries *chats_queries.ChatsQueries
	w            http.ResponseWriter
	req          *http.Request
	chatId       string
}

type chatSocketInitParams struct {
	chatsQueries *chats_queries.ChatsQueries
	w            http.ResponseWriter
	req          *http.Request
	chatId       string
}

func NewChatSocket(p chatSocketInitParams) *ChatSocket {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, e := upgrader.Upgrade(p.w, p.req, nil)
	if e != nil {
		panic(e)
	}

	return &ChatSocket{
		conn:         conn,
		chatsQueries: p.chatsQueries,
		w:            p.w,
		req:          p.req,
		chatId:       p.chatId,
	}
}

func (s *ChatSocket) Connect() {
	s.Broadcast()
}

func (s *ChatSocket) Disconnect() {
	s.conn.Close()
}

func (s *ChatSocket) Broadcast() {
	for {
		messageType, msgBytes, err := s.conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		var msgStruct chats_queries.MessageFromDB
		err = json.Unmarshal(msgBytes, &msgStruct)
		if err != nil {
			s.conn.WriteMessage(messageType, api_main.FailBytesResponse(err.Error()))
			continue
		}

		queryErr := s.chatsQueries.AddMessage(s.chatId, msgStruct.FromEmail, msgStruct.Message)
		if queryErr != nil {
			s.conn.WriteMessage(messageType, api_main.FailBytesResponse(queryErr.Error()))
			continue
		}

		s.conn.WriteMessage(messageType, msgBytes)
	}
}

var _ interfaces_module.Socket = (*ChatSocket)(nil)
