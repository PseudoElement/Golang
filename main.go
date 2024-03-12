package main

import (
	oneinch "go-server/src/modules/1inch"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//use deps from another module
	// filtered := utils.Filter(slice, func(num int, ind int) bool {
	// 	return num > 1;
	// })
	router := mux.NewRouter().StrictSlash(true);
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/oneinch/quote", oneinch.QuoteController).Methods("GET");

	log.Fatal(http.ListenAndServe(":8080", api))

}
