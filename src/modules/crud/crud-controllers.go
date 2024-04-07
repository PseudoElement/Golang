package crud

import (
	"encoding/json"
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
)

func _healthcheckController(w http.ResponseWriter, req *http.Request) {
	api_main.SetResponseHeaders(w, req);
	w.WriteHeader(http.StatusOK);
	w.Write([]byte("Server works!"));
}

func _addPostController(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(req.Body);
}