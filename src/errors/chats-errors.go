package errors_module

import (
	"fmt"
	"net/http"
)

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

func ForbiddenConnectionToChat() ErrorWithStatus {
	return &ApiError{
		message: "You can't connect to this chat!",
		status:  http.StatusBadRequest,
	}
}

func ChatDefaultError(msg string) ErrorWithStatus {
	return &ApiError{
		message: msg,
		status:  http.StatusBadRequest,
	}
}

func SocketConnectionError(socketName string) ErrorWithStatus {
	return &ApiError{
		message: fmt.Sprintf("Can't init connection to socket `%v`!", socketName),
		status:  http.StatusBadRequest,
	}
}
