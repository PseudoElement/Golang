package main

import (
	"fmt"
	ApiService "go-server/src/api"
	Middlewares "go-server/src/api/middlewares"
	Oneinch "go-server/src/modules/1inch"
	Crud "go-server/src/modules/crud"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
    UserID string `json:"userId"`
    ID     int `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true);
	r := router.PathPrefix("/api/v1").Subrouter();
	
	Middlewares.AllowOriginsMiddleware(r);

	Oneinch.SetOneinchRoutes(r);
	Crud.SetCrudRoutes(r);

	user, err := ApiService.Post[User]("https://jsonplaceholder.typicode.com/posts", User{"229", 230, "Hello Borroq", "Sintolidze"}, map[string]string{})
	if err != nil{
		panic(err);
	}

	fmt.Println("POST_RESPONSE - ", user);

	fmt.Println("Listening port 8080...");
	log.Fatal(http.ListenAndServe(":8080", r))
}
