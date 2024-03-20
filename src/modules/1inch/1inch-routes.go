package Oneinch

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetOneinchRoutes(router *mux.Router) {
	router.HandleFunc("/oneinch/quote", _quoteController).Methods(http.MethodGet);
	router.HandleFunc("/oneinch/swap", _swapController).Methods(http.MethodGet);
	router.HandleFunc("/oneinch/allowance", _getAllowanceController).Methods(http.MethodGet);
	router.HandleFunc("/oneinch/approve", _getApproveConfigController).Methods(http.MethodGet);
	router.HandleFunc("/hello", _helloController).Methods(http.MethodGet);
}