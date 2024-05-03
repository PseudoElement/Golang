package auth

import (
	"net/http"

	api_main "github.com/pseudoelement/go-server/src/api"
	auth_models "github.com/pseudoelement/go-server/src/modules/auth/models"
)

func (m *AuthModule) _healthcheckController(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is alive!"))
}

func (m *AuthModule) _registrationController(w http.ResponseWriter, req *http.Request) {
	body, err := api_main.ParseReqBody[auth_models.UserRegister](w, req)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	tokenStruct, err := m.handleRegistration(w, body)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, tokenStruct, http.StatusCreated)
}

func (m *AuthModule) _loginController(w http.ResponseWriter, req *http.Request) {
	body, err := api_main.ParseReqBody[auth_models.UserLogin](w, req)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	userStruct, err := m.handleLogin(w, body)
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, userStruct, http.StatusOK)
}

func (m *AuthModule) _userController(w http.ResponseWriter, req *http.Request) {
	params, err := api_main.MapQueryParams(req, "email")
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	user, err := m.authDb.GetUser(params["email"])
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, user, http.StatusOK)
}

func (m *AuthModule) _allUsersController(w http.ResponseWriter, req *http.Request) {
	users, err := m.redis.GetAllUsers()
	if err != nil {
		api_main.FailResponse(w, err.Error(), err.Status())
		return
	}

	api_main.SuccessResponse(w, users, http.StatusOK)
}
