package middleware

import (
	"time"
	"workuo/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_KEY))
}

func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims = user.Claims.(jwt.MapClaims)
	}

	return
}
