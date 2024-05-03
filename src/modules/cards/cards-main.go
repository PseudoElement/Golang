package cards

import (
	"github.com/gorilla/mux"
	cards_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/cards"
)

type CardsModule struct {
	cardsQueries *cards_queries.CardsQueries
	router       *mux.Router
}

func NewModule(cardsQueries *cards_queries.CardsQueries, router *mux.Router) *CardsModule {
	return &CardsModule{cardsQueries: cardsQueries, router: router}
}
