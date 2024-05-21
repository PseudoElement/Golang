package chats

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type ChatsUpdatesSocket struct {
	writer  http.ResponseWriter
	req     *http.Request
	email   string
	conn    *websocket.Conn
	clients map[string]*websocket.Conn
}

type chatsUpdatesSocketInitParams struct {
	writer  http.ResponseWriter
	req     *http.Request
	email   string
	clients map[string]*websocket.Conn
}

func NewChatsUpdatesSocket(p chatsUpdatesSocketInitParams) *ChatsUpdatesSocket {
	return &ChatsUpdatesSocket{
		writer:  p.writer,
		req:     p.req,
		email:   p.email,
		clients: p.clients,
	}
}

func (c *ChatsUpdatesSocket) Connect() errors_module.ErrorWithStatus {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(c.writer, c.req, nil)
	if err != nil {
		return errors_module.SocketConnectionError("ChatsUpdatesSocket")
	}

	c.conn = conn
	c.clients[c.email] = conn

	return nil
}

func (c *ChatsUpdatesSocket) Disconnect() errors_module.ErrorWithStatus {
	err := c.conn.Close()
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}
	delete(c.clients, c.email)

	return nil
}

func (c *ChatsUpdatesSocket) Clients() map[string]*websocket.Conn {
	return c.clients
}

func (c *ChatsUpdatesSocket) Conn() *websocket.Conn {
	return c.conn
}

func (c *ChatsUpdatesSocket) Broadcast() {
	defer c.Disconnect()
	for {
		_, msgBytes, err := c.Conn().ReadMessage()
		if err != nil {
			return
		}

		value, err := c.unmarshalBytes(msgBytes)
		if err != nil {
			return
		}

		for _, client := range c.Clients() {
			err = client.WriteJSON(value)
			if err != nil {
				return
			}
		}
	}
}

func (c *ChatsUpdatesSocket) unmarshalBytes(msgBytes []byte) (NewChatUpdatesAction, error) {
	var obj NewChatUpdatesAction
	if err := json.Unmarshal(msgBytes, &obj); err != nil {
		return obj, errors_module.UnmarshalError("NewChatUpdatesAction")
	}
	return obj, nil
}

var _ interfaces_module.Socket = (*ChatsUpdatesSocket)(nil)
