package errors_module

import "net/http"

func ChatAlreadyCreated() ErrorWithStatus {
	return &ApiError{
		message: "Chat already exists!",
		status:  http.StatusBadRequest,
	}
}

func ChatNotFound() ErrorWithStatus {
	return &ApiError{
		message: "Chat not found by id!",
		status:  http.StatusBadRequest,
	}
}

func ChatDefaultError(msg string) ErrorWithStatus {
	return &ApiError{
		message: msg,
		status:  http.StatusBadRequest,
	}
}
