package auth_errors

import (
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

type AuthError struct {
	message string
	status  int
}

func (e *AuthError) Error() string {
	return e.message
}

func (e *AuthError) Status() int {
	return e.status
}

func UserAlreadyRegistered() errors_module.ErrorWithStatus {
	return &AuthError{message: "User is already registered!", status: http.StatusBadRequest}
}

func CantCreateUser() errors_module.ErrorWithStatus {
	return &AuthError{message: "Error occured creating new user!", status: http.StatusBadRequest}
}

func CantFindUser() errors_module.ErrorWithStatus {
	return &AuthError{message: "Can't find user!", status: http.StatusBadRequest}
}