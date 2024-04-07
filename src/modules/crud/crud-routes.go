package crud

import "github.com/gorilla/mux"

func SetcrudRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", _healthcheckController);
	router.HandleFunc("/posts/add-post", _addPostController);
}