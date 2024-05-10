package chats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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
	return &ChatSocket{
		chatsQueries:   p.chatsQueries,
		writer:         p.writer,
		req:            p.req,
		chatId:         p.chatId,
		isBroadcasting: false,
	}
}

func (s *ChatSocket) Connect() errors_module.ErrorWithStatus {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(s.writer, s.req, nil)
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}

	s.conn = conn

	return nil
}

func (s *ChatSocket) Disconnect() errors_module.ErrorWithStatus {
	err := s.conn.Close()
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}

	return nil
}

func (s *ChatSocket) Broadcast(email string) {
	defer s.conn.Close()
	go func() {
		for {
			s.isBroadcasting = true
			messageType, msgBytes, err := s.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					break
				}
				fmt.Printf("Broadcast ReadMessage error: %v\n", err)
				return
			}

			var msgStruct chats_queries.MessageFromDB
			err = json.Unmarshal(msgBytes, &msgStruct)
			if err != nil {
				// s.conn.WriteMessage(messageType, api_main.FailBytesResponse(err.Error()))
				fmt.Printf("Broadcast Unmarshal error - %v", err.Error())
				return
			}

			queryErr := s.chatsQueries.AddMessage(s.chatId, msgStruct.FromEmail, msgStruct.Message)
			if queryErr != nil {
				// s.conn.WriteMessage(messageType, api_main.FailBytesResponse(queryErr.Error()))
				fmt.Printf("Broadcast QueryErr error - %v", queryErr)
				return
			}

			err = s.conn.WriteMessage(messageType, msgBytes)
			if err != nil {
				fmt.Println("Broadcast WriteMessage error - %v", err)
				return
			}
		}
	}()
}

var _ interfaces_module.Socket = (*ChatSocket)(nil)
