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
	
	tokenStruct, err := auth_services.HandleRegistration(w, body);
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	api_main.SuccessResponse(w, tokenStruct, http.StatusCreated)
}

func _loginController(w http.ResponseWriter, req *http.Request){
	body, err := api_main.ParseReqBody[auth_models.UserLogin](w, req);
	if err != nil{
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	userStruct, err := auth_services.HandleLogin(w, body)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	api_main.SuccessResponse(w, userStruct, http.StatusOK);
}

func _userController(w http.ResponseWriter, req *http.Request){
	params, err := api_main.MapQueryParams(req, "email")
	if err != nil{
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	user, err := auth_db.GetUser(params["email"]);
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return;
	}

	api_main.SuccessResponse(w, user, http.StatusOK);
}

func _allUsersController(w http.ResponseWriter, req *http.Request) {
	
}