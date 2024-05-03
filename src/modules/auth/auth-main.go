package auth

import (
	"github.com/gorilla/mux"
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	auth_db "github.com/pseudoelement/go-server/src/modules/auth/db"
)

type AuthModule struct {
	redis  *redis_main.RedisDB
	authDb auth_db.AuthDBService
	router *mux.Router
}

func NewModule(redis *redis_main.RedisDB, router *mux.Router) *AuthModule {
	return &AuthModule{
		redis:  redis,
		authDb: auth_db.NewModule(redis),
		router: router,
	}
}
