package usecase

import "simple-note-api/domain"

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
}
