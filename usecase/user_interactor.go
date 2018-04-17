package usecase

import "simple-note-api/domain"

type UserInteractor struct {
	UserRepository UserRepositoryInterface
}

type UserCreateError struct {
	Msg string
}

func (_ *UserCreateError) Error() string {
	return "Failed to create user"
}

func (interactor *UserInteractor) Users(sender domain.User) ([]domain.User, error) {
	users, err := interactor.UserRepository.FindAll()
	return users, err
}

func (interactor *UserInteractor) Create(sender domain.User, param domain.User) (domain.User, error) {
	if !sender.Admin {
		return domain.User{}, &NotPermittedError{}
	}

	user, err := interactor.UserRepository.FindByName(param.Name)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID != 0 {
		return domain.User{}, &UserCreateError{"Specified name already exists"}
	}

	user, err = interactor.UserRepository.Add(param)
	return user, err
}
