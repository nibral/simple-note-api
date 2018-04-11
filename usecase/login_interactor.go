package usecase

import (
	"simple-note-api/domain"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginInteractor struct {
	UserRepository UserRepositoryInterface
}

func (interactor *LoginInteractor) Login(name string, password string) (domain.User, string, error) {
	user, err := interactor.UserRepository.FindByName(name)
	if err != nil {
		return domain.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, "", errors.New("name or password is incorrect")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return domain.User{}, "", err
	}

	return user, t, nil
}
