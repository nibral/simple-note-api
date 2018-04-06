package usecase

import "simple-note-api/domain"

type UserInteractor struct {
	UserRepository UserRepositoryInterface
}

func (interactor *UserInteractor) Users() ([]domain.User, error) {
	users, err := interactor.UserRepository.FindAll()
	return users, err
}
