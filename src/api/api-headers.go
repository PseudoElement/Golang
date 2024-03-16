package ApiService

import "net/http"

func SetResponseHeaders(w http.ResponseWriter, req *http.Request) {
	origin := req.Header.Get("Origin")
	if origin == "http://localhost:5173" || origin == "http://localhost:3000" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}