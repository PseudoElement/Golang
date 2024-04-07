package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	redis_service "github.com/pseudoelement/go-server/src/db/redis"
	oneinch "github.com/pseudoelement/go-server/src/modules/1inch"
	crud "github.com/pseudoelement/go-server/src/modules/crud"
)


func main() {
	router := mux.NewRouter().StrictSlash(true);
	r := router.PathPrefix("/api/v1").Subrouter();
	
	redis_service.Init();
	
	oneinch.SetOneinchRoutes(r);
	crud.SetcrudRoutes(r);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(":8080", r))
}
