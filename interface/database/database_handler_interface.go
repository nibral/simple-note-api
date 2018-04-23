package database

import "simple-note-api/domain"

type HandlerInterface interface {
	GetAllUsers() ([]domain.User, error)
	GetNewUserID() (int, error)
	GetUserByID(id int) (domain.User, error)
	GetUserByName(name string) (domain.User, error)
	GetUserCountByID(id int) (int, error)
	GetUserCountByName(name string) (int, error)
	PutUser(user domain.User) error
	UpdateUser(id int, user domain.User) error
	DeleteUser(id int) error
}
