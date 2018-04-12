package usecase

import (
	"simple-note-api/domain"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginInteractor struct {
	Config         domain.Config
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
	claims["exp"] = time.Now().Add(time.Duration(interactor.Config.JwtLifetime) * time.Second).Unix()
	t, err := token.SignedString([]byte(interactor.Config.JwtSecret))
	if err != nil {
		return domain.User{}, "", err
	}

	return user, t, nil
}
