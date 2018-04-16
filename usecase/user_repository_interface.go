package usecase

import "simple-note-api/domain"

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
	FindByName(name string) (domain.User, error)
	Add(param domain.User) (domain.User, error)
}
