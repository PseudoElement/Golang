package chats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
	errors_module "github.com/pseudoelement/go-server/src/errors"
	"github.com/pseudoelement/go-server/src/utils"
)

type ChatClient struct {
	chatsQueries *chats_queries.ChatsQueries
	writer       http.ResponseWriter
	req          *http.Request
	chatId       string
	email        string
	conn         *websocket.Conn
	chats        map[string]map[string]*ChatClient
}

type chatClientInitParams struct {
	chatsQueries *chats_queries.ChatsQueries
	writer       http.ResponseWriter
	req          *http.Request
	chatId       string
	email        string
	chats        map[string]map[string]*ChatClient
}

func NewChatClient(p chatClientInitParams) *ChatClient {
	p.chats[p.chatId][p.email] = &ChatClient{
		chatsQueries: p.chatsQueries,
		writer:       p.writer,
		req:          p.req,
		chatId:       p.chatId,
		email:        p.email,
		chats:        p.chats,
	}

	return p.chats[p.chatId][p.email]
}

func (cl *ChatClient) Connect() errors_module.ErrorWithStatus {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	upgrader.CheckOrigin = func(req *http.Request) bool {
		origin := req.Header.Get("Origin")
		return utils.Contains(ALLOWED_ORIGINS, origin)
	}

	conn, err := upgrader.Upgrade(cl.writer, cl.req, nil)
	if err != nil {
		return errors_module.SocketConnectionError("ChatClient")
	}

	cl.conn = conn
	cl.handleClientConnection()

	return nil
}

func (cl *ChatClient) Chat() map[string]*ChatClient {
	return cl.chats[cl.chatId]
}

func (cl *ChatClient) Email() string {
	return cl.email
}

func (cl *ChatClient) Conn() *websocket.Conn {
	return cl.conn
}

func (cl *ChatClient) isSomeoneAlreadyConnected() bool {
	return len(cl.Chat()) > 0
}

func (cl *ChatClient) Disconnect() errors_module.ErrorWithStatus {
	err := cl.Conn().Close()
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}
	delete(cl.Chat(), cl.Email())

	return nil
}

func (cl *ChatClient) Broadcast() {
	defer cl.handleClientDisconnection()
	for {
		_, msgBytes, err := cl.Conn().ReadMessage()
		if err != nil {
			fmt.Println("Read error: ", err)
			return
		}

		msg, err := cl.parseMessageFromClient(msgBytes)
		if err != nil {
			return
		}

		switch msg.Type {
		case MSG_MESSAGE_TYPE:
			err = cl.handleNewMessage(msg)
		case MSG_BAN_TYPE:
			err = cl.handleUserBlock(msg)
		}

		if err != nil {
			fmt.Println("Write response failed: ", err.Error())
			return
		}

	}
}

func (cl *ChatClient) handleClientConnection() {
	if !cl.isSomeoneAlreadyConnected() {
		return
	}
	msg := ConnectActionToClient{
		Message: fmt.Sprintf("%v connected!", cl.Email()),
		Type:    MSG_CONNECT_TYPE,
	}

	for _, client := range cl.Chat() {
		if client.Email() == cl.Email() {
			continue
		}
		err := client.Conn().WriteJSON(msg)
		if err != nil {
			fmt.Println("Connect error: ", err)
			return
		}
	}
}

func (cl *ChatClient) handleClientDisconnection() {
	cl.Disconnect()
	msg := DisconnectActionToClient{
		Message: fmt.Sprintf("%v disconnected!", cl.Email()),
		Type:    MSG_DISCONNECT_TYPE,
	}

	for _, client := range cl.Chat() {
		err := client.Conn().WriteJSON(msg)
		if err != nil {
			fmt.Println("Disconnect error: ", err)
			return
		}
	}
}

func (cl *ChatClient) handleNewMessage(msg MessageFromClient) error {
	queryErr := cl.chatsQueries.AddMessage(cl.chatId, msg.Email, msg.Message)
	if queryErr != nil {
		return queryErr
	}

	for _, client := range cl.Chat() {
		err := client.Conn().WriteJSON(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cl *ChatClient) handleUserBlock(msg MessageFromClient) error {
	blockedUser := cl.Chat()[msg.Message]
	blockedUser.Disconnect()
	delete(cl.Chat(), blockedUser.Email())

	message := fmt.Sprintf("User %v has been blocked", msg.Message)

	for _, client := range cl.Chat() {
		err := client.Conn().WriteJSON(MessageFromClient{
			Message: message,
			Email:   msg.Email,
			Type:    MSG_BAN_TYPE,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (cl *ChatClient) parseMessageFromClient(msgBytes []byte) (MessageFromClient, error) {
	var msg MessageFromClient
	err := json.Unmarshal(msgBytes, &msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

var _ interfaces_module.Socket = (*ChatClient)(nil)
