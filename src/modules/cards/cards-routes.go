package cards

import (
	"fmt"
	"net/http"
)

func (m *CardsModule) SetRoutes() {
	m.router.HandleFunc("/cards/get", m._getCarByIdController).Methods(http.MethodGet)
	m.router.HandleFunc("/cards/get-all", m._getAllSortedCardController).Methods(http.MethodGet)
	m.router.HandleFunc("/cards/add", m._addCardController).Methods(http.MethodPost)
	m.router.HandleFunc("/cards/update", m._updateCardController).Methods(http.MethodPut)
	m.router.HandleFunc("/cards/delete", m._addCardController).Methods(http.MethodDelete)

	fmt.Println("CardsModule started!")
}
