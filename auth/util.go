package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type BearerClaims struct {
	jwt.RegisteredClaims
}

func GetJwtToken(c echo.Context) *jwt.Token {
	cookie, err := c.Cookie("Bearer")
	if err != nil {
		return nil
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &BearerClaims{}, noKeyFunc)
	if err != nil {
		return nil
	}
	return token
}

func noKeyFunc(t *jwt.Token) (interface{}, error) {
	return []byte{}, nil
}
