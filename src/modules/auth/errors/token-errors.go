package auth_errors

import (
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
)

func CantCreateToken() errors_module.ErrorWithStatus {
	return &AuthError{message: "Error occured creating  jwt-token!", status: http.StatusBadRequest}
}