package main

import (
	"fmt"
	"os"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	oneinch "github.com/pseudoelement/go-server/src/modules/1inch"
	auth_main "github.com/pseudoelement/go-server/src/modules/auth"
	crud "github.com/pseudoelement/go-server/src/modules/crud"
)


func main() {
	router := mux.NewRouter().StrictSlash(true);
	r := router.PathPrefix("/api/v1").Subrouter();

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	
	redis_main.Init();

	oneinch.SetOneinchRoutes(r);
	crud.SetCrudRoutes(r);
	auth_main.SetAuthRoutes(r);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
