package cards

import (
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
)

func (m *CardsModule) _addCardController(w http.ResponseWriter, req *http.Request) {
	body, err := api_main.ParseReqBody[NewCard](w, req)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.cq.AddCard(body.Author, body.Info)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	res := CardActionSuccess{
		Message: "Card successfully saved.",
	}

	api_main.SuccessResponse(w, res, http.StatusCreated)
}

func (m *CardsModule) _updateCardController(w http.ResponseWriter, req *http.Request) {
	body, err := api_main.ParseReqBody[CardUpdate](w, req)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.cq.UpdateCard(body.Id, body.Author, body.Info)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	res := CardActionSuccess{
		Message: "Card successfully updated.",
	}

	api_main.SuccessResponse(w, res, http.StatusCreated)
}

func (m *CardsModule) _deleteCardController(w http.ResponseWriter, req *http.Request) {
	body, err := api_main.ParseReqBody[CardDelete](w, req)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.cq.DeleteCard(body.Id)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	res := CardActionSuccess{
		Message: "Card successfully deleted.",
	}

	api_main.SuccessResponse(w, res, http.StatusCreated)
}

func (m *CardsModule) _getCarByIdController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "id")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	cardDb, err := m.cq.GetCard(params["id"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	cardToClient := m.convertCardToClient(cardDb)

	api_main.SuccessResponse(w, cardToClient, http.StatusOK)
}

func (m *CardsModule) _getAllSortedCardController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "sortBy", "sortDir", "page", "limitPerPage")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	cards, err := m.getSortedCards(params)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, cards, http.StatusOK)
}
