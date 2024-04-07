package redis_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Context = context.Background();
var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0, 
})

func Init() *redis.Client {
	if err := Client.Set(Context, "user", "Borrow", 0).Err(); err != nil{
		panic(err);
	}

	user, err := Client.Get(Context, "user").Result()
	if err != nil{
		panic(err)
	}

	fmt.Println("User is ", user);

	return Client;
}

func GetValue[T any](key string) (T, error) {
	val, err := Client.Get(Context, key).Result()
	res_generic := new(T);

	if(err == redis.Nil){
		return *res_generic, errors.New("value equals to null")
	}else if(val == ""){
		return *res_generic, errors.New("value is empty")
	}else if(err != nil){
		return *res_generic, err;
	}

	if err := json.Unmarshal([]byte(val), &res_generic); err != nil {
		return *res_generic, err;
	}
	return *res_generic, nil;
} 