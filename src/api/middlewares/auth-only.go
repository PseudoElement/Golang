package middlewares

import (
	"net/http"

	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_services "github.com/pseudoelement/go-server/src/modules/auth/services"
)

func AuthOnly(w http.ResponseWriter, req *http.Request)  {
	if(!auth_services.IsTokenValid(req)){
		errors_module.Unauthorized();
	}
}