package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	redis_module "github.com/pseudoelement/go-server/src/db/redis"
	oneinch "github.com/pseudoelement/go-server/src/modules/1inch"
	auth_module "github.com/pseudoelement/go-server/src/modules/auth"
	crud "github.com/pseudoelement/go-server/src/modules/crud"
)


func main() {
	router := mux.NewRouter().StrictSlash(true);
	r := router.PathPrefix("/api/v1").Subrouter();

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	
	redis_module.Init();

	token, _ := auth_module.CreateToken(10);

	fmt.Println("Token is ", token);
	
	oneinch.SetOneinchRoutes(r);
	crud.SetcrudRoutes(r);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(":8080", r))
}
