package controller

import (
	"simple-note-api/domain"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(token *jwt.Token) (domain.User) {
	claims := token.Claims.(jwt.MapClaims)
	user := domain.User{
		ID:    int(claims["id"].(float64)),
		Name:  claims["name"].(string),
		Admin: claims["admin"].(bool),
	}

	return user
}
