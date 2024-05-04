package chats

import (
	"fmt"
	"net/http"
	"sync"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	"github.com/pseudoelement/go-server/src/utils"
)

func (m *ChatsModule) handleChatCreation(w http.ResponseWriter, req *http.Request, fromEmail string, toEmail string) errors_module.ErrorWithStatus {
	if m.isChatExists(fromEmail, toEmail) {
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
	m.chats = append(m.chats, newChat)
	go newChat.Broadcast()

	m.actionChan <- ChatAction{
		ActionType: "create",
		ChatId:     chatId,
	}

	return nil
}

func (m *ChatsModule) handleChatDeletion(chatId string) errors_module.ErrorWithStatus {
	err := m.chatsQueries.DeleteChatById(chatId)
	if err != nil {
		return err
	}

	found, _ := utils.Find(m.chats, func(_chat ChatSocket) bool {
		return _chat.chatId == chatId
	})
	chat, ok := found.(ChatSocket)
	if !ok {
		return errors_module.ChatNotFound()
	}

	err = chat.Disconnect()
	if err != nil {
		return err
	}

	m.chats = utils.Filter(m.chats, func(_chat ChatSocket, i int) bool {
		return _chat.chatId != chatId
	})

	m.actionChan <- ChatAction{
		ActionType: "delete",
		ChatId:     chatId,
	}

	return nil
}

func (m *ChatsModule) initChatListening(w http.ResponseWriter, req *http.Request, email string) errors_module.ErrorWithStatus {
	chats, err := m.chatsQueries.GetAllChatsOfUser(email)
	if err != nil {
		return err
	}

	for _, chat := range chats {
		fromEmail := chat.Members[0]
		toEmail := chat.Members[1]
		m.handleChatCreation(w, req, fromEmail, toEmail)
	}

	return nil
}

func (m *ChatsModule) listenAllChatsOfUser() {
	for _, chat := range m.chats {
		if !chat.isBroadcasting {
			go chat.Broadcast()
		}
	}
}

func (m *ChatsModule) subscribeOnChatsUpdates() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-m.actionChan:
				m.listenAllChatsOfUser()
			case <-m.fullDisconnectionChan:
				wg.Done()
				break
			default:
				fmt.Println("HUI")
			}
		}
	}()
	wg.Wait()
}

func (m *ChatsModule) isChatExists(fromEmail string, toEmail string) bool {
	_, err := m.chatsQueries.GetChatByMembers(fromEmail, toEmail)
	return err == nil
}
