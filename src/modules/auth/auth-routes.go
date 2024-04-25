package auth_main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", _healthcheckController).Methods(http.MethodGet)
	router.HandleFunc("/auth/user", _userController).Methods(http.MethodGet)
	router.HandleFunc("/auth/users", _allUsersController).Methods(http.MethodGet)
	router.HandleFunc("/auth/register", _registrationController).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", _loginController).Methods(http.MethodPost)
}
