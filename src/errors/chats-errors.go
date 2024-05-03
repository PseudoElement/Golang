package errors_module

import "net/http"

func ChatAlreadyCreated() ErrorWithStatus {
	return &ApiError{
		message: "Chat already exists!",
		status:  http.StatusBadRequest,
	}
}
