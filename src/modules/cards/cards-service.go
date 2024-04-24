package cards

import (
	"strconv"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func (m *CardsModule) getSortedCards(params map[string]string) ([]CardToClient, errors_module.ErrorWithStatus) {
	page, _ := strconv.Atoi(params["page"])
	limitPerPage, _ := strconv.Atoi(params["limitPerPage"])

	cardsDB, err := m.cq.GetAllSortedCard(
		params["sortBy"],
		params["sortDir"],
		page,
		limitPerPage)
	if err != nil {
		return nil, err
	}

	var cardsToClient []CardToClient
	for _, cardDb := range cardsDB {
		cardsToClient = append(cardsToClient, m.convertCardToClient(cardDb))
	}

	return cardsToClient, nil
}
