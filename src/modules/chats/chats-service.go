package chats

import (
	"fmt"
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func (m *ChatsModule) handleChatCreation(fromEmail string, toEmail string) errors_module.ErrorWithStatus {
	if m.isChatExists(fromEmail, toEmail) {
		return errors_module.ChatAlreadyCreated()
	}
	chatId, err := m.chatsQueries.CreateChat()
	if err != nil {
		return err
	}

	m.actionChan <- ChatAction{
		ActionType: "Creation",
		ChatId:     chatId,
	}

	return nil
}

func (m *ChatsModule) handleChatDeletion(chatId string) errors_module.ErrorWithStatus {
	err := m.chatsQueries.DeleteChatById(chatId)
	if err != nil {
		return err
	}

	m.actionChan <- ChatAction{
		ActionType: "Deletion",
		ChatId:     chatId,
	}

	return nil
}

func (m *ChatsModule) handleChatListening(w http.ResponseWriter, req *http.Request, fromEmail string) errors_module.ErrorWithStatus {
	chats, err := m.chatsQueries.GetAllChatsOfUser(fromEmail)
	if err != nil {
		return err
	}

	for _, chat := range chats {
		m.CreateChat(w, req, chat.Id)
	}

	return nil
}

func (m *ChatsModule) handleChannelActions() {
	for {
		select {
		case action := <-m.actionChan:
			fmt.Println("INCOMING_ACTION - ", action)
		}
	}
}

func (m *ChatsModule) isChatExists(fromEmail string, toEmail string) bool {
	_, err := m.chatsQueries.GetChatByMembers(fromEmail, toEmail)
	return err == nil
}
