package usecase

import "simple-note-api/domain"

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
	FindByID(id int) (domain.User, error)
	FindByName(name string) (domain.User, error)
	Add(user domain.User) (domain.User, error)
	Update(id int, user domain.User) (domain.User, error)
	Delete(id int) error
}
