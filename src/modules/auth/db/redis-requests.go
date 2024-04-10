package auth_db

import (
	redis_main "github.com/pseudoelement/go-server/src/db/redis"
	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_errors "github.com/pseudoelement/go-server/src/modules/auth/errors"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
)

func SaveNewUser(body auth_models.UserRegister) (errors_module.ErrorWithStatus){
	if IsUserAlreadyRegistered(body.Email) {
		return auth_errors.UserAlreadyRegistered()
	}

	err := redis_main.SetStruct(body.Email, body);
	if err != nil{
		return auth_errors.CantCreateUser();
	}

	return nil;
}

func GetUser(email string) (auth_models.UserRegister, errors_module.ErrorWithStatus) {
	user, err := redis_main.GetStruct[auth_models.UserRegister](email);
	if err != nil{
		return auth_models.UserRegister{}, auth_errors.CantFindUser();
	}
	
	return user, nil;
}

func IsUserAlreadyRegistered(email string) bool {
	 _, err := redis_main.GetStruct[auth_models.UserRegister](email);
	 return err == nil;
}