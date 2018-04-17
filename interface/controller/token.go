package controller

import (
	"simple-note-api/domain"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

func ParseToken(context echo.Context) (domain.User) {
	token := context.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user := domain.User{
		ID:    int(claims["id"].(float64)),
		Name:  claims["name"].(string),
		Admin: claims["admin"].(bool),
	}

	return user
}
