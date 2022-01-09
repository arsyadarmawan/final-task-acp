package middlewares

import (
	"time"

	_jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(id uint) (string, error) {
	claims := _jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Hour * 6)
	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("ThisisSecretGais"))
}

func extractToken(e echo.Context) uint {
	user := e.Get("user").(*_jwt.Token)
	if user.Valid {
		claims := user.Claims.(_jwt.MapClaims)
		userId := claims["userId"].(uint)
		return userId
	}
	return 0

}
