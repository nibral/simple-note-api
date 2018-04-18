package database

import "simple-note-api/domain"

type HandlerInterface interface {
	GetAllUsers() ([]domain.User, error)
	GetNewUserID() (int, error)
	GetUserByName(name string) (domain.User, error)
	GetUserCountByName(name string) (int, error)
	AddUser(param domain.User) error
}
