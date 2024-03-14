package oneinch

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetOneInchRoutes(router *mux.Router) {
	router.HandleFunc("/oneinch/quote", _quoteController).Methods(http.MethodGet);
	router.HandleFunc("/hello", _helloController).Methods(http.MethodGet);
}