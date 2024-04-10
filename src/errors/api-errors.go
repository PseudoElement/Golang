package api_errors

import (
	"encoding/json"
	"net/http"
)

type ApiError struct{
	Message string
}

func(e *ApiError) Error() string{
	return e.Message;
}

func IncorrectQueryParams(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest);
	json.NewEncoder(w).Encode(ApiError{Message: "Incorrect query parameters!"})
}

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized);
	json.NewEncoder(w).Encode(ApiError{Message: "Unauthorized!"})
}

func IncorrectBody(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest);
	json.NewEncoder(w).Encode(ApiError{Message: "Incorrect request body!"})
}