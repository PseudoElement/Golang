package auth_main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth/register", _registrationController).Methods(http.MethodPost);
	router.HandleFunc("/auth/login", _loginController).Methods(http.MethodPost);
	// router.HandleFunc("/auth/refresh-token", _getAllowanceController).Methods(http.MethodGet);
}