package auth

import (
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	auth_db "github.com/pseudoelement/go-server/src/modules/auth/db"
)

type AuthModule struct {
	redis *redis_main.RedisDB
	dbSrv auth_db.AuthDBService
}

func NewModule(redis *redis_main.RedisDB) *AuthModule {
	return &AuthModule{
		redis: redis,
		dbSrv: auth_db.NewModule(redis),
	}
}
