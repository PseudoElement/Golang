package auth

import (
	"fmt"
	"net/http"
)

func (m *AuthModule) SetRoutes() {
	m.router.HandleFunc("/healthcheck", m._healthcheckController).Methods(http.MethodGet)
	m.router.HandleFunc("/auth/user", m._userController).Methods(http.MethodGet)
	m.router.HandleFunc("/auth/users", m._allUsersController).Methods(http.MethodGet)
	m.router.HandleFunc("/auth/register", m._registrationController).Methods(http.MethodPost)
	m.router.HandleFunc("/auth/login", m._loginController).Methods(http.MethodPost)

	fmt.Println("AuthModule started!")
}
