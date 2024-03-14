package main

import (
	"fmt"
	oneinch "go-server/src/modules/1inch"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true);
	r := router.PathPrefix("/api/v1").Subrouter();
	
	// ApiService.AllowOriginsMiddleware(r);

	oneinch.SetOneInchRoutes(r);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(":8080", r))

}
