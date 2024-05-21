package chats

import (
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	slice_utils "github.com/pseudoelement/go-server/src/utils/slices"
)

func (m *ChatsModule) isChatExistsByMembers(fromEmail string, toEmail string) bool {
	_, err := m.chatsQueries.GetChatByMembers(fromEmail, toEmail)
	return err == nil
}

func (m *ChatsModule) isAvailableChat(chatId string, email string) bool {
	chat, err := m.chatsQueries.GetChatById(chatId)
	if err != nil {
		return false
	}
	if !slice_utils.Contains(chat.Members, email) {
		return false
	}
	return err == nil
}

func (m *ChatsModule) createNewChat(w http.ResponseWriter, req *http.Request, fromEmail string, toEmail string) (string, errors_module.ErrorWithStatus) {
	if m.isChatExistsByMembers(fromEmail, toEmail) {
		return "", errors_module.ChatAlreadyCreated()
	}

	chatId, err := m.chatsQueries.CreateChat(fromEmail, toEmail)
	if err != nil {
		return "", err
	}

	return chatId, nil
}

func (m *ChatsModule) connectNewClientToChat(w http.ResponseWriter, req *http.Request, chatId string, email string) errors_module.ErrorWithStatus {
	if !m.isAvailableChat(chatId, email) {
		return errors_module.ForbiddenConnectionToChat()
	}

	_, ok := m.chats[chatId]
	if !ok {
		m.chats[chatId] = make(map[string]*ChatClient)
	}

	client := NewChatClient(chatClientInitParams{
		chatsQueries: m.chatsQueries,
		writer:       w,
		req:          req,
		chatId:       chatId,
		email:        email,
		chats:        m.chats,
	})

	if err := client.Connect(); err != nil {
		return err
	}
	go client.Broadcast()

	return nil
}

func (m *ChatsModule) disconnectChatById(email string, chatId string) errors_module.ErrorWithStatus {
	client, ok := m.chats[chatId][email]
	if !ok {
		return errors_module.ChatNotFound()
	}

	err := client.Disconnect()
	if err != nil {
		return err
	}

	err = m.chatsQueries.DeleteMemberFromChat(email, chatId)
	if err != nil {
		return err
	}

	return nil
}

func (m *ChatsModule) listenToChatCreationDeletion(w http.ResponseWriter, req *http.Request, email string) errors_module.ErrorWithStatus {
	updates := NewChatsUpdatesSocket(chatsUpdatesSocketInitParams{
		writer:  w,
		req:     req,
		email:   email,
		clients: m.updatesListeners,
	})

	if err := updates.Connect(); err != nil {
		return err
	}

	go updates.Broadcast()

	return nil
}
