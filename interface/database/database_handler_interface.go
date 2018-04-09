package database

import "simple-note-api/domain"

type HandlerInterface interface {
	GetAllUsers() ([]domain.User, error)
}
