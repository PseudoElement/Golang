package middlewares

import (
	"net/http"

	api_errors "github.com/pseudoelement/go-server/src/errors"
	auth_module "github.com/pseudoelement/go-server/src/modules/auth"
)

func AuthOnly(w http.ResponseWriter, req *http.Request)  {
	if(!auth_module.IsTokenValid(req)){
		api_errors.Unauthorized(w);
	}
}