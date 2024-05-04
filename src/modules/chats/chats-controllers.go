package chats

import (
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	types_module "github.com/pseudoelement/go-server/src/common/types"
)

func (m *ChatsModule) _disconnectAllChats(w http.ResponseWriter, req *http.Request) {
	m.fullDisconnectionChan <- true

	msg := types_module.MessageToClient{
		Message: "All chats disconnected!",
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _createChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "from_email", "to_email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.handleChatCreation(w, req, params["from_email"], params["to_email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: "Chat successully created!",
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _deleteChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "chatId")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.handleChatDeletion(params["chatId"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	msg := types_module.MessageToClient{
		Message: "Chat successfully deleted!",
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}

func (m *ChatsModule) _listenChatsController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "from_email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.initChatListening(w, req, params["from_email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}
	m.listenAllChatsOfUser()

	msg := types_module.MessageToClient{
		Message: "All chats are listening!",
	}

	api_main.SuccessResponse(w, msg, http.StatusOK)
}
