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
	"github.com/pseudoelement/go-server/src/db/postgres/queries"
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	oneinch "github.com/pseudoelement/go-server/src/modules/1inch"
	auth_main "github.com/pseudoelement/go-server/src/modules/auth"
	crud "github.com/pseudoelement/go-server/src/modules/crud"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	r := router.PathPrefix("/api/v1").Subrouter()

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started!")
	time.Sleep(1 * time.Second)

	redis_main.Init()
	fmt.Println("Redis started!")

	pg := postgres_main.GetInstance()
	pg.Init()
	db := pg.GetDB()
	fmt.Println("PostgreSQL started!")

	q := queries.GetInstance(db)
	err = q.CreateTable()
	if err != nil {
		panic(err)
	}
	q.AddCard("Korben Dallas", "Tanki govno!!!")

	oneinch.SetOneinchRoutes(r)
	crud.SetCrudRoutes(r)
	auth_main.SetAuthRoutes(r)

	fmt.Println("Listening port 8080...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
