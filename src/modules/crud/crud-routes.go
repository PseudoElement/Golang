package Crud

import "github.com/gorilla/mux"

func SetCrudRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", _healthcheckController)
}