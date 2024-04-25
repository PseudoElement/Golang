package cards

import (
	cards_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/cards"
)

func (m *CardsModule) convertCardToClient(cardFromDB cards_queries.CardFromDB) CardToClient {
	return CardToClient{
		Author:    cardFromDB.Author,
		Info:      cardFromDB.Info,
		CreatedAt: cardFromDB.CreatedAt,
		UpdatedAt: cardFromDB.UpdatedAt,
		Id:        cardFromDB.Id,
	}
}
