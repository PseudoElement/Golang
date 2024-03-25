package Crud

import "github.com/gorilla/mux"

func SetCrudRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", _healthcheckController);
	router.HandleFunc("/posts/add-post", _addPostController);
}