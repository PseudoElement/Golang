package chats

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	api_main "github.com/pseudoelement/go-server/src/api"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
	types_module "github.com/pseudoelement/go-server/src/common/types"
	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type ChatsUpdatesSocket struct {
	conn           *websocket.Conn
	connectChan    chan ConnectAction
	disconnectChan chan DisconnectAction
	createChan     chan CreateAction
}

type chatsUpdatesSocketInitParams struct {
	writer         http.ResponseWriter
	req            *http.Request
	connectChan    chan ConnectAction
	disconnectChan chan DisconnectAction
	createChan     chan CreateAction
}

func NewChatsUpdatesSocket(p chatsUpdatesSocketInitParams) *ChatsUpdatesSocket {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, e := upgrader.Upgrade(p.writer, p.req, nil)
	if e != nil {
		panic(e)
	}

	return &ChatsUpdatesSocket{
		conn:           conn,
		connectChan:    p.connectChan,
		disconnectChan: p.disconnectChan,
		createChan:     p.createChan,
	}
}

func (s *ChatsUpdatesSocket) Disconnect() errors_module.ErrorWithStatus {
	err := s.conn.Close()
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}

	return nil
}

func (s *ChatsUpdatesSocket) Connect() {}

func (s *ChatsUpdatesSocket) Broadcast(email string) {
	for {
		select {
		case connect := <-s.connectChan:
			msg := types_module.MessageToClient{
				Message: fmt.Sprintf("User %v is connected!", connect.Email),
			}
			s.conn.WriteMessage(websocket.BinaryMessage, api_main.SuccessBytesResponse(msg))
		case disconnect := <-s.disconnectChan:
			msg := types_module.MessageToClient{
				Message: fmt.Sprintf("User %v is disconnected!", disconnect.Email),
			}
			s.conn.WriteMessage(websocket.BinaryMessage, api_main.SuccessBytesResponse(msg))
		case create := <-s.createChan:
			if create.FromEmail != email || create.ToEmail != email {
				continue
			}
			msg := types_module.MessageToClient{
				Message: fmt.Sprintf("Chat created between %v and %v.", create.FromEmail, create.ToEmail),
			}
			err := s.conn.WriteJSON(msg)
			if err != nil {
				log.Fatal("WriteJSON error - ", err)
				break
			}
		}
	}
}

var _ interfaces_module.Socket = (*ChatsUpdatesSocket)(nil)
