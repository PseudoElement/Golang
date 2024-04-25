package redis_main

import (
	"context"

	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Address  string
	Password string
	DB       int
	ctx      context.Context
	client   *redis.Client
}

func GetInstance() *RedisDB {
	return &RedisDB{
		Address:  "redis:6379",
		Password: "",
		DB:       0,
	}
}

func (r *RedisDB) Init() {
	client := redis.NewClient(&redis.Options{
		// without docker-compose.yml Addr:     "localhost:6379",
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.DB,
	})
	r.client = client
	r.ctx = context.Background()

	_, err := r.client.Ping(r.ctx).Result()
	if err != nil {
		panic(err)
	}

	r.client.Set(r.ctx, "user", auth_models.UserRegister{
		Name:     "Chma",
		Email:    "4ma@mail.ru",
		Password: "4ma228",
	}, 0)

	_, err = client.Get(r.ctx, "user").Result()
	if err != nil {
		panic(err)
	}
}
