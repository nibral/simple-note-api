package usecase

import (
	"simple-note-api/domain"
	"errors"
)

type LoginInteractor struct {
	UserRepository UserRepositoryInterface
}

func (interactor *LoginInteractor) Login(name string, password string) (domain.User, string, error) {
	user, err := interactor.UserRepository.FindByName(name)
	if err != nil {
		return domain.User{}, "", err
	}

	if password != user.Password {
		return domain.User{}, "", errors.New("name or password is incorrect")
	}

	return user, "token", nil
}
