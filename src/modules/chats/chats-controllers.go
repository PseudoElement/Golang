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
	params, err := api_main.MapQueryParams(req, "chatId", "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.disconnectChatById(params["email"], params["chatId"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: fmt.Sprintf("Chat `%v` is disconnected!", params["chatId"]),
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _conectChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chatId", "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.connectToChatById(w, req, params["chatId"], params["email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: fmt.Sprintf("Chat `%v` is listening!", params["chatId"]),
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _listenToUpdatesController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	go m.listenToUpdates(w, req, params["email"])

	msg := types_module.MessageToClient{
		Message: "Updates are listening!",
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}
