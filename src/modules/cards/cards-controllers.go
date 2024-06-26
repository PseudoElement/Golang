package cards

import (
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	types_module "github.com/pseudoelement/go-server/src/common/types"
)

func (m *CardsModule) _addCardController(w http.ResponseWriter, req *http.Request) {
	body, err := api_main.ParseReqBody[NewCard](w, req)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	err = m.cardsQueries.AddCard(body.Author, body.Info)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	res := types_module.MessageToClient{
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

	err = m.cardsQueries.UpdateCard(body.Id, body.Author, body.Info)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	res := types_module.MessageToClient{
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

	err = m.cardsQueries.DeleteCard(body.Id)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	res := types_module.MessageToClient{
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

	cardDb, err := m.cardsQueries.GetCard(params["id"])
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
