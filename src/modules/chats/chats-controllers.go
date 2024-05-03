package chats

import (
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
)

func (m *ChatsModule) _createChatController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "from_email", "to_email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.chatQueries.CreateChat(params["from_email"], params["to_email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, "Chat created!", http.StatusOK)
}

func (m *ChatsModule) _listenChatController(w http.ResponseWriter, req *http.Request) {
	socket := NewChatSocket(w, req, m.chatQueries)
}
