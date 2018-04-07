package database

import "simple-note-api/domain"

type DatabaseHandlerInterface interface {
	GetAllUsers() ([]domain.User, error)
}
