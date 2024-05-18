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
		return errors_module.ChatDefaultError(err.Error())
	}

	cl.conn = conn
	cl.handleClientConnection()

	return nil
}

func (cl *ChatClient) Clients() map[string]*ChatClient {
	return cl.chats[cl.chatId]
}

func (cl *ChatClient) Conn() *websocket.Conn {
	return cl.conn
}

func (cl *ChatClient) IsSomeoneAlreadyConnected() bool {
	return len(cl.Clients()) > 0
}

func (cl *ChatClient) Disconnect() errors_module.ErrorWithStatus {
	err := cl.Conn().Close()
	if err != nil {
		return errors_module.ChatDefaultError(err.Error())
	}
	delete(cl.Clients(), cl.email)

	return nil
}

func (cl *ChatClient) Broadcast() {
	for {
		msgType, msgBytes, err := cl.Conn().ReadMessage()
		if err != nil {
			fmt.Println("Read error: ", err)
			cl.handleClientDisconnection()
			return
		}

		err = cl.handleNewMessage(msgType, msgBytes)
		if err != nil {
			fmt.Println("New message error: ", err.Error())
			cl.handleClientDisconnection()
			return
		}
	}
}

func (cl *ChatClient) handleClientConnection() {
	if !cl.IsSomeoneAlreadyConnected() {
		return
	}
	msg := ConnectActionToClient{
		Message: fmt.Sprintf("%v connected!", cl.email),
		Type:    MSG_CONNECT_TYPE,
	}

	for _, client := range cl.Clients() {
		if client.email == cl.email {
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
		Message: fmt.Sprintf("%v disconnected!", cl.email),
		Type:    MSG_DISCONNECT_TYPE,
	}

	for _, client := range cl.Clients() {
		err := client.Conn().WriteJSON(msg)
		if err != nil {
			fmt.Println("Disconnect error: ", err)
			return
		}
	}
}

func (cl *ChatClient) handleNewMessage(msgType int, msgBytes []byte) error {
	var msgFromClient MessageFromClient
	err := json.Unmarshal(msgBytes, &msgFromClient)
	if err != nil {
		return err
	}

	// queryErr := cl.chatsQueries.AddMessage(cl.chatId, msgFromClient.Email, msgFromClient.Message)
	// if queryErr != nil {
	// 	return queryErr
	// }

	msgToClient := MessageActionToClient{
		Message: msgFromClient.Message,
		Email:   msgFromClient.Email,
		Type:    MSG_MESSAGE_TYPE,
	}

	for _, client := range cl.Clients() {
		err = client.Conn().WriteJSON(msgToClient)
		if err != nil {
			return err
		}
	}

	return nil
}

var _ interfaces_module.Socket = (*ChatClient)(nil)
