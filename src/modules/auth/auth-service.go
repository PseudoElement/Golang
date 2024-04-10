package auth_module

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(expirationMin time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims);
	claims["exp"] = time.Now().Add(expirationMin * time.Minute);
	claims["authorized"] = true

	secret := []byte(os.Getenv("JWT_SECRET_KEY"));

	token_string, err := token.SignedString(secret);
	if err != nil {
		fmt.Println(err)
    	return "", err
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

	 if err != nil {
		return false;
	 }
	
	 return true;
}