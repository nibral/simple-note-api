package database

import "simple-note-api/domain"

type SimpleNoteDatabaseInterface interface {
	GetAllUsers() ([]domain.User, error)
}
