package RedisService

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Context = context.Background();

func Init() {
	redisApi := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, 
	})

	if err := redisApi.Set(Context, "user", "Borrow", 0).Err(); err != nil{
		panic(err);
	}

	user, err := redisApi.Get(Context, "user").Result()
	if err != nil{
		panic(err)
	}

	fmt.Println("User is ", user);
}