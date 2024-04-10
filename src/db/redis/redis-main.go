package redis_module

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background();
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0, 
})

func Init(){
	_, err := client.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }

	client.Set(ctx, "user", "4MA", 0)

	_, err = client.Get(ctx, "user").Result()
	if err != nil{
		panic(err)
	}
}

