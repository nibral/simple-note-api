package usecase

import "simple-note-api/domain"

type UserInteractor struct {
	UserRepository UserRepositoryInterface
}

type UserCreateError struct {
	Msg string
}

func (_ *UserCreateError) Error() string {
	return "Specified name already exists"
}

func (interactor *UserInteractor) Users() ([]domain.User, error) {
	users, err := interactor.UserRepository.FindAll()
	return users, err
}

func (interactor *UserInteractor) Create(param domain.User) (domain.User, error) {
	user, err := interactor.UserRepository.FindByName(param.Name)
	if err != nil {
		return domain.User{}, nil
	}
	if user.ID != 0 {
		return domain.User{}, &UserCreateError{}
	}

	user, err = interactor.UserRepository.Add(param)
	return user, err
}
