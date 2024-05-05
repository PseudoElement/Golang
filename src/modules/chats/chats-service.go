package chats

import (
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func (m *ChatsModule) isChatExistsByMembers(fromEmail string, toEmail string) bool {
	_, err := m.chatsQueries.GetChatByMembers(fromEmail, toEmail)
	return err == nil
}

func (m *ChatsModule) isChatExistsById(chatId string) bool {
	_, err := m.chatsQueries.GetChatById(chatId)
	return err == nil
}

func (m *ChatsModule) createNewChat(w http.ResponseWriter, req *http.Request, fromEmail string, toEmail string) errors_module.ErrorWithStatus {
	if m.isChatExistsByMembers(fromEmail, toEmail) {
		return errors_module.ChatAlreadyCreated()
	}

	chatId, err := m.chatsQueries.CreateChat()
	if err != nil {
		return err
	}

	newChat := NewChatSocket(chatSocketInitParams{
		chatsQueries: m.chatsQueries,
		writer:       w,
		req:          req,
		chatId:       chatId,
	})
	m.chats[chatId] = newChat
	m.createChan <- CreateAction{
		FromEmail: fromEmail,
		ToEmail:   toEmail,
	}

	go newChat.Broadcast(fromEmail)

	return nil
}

func (m *ChatsModule) connectToChatById(w http.ResponseWriter, req *http.Request, chatId string, email string) errors_module.ErrorWithStatus {
	if !m.isChatExistsById(chatId) {
		return errors_module.ChatNotFound()
	}

	newChat := NewChatSocket(chatSocketInitParams{
		chatsQueries: m.chatsQueries,
		writer:       w,
		req:          req,
		chatId:       chatId,
	})
	m.chats[chatId] = newChat
	m.connectChan <- ConnectAction{
		ChatId: chatId,
		Email:  email,
	}
	go newChat.Broadcast(email)

	return nil
}

func (m *ChatsModule) disconnectChatById(email string, chatId string) errors_module.ErrorWithStatus {
	chat, ok := m.chats[chatId]
	if !ok {
		return errors_module.ChatNotFound()
	}

	err := chat.Disconnect()
	if err != nil {
		return err
	}

	err = m.chatsQueries.DeleteMemberFromChat(email, chatId)
	if err != nil {
		return err
	}
	delete(m.chats, chatId)
	m.disconnectChan <- DisconnectAction{
		ChatId: chatId,
		Email:  email,
	}

	return nil
}

func (m *ChatsModule) listenToUpdates(w http.ResponseWriter, req *http.Request, email string) {
	updates := NewChatsUpdatesSocket(chatsUpdatesSocketInitParams{
		writer:         w,
		req:            req,
		connectChan:    m.connectChan,
		disconnectChan: m.disconnectChan,
		createChan:     m.createChan,
	})
	go updates.Broadcast(email)
}
