package database

import "simple-note-api/domain"

type HandlerInterface interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByName(name string) (domain.User, error)
	GetNewUserID() (int, error)
	AddUser(param domain.User) error
}
