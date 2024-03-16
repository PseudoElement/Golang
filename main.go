package main

import (
	"fmt"
	Middlewares "go-server/src/api/middlewares"
	Oneinch "go-server/src/modules/1inch"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true);
	r := router.PathPrefix("/api/v1").Subrouter();
	
	Middlewares.AllowOriginsMiddleware(r);

	Oneinch.SetOneinchRoutes(r);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(":8080", r))

}
