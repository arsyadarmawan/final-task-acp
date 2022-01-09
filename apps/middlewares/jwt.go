package middlewares

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func generateToken(id uint) (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["userid"] = id
	token, err := sign.SignedString([]byte("thisissecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func Auth(c *echo.Context) bool {

	tokenString := echo.HeaderAuthorization
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("thisissecret"), nil
	})

	if token != nil && err == nil {
		return true
	} else {
		return false
	}
}

func ExtractClaims() (jwt.MapClaims, bool) {
	tokenString := echo.HeaderAuthorization
	hmacSecretString := "thisissecret" // Value
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		// log.Printf("Invalid JWT Token")
		return nil, false
	}
}
