package usecase

import (
	"fmt"

	"simple-note-api/domain"

	"golang.org/x/crypto/bcrypt"
)

type LoginInteractor struct {
	UserRepository UserRepositoryInterface
}

func (interactor *LoginInteractor) Login(name string, password string) (domain.User, error) {
	user, err := interactor.UserRepository.FindByName(name)
	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, fmt.Errorf("name or password is incorrect: %v, %v", name, password)
	}

	return user, nil
}
