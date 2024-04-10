package auth_main

import (
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	auth_db "github.com/pseudoelement/go-server/src/modules/auth/db"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
	auth_services "github.com/pseudoelement/go-server/src/modules/auth/services"
)

func _registrationController(w http.ResponseWriter, req *http.Request){
	body, err := api_main.ParseReqBody[auth_models.UserRegister](w, req);
	if err != nil{
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}
	
	err = auth_db.SaveNewUser(body);
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	token, err := auth_services.CreateToken(40);
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	tokenStruct := auth_models.Token{
		Token: token,
	}

	api_main.SuccessResponse(w, tokenStruct, http.StatusCreated)
}

func _loginController(w http.ResponseWriter, req *http.Request){
	body, err := api_main.ParseReqBody[auth_models.UserLogin](w, req);
	if err != nil{
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	user, err := auth_db.GetUser(body.Email)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	token, err := auth_services.CreateToken(40);
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	userStruct := auth_models.UserWithToken{
		Token: token,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	api_main.SuccessResponse(w, userStruct, http.StatusOK);
}