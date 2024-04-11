package auth_services

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	errors_module "github.com/pseudoelement/go-server/src/errors"
	auth_errors "github.com/pseudoelement/go-server/src/modules/auth/errors"
)

func CreateToken(expirationMin time.Duration) (string, errors_module.ErrorWithStatus) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims);
	claims["exp"] = time.Now().Add(expirationMin * time.Minute);
	claims["authorized"] = true

	secret := []byte(os.Getenv("JWT_SECRET"));

	token_string, err := token.SignedString(secret);
	if err != nil {
		fmt.Println(err)
    	return "", auth_errors.CantCreateToken();
 	}

	return token_string, nil
}

/* Returns true if token valid, otherwise - false */
func IsTokenValid(req *http.Request) bool {
	token_string, ok := req.Header["Authorization"];
	if !ok{
		return false;
	}

	_, err := jwt.Parse(token_string[0], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return "", errors.New("Unauthorized!");
		}
		return "", nil

	 })

	 return err == nil;
}

func IsTokenValid2(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return "", errors.New("unauthorized");
		}
		return "", nil

	 })

	 return err == nil;
}