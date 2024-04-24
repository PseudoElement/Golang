package cards

import (
	"github.com/pseudoelement/go-server/src/db/postgres/queries"
)

func (m *CardsModule) convertCardToClient(cardFromDB queries.CardFromDB) CardToClient {
	return CardToClient{
		Author:    cardFromDB.Author,
		Info:      cardFromDB.Info,
		CreatedAt: cardFromDB.CreatedAt,
		UpdatedAt: cardFromDB.UpdatedAt,
		Id:        cardFromDB.Id,
	}
}
