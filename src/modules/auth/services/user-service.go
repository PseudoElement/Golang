package auth_services

import (
	"fmt"
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_db "github.com/pseudoelement/go-server/src/modules/auth/db"
	auth_errors "github.com/pseudoelement/go-server/src/modules/auth/errors"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
)

func HandleRegistration(w http.ResponseWriter, user auth_models.UserRegister) (auth_models.Token, errors_module.ErrorWithStatus) {
	encryptedPassword := EncryptPassword(user.Password);
	fmt.Println("Encrypted - ", encryptedPassword);

	err := auth_db.SaveNewUser(auth_models.UserRegister{
		Name: user.Name, 
		Email: user.Email, 
		Password:  encryptedPassword,
	});
	if err != nil {
		return auth_models.Token{}, err;
	}

	token, err := CreateToken(40);
	if err != nil {
		return auth_models.Token{}, err;
	}

	tokenStruct := auth_models.Token{
		Token: token,
	}

	return tokenStruct, nil;
}

func HandleLogin(w http.ResponseWriter, body auth_models.UserLogin) (auth_models.UserWithToken, errors_module.ErrorWithStatus) {
	user, err := auth_db.GetUser(body.Email)
	if err != nil {
		return auth_models.UserWithToken{}, err;
	}
	fmt.Println("Hash - ", user.Password)
	fmt.Println("Password - ", body.Password)

	if !IsPasswordValid(user.Password, body.Password) {
		return auth_models.UserWithToken{}, auth_errors.InvalidPassword();
	}

	token, err := CreateToken(40);
	if err != nil {
		return auth_models.UserWithToken{}, err;
	}

	userStruct := auth_models.UserWithToken{
		Token: token,
		Name: user.Name,
		Email: user.Email,
	}

	return userStruct, nil;
}