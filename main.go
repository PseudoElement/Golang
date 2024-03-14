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
	api := router.PathPrefix("/api/v1").Subrouter();

	oneinch.SetOneInchRoutes(api);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(":8080", api))

}
