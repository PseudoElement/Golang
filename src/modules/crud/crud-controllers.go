package Crud

import (
	"encoding/json"
	ApiService "go-server/src/api"
	"net/http"
)

func _healthcheckController(w http.ResponseWriter, req *http.Request) {
	ApiService.SetResponseHeaders(w, req);
	w.WriteHeader(http.StatusOK);
	w.Write([]byte("Server works!"));
}

func _addPostController(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(req.Body);
}