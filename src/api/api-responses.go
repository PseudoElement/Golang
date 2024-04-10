package api_main

import (
	"encoding/json"
	"net/http"
)

type failRes struct{
	Message string `json:"message"`
}

func SuccessResponse[T any](w http.ResponseWriter, body T, status int) {
	w.WriteHeader(status);
	json.NewEncoder(w).Encode(body);
}

func FailResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status);
	var res = &failRes{Message: message}
	json.NewEncoder(w).Encode(res)
}