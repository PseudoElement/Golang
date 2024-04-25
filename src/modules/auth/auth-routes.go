package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (m *AuthModule) SetRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", m._healthcheckController).Methods(http.MethodGet)
	router.HandleFunc("/auth/user", m._userController).Methods(http.MethodGet)
	router.HandleFunc("/auth/users", m._allUsersController).Methods(http.MethodGet)
	router.HandleFunc("/auth/register", m._registrationController).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", m._loginController).Methods(http.MethodPost)

	fmt.Println("AuthModule started!")
}
