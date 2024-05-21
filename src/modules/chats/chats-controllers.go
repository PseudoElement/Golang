package chats

import (
	"fmt"
	"net/http"
	"os"

	api_main "github.com/pseudoelement/go-server/src/api"
	types_module "github.com/pseudoelement/go-server/src/common/types"
)

func (m *ChatsModule) _createChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "from_email", "to_email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	chatId, err := m.createNewChat(w, req, params["from_email"], params["to_email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := ChatCreatedMessage{
		ChatId: chatId,
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _deleteChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chat_id", "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: fmt.Sprintf("Chat `%v` is disconnected!", params["chat_id"]),
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _listenToChatCreationOrDeletionController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.listenToChatCreationDeletion(w, req, params["email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}
}

func (m *ChatsModule) _getMessagesInChatByIdController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chat_id")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	messages, err := m.chatsQueries.GetChatMessages(params["chat_id"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, messages, http.StatusOK)
}

func (m *ChatsModule) _htmlTemplateController(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("/app/src/views/chat.html")
	if err != nil {
		api_main.FailResponse(w, err.Error(), 400)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		api_main.FailResponse(w, err.Error(), 400)
		return
	}

	http.ServeContent(w, req, fileInfo.Name(), fileInfo.ModTime(), file)
	return
}

func (m *ChatsModule) _conectToChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chatId", "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.connectNewClientToChat(w, req, params["chatId"], params["email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}
}
