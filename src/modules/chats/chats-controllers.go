package chats

import (
	"fmt"
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	types_module "github.com/pseudoelement/go-server/src/common/types"
)

func (m *ChatsModule) _createChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "from_email", "to_email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.createNewChat(w, req, params["from_email"], params["to_email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: fmt.Sprintf("Chat is created! Members: %v, %v.", params["from_email"], params["to_email"]),
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _deleteChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chat_id", "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.disconnectChatById(params["email"], params["chat_id"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: fmt.Sprintf("Chat `%v` is disconnected!", params["chat_id"]),
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _conectChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chat_id", "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.connectToChatById(w, req, params["chat_id"], params["email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: fmt.Sprintf("Chat `%v` is listening!", params["chat_id"]),
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _listenToUpdatesController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.listenToUpdates(w, req, params["email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: "Updates are listening!",
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
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
