package cards

import "net/http"

func (m *CardsModule) SetCardsRoutes() {
	m.router.HandleFunc("/healthcheck", m._healthcheckController).Methods(http.MethodGet)
	m.router.HandleFunc("/cards/get", m._getCarByIdController).Methods(http.MethodGet)
	m.router.HandleFunc("/cards/add", m._addCardController).Methods(http.MethodPost)
	m.router.HandleFunc("/cards/update", m._updateCardController).Methods(http.MethodPut)
	m.router.HandleFunc("/cards/delete", m._addCardController).Methods(http.MethodDelete)
}
