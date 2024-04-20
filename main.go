package main

import (
	"fmt"
	"os"
	"time"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	postgres_main "github.com/pseudoelement/go-server/src/db/postgres"
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

	fmt.Println("Server started!");
	
	time.Sleep(2 * time.Second);
	redis_main.Init();
	fmt.Println("Redis started!");
	postgres_main.Init();
	fmt.Println("PostgreSQL started!");

	oneinch.SetOneinchRoutes(r);
	crud.SetCrudRoutes(r);
	auth_main.SetAuthRoutes(r);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
