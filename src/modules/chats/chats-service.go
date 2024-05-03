package chats

import (
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func (m *ChatsModule) handleChatCreation(fromEmail string, toEmail string) errors_module.ErrorWithStatus {
	if m.isChatExists(fromEmail, toEmail) {
		return errors_module.ChatAlreadyCreated()
	}
	_, err := m.chatsQueries.CreateChat()
	if err != nil {
		return err
	}

	return nil
}

func (m *ChatsModule) handleChatListening(w http.ResponseWriter, req *http.Request, fromEmail string) errors_module.ErrorWithStatus {
	chats, err := m.chatsQueries.GetAllChatsOfUser(fromEmail)
	if err != nil {
		return err
	}

	for _, chat := range chats {
		socket := NewChatSocket(chatSocketInitParams{
			chatsQueries: m.chatsQueries,
			w:            w,
			req:          req,
			chatId:       chat.Id,
		})
		m.AddChat(socket)
		go socket.Broadcast()
	}

	return nil
}

func (m *ChatsModule) isChatExists(fromEmail string, toEmail string) bool {
	_, err := m.chatsQueries.GetChatByMembers(fromEmail, toEmail)
	return err == nil
}
