package auth_db

import (
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_errors "github.com/pseudoelement/go-server/src/modules/auth/errors"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
)

type AuthDBService struct {
	redis *redis_main.RedisDB
}

func NewModule(redis *redis_main.RedisDB) AuthDBService {
	return AuthDBService{
		redis: redis,
	}
}

func (s *AuthDBService) SaveNewUser(body auth_models.UserRegister) errors_module.ErrorWithStatus {
	if s.isUserAlreadyRegistered(body.Email) {
		return auth_errors.UserAlreadyRegistered()
	}

	err := s.redis.SetStruct(body.Email, body)
	if err != nil {
		return auth_errors.CantCreateUser()
	}

	return nil
}

func (s *AuthDBService) GetUser(email string) (auth_models.UserRegister, errors_module.ErrorWithStatus) {
	var user auth_models.UserRegister
	if err := s.redis.GetStruct(email, &user); err != nil {
		return auth_models.UserRegister{}, auth_errors.CantFindUser()
	}

	return user, nil
}

func (s *AuthDBService) isUserAlreadyRegistered(email string) bool {
	var user auth_models.UserRegister
	err := s.redis.GetStruct(email, &user)
	return err == nil
}
