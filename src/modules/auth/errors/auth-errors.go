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

func InvalidPassword() errors_module.ErrorWithStatus{
	return &AuthError{message: "Invalid password!", status: http.StatusBadRequest}
}