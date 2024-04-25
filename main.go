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
	cards_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/cards"
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	oneinch "github.com/pseudoelement/go-server/src/modules/1inch"
	"github.com/pseudoelement/go-server/src/modules/auth"
	"github.com/pseudoelement/go-server/src/modules/cards"
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

	redisInstance := redis_main.GetInstance()
	redisInstance.Init()
	fmt.Println("Redis started!")

	pg := postgres_main.GetInstance()
	pg.Init()
	db := pg.GetDB()
	fmt.Println("PostgreSQL started!")

	//queries
	cardsQueries := cards_queries.NewCardsQueries(db)
	//

	//modules
	cardsModule := cards.NewModule(cardsQueries, r)
	authModule := auth.NewModule(redisInstance)
	//

	initErr := initAllTables([]postgres_main.TableCreator{cardsQueries})
	if initErr != nil {
		panic(initErr)
	}

	cardsModule.SetRoutes()
	authModule.SetRoutes(r)
	oneinch.SetRoutes(r)

	fmt.Println("Listening port 8080...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
