package chats

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	api_main "github.com/pseudoelement/go-server/src/api"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type ChatSocket struct {
	conn           *websocket.Conn
	chatsQueries   *chats_queries.ChatsQueries
	writer         http.ResponseWriter
	req            *http.Request
	chatId         string
	isBroadcasting bool
}

type chatSocketInitParams struct {
	chatsQueries *chats_queries.ChatsQueries
	writer       http.ResponseWriter
	req          *http.Request
	chatId       string
}

func NewChatSocket(p chatSocketInitParams) *ChatSocket {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, e := upgrader.Upgrade(p.writer, p.req, nil)
	if e != nil {
		panic(e)
	}

	return &ChatSocket{
		conn:           conn,
		chatsQueries:   p.chatsQueries,
		writer:         p.writer,
		req:            p.req,
		chatId:         p.chatId,
		isBroadcasting: false,
	}
}

func (s *ChatSocket) Connect() {
}

func (s *ChatSocket) Disconnect() errors_module.ErrorWithStatus {
	err := s.conn.Close()
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}

	return nil
}

func (s *ChatSocket) Broadcast(email string) {
	for {
		s.isBroadcasting = true
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
