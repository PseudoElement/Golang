package main

import (
	"fmt"
	"os"
	"time"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	interfaces_module "github.com/pseudoelement/go-server/src/common/interfaces"
	postgres_main "github.com/pseudoelement/go-server/src/db/postgres"
	cards_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/cards"
	chats_queries "github.com/pseudoelement/go-server/src/db/postgres/queries/chats"
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	oneinch "github.com/pseudoelement/go-server/src/modules/1inch"
	"github.com/pseudoelement/go-server/src/modules/auth"
	"github.com/pseudoelement/go-server/src/modules/cards"
	"github.com/pseudoelement/go-server/src/modules/chats"
	"github.com/rs/cors"
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
	chatsQueries := chats_queries.NewChatsQueries(db)
	//

	//modules
	cardsModule := cards.NewModule(cardsQueries, r)
	chatsModule := chats.NewModule(chatsQueries, r)
	authModule := auth.NewModule(redisInstance, r)
	oneinchModule := oneinch.NewModule(r)
	//

	initErr := initAllTables([]postgres_main.TableCreator{
		cardsQueries,
		chatsQueries,
	})
	if initErr != nil {
		panic(initErr)
	}
	initRoutes([]interfaces_module.ModuleWithRoutes{
		cardsModule,
		chatsModule,
		authModule,
		oneinchModule,
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*", "https://websocketking.com"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		MaxAge:           10,
		Debug:            false,
	})
	handler := c.Handler(router)

	fmt.Println("Listening port 8080...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), handler))
}
