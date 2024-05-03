package chats

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	api_main "github.com/pseudoelement/go-server/src/api"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
)

type ChatSocket struct {
	conn         *websocket.Conn
	chatsQueries *chats_queries.ChatsQueries
}

func NewChatSocket(w http.ResponseWriter, req *http.Request, chatsQueries *chats_queries.ChatsQueries) *ChatSocket {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, e := upgrader.Upgrade(w, req, nil)
	if e != nil {
		panic(e)
	}

	return &ChatSocket{
		conn:         conn,
		chatsQueries: chatsQueries,
	}
}

func (s *ChatSocket) Connect() {
	s.Listen()
}

func (s *ChatSocket) Disconnect() {
	s.conn.Close()
}

func (s *ChatSocket) Listen() {
	for {
		messageType, msgBytes, err := s.conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		var messageStruct chats_queries.MessageFromDB
		err = json.Unmarshal(msgBytes, &messageStruct)
		if err != nil {
			s.conn.WriteMessage(messageType, api_main.FailBytesResponse(err.Error()))
		} else {
			s.conn.WriteMessage(messageType, api_main.SuccessBytesResponse(messageStruct))
		}
	}
}
