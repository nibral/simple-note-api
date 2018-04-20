package usecase

import (
	"simple-note-api/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository UserRepositoryInterface
}

func (interactor *UserInteractor) List(sender domain.User) ([]domain.User, error) {
	return interactor.UserRepository.FindAll()
}

func (interactor *UserInteractor) Get(sender domain.User, id int) (domain.User, error) {
	if !sender.Admin && sender.ID != id {
		return domain.User{}, &NotPermittedError{}
	}

	return interactor.UserRepository.FindByID(id)
}

func (interactor *UserInteractor) Create(sender domain.User, param domain.User) (domain.User, error) {
	if !sender.Admin {
		return domain.User{}, &NotPermittedError{}
	}

	// can't create same name user
	user, err := interactor.UserRepository.FindByName(param.Name)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID != 0 {
		return domain.User{}, &InvalidParameterError{"Specified name already exists"}
	}

	// password save as hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}
	param.Password = string(hashedPassword)

	return interactor.UserRepository.Add(param)
}

func (interactor *UserInteractor) Update(sender domain.User, id int, param domain.User) (domain.User, error) {
	if !sender.Admin && sender.ID != id {
		return domain.User{}, &NotPermittedError{}
	}

	// get stored user
	idUser, err := interactor.UserRepository.FindByID(id)
	if err != nil {
		return domain.User{}, err
	}
	if idUser.ID == 0 {
		return domain.User{}, &InvalidParameterError{"Specified user doesn't exists"}
	}

	// can't create same name user
	nameUser, err := interactor.UserRepository.FindByName(param.Name)
	if err != nil {
		return domain.User{}, err
	}
	if nameUser.ID != id && nameUser.ID != 0 {
		return domain.User{}, &InvalidParameterError{"Specified name already used"}
	}

	// can't degrade myself
	if sender.ID == id && idUser.Admin == true && param.Admin == false {
		return domain.User{}, &InvalidParameterError{"Can't degrade myself"}
	}

	param.ID = id
	if param.Password == "" {
		param.Password = idUser.Password
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
		if err != nil {
			return domain.User{}, err
		}
		param.Password = string(hashedPassword)
	}

	return interactor.UserRepository.Update(id, param)
}
